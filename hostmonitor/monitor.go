package main

import (
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

func collect() *Host {
	data := &Host{}
	hostInfo(data)
	cpuPercent(data)
	memPercent(data)
	netStat(data)
	diskStat(data)
	return data
}

func hostInfo(h *Host) {
	v, _ := host.Info()
	h.Hostname = v.Hostname
	h.HostID = v.HostID
}

func cpuPercent(h *Host) {
	values, _ := cpu.Percent(0, false)
	if values != nil {
		h.CPUPercent = values[0]
	}
}
func memPercent(h *Host) {
	vm, _ := mem.VirtualMemory()
	if vm != nil {
		h.MemPercent = vm.UsedPercent
	}
}

var (
	lastTime, lastSent, lastRecv uint64
)

func netStat(h *Host) {
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
			rate := float64(subAbs(lastSent, v.BytesSent)+subAbs(lastRecv, v.BytesRecv)) / float64(usedTime)
			speed := float64(1000 * 1000)
			h.NetPercent = rate * 800.0 / speed
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

func diskStat(h *Host) {
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
			h.DiskPercent = 100.0 * float64(total-free) / float64(total)
		}
	}
}
