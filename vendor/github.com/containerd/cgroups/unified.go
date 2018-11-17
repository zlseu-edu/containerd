package cgroups

import (
	"path/filepath"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

var (
	supportedSubsystems = supportedSubsystemSet{
		"blkio",
		"memory",
		"cpu",
	}
)

type supportedSubsystemSet []string

func NewUnified(root string) *unifiedController {
	return &unifiedController{
		root: filepath.Join(root, string(Unified)),
	}
}

type unifiedController struct {
	root string
}

func (u *unifiedController) Name() Name {
	return Unified
}

func (u *unifiedController) Path(path string) string {
	return filepath.Join(u.root, path)
}

func (u *unifiedController) Create(path string, resources *specs.LinuxResources) error {
	for _, sys := range supportedSubsystems {
		switch sys {
		case "blkio":
			createBlkio(path, resources)
		case "memory":
			createMemory(path, resources)
		case "cpu":
			createCpu(path, resources)
		}
	}
	return nil
}

func createCpu(path string, resources *specs.LinuxResources) error {
	return nil
}

func createBlkio(path string, resources *specs.LinuxResources) error {
	return nil
}

func createMemory(path string, resources *specs.LinuxResources) error {
	return nil
}

func (u *unifiedController) Update(path string, resources *specs.LinuxResources) error {
	return u.Create(path, resources)
}

func (u *unifiedController) Stat(path string, stats *Metrics) error {
	return nil
}
