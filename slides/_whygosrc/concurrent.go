package main

import "fmt"

func main() {
	// THREAD_S OMIT
	var ca, cb = make(chan bool), make(chan bool)
	for {
		go func() {
			fmt.Println(<-ca)
			cb <- true
		}()
		go func() {
			ca <- false
			fmt.Println(<-cb)
		}()
	}
	// THREAD_E OMIT
}
