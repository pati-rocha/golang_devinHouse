package main

import "fmt"

func Hello() {
	fmt.Println("Hello!!!")
}

func main() {

	Hello()

	func () {

		for i := 1; i <= 4; i++ {
			fmt.Println("Fluxo de dados batch number: ",i)
		}
	}()

}