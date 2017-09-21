package main

import (
	"strings"
	"fmt"
	"bytes"
)

//Host host object
type Host struct{
	buffer bytes.Buffer
}

func (h *Host) addString(name, value string){
	h.buffer.WriteString(fmt.Sprintf(",\"%s\":\"%s\"\n", name, value))
}

func (h *Host) addFloat(name string, value float64){
	h.buffer.WriteString(fmt.Sprintf(",\"%s\":%f\n", name, value))
}

func (h *Host) ToJson()string{
	json:=h.buffer.String()
	return strings.Trim(json, ",")
}

func (h *Host) Type()string{
	return "SimpleHost"
}