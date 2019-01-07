package main

import "fmt"

type BitbucketProcessor struct {
	Processor
}

func (BitbucketProcessor) process(interface{}) (NotifyData, error) {
	fmt.Print("yo")
	return NotifyData{}, nil
}
