package main

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type Info interface {
	Host() (string, error)
	Memory() (uint64, uint64, float64, error)
}

func NewInfo() Info {
	return &Host{}
}

type Host struct{}

func (h *Host) Host() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}

	return info.OS, nil
}

func (h *Host) Memory() (uint64, uint64, float64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}

	return v.Total, v.Free, v.UsedPercent, nil
}
