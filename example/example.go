package main

import (
	"fmt"

	"github.com/renan061/gollup"
)

var counter = 0

func handleAdvance(emitter *gollup.Emitter, input *gollup.Input) {
	fmt.Println("---------- Advance ----------")

	data := string(input.Data)
	fmt.Println("Data =", data)
	fmt.Println("Counter = ", counter)

	n := 1
	if data == "multiple" {
		n = 5
	}

	for i := 0; i < n; i++ {
		err := emitter.EmitNotice([]byte(fmt.Sprintf("%s (%d)", input.Data, counter)))
		assert(err != nil, err)
		counter += 1
	}
}

func handleInspect(emitter *gollup.Emitter, query *gollup.Query) {
	fmt.Println("---------- Inspect ----------")

	data := string(query.Data)
	fmt.Println("Data =", data)
	fmt.Println("Counter = ", counter)

	n := 1
	if data == "multiple" {
		n = 10
	}

	for i := 0; i < n; i++ {
		err := emitter.EmitReport([]byte(fmt.Sprintf("%d", counter)))
		assert(err != nil, err)
	}
}

func main() {
	fmt.Println("---------- Started the dapp ----------")

	rollup, err := gollup.NewRollup(handleAdvance, handleInspect)
	assert(err != nil, err)
	defer rollup.Destroy()

	rollup.Run()
}

func assert(b bool, err error) {
	if b {
		panic(err)
	}
}
