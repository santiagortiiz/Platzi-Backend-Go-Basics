package main

import "fmt"

// text is the input of the channel
func speak(text string, c chan<- string) {
	c <- text
}

/* text in the output of the channel
func speak(text string, c <-chan string) {
	text = <- c
}
*/

func goroutines_with_channels() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go speak("Bye", c)

	fmt.Println(<-c)
}
