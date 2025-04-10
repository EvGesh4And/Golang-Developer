package main

type MetricCollector interface {
	Record()
}

type DummyRecorder struct{}

func (DummyRecorder) Record() {}

func main() {
	var v1 MetricCollector
	var v2 DummyRecorder = v1
	_ = v2
}
