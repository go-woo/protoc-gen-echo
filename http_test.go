package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func RegxTest(t *testing.T) {

}

func Test_buildPathVars(t *testing.T) {
	md := buildPathVars("/hello/{name}/asd/{zxc}")
	sm, _ := json.Marshal(md)
	fmt.Fprintf(os.Stderr, "3=== = %v\n", string(sm))

}

func Test_replacePath(t *testing.T) {
	s := replacePath("name", "name", "/hello/{name}/asd/{zxc}")
	fmt.Fprintf(os.Stderr, "3=== = %v\n", s)
}
