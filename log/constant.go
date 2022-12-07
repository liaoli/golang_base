package log

/**
*@author: 廖理
*@date:2022/12/7
**/

type Level int8

const (
	info Level = iota
	error
)

func GetLevelDes(l Level) string {
	switch l {
	case info:
		return "info"
	case error:
		return "error"
	}
	return ""
}
