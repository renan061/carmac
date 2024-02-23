// (c) Cartesi and individual authors (see AUTHORS)
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package llf

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

func (rollup *Rollup) Run(accept bool) {
	for {
		finish, err := rollup.binding.Finish(accept)
		if err != nil {
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