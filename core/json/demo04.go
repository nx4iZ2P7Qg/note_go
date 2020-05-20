package main

import (
	"encoding/json"
	"log"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func demo0401() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var m FamilyMember
	// unmarshal 分配了一段空间给 m
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatal("fatal")
	}
	log.Println(m)
}
