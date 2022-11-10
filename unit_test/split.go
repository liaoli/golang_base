package unit_test

import "strings"

/**
*@author: 廖理
*@date:2022/11/10
**/

// Split 把字符串s按照给定的分隔符sep进行分割返回字符串切片
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		if i > 0 {
			result = append(result, s[:i])
		}
		//s = s[i+1:]
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func Split2(s, sep string) (result []string) {

	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)

	for i > -1 {
		if i > 0 {
			result = append(result, s[:i])
		}
		//s = s[i+1:]
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
