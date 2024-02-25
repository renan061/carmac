// (c) Cartesi and individual authors (see AUTHORS)
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package llf

import "fmt"

type Emitter struct {
	binding *Binding
}

func (emitter *Emitter) Voucher(address [20]byte, value []byte, data []byte) error {
	return emitter.binding.EmitVoucher(address, value, data)
}

func (emitter *Emitter) Notice(data []byte) error {
	return emitter.binding.EmitNotice(data)
}

func (emitter *Emitter) Report(data []byte) error {
	return emitter.binding.EmitReport(data)
}

// ------------------------------------------------------------------------------------------------

type Rollup struct {
	Emitter

	handleAdvance func(*Emitter, *Advance)
	handleInspect func(*Emitter, *Inspect)
}

func NewRollup(
	handleAdvance func(*Emitter, *Advance),
	handleInspect func(*Emitter, *Inspect)) (*Rollup, error) {

	binding, err := NewBinding()
	if err != nil {
		return nil, err
	}

	return &Rollup{
		Emitter:       Emitter{binding},
		handleAdvance: handleAdvance,
		handleInspect: handleInspect,
	}, nil
}

func (rollup *Rollup) Run() {
	accept := true
	for {
		finish, err := rollup.binding.Finish(accept)
		if err != nil {
			fmt.Println("err")
			panic(err)
		}

		switch finish.NextRequestType {
		case AdvanceStateRequest:
			advance, err := rollup.binding.ReadAdvanceState()
			if err != nil {
				panic(err)
			}
			rollup.handleAdvance(&rollup.Emitter, advance)
		case InspectStateRequest:
			inspect, err := rollup.binding.ReadInspectState()
			if err != nil {
				panic(err)
			}
			rollup.handleInspect(&rollup.Emitter, inspect)
		default:
			panic("unreachable")
		}
	}
}
