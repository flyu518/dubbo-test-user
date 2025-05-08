package util

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析时间，支持：1h, 1d, 1m, 1s等
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

// SplitHostPort 解析地址和端口，简单的取最后一个冒号
func SplitHostPort(raw string) (host string, port uint64) {
	idx := strings.LastIndex(raw, ":")
	if idx == -1 {
		return raw, 0
	}

	port, err := strconv.ParseUint(raw[idx+1:], 10, 64)
	if err != nil {
		return raw, 0
	}

	return raw[:idx], port
}
