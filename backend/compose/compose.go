package compose

import (
	"github.com/sjkyspa/stacks/controller/api/api"
	"gopkg.in/yaml.v2"
	"fmt"
	"github.com/sjkyspa/stacks/client/backend"
	"strings"
	"path/filepath"
	"os"
)

type ComposeBackend struct {

}
type Service struct {
	Image      string `json:"image" yaml:"image"`
	Entrypoint string `json:"entrypoint" yaml:"entrypoint,omitempty"`
	Command    string `json:"command" yaml:"command,omitempty"`
	Volumes    []string `json:"volumes" yaml:"volumes,omitempty"`
	Links      []string `json:"links" yaml:"links,omitempty"`
	Ports      []string `json:"ports" yaml:"ports,omitempty"`
}

type ComposeFile struct {
	Version  string `json:"version"`
	Services map[string]Service `json:"services"`
}

func (cb ComposeBackend) Up() {

}

func toString(v api.Volume) string {
	if "" == v.HostPath {
		return fmt.Sprintf("%s", v.ContainerPath)
	}

	if "" == v.Mode {
		return fmt.Sprintf("%s:%s", v.HostPath, v.ContainerPath)
	}

	var hostpath string
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Please ensure the current dir can be accessed")
	}

	if !filepath.IsAbs(v.HostPath) {
		hostpath = filepath.Join("/Mac", dir, ".local", v.HostPath)
	}

	return fmt.Sprintf("%s:%s:%s", hostpath, v.ContainerPath, strings.ToLower(v.Mode))
}

func (cb ComposeBackend) ToComposeFile(s api.Stack) string {
	services := s.GetServices()
	composeServices := make(map[string]Service, 0)

	for name, service := range services {
		if !service.IsBuildable() {
			volumes := make([]string, 0)
			for _, v:= range service.GetVolumes() {
				volumes = append(volumes, toString(v))
			}
			composeServices[name] = Service{
				Image: service.GetImage(),
				Links: service.GetLinks(),
				Volumes: volumes,
			}
		} else {
			volumes := make([]string, 0)
			for _, v:= range service.GetVolumes() {
				volumes = append(volumes, toString(v))
			}

			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				panic("Please ensure the current dir can be accessed")
			}

			volumes = append(volumes, "/var/run/docker.sock:/var/run/docker.sock")
			volumes = append(volumes, fmt.Sprintf("%s:/codebase", filepath.Join("/Mac", dir)))

			composeServices["runtime"] = Service{
				Image: service.GetBuild().Name,
				Entrypoint: "/bin/sh",
				Command: "-c 'tail -f /dev/null'",
				Volumes: volumes,
			}
		}
	}
	composeFile := ComposeFile{
		Version: "2",
		Services: composeServices,
	}

	out, err := yaml.Marshal(composeFile)
	if err != nil {
		panic(fmt.Sprintf("Error happend when translate to yaml %v", err))
	}

	return string(out)
}

func NewComposeBackend() backend.Runtime {
	return ComposeBackend{}
}