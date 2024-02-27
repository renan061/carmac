package main

import (
	"fmt"

	"github.com/renan061/gollup"
)

func handleAdvance(emitter *rollup.Emitter, input *rollup.Input) {
	fmt.Println("Advance!")
	fmt.Println("Received input", input)
	fmt.Println("Data =", string(input.Data))
	err := emitter.EmitNotice(input.Data)
	if err != nil {
		panic(fmt.Sprintf("EmitNotice error!", err))
	}
}

func handleInspect(emitter *rollup.Emitter, query *rollup.Query) {
	fmt.Println("Inspect!")
}

func main() {
	fmt.Println("---------- Started the dapp ----------")
	rollup, err := rollup.NewRollup(handleAdvance, handleInspect)
	if err != nil {
		panic(err)
	}
	defer rollup.Destroy()
	rollup.Run()
}
