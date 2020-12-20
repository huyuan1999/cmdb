package client

import (
	"encoding/json"
	"testing"
)

func TestCPU_GetNumber(t *testing.T) {
	cpu, err := NewCPU("test/cpuinfo")
	if err != nil {
		t.Fatal(err.Error())
	}
	cpu.GetNumber()
	t.Log(cpu.Number)
}

func TestCPU_GetCore(t *testing.T) {
	cpu, err := NewCPU("test/cpuinfo")
	if err != nil {
		t.Fatal(err.Error())
	}
	cpu.GetCore()
	t.Log(cpu.Core)
}

func TestCPU_GetSibling(t *testing.T) {
	cpu, err := NewCPU("test/cpuinfo")
	if err != nil {
		t.Fatal(err.Error())
	}
	cpu.GetSibling()
	t.Log(cpu.Sibling)
}

func TestCPU_GetProcessor(t *testing.T) {
	cpu, err := NewCPU("test/cpuinfo")
	if err != nil {
		t.Fatal(err.Error())
	}
	cpu.GetProcessor()
	t.Log(cpu.Processor)
}

func TestCPU_GetModelName(t *testing.T) {
	cpu, err := NewCPU("test/cpuinfo")
	if err != nil {
		t.Fatal(err.Error())
	}
	cpu.GetModelName()
	t.Log(cpu.ModelName)
}

func TestCPU(t *testing.T)  {
	cpu, err := NewCPU("test/cpuinfo")
	if err != nil {
		t.Fatal(err.Error())
	}
	cpu.GetModelName()
	cpu.GetProcessor()
	cpu.GetSibling()
	cpu.GetNumber()
	cpu.GetCore()
	data, err := json.Marshal(cpu)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(data))
}
