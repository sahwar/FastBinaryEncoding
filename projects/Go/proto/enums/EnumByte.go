// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "encoding/json"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// EnumByte enum type
type EnumByte byte

// EnumByte enum values
//noinspection GoSnakeCaseUsage
const (
    EnumByte_ENUM_VALUE_0 = EnumByte(0 + 0)
    EnumByte_ENUM_VALUE_1 = EnumByte(0 + 0)
    EnumByte_ENUM_VALUE_2 = EnumByte(0 + 1)
    EnumByte_ENUM_VALUE_3 = EnumByte(254 + 0)
    EnumByte_ENUM_VALUE_4 = EnumByte(254 + 1)
    EnumByte_ENUM_VALUE_5 = EnumByte(EnumByte_ENUM_VALUE_3)
)

// Convert enum to string
func (e EnumByte) String() string {
    if e == EnumByte_ENUM_VALUE_0 {
        return "ENUM_VALUE_0"
    }
    if e == EnumByte_ENUM_VALUE_1 {
        return "ENUM_VALUE_1"
    }
    if e == EnumByte_ENUM_VALUE_2 {
        return "ENUM_VALUE_2"
    }
    if e == EnumByte_ENUM_VALUE_3 {
        return "ENUM_VALUE_3"
    }
    if e == EnumByte_ENUM_VALUE_4 {
        return "ENUM_VALUE_4"
    }
    if e == EnumByte_ENUM_VALUE_5 {
        return "ENUM_VALUE_5"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e EnumByte) MarshalJSON() ([]byte, error) {
    return json.Marshal(byte(e))
}

// Convert JSON to enum
func (e *EnumByte) UnmarshalJSON(b []byte) error {
    var value byte
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    *e = EnumByte(value)
    return nil
}
