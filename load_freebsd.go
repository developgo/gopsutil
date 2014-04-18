// +build freebsd

package main

import (
	"exec"
	"strconv"
	"strings"
)

func (l Load) LoadAvg() (LoadAvg, error) {
	out, err := exec.Command("/sbin/sysctl", "-n", "vm.loadavg").Output()
	if err != nil {
		return LoadAvg{}, err
	}
	v := strings.Replace(string(out), "{ ", "", 1)
	v = strings.Replace(string(v), " }", "", 1)
	values := strings.Fields(string(v))

	load1, err := strconv.ParseFloat(values[0], 32)
	if err != nil {
		return LoadAvg{}, err
	}
	load5, err := strconv.ParseFloat(values[1], 32)
	if err != nil {
		return LoadAvg{}, err
	}
	load15, err := strconv.ParseFloat(values[2], 32)
	if err != nil {
		return LoadAvg{}, err
	}

	ret := LoadAvg{
		Load1:  float32(load1),
		Load5:  float32(load5),
		Load15: float32(load15),
	}

	return ret, nil
}
