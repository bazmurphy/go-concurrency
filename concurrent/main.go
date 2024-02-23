package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// define a struct for the single update that gets sent to the channel
type SingleUpdate struct {
	Key   string
	Value string
}

func mapLetters(char string, updatesChannel chan SingleUpdate, wg *sync.WaitGroup, waitGroupCounter *int) {
	// execute wg.Done at the end of the function
	defer wg.Done()

	// time tracking
	currentTime := time.Now()
	microseconds := currentTime.UnixNano() / int64(time.Microsecond)
	fmt.Println(microseconds, "| newly spawned goroutine sending a SingleUpdate to the updatesChannel")

	// Increment the counter and print the current count
	*waitGroupCounter++
	fmt.Println("WaitGroup size:", *waitGroupCounter)

	result := SingleUpdate{char, string(char)}
	fmt.Println(microseconds, "|", result, "| --> SingleUpdate created and sending to the updatesChannel")

	// send the package to the updatesChannel channel
	updatesChannel <- result
}

func main() {
	// total duration tracking
	start := time.Now()
	startMicroseconds := start.UnixNano() / int64(time.Microsecond)

	someString := "BcCBabCaBBcAbCAACAcBbaCAABbBaBCAcaBCBcCaBAbcABCcBBcaACBCbCaCabCBCbAcBCBaCabcCABcBCacBCACcABaBCabCaBBcAABaBCBCACbCaABcCBacABaBCBaCcBCABAcBcABaCBbCABCabCcABABcbBCABcACBCbAcBACBCabcCaBCbAcaBCABacBCBCAcbBCABcaBABCABcABcbCBAcACABcABaBCbcACACBaBCBCaBaCABcCBACbCBABCAbCABCabcBCACBaCABcbABACabcACBCBCAABcbBCbCABAcBCABABcabABCabBCbcABABcABCbCABcaBCabBCABCAbCABCbACABACBcabBcCABABCABCaBCabcBcACBcabCABCABcabCBAcABaCABCaBCBcaABCBcABACcbACBcaABACcbCaBCABaBCabcBACBCABcaBCabcBaCABcaBCBcAABCABCBaBCACBcabCABcABCBACabCABCBcABaCABcABCBabcABcABCBCABcABcaBCABCABcBCAcbCABCAbCBABcaBcABCabcACBABCABcabCABCABcabCBAcaBCABcABCabcABCAbCABCABCabCABCABcABCabcCABcaBcabCABCaBCABcABcaCBAcABcABCabCABCABcaBCABcABcabCABcABCabcACBABCABCABcabCABCabcABcABcABcABCABCABcABCabCABcabCABcABcabCABCabcABCABcaBCabcABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABcaBCABcABcABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABcaBCABcABcABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABCABCABcabCABCabcABcABcABcABCABCABcaBCabcABcabCABCabcABcABcABcABC"

	// make a map
	charMap := make(map[string]string)

	// make a channel
	updatesChannel := make(chan SingleUpdate)

	// create a wait group
	var wg sync.WaitGroup

	// keep track of how many goroutines get spawned
	var waitGroupCounter int

	// loop over the string
	for _, char := range someString {
		// add +1 to the waitgroup counter because we will spawn a new go routine
		wg.Add(1)

		// make a new go routine and run the function mapLetters
		go mapLetters(string(char), updatesChannel, &wg, &waitGroupCounter)
	}

	// close the "updatesChannel" once all the goroutines are finished
	go func() {
		// wait for all goroutines to finish
		wg.Wait()

		// close the "updatesChannel" channel
		close(updatesChannel)
	}()

	// (!!!)
	// the receive from channel operator (<-) is implicitly used in the for loop
	// when you "range" over a channel like this, it automatically handles receiving values from the channel until it's closed.
	// the loop will continue until the channel is closed, at which point the loop will exit.
	// you don't need to explicitly use the receive operator (<-) because the "range" construct handles that internally

	// RANGE SYNTAX

	// // read from the updatesChannel as the SingleUpdates are received (live)
	// for update := range updatesChannel {
	// 	// time tracking
	// 	currentTime := time.Now()
	// 	microseconds := currentTime.UnixNano() / int64(time.Microsecond)
	// 	fmt.Println(microseconds, "|", update, "| <-- SingleUpdate received in the updatesChannel")

	// 	// update the map ("a" and "A" will be stored in the same key "a")
	// 	charMap[strings.ToLower(update.Key)] += update.Value
	// }

	// to explicitly use the receive from channel operator (<-) instead of the range keyword
	// you can rewrite the loop like this:

	// INFINITE FOR (WHILE) LOOP SYNTAX

	// read from the updatesChannel as the SingleUpdates are received (live)
	for {
		// receive a SingleUpdate from the channel
		update, ok := <-updatesChannel

		// time tracking
		currentTime := time.Now()
		microseconds := currentTime.UnixNano() / int64(time.Microsecond)
		fmt.Println(microseconds, "|", update, "| <-- SingleUpdate received in the updatesChannel")

		// check if the channel has been closed
		if !ok {
			break
		}

		// update the map ("a" and "A" will be stored in the same key "a")
		charMap[strings.ToLower(update.Key)] += update.Value
	}

	// print the map
	fmt.Println("final charMap:", charMap)
	// check these two values
	fmt.Println("string length:", len(someString), "| waitGroupCounter:", waitGroupCounter)

	// total duration tracking
	finish := time.Now()
	finishMicroseconds := finish.UnixNano() / int64(time.Microsecond)
	duration := finishMicroseconds - startMicroseconds
	fmt.Println("total duration |", duration)
}
