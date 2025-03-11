package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	var countSend = 10

	var countRec = 5

	var countMess = 100

	stopCh := make(chan struct{})

	dataCh := make(chan int)

	signCh := make(chan struct{}, 1)

	wg := sync.WaitGroup{}

	wg.Add(countRec + countSend)

	// senders
	for i := 0; i < countSend; i++ {
		go func() {
			defer func() {
				wg.Done()
				fmt.Println("send", i, "close")
			}()
			for {
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case dataCh <- i:
				case <-stopCh:
					return
				}
			}
		}()
	}

	// the receiver
	for i := 0; i < countRec; i++ {
		go func() {
			defer func() {
				wg.Done()
				fmt.Println("receiver", i, "close")
			}()
			var v int

			for {

				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case d := <-dataCh:
					v += d
				}

				log.Printf("Rec %d: %d", i, v)
				time.Sleep(time.Millisecond)
				if v > countMess {
					select {
					case signCh <- struct{}{}:
					default:
					}
					return
				}
			}
		}()
	}

	go func() {
		select {
		case <-signCh:
			close(stopCh)
			return
		}
	}()
	wg.Wait()
}
