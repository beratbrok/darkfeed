// automatically generated by the FlatBuffers compiler, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OptionQuote struct {
	_tab flatbuffers.Table
}

func GetRootAsOptionQuote(buf []byte, offset flatbuffers.UOffsetT) *OptionQuote {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OptionQuote{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OptionQuote) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OptionQuote) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OptionQuote) SeqNum() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OptionQuote) MutateSeqNum(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *OptionQuote) Symbol(obj *Symbol) *Symbol {
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

func (rcv *OptionQuote) OptionType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OptionQuote) MutateOptionType(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *OptionQuote) ReportingExg() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OptionQuote) MutateReportingExg(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *OptionQuote) Ts(obj *Timestamp) *Timestamp {
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

func (rcv *OptionQuote) Expiration(obj *Timestamp) *Timestamp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
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

func (rcv *OptionQuote) Strike(obj *Price) *Price {
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

func (rcv *OptionQuote) Bid(obj *Price) *Price {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
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

func (rcv *OptionQuote) Ask(obj *Price) *Price {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
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

func (rcv *OptionQuote) BidSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OptionQuote) MutateBidSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(22, n)
}

func (rcv *OptionQuote) AskSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OptionQuote) MutateAskSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(24, n)
}

func (rcv *OptionQuote) Condition() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OptionQuote) MutateCondition(n byte) bool {
	return rcv._tab.MutateByteSlot(26, n)
}

func OptionQuoteStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func OptionQuoteAddSeqNum(builder *flatbuffers.Builder, seqNum uint64) {
	builder.PrependUint64Slot(0, seqNum, 0)
}
func OptionQuoteAddSymbol(builder *flatbuffers.Builder, symbol flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(symbol), 0)
}
func OptionQuoteAddOptionType(builder *flatbuffers.Builder, optionType byte) {
	builder.PrependByteSlot(2, optionType, 0)
}
func OptionQuoteAddReportingExg(builder *flatbuffers.Builder, reportingExg byte) {
	builder.PrependByteSlot(3, reportingExg, 0)
}
func OptionQuoteAddTs(builder *flatbuffers.Builder, ts flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(ts), 0)
}
func OptionQuoteAddExpiration(builder *flatbuffers.Builder, expiration flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(expiration), 0)
}
func OptionQuoteAddStrike(builder *flatbuffers.Builder, strike flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(strike), 0)
}
func OptionQuoteAddBid(builder *flatbuffers.Builder, bid flatbuffers.UOffsetT) {
	builder.PrependStructSlot(7, flatbuffers.UOffsetT(bid), 0)
}
func OptionQuoteAddAsk(builder *flatbuffers.Builder, ask flatbuffers.UOffsetT) {
	builder.PrependStructSlot(8, flatbuffers.UOffsetT(ask), 0)
}
func OptionQuoteAddBidSize(builder *flatbuffers.Builder, bidSize uint32) {
	builder.PrependUint32Slot(9, bidSize, 0)
}
func OptionQuoteAddAskSize(builder *flatbuffers.Builder, askSize uint32) {
	builder.PrependUint32Slot(10, askSize, 0)
}
func OptionQuoteAddCondition(builder *flatbuffers.Builder, condition byte) {
	builder.PrependByteSlot(11, condition, 0)
}
func OptionQuoteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
