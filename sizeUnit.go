package util

import (
	"math"
	"strconv"
)

const gigaByte = 1024 * 1024 * 1024 * 1024
const teraByte = 1024 * 1024 * 1024
const megaByte = 1024 * 1024
const kiloByte = 1024

func SizeToByteUnit(value int64) string {

	if value == -1 {
		return "0 B"
	}

	if value > gigaByte {
		return strconv.FormatInt(int64(math.Round(float64(value/gigaByte))), 10) + " GB"
	}

	if value > teraByte {
		return strconv.FormatInt(int64(math.Round(float64(value/teraByte))), 10) + " TB"
	}

	if value > megaByte {
		return strconv.FormatInt(int64(math.Round(float64(value/megaByte))), 10) + " MB"
	}

	if value > kiloByte {
		return strconv.FormatInt(int64(math.Round(float64(value/kiloByte))), 10) + " KB"
	}

	return strconv.FormatInt(value, 10) + " B"
}
