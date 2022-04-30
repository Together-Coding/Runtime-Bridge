package utils

import (
	"os"
	"strconv"
)

// GetConfig get env var named `name` of type `type_`
func GetConfig(name string, type_ string) (vv interface{}) {
	v := os.Getenv(name)

	switch type_ {
	case "bool":
		vv, _ = strconv.ParseBool(v)
	case "int":
		vv, _ = strconv.ParseInt(v, 10, 64)
	case "uint":
		vv, _ = strconv.ParseUint(v, 10, 64)
	case "float":
		vv, _ = strconv.ParseFloat(v, 64)
	}
	return
}
func GetConfigBool(name string) bool {
	return GetConfig(name, "bool").(bool)
}

func GetConfigUint8(name string) uint8 {
	return uint8(GetConfig(name, "uint").(uint64))
}
func GetConfigUint16(name string) uint16 {
	return uint16(GetConfig(name, "uint").(uint64))
}
func GetConfigUint32(name string) uint32 {
	return uint32(GetConfig(name, "uint").(uint64))
}
func GetConfigUint(name string) uint64 {
	return GetConfig(name, "uint").(uint64)
}

func GetConfigInt8(name string) int8 {
	return int8(GetConfig(name, "int").(int64))
}
func GetConfigInt16(name string) int16 {
	return int16(GetConfig(name, "int").(int64))
}
func GetConfigInt32(name string) int32 {
	return int32(GetConfig(name, "int").(int64))
}
func GetConfigInt(name string) int64 {
	return GetConfig(name, "int").(int64)
}
