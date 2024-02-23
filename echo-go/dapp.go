package main

// #cgo CFLAGS: -I/home/renan/.local/outuni/libcmt
// #cgo LDFLAGS: -L/home/renan/.local/outuni/libcmt -lcmt
// #include "rollup.h"
import "C"

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Input struct {
	RequestType string `json:"request_type"`
	Data        struct {
		Payload  string `json:"payload"`
		Metadata struct {
			Sender      string `json:"msg_sender"`
			EpochIndex  uint64 `json:"epoch_index"`
			InputIndex  uint64 `json:"input_index"`
			BlockNumber uint64 `json:"block_number"`
			Timestamp   uint64 `json:"timestamp"`
		} `json:"metadata"`
	}

	Payload string
}

const server = "http://127.0.0.1:5004"

func finish(status string) (*Input, error) {
	finish, err := json.Marshal(map[string]string{"status": status})
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(finish)
	resp, err := http.Post(server+"/finish", "application/json", body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		panic("resp.Status != http.StatusOk")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	input := Input{}
	err = json.Unmarshal(bytes, &input)
	if err != nil {
		return nil, err
	}
	if input.RequestType != "advance_state" {
		panic("not advance state")
	}

	s, _ := hex.DecodeString(input.Data.Payload[2:])
	input.Payload = string(s)

	return &input, nil
}

func notice(payload string) (*Input, error) {
	hexy := hexutil.Encode([]byte(payload))
	notice, err := json.Marshal(map[string]string{"payload": hexy})
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(notice)
	resp, err := http.Post(server+"/notice", "application/json", body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		panic("resp.Status != http.StatusOk")
	}

	// TODO: get index
	return nil, nil
}

func main() {
	fmt.Println("---------- Hello ----------")

	// binding := binding.New()

	/*
		for {
			input, err := finish("accept")
			if err != nil {
				panic(err)
			}

			fmt.Println("---")
			fmt.Println(input)
			notice(input.Payload)
			notice(input.Payload)
			notice(input.Payload)
			notice(input.Payload)
			notice(input.Payload)
			fmt.Println("---")
		}
	*/
}

// voucher
//   emit a voucher read from stdin as a JSON object in the format
//     {"destination": <address>, "payload": <string>}
//   where field "destination" contains a 20-byte EVM address in hex.
//   if successful, prints to stdout a JSON object in the format
//     {"index": <number> }
//   where field "index" is the index allocated for the voucher
//   in the voucher hashes array.
//
// notice
//   emit a notice read from stdin as a JSON object in the format
//     {"payload": <string> }
//   if successful, prints to stdout a JSON object in the format
//     {"index": <number> }
//   where field "index" is the index allocated for the notice
//   in the voucher hashes array.
//
// report
//   emit a report read from stdin as a JSON object in the format
//     {"payload": <string> }
