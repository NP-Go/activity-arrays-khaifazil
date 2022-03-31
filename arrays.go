package main

import (
	"fmt"
)

func main() {
	//Insert Code here
	array := [7]string{
		"Operating System List",
		"Windows XP",
		"Linux 1.0",
		"Raspbian Teensy",
		"MacOS",
		"IOS",
		"Google Android"}

	for _, item := range array {
		fmt.Println(len(item))
	}

	array[1] = "Windows 10"
	array[2] = "Linux 16.04"
	array[3] = "Raspbian Buster"

	fmt.Println(array)
	
	var newArray [10]string

	for i, _ := range array {
		newArray[i] = array[i]
	}

	newArray[7] = "Ubuntu"
	newArray[8] = "MS-Dos"
	newArray[9] = "Solaris"

	fmt.Println(newArray)

	fmt.Println(newArray[0:3])
	fmt.Println(newArray[3:6])
	fmt.Println(newArray[6:])

}
