// Code generated by usestdlibvars, DO NOT EDIT.

package os_test

import "fmt"

var (
	_ = "/dev/null" // want `"/dev/null" can be replaced by os\.DevNull`
)

const (
	_ = "/dev/null" // want `"/dev/null" can be replaced by os\.DevNull`
)

func _() {
	_ = func(s string) string { return s }("/dev/null") // want `"/dev/null" can be replaced by os\.DevNull`
	_ = func(s string) string { return s }("text before key /dev/null")
	_ = func(s string) string { return s }("/dev/null text after key")
}

func _() {
	_ = fmt.Sprint("/dev/null") // want `"/dev/null" can be replaced by os\.DevNull`
	_ = fmt.Sprint("text before key /dev/null")
	_ = fmt.Sprint("/dev/null text after key")
}
