package main

import (
	"fmt"

	"github.com/renan061/llf/pkg/llf"
)

func advance(emitter *llf.Emitter, advance *llf.Advance) {
	fmt.Println("Advance!")
}

func inspect(emitter *llf.Emitter, inspect *llf.Inspect) {
	fmt.Println("Inspect!")
}

func main() {
	fmt.Println("---------- Hello ----------")
	rollup, err := llf.NewRollup(advance, inspect)
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

// const server = "http://127.0.0.1:5004"
//
// func finish(status string) (*Input, error) {
// 	finish, err := json.Marshal(map[string]string{"status": status})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	body := bytes.NewBuffer(finish)
// 	resp, err := http.Post(server+"/finish", "application/json", body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		panic("resp.Status != http.StatusOk")
// 	}
//
// 	bytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	input := Input{}
// 	err = json.Unmarshal(bytes, &input)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if input.RequestType != "advance_state" {
// 		panic("not advance state")
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
