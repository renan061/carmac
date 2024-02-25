package main

import (
	"fmt"

	"github.com/renan061/carmac/golang/rollup"
)

func advance(emitter *rollup.Emitter, advance *rollup.Advance) {
	fmt.Println("Advance!")
}

func inspect(emitter *rollup.Emitter, inspect *rollup.Inspect) {
	fmt.Println("Inspect!")
}

func main() {
	fmt.Println("---------- Started the dapp ----------")
	rollup, err := rollup.NewRollup(advance, inspect)
	if err != nil {
		panic(err)
	}
	rollup.Run()
}

// type Input struct {
// 	RequestType string `json:"request_type"`
// 	Data        struct {
// 		Payload  string `json:"payload"`
// 		Metadata struct {
// 			Sender      string `json:"msg_sender"`
// 			EpochIndex  uint64 `json:"epoch_index"`
// 			InputIndex  uint64 `json:"input_index"`
// 			BlockNumber uint64 `json:"block_number"`
// 			Timestamp   uint64 `json:"timestamp"`
// 		} `json:"metadata"`
// 	}
//
// 	Payload string
// }

// func finish(status string) (*Input, error) {
// 	input := Input{}
// 	err = json.Unmarshal(bytes, &input)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	s, _ := hex.DecodeString(input.Data.Payload[2:])
// 	input.Payload = string(s)
//
// 	return &input, nil
// }
//
// func notice(payload string) (*Input, error) {
// 	hexy := hexutil.Encode([]byte(payload))
// 	notice, err := json.Marshal(map[string]string{"payload": hexy})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	body := bytes.NewBuffer(notice)
// 	resp, err := http.Post(server+"/notice", "application/json", body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusCreated {
// 		panic("resp.Status != http.StatusOk")
// 	}
//
// 	// TODO: get index
// 	return nil, nil
// }
