package attribute

import (
	"errors"
	"jvm/class/class_file_commons"
)

/**
union stack_map_frame {
    same_frame; // 0-63
    same_locals_1_stack_item_frame; // 64-127
    same_locals_1_stack_item_frame_extended; // 247
    chop_frame; // 248-250
    same_frame_extended; // 251
    append_frame; // 252-254
    full_frame; // 255
}
*/
type StackMapFrame_Info interface {
	ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info
}

func newStackMapFrame(frameType uint8) StackMapFrame_Info {
	if frameType >= 0 && frameType <= 63 { // same_frame
		return &SameFrame{FrameType: frameType}
	} else if frameType >= 64 && frameType <= 127 { // same_locals_1_stack_item_frame
		return &SameLocals1StackItemFrame{FrameType: frameType}
	} else if frameType == 247 { // same_locals_1_stack_item_frame_extended
		return &SameLocals1StackItemFrameExtended{FrameType: frameType}
	} else if frameType >= 248 && frameType <= 250 { // chop_frame
		return &ChopFrame{FrameType: frameType}
	} else if frameType == 251 { // same_frame_extended
		return &SameFrameExtended{FrameType: frameType}
	} else if frameType >= 252 && frameType <= 254 { // append_frame
		return &AppendFrame{FrameType: frameType}
	} else if frameType == 255 { // full_frame
		return &FullFrame{FrameType: frameType}
	}
	panic(errors.New("runtime error: not fount stack_map_frame type " + string(frameType)))
	return nil
}

/**
same_frame {
    u1 frame_type = SAME; // 0-63
}
*/
type SameFrame struct {
	FrameType uint8
}

func (s *SameFrame) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	return s
}

/**
same_locals_1_stack_item_frame {
    u1 frame_type = SAME_LOCALS_1_STACK_ITEM; // 64-127
	verification_type_info stack[1];
}
*/
type SameLocals1StackItemFrame struct {
	FrameType uint8
	Stack     [1]VerificationTypeInfo
}

func (s *SameLocals1StackItemFrame) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	tag := reader.ReadUint8()
	s.Stack[0] = newVerificationTypeInfo(tag).ReadVerficationTypeInfo(reader)
	return s
}

/**
same_locals_1_stack_item_frame_extended {
    u1 frame_type = SAME_LOCALS_1_STACK_ITEM_EXTENDED; // 247
	u2 offset_delta;
	verification_type_info stack[1];
}
*/
type SameLocals1StackItemFrameExtended struct {
	FrameType   uint8
	OffsetDelta uint16
	Stack       [1]VerificationTypeInfo
}

func (s *SameLocals1StackItemFrameExtended) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	s.OffsetDelta = reader.ReadUint16()
	tag := reader.ReadUint8()
	s.Stack[0] = newVerificationTypeInfo(tag).ReadVerficationTypeInfo(reader)
	return s
}

/**
chop_frame {
    u1 frame_type = CHOP; // 248-250
	u2 offset_delta;
}
*/
type ChopFrame struct {
	FrameType   uint8
	OffsetDelta uint16
}

func (c *ChopFrame) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	c.OffsetDelta = reader.ReadUint16()
	return c
}

/**
same_frame_extended {
    u1 frame_type = SAME_FRAME_EXTENDED; // 251
	u2 offset_delta;
}
*/
type SameFrameExtended struct {
	FrameType   uint8
	OffsetDelta uint16
}

func (s *SameFrameExtended) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	s.OffsetDelta = reader.ReadUint16()
	return s
}

/**
append_frame {
    u1 frame_type = APPEND; // 252-254
	u2 offset_delta;
	verification_type_info locals[frame_type - 251];
}
*/
type AppendFrame struct {
	FrameType   uint8
	OffsetDelta uint16
	Locals      []VerificationTypeInfo
}

func (a *AppendFrame) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	a.OffsetDelta = reader.ReadUint16()
	size := a.FrameType - 251
	for i := 0; i < int(size); i++ {
		tag := reader.ReadUint8()
		a.Locals = append(a.Locals, newVerificationTypeInfo(tag).ReadVerficationTypeInfo(reader))
	}
	return a
}

/**
full_frame {
    u1 frame_type = FULL_FRAME; // 255
	u2 offset_delta;
	u2 number_of_locals;
	verification_type_info locals[number_of_locals];
	u2 number_of_stack_items;
	verification_type_info stack[number_of_stack_items];
}
*/
type FullFrame struct {
	FrameType          uint8
	OffsetDelta        uint16
	NumberOfLocals     uint16
	Locals             []VerificationTypeInfo // size of NumberOfLocals
	NumberOfStackItems uint16
	Stack              []VerificationTypeInfo // size of NumberOfStackItems
}

func (f *FullFrame) ReadStackMapFrame(reader class_file_commons.Reader) StackMapFrame_Info {
	f.OffsetDelta = reader.ReadUint16()

	f.NumberOfLocals = reader.ReadUint16()
	size := f.NumberOfLocals
	for i := 0; i < int(size); i++ {
		tag := reader.ReadUint8()
		f.Locals = append(f.Locals, newVerificationTypeInfo(tag).ReadVerficationTypeInfo(reader))
	}

	f.NumberOfStackItems = reader.ReadUint16()
	size = f.NumberOfStackItems
	for i := 0; i < int(size); i++ {
		tag := reader.ReadUint8()
		f.Stack = append(f.Stack, newVerificationTypeInfo(tag).ReadVerficationTypeInfo(reader))
	}
	return f
}
