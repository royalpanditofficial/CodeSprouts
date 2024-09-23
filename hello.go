package main

import "fmt"

func main() {
fmt.Print("What is your name? ")

var name string

fmt.Scanln(&name)

fmt.Printf("Hello, %v! I see you are learning Go lang. That's great!\n", name)
}
