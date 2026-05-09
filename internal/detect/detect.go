// Package detect verifies if the tools and applications needed for tb-override to work are available and installed or not
package detect

import (
	"os/exec"
)

type Platform struct {
	Proxy Proxy
}

type Proxy struct {
	Type      string
	Supported bool
}

func PlatformInfo() (Platform, error) {
	_, err := exec.LookPath("nginx")
	if err != nil {
		return Platform{}, err
	} else {
		proxy := Proxy{
			Type:      "nginx",
			Supported: true,
		}

		platform := Platform{
			Proxy: proxy,
		}

		return platform, nil
	}
}
