// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "encoding/json"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// EnumInt8 enum type
type EnumInt8 int8

// EnumInt8 enum values
//noinspection GoSnakeCaseUsage
const (
    EnumInt8_ENUM_VALUE_0 = EnumInt8(0 + 0)
    EnumInt8_ENUM_VALUE_1 = EnumInt8(-128 + 0)
    EnumInt8_ENUM_VALUE_2 = EnumInt8(-128 + 1)
    EnumInt8_ENUM_VALUE_3 = EnumInt8(126 + 0)
    EnumInt8_ENUM_VALUE_4 = EnumInt8(126 + 1)
    EnumInt8_ENUM_VALUE_5 = EnumInt8(EnumInt8_ENUM_VALUE_3)
)

// Convert enum to string
func (e EnumInt8) String() string {
    if e == EnumInt8_ENUM_VALUE_0 {
        return "ENUM_VALUE_0"
    }
    if e == EnumInt8_ENUM_VALUE_1 {
        return "ENUM_VALUE_1"
    }
    if e == EnumInt8_ENUM_VALUE_2 {
        return "ENUM_VALUE_2"
    }
    if e == EnumInt8_ENUM_VALUE_3 {
        return "ENUM_VALUE_3"
    }
    if e == EnumInt8_ENUM_VALUE_4 {
        return "ENUM_VALUE_4"
    }
    if e == EnumInt8_ENUM_VALUE_5 {
        return "ENUM_VALUE_5"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e EnumInt8) MarshalJSON() ([]byte, error) {
    return json.Marshal(int8(e))
}

// Convert JSON to enum
func (e *EnumInt8) UnmarshalJSON(b []byte) error {
    var value int8
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    *e = EnumInt8(value)
    return nil
}
