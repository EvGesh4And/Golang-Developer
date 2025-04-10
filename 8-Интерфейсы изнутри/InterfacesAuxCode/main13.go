package main

import "fmt"

type MetricCollector interface {
	Record()
}

type AudioRecorder interface {
	Record()
	Play()
}

type DummyRecorder struct{}

func (DummyRecorder) Record() {}

func main() {
	var v1 MetricCollector = DummyRecorder{}
	var v2 AudioRecorder = v1
	_ = v2

	fmt.Println("OK")
}
