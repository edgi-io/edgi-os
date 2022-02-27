package sysctl

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/edgi-io/edgi-os/pkg/config"
)

func ConfigureSysctl(cfg *config.CloudConfig) error {
	for k, v := range cfg.EDGI.Sysctls {
		elements := []string{"/proc", "sys"}
		elements = append(elements, strings.Split(k, ".")...)
		path := path.Join(elements...)
		if err := ioutil.WriteFile(path, []byte(v), 0644); err != nil {
			return err
		}
	}
	return nil
}
