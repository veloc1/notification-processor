package main

type Processor interface {
	process(interface{}) (NotifyData, error)
}
