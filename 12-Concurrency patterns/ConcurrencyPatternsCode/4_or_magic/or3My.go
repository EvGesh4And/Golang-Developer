package main

func or3(channels ...<-chan struct{}) <-chan struct{} {

	orDone := make(chan struct{})

	go func() {
		defer close(orDone)
		sig := make(chan struct{}, 1)
		for _, ch := range channels {
			go func() {
				select {
				case <-ch:
					sig <- struct{}{}
				case <-orDone:
				}
			}()
		}
		<-sig
	}()

	return orDone
}
