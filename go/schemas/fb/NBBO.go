// automatically generated by the FlatBuffers compiler, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NBBO struct {
	_tab flatbuffers.Table
}

func GetRootAsNBBO(buf []byte, offset flatbuffers.UOffsetT) *NBBO {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NBBO{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NBBO) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NBBO) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NBBO) SeqNum() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateSeqNum(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *NBBO) Symbol(obj *Symbol) *Symbol {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Symbol)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NBBO) BestBidExg() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateBestBidExg(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *NBBO) BestAskExg() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateBestAskExg(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *NBBO) Ts(obj *Timestamp) *Timestamp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Timestamp)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NBBO) BestBid(obj *Price) *Price {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Price)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NBBO) BestAsk(obj *Price) *Price {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Price)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NBBO) BestBidSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateBestBidSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

func (rcv *NBBO) BestAskSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateBestAskSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

func (rcv *NBBO) BestBidCondition() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateBestBidCondition(n byte) bool {
	return rcv._tab.MutateByteSlot(22, n)
}

func (rcv *NBBO) BestAskCondition() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NBBO) MutateBestAskCondition(n byte) bool {
	return rcv._tab.MutateByteSlot(24, n)
}

func NBBOStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}
func NBBOAddSeqNum(builder *flatbuffers.Builder, seqNum uint64) {
	builder.PrependUint64Slot(0, seqNum, 0)
}
func NBBOAddSymbol(builder *flatbuffers.Builder, symbol flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(symbol), 0)
}
func NBBOAddBestBidExg(builder *flatbuffers.Builder, bestBidExg byte) {
	builder.PrependByteSlot(2, bestBidExg, 0)
}
func NBBOAddBestAskExg(builder *flatbuffers.Builder, bestAskExg byte) {
	builder.PrependByteSlot(3, bestAskExg, 0)
}
func NBBOAddTs(builder *flatbuffers.Builder, ts flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(ts), 0)
}
func NBBOAddBestBid(builder *flatbuffers.Builder, bestBid flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(bestBid), 0)
}
func NBBOAddBestAsk(builder *flatbuffers.Builder, bestAsk flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(bestAsk), 0)
}
func NBBOAddBestBidSize(builder *flatbuffers.Builder, bestBidSize uint32) {
	builder.PrependUint32Slot(7, bestBidSize, 0)
}
func NBBOAddBestAskSize(builder *flatbuffers.Builder, bestAskSize uint32) {
	builder.PrependUint32Slot(8, bestAskSize, 0)
}
func NBBOAddBestBidCondition(builder *flatbuffers.Builder, bestBidCondition byte) {
	builder.PrependByteSlot(9, bestBidCondition, 0)
}
func NBBOAddBestAskCondition(builder *flatbuffers.Builder, bestAskCondition byte) {
	builder.PrependByteSlot(10, bestAskCondition, 0)
}
func NBBOEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}