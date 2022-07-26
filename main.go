//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(net chan<- *Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			//menutup chanel
			close(net)
			break
		}

		net <- tweet
	}
}

func consumer(net <-chan *Tweet, finish chan<- bool) {
	for t := range net {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	finish <- true
}

func main() {

	start := time.Now()
	stream := GetMockStream()

	// Membuat channel dan menentukan panjang buffer chanel
	net := make(chan *Tweet, 8)
	finish := make(chan bool)

	// Producer
	go producer(net, stream)

	// Consumer
	go consumer(net, finish)

	<-finish

	fmt.Printf("Process took %s\n", time.Since(start))
}
