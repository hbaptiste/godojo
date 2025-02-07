package main

import "fmt"

func chanSend(channel chan string, msg string) {
	channel <- msg
}

// to do, make send and receive
/*func chanRecieve(channel chan interface{}, dest interface{}) {
	dest := <-dest
}*/

func main() {
	messages := make(chan string)
	result := make(chan bool)

	// send
	go func() {
		chanSend(messages, "partenHarris")
		chanSend(messages, "Plus difficile")
		chanSend(messages, "Strange last")
		result <- true
		fmt.Println("Inside Strange")
	}()

	// recieve
	go func() {
		msg := <-messages
		fmt.Println(msg)
	}()
	test := <-result
	fmt.Println(test)
	//fmt.Println(result)
}
