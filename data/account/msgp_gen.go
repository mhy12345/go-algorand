package account

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"sort"

	"github.com/algorand/msgp/msgp"
)

// The following msgp objects are implemented in this file:
// ParticipationKeyIdentity
//             |-----> (*) MarshalMsg
//             |-----> (*) CanMarshalMsg
//             |-----> (*) UnmarshalMsg
//             |-----> (*) CanUnmarshalMsg
//             |-----> (*) Msgsize
//             |-----> (*) MsgIsZero
//
// StateProofKey
//       |-----> MarshalMsg
//       |-----> CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> Msgsize
//       |-----> MsgIsZero
//
// StateProofKeys
//        |-----> MarshalMsg
//        |-----> CanMarshalMsg
//        |-----> (*) UnmarshalMsg
//        |-----> (*) CanUnmarshalMsg
//        |-----> Msgsize
//        |-----> MsgIsZero
//

// MarshalMsg implements msgp.Marshaler
func (z *ParticipationKeyIdentity) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(6)
	var zb0001Mask uint8 /* 7 bits */
	if (*z).Parent.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).FirstValid.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).KeyDilution == 0 {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).LastValid.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	if (*z).VoteID.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x20
	}
	if (*z).VRFSK.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x40
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "addr"
			o = append(o, 0xa4, 0x61, 0x64, 0x64, 0x72)
			o = (*z).Parent.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "fv"
			o = append(o, 0xa2, 0x66, 0x76)
			o = (*z).FirstValid.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "kd"
			o = append(o, 0xa2, 0x6b, 0x64)
			o = msgp.AppendUint64(o, (*z).KeyDilution)
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "lv"
			o = append(o, 0xa2, 0x6c, 0x76)
			o = (*z).LastValid.MarshalMsg(o)
		}
		if (zb0001Mask & 0x20) == 0 { // if not empty
			// string "vote-id"
			o = append(o, 0xa7, 0x76, 0x6f, 0x74, 0x65, 0x2d, 0x69, 0x64)
			o = (*z).VoteID.MarshalMsg(o)
		}
		if (zb0001Mask & 0x40) == 0 { // if not empty
			// string "vrfsk"
			o = append(o, 0xa5, 0x76, 0x72, 0x66, 0x73, 0x6b)
			o = (*z).VRFSK.MarshalMsg(o)
		}
	}
	return
}

func (_ *ParticipationKeyIdentity) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*ParticipationKeyIdentity)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ParticipationKeyIdentity) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0002 bool
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Parent.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Parent")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).VRFSK.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VRFSK")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).VoteID.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VoteID")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).FirstValid.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "FirstValid")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).LastValid.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "LastValid")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).KeyDilution, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "KeyDilution")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = ParticipationKeyIdentity{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "addr":
				bts, err = (*z).Parent.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Parent")
					return
				}
			case "vrfsk":
				bts, err = (*z).VRFSK.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VRFSK")
					return
				}
			case "vote-id":
				bts, err = (*z).VoteID.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VoteID")
					return
				}
			case "fv":
				bts, err = (*z).FirstValid.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "FirstValid")
					return
				}
			case "lv":
				bts, err = (*z).LastValid.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "LastValid")
					return
				}
			case "kd":
				(*z).KeyDilution, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "KeyDilution")
					return
				}
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
		}
	}
	o = bts
	return
}

func (_ *ParticipationKeyIdentity) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*ParticipationKeyIdentity)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ParticipationKeyIdentity) Msgsize() (s int) {
	s = 1 + 5 + (*z).Parent.Msgsize() + 6 + (*z).VRFSK.Msgsize() + 8 + (*z).VoteID.Msgsize() + 3 + (*z).FirstValid.Msgsize() + 3 + (*z).LastValid.Msgsize() + 3 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *ParticipationKeyIdentity) MsgIsZero() bool {
	return ((*z).Parent.MsgIsZero()) && ((*z).VRFSK.MsgIsZero()) && ((*z).VoteID.MsgIsZero()) && ((*z).FirstValid.MsgIsZero()) && ((*z).LastValid.MsgIsZero()) && ((*z).KeyDilution == 0)
}

// MarshalMsg implements msgp.Marshaler
func (z StateProofKey) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

func (_ StateProofKey) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(StateProofKey)
	if !ok {
		_, ok = (z).(*StateProofKey)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateProofKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 []byte
		zb0001, bts, err = msgp.ReadBytesBytes(bts, []byte((*z)))
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = StateProofKey(zb0001)
	}
	o = bts
	return
}

func (_ *StateProofKey) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*StateProofKey)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StateProofKey) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z StateProofKey) MsgIsZero() bool {
	return len(z) == 0
}

// MarshalMsg implements msgp.Marshaler
func (z StateProofKeys) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	if z == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendMapHeader(o, uint32(len(z)))
	}
	za0001_keys := make([]uint64, 0, len(z))
	for za0001 := range z {
		za0001_keys = append(za0001_keys, za0001)
	}
	sort.Sort(SortUint64(za0001_keys))
	for _, za0001 := range za0001_keys {
		za0002 := z[za0001]
		_ = za0002
		o = msgp.AppendUint64(o, za0001)
		o = msgp.AppendBytes(o, []byte(za0002))
	}
	return
}

func (_ StateProofKeys) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(StateProofKeys)
	if !ok {
		_, ok = (z).(*StateProofKeys)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateProofKeys) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 int
	var zb0004 bool
	zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0003 > 1000 {
		err = msgp.ErrOverflow(uint64(zb0003), uint64(1000))
		err = msgp.WrapError(err)
		return
	}
	if zb0004 {
		(*z) = nil
	} else if (*z) == nil {
		(*z) = make(StateProofKeys, zb0003)
	}
	for zb0003 > 0 {
		var zb0001 uint64
		var zb0002 StateProofKey
		zb0003--
		zb0001, bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		{
			var zb0005 []byte
			zb0005, bts, err = msgp.ReadBytesBytes(bts, []byte(zb0002))
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			zb0002 = StateProofKey(zb0005)
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

func (_ *StateProofKeys) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*StateProofKeys)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StateProofKeys) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for za0001, za0002 := range z {
			_ = za0001
			_ = za0002
			s += 0 + msgp.Uint64Size + msgp.BytesPrefixSize + len([]byte(za0002))
		}
	}
	return
}

// MsgIsZero returns whether this is a zero value
func (z StateProofKeys) MsgIsZero() bool {
	return len(z) == 0
}
