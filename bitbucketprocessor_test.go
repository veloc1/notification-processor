package main

import "testing"

func TestEmptyData(t *testing.T) {
	processor := BitbucketProcessor{}

	processor.process(0)
}
