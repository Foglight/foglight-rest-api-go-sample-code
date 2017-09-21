package main

import (
	"bytes"
	"text/template"
)

//Host Host topology object
type Host struct {
	Hostname    string
	HostID      string
	CPUPercent  float64
	MemPercent  float64
	DiskPercent float64
	NetPercent  float64
}

//ToJSON convert host to json format
func (c *Host) ToJSON() string {
	tplString := `
	{
		"name":"{{.Hostname}}",
		"hostId":"{{.HostID}}",
		"systemId":"{{.HostID}}",
		"cpus":{
			"host":{"name":"{{.Hostname}}","hostId":"{{.HostID}}","systemId":"{{.HostID}}"},
			"utilization" : {{.CPUPercent}}
		},
		"memory":{
			"host":{"name":"{{.Hostname}}","hostId":"{{.HostID}}","systemId":"{{.HostID}}"},
			"utilization" : {{.MemPercent}}
		},
		"storage":{
			"host":{"name":"{{.Hostname}}","hostId":"{{.HostID}}","systemId":"{{.HostID}}"},
			"diskUtilization": {{.DiskPercent}}
		},
		"network":{
			"host":{"name":"{{.Hostname}}","hostId":"{{.HostID}}","systemId":"{{.HostID}}"},
			"utilization": {{.NetPercent}}
		}
	}
 `
	tmpl, _ := template.New("data").Parse(tplString)
	buf := bytes.NewBufferString("")
	tmpl.Execute(buf, c)
	return buf.String()
}

//Type get type of Host
func (c *Host) Type() string {
	return "Host"
}
