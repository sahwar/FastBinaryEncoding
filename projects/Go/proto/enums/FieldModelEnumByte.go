// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumByte field model class
type FieldModelEnumByte struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelEnumByte) FBESize() int { return 1 }
// Get the field extra size
func (fm FieldModelEnumByte) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelEnumByte) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumByte) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumByte) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumByte) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelEnumByte(buffer *fbe.Buffer, offset int) *FieldModelEnumByte {
    return &FieldModelEnumByte{ buffer: buffer, offset: offset }
}

// Check if the value is valid
func (fm FieldModelEnumByte) Verify() bool { return true }

// Get the value
func (fm FieldModelEnumByte) Get() (EnumByte, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelEnumByte) GetDefault(defaults EnumByte) (EnumByte, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumByte(0), nil
    }

    return EnumByte(fbe.ReadByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelEnumByte) Set(value EnumByte) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), byte(value))
    return nil
}
