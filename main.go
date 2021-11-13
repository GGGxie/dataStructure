package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Error: failed to parse /data/catalog/values.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal number into Go value of type map[string]interface {}"
	fmt.Println(ShowMsgFromError(s))
}

// check the helm lint error to front message.
func ShowMsgFromError(s string) string {
	strs := strings.Split(s, "Error")
	results := make([]string, 0)
	for _, str := range strs {
		if len(str) == 0 || strings.TrimSpace(str) == "" {
			continue
		} else if strings.Contains(str, "nil pointer evaluating") {
			str = fmt.Sprintf("[必备字段值为空] %v", str)
		} else if strings.Contains(str, "The system cannot find the path") {
			str = fmt.Sprintf("[找不到路径] %v", str)
		} else if strings.Contains(str, "wrong type for value") {
			str = fmt.Sprintf("[值类型不对] %v", str)
		} else if strings.Contains(str, "did not find expected key") {
			str = fmt.Sprintf("[找不到必备字段] %v", str)
		} else if strings.Contains(str, "wrong type for value") {
			str = fmt.Sprintf("[值类型不对] %v", str)
		} else if strings.Contains(str, "error unmarshaling JSON") {
			str = fmt.Sprintf("[值类型不对] %v", str)
		} else {
			str = fmt.Sprintf("[未知错误] %v", str)
		}
		results = append(results, str)
	}
	return strings.Join(results, "\n")
}
