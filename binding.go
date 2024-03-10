// (c) Cartesi and individual authors (see AUTHORS)
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package gollup

// #cgo CFLAGS: -I/home/renan/Local/install/include
// #cgo LDFLAGS: -L/home/renan/Local/install/lib -lcmt
// #include <stdlib.h>
// #include <string.h>
// #include "libcmt/rollup.h"
// #include "libcmt/io.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Finish struct {
	AcceptPreviousRequest    bool
	NextRequestType          RequestType
	NextRequestPayloadLength uint32
}

type Input struct {
	Sender         [20]byte
	BlockNumber    uint64
	BlockTimestamp uint64
	Index          uint64
	Data           []byte
}

type Query struct {
	Data []byte
}

// ------------------------------------------------------------------------------------------------

type RequestType = uint8

const AdvanceStateRequest RequestType = C.CMT_IO_REASON_ADVANCE
const InspectStateRequest RequestType = C.CMT_IO_REASON_INSPECT

// ------------------------------------------------------------------------------------------------

var (
	CErrRollupInit       = errors.New("cmt_rollup_init error")
	CErrRollupFinish     = errors.New("cmt_rollup_finish error")
	CErrReadAdvanceState = errors.New("cmt_rollup_read_advance_state error")
	CErrReadInspectState = errors.New("cmt_rollup_read_inspect_state error")
	CErrEmitVoucher      = errors.New("cmt_rollup_emit_voucher error")
	CErrEmitNotice       = errors.New("cmt_rollup_emit_notice error")
	CErrEmitReport       = errors.New("cmt_rollup_emit_report error")
)

type Binding struct {
	rollup *C.cmt_rollup_t
}

func NewBinding() (*Binding, error) {
	var rollup C.cmt_rollup_t
	result := C.cmt_rollup_init(&rollup)
	if err := toError(result, CErrRollupInit); err != nil {
		return nil, err
	} else {
		return &Binding{rollup: &rollup}, nil
	}
}

func (binding *Binding) Destroy() {
	C.cmt_rollup_fini(binding.rollup)
}

func (binding *Binding) Finish(accept bool) (*Finish, error) {
	inner := C.cmt_rollup_finish_t{
		accept_previous_request: C.bool(accept),
	}

	result := C.cmt_rollup_finish(binding.rollup, &inner)
	if err := toError(result, CErrRollupFinish); err != nil {
		return nil, err
	}

	fmt.Println("========= accept_previous_request = ", inner.accept_previous_request)
	fmt.Println("========= next_request_type = ", inner.next_request_type)
	fmt.Println("========= next_request_payload_length = ", inner.next_request_payload_length)

	finish := &Finish{
		AcceptPreviousRequest:    bool(inner.accept_previous_request),
		NextRequestType:          RequestType(inner.next_request_payload_length), // TODO
		NextRequestPayloadLength: uint32(inner.next_request_payload_length),
	}
	return finish, nil
}

func (binding *Binding) ReadAdvanceState() (*Input, error) {
	var inner C.cmt_rollup_advance_t
	result := C.cmt_rollup_read_advance_state(binding.rollup, &inner)
	if err := toError(result, CErrReadAdvanceState); err != nil {
		return nil, err
	}
	// TODO: should I free inner.data?

	var sender [20]byte
	for i, v := range inner.sender {
		sender[i] = byte(v)
	}

	return &Input{
		Data:           C.GoBytes(inner.data, C.int(inner.length)),
		Sender:         sender,
		BlockNumber:    uint64(inner.block_number),
		BlockTimestamp: uint64(inner.block_timestamp),
		Index:          uint64(inner.index),
	}, nil
}

func (binding *Binding) ReadInspectState() (*Query, error) {
	var inner C.cmt_rollup_inspect_t
	result := C.cmt_rollup_read_inspect_state(binding.rollup, &inner)
	if err := toError(result, CErrReadInspectState); err != nil {
		return nil, err
	}
	// TODO: should I free inner.data?

	return &Query{
		Data: C.GoBytes(inner.data, C.int(inner.length)),
	}, nil
}

func (binding *Binding) EmitVoucher(address [20]byte, value []byte, voucher []byte) error {
	addressLength := C.uint(20)
	addressData := C.CBytes(address[:])
	defer C.free(addressData)

	valueLength := C.uint(len(value))
	valueData := C.CBytes(value)
	defer C.free(valueData)

	voucherLength := C.uint(len(voucher))
	voucherData := C.CBytes(voucher)
	defer C.free(voucherData)

	result := C.cmt_rollup_emit_voucher(binding.rollup,
		addressLength, addressData,
		valueLength, valueData,
		voucherLength, voucherData,
	)
	return toError(result, CErrEmitNotice)
}

func (binding *Binding) EmitNotice(notice []byte) error {
	length := C.uint(len(notice))

	data := C.CBytes(notice)
	defer C.free(data)
	fmt.Println("data", data)
	fmt.Println("length", length)

	data = unsafe.Pointer(&notice[0]) // TODO

	result := C.cmt_rollup_emit_notice(binding.rollup, length, data)
	return toError(result, CErrEmitNotice)
}

func (binding *Binding) EmitReport(report []byte) error {
	length := C.uint(len(report))
	data := C.CBytes(report)
	defer C.free(data)

	data = unsafe.Pointer(&report[0]) // TODO

	result := C.cmt_rollup_emit_report(binding.rollup, length, data)
	return toError(result, CErrEmitReport)
}

// ------------------------------------------------------------------------------------------------

func toError(errno C.int, err error) error {
	if errno != 0 {
		cs := C.strerror(errno)
		defer C.free(unsafe.Pointer(cs))
		s := C.GoString(cs)
		return errors.New(s)
		// return err
	} else {
		return nil
	}
}
