package myscript

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"regexp"
	"strings"
)

// 给 struct 各变量添加小写字母开头的 json flag
func AppendJsonFlag() {

	s := `
type TimeWindow struct {
	StartTime string
	StopTime  string
}

type PlmnId struct {
	Mcc string
	Mnc string
}
`
	for _, line := range strings.Split(s, "\n") {
		r, _ := regexp.Compile("\t(.*)( +)(.*)")
		if r.MatchString(line) {
			// 取出 struct 内变量
			rep := "${1}"
			field := fmt.Sprintf("%v", r.ReplaceAllString(line, rep))
			field = strcase.ToLowerCamel(field)
			fmt.Printf("%v %v\n", line, fmt.Sprintf("`json:\"%v\"`", field))
		} else {
			fmt.Printf("%v\n", line)
		}
	}
}
