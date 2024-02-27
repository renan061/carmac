package main

import (
	"fmt"

	"github.com/renan061/gollup"
)

func handleAdvance(emitter *gollup.Emitter, input *gollup.Input) {
	fmt.Println("Advance!")
	fmt.Println("Received input", input)
	fmt.Println("Data =", string(input.Data))
	err := emitter.EmitNotice(input.Data)
	if err != nil {
		panic(fmt.Sprintf("EmitNotice error!", err))
	}
}

func handleInspect(emitter *gollup.Emitter, query *gollup.Query) {
	fmt.Println("Inspect!")
}

func main() {
	fmt.Println("---------- Started the dapp ----------")
	rollup, err := gollup.NewRollup(handleAdvance, handleInspect)
	if err != nil {
		panic(err)
	}
	defer rollup.Destroy()
	rollup.Run()
}
