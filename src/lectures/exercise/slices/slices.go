//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func printContents(contents []Part) {
	fmt.Println("---Parts on the line:---")
	for i := 0; i < len(contents); i++ {
		parts := contents[i]
		fmt.Println(parts)

	}

}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"seat", "wheel", "battery"}
	printContents(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "radiator", "disk brake")
	printContents(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	printContents(assemblyLine)
}
