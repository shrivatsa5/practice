package main

type Request struct {
	memory int
	cpu    int
	id     int
	nodeId int
}

func NewRequest(memory, cpu, id int) *Request {
	request := Request{
		memory: memory,
		cpu:    cpu,
		id:     id,
	}
	return &request
}
