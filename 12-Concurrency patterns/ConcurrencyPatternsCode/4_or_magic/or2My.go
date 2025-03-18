package main

func or2(first bool, channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0] // <-or(ch) == <-ch
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			channelsNext := channels[3:]
			if first {
				first = false
				channelsNext = append(channelsNext, orDone)
			}
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or2(first, channelsNext...):
			}
		}
	}()
	return orDone
}
