package main

import "fmt

// App - the struct with contains things like pointers to
// database connections
type App struct {

}



func main() {
	fmt.Println("Go REST API Course")
	app :- App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up our REST API")
		fmt.Println(err)
	}