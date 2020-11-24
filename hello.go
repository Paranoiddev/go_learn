package main

import (
	"fmt"
	"strings"
)

// this is a fuckin comment.
func main() {
	broken := "G### rocks! s#meb#dy #nce t#ld me I aint the sharpest t##l in the shed"
	valueassignedtovar := strings.NewReplacer("#", "o")
	fixed := valueassignedtovar.Replace(broken)
	fmt.Println(fixed)
}