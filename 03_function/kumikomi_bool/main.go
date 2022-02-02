package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func main() {
	/*
		var a, b, c bool
		if a && b || !c {
			println("true")
		} else {
			println("false")
		}
	*/

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"a", "b", "c", "a && b || !c"})

	list := []bool{false, true}
	for _, a := range list {
		for _, b := range list {
			for _, c := range list {
				if a && b || !c {
					table.Append([]string{strconv.FormatBool(a), strconv.FormatBool(b), strconv.FormatBool(c), "true"})
				} else {
					table.Append([]string{strconv.FormatBool(a), strconv.FormatBool(b), strconv.FormatBool(c), "false"})
				}
			}
		}
	}
	table.Render()
}
