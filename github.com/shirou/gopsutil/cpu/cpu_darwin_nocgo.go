// +build darwin
// +build !cgo

package cpu

import "github.com/general252/EasyDarwinLib/github.com/shirou/gopsutil/internal/common"

func perCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func allCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}
