package main

import (
	"./ibc"
	"./gha"
	"log"
)

func main() {
	g := gha.NewGHA("", "secret", "IBCBoiler")


	m := ibc.NewMonitor("192.168.2.27")
	d, err := m.GetExtendedDetails()
	if err != nil {
		panic(err)
	}
	log.Print("Details: ", d)
}
