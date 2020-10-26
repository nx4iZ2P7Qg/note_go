package myscript

import (
	"fmt"
	"strings"
)

// Compare 对比 hash 值
func Compare() {
	s1 := "0265ccd6dcc9726cb2f1aee1f89cee79f5e612bd350ed20f774cdf64fef36d83"
	s2 := "0265CCD6DCC9726CB2F1AEE1F89CEE79F5E612BD350ED20F774CDF64FEF36D83"
	fmt.Print(strings.ToUpper(s1) == strings.ToUpper(s2))
}
