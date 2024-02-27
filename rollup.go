// (c) Cartesi and individual authors (see AUTHORS)
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package gollup

type Emitter struct {
	binding *Binding
}

func (emitter *Emitter) EmitVoucher(address [20]byte, value []byte, data []byte) error {
	return emitter.binding.EmitVoucher(address, value, data)
}

func (emitter *Emitter) EmitNotice(data []byte) error {
	return emitter.binding.EmitNotice(data)
}

func (emitter *Emitter) EmitReport(data []byte) error {
	return emitter.binding.EmitReport(data)
}

// ------------------------------------------------------------------------------------------------

type Rollup struct {
	Emitter

	handleAdvance func(*Emitter, *Input)
	handleInspect func(*Emitter, *Query)
}

func NewRollup(
	handleAdvance func(*Emitter, *Input),
	handleInspect func(*Emitter, *Query)) (*Rollup, error) {

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

func (rollup *Rollup) Destroy() {
	rollup.binding.Destroy()
}

func (rollup *Rollup) Run() {
	accept := true
	for {
		finish, err := rollup.binding.Finish(accept)
		if err != nil {
			panic(err)
		}

		switch finish.NextRequestType {
		case AdvanceStateRequest:
			input, err := rollup.binding.ReadAdvanceState()
			if err != nil {
				panic(err)
			}
			rollup.handleAdvance(&rollup.Emitter, input)
		case InspectStateRequest:
			query, err := rollup.binding.ReadInspectState()
			if err != nil {
				panic(err)
			}
			rollup.handleInspect(&rollup.Emitter, query)
		default:
			panic("unreachable")
		}
	}
}
