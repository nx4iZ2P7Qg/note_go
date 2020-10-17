package common

import "log"

// CheckError 异常则抛出
func CheckError(i interface{}) {
	switch v := i.(type) {
	case error:
		if v != nil {
			log.Fatal("[checkError-fatal]", v)
		}
	case bool:
		if !v {
			log.Fatal("[checkError-fatal]", "false")
		}
	}

}
