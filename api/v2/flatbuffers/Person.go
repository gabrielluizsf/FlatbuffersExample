// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Person struct {
	_tab flatbuffers.Table
}

func GetRootAsPerson(buf []byte, offset flatbuffers.UOffsetT) *Person {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Person{}
	x.Init(buf, n+offset)
	return x
}

func FinishPersonBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsPerson(buf []byte, offset flatbuffers.UOffsetT) *Person {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Person{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedPersonBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Person) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Person) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Person) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Person) Age() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Person) MutateAge(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func PersonStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PersonAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func PersonAddAge(builder *flatbuffers.Builder, age int32) {
	builder.PrependInt32Slot(1, age, 0)
}
func PersonEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}