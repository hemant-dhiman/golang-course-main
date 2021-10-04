package main

import (
	"fmt"
)

func main() {
	/* colours := map[string]string{ //one way of declaring a map in go
		//key: value
		"red":   "#ff0000",
		"green": "#ff0001",
		"blue":  "#ff0002",
	} */

	//var colours map[string]string // second way of declaring a map in go

	//third way of declaring a map in go using function

	colours := make(map[string]string)

	colours["white"] = "#ffffff"
	colours["black"] = "#000000"

	fmt.Println(colours)

	delete(colours, "white")

	fmt.Printf("%v \n\n", colours)

	newMap := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0001",
		"blue":  "#ff0002",
		"white": "#ffffff",
	}

	fmt.Println("iterating a map...")
	printMap(newMap)
}

func printMap(c map[string]string) {
	for colour, hexValue := range c {
		fmt.Print("\nColour Name: '" + colour + "' Hex Value: '" + hexValue + "'")
	}
}
