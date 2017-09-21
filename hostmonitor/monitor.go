package main

import (
	"fmt"
	"github.com/Foglight/foglight-rest-api-go-sample-code/rest"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func main() {
	conf := rest.NewConfig()
	client := rest.NewClient(conf)
	var lastSubmissionMs uint64
	for {
		now := rest.Now()
		data := collect()
		if lastSubmissionMs <= 0 {
			lastSubmissionMs = now - uint64(conf.CollectionIntervalMinutes*600000)
		}
		client.Submit(data, lastSubmissionMs, now)
		lastSubmissionMs = now
		time.Sleep(time.Duration(conf.CollectionIntervalMinutes) * time.Minute)
	}
}

func collect() *rest.SimpleTopology {
	data := rest.NewSimpleTopology("SimpleHost")
	hostInfo(data)
	cpuPercent(data)
	memPercent(data)
	netStat(data)
	diskStat(data)
	return data
}

func hostInfo(h *rest.SimpleTopology) {
	v, _ := host.Info()
	h.AddString("hostname", v.Hostname)
	h.AddString("name", v.Hostname)
	h.AddString("os", fmt.Sprintf("%s - %s(%s)", v.Platform, v.OS, v.KernelVersion))
}

func cpuPercent(h *rest.SimpleTopology) {
	values, _ := cpu.Percent(0, false)
	if values != nil {
		h.AddFloat("cpuPercent", values[0])
	}
}
func memPercent(h *rest.SimpleTopology) {
	vm, _ := mem.VirtualMemory()
	if vm != nil {
		h.AddFloat("memPercent", vm.UsedPercent)
	}
}

var (
	lastTime, lastSent, lastRecv uint64
)

func netStat(h *rest.SimpleTopology) {
	vs, err := net.IOCounters(false)
	if err == nil {
		v := vs[0]
		currTime := rest.Now()
		if lastTime == 0 {
			lastTime, lastSent, lastRecv = currTime, v.BytesSent, v.BytesRecv
			return
		}
		usedTime := currTime - lastTime
		if usedTime >= 0 {
			h.AddFloat("netSentRate", float64(subAbs(lastSent, v.BytesSent))/float64(usedTime))
			h.AddFloat("netRecvRate", float64(subAbs(lastRecv, v.BytesRecv))/float64(usedTime))
		}
		lastTime, lastSent, lastRecv = currTime, v.BytesSent, v.BytesRecv
	}
}
func subAbs(a uint64, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}

func diskStat(h *rest.SimpleTopology) {
	var total, free uint64
	p, err := disk.Partitions(false)
	if err == nil {
		for _, d := range p {
			usage, err := disk.Usage(d.Mountpoint)
			if err == nil {
				total += usage.Total
				free += usage.Free
			}
		}
		if total > 0 {
			h.AddFloat("diskTotal", float64(total))
			h.AddFloat("diskFree", float64(free))
		}
	}
}
