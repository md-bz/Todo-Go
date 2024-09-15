package main

type todo struct {
	Description string `json:"description"`
	Done        bool
}

var database = map[string][]todo{
	"user1": make([]todo, 0, 10),
	"user2": make([]todo, 0, 10),
}
