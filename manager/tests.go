package manager

import (
	"fmt"
	tmp "github.com/Cguilliman/terminal-file-browser/manager/temporary"
	// tmp "github.com/Cguilliman/terminal-file-browser/manager/zipping"
)

const PATH = "/home/guilliman/go/src/github.com/Cguilliman/terminal-file-browser/"

func assertEqual(val, val1 interface{}) {
	if val != val1 {
		fmt.Printf("%v and %v not Equal\n", val, val1)
		// panic(fmt.Sprintf("%v and %v not Equal", val, val1))
	} else {
		fmt.Printf("*")
	}
}

func assertManyEqual(val, val1 []interface{}) {
	assertEqual(len(val), len(val1))
	for n, _ := range val {
		assertEqual(val[n], val1[n])
	}
}

func stringArrayToInterface(array ...string) []interface{} {
	result := make([]interface{}, len(array))
	for n, item := range array {
		result[n] = item
	}
	return result
}

func testConvertFunc() {
	command, args := convert("command some value", "path")
	assertEqual(command, "command")
	assertManyEqual(
		stringArrayToInterface(args...),
		[]interface{}{"some", "value", "path"},
	)

	command, args = convert("command some value ", "path")
	assertEqual(command, "command")
	assertManyEqual(
		stringArrayToInterface(args...),
		[]interface{}{"some", "value", "path"},
	)

	command, args = convert("command !! value", "path")
	assertEqual(command, "command")
	assertManyEqual(
		stringArrayToInterface(args...),
		[]interface{}{"path", "value"},
	)
}

// func testGetNested() {
// 	fmt.Println(GetNested(PATH+"inputs/"))
// }

// func testUnzipping() {
// 	zipping.Unzip(PATH+"new.zip", "testme")
// }

func testCopping() {
	tmp.DirCopping(
		PATH+"inputs/",
		PATH+"scripts/inputs/",
	)
	tmp.FileCopping(
		PATH+"inputs/main.go",
		PATH+"scripts/inputs/main.go",
	)
}

func RunTests() {
	// testConvertFunc()
	// testGetNested()
	// testUnzipping()
	// testCopping()
}
