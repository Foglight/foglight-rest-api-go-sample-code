package rest

import (
	"bytes"
	"fmt"
	"strings"
)

//SimpleTopology for simple data submission
type SimpleTopology struct {
	buffer   bytes.Buffer
	TypeName string
}

//NewSimpleTopology create a new simple topology type
func NewSimpleTopology(typeName string) *SimpleTopology {
	return &SimpleTopology{TypeName: typeName}
}

//AddString add a string property
func (h *SimpleTopology) AddString(name, value string) {
	h.buffer.WriteString(fmt.Sprintf(",\"%s\":\"%s\"\n", name, value))
}

//AddFloat add a float property
func (h *SimpleTopology) AddFloat(name string, value float64) {
	h.buffer.WriteString(fmt.Sprintf(",\"%s\":%f\n", name, value))
}

//ToJSON convert data to json format, note don't need the outter {}
func (h *SimpleTopology) ToJSON() string {
	json := h.buffer.String()
	return strings.Trim(json, ",")
}

//Type get type of this topology object
func (h *SimpleTopology) Type() string {
	return h.TypeName
}
