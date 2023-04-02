package main

import "fmt"

func Run() error {
	fmt.Println("starting up...")
	return nil
}

func main() {
	fmt.Println("Go Rest API Comments")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
