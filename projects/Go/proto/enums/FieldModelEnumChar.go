// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumChar field model class
type FieldModelEnumChar struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelEnumChar) FBESize() int { return 1 }
// Get the field extra size
func (fm FieldModelEnumChar) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelEnumChar) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumChar) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumChar) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelEnumChar(buffer *fbe.Buffer, offset int) *FieldModelEnumChar {
    return &FieldModelEnumChar{ buffer: buffer, offset: offset }
}

// Check if the value is valid
func (fm FieldModelEnumChar) Verify() bool { return true }

// Get the value
func (fm FieldModelEnumChar) Get() (EnumChar, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelEnumChar) GetDefault(defaults EnumChar) (EnumChar, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumChar(0), nil
    }

    return EnumChar(fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelEnumChar) Set(value EnumChar) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint8(value))
    return nil
}
