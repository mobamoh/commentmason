package main

import "fmt"

// Run is responsible for the instantiation
// and startup of our application
func Run() error {
	fmt.Println("starting up the app...")
	return nil
}

func main() {
	fmt.Println("### CommentMason API ###")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
