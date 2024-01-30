package media

import (
	"fmt"
	"math"
	"mime"
	"reflect"
	"strconv"
	"strings"
)

type Type struct {
	prefix  string
	subType string
	suffix  string
	params  map[string]string
}

var Wildcard = Type{
	prefix:  "*",
	subType: "*",
}

func NewType(v string) *Type {
	v, params, err := mime.ParseMediaType(v)
	if err != nil {
		return nil
	}

	parts := strings.SplitN(v, "/", 2)
	if len(parts) == 1 {
		return nil
	}

	if parts[0] == "*" && !strings.HasPrefix(parts[1], "*") {
		return nil
	}

	mt := Type{
		prefix:  parts[0],
		subType: parts[1],
		params:  params,
	}

	suffixParts := strings.SplitN(mt.subType, "+", 2)
	if len(suffixParts) > 1 {
		mt.suffix = suffixParts[1]
	}

	return &mt
}
func (m *Type) GetType() string    { return m.prefix }
func (m *Type) GetSubType() string { return m.subType }
func (m *Type) GetSuffix() string  { return m.suffix }

func (m *Type) IsWildcardType() bool { return m.prefix == "*" }
func (m *Type) IsWildcardSubType() bool {
	return m.subType == "*" || strings.HasPrefix(m.subType, "*+")
}
func (m *Type) GetParameters() map[string]string {
	q := m.GetQualityValue()
	if q == 1 {
		delete(m.params, "q")
	}
	if m.params == nil {
		return make(map[string]string)
	}
	return m.params
}
func (m *Type) GetQualityValue() float64 {
	q, ok := m.params["q"]
	if !ok {
		return 1
	}
	parsed, err := strconv.ParseFloat(q, 32)
	if err != nil {
		return 1
	}
	if parsed > 1 || parsed < 0 {
		return 1
	}
	parsed = threeDecimalPlaces(parsed)

	return parsed
}
func (m *Type) IsCompatibleWith(other *Type) bool {
	if m == nil || other == nil {
		return false
	}
	if m.IsWildcardType() || other.IsWildcardType() {
		return true
	}
	if m.GetType() == other.GetType() {
		if m.GetSubType() == other.GetSubType() {
			return true
		}
		if m.IsWildcardSubType() || other.IsWildcardSubType() {
			mSuffix := m.GetSuffix()
			otherSuffix := other.GetSuffix()

			if m.GetSubType() == Wildcard.GetSubType() || other.GetSubType() == Wildcard.GetSubType() {
				return true
			} else if m.IsWildcardSubType() && mSuffix != "" {
				return mSuffix == other.GetSubType() || mSuffix == otherSuffix
			} else if other.IsWildcardSubType() && otherSuffix != "" {
				return m.GetSubType() == otherSuffix || otherSuffix == mSuffix
			}
		}
	}

	return false
}
func (m *Type) Equals(other *Type) bool {
	if m == nil || other == nil {
		return false
	}
	return m.GetType() == other.GetType() && m.GetSubType() == other.GetSubType() && reflect.DeepEqual(m.GetParameters(), other.GetParameters())
}
func (m *Type) String() string {
	return mime.FormatMediaType(fmt.Sprintf("%s/%s", m.prefix, m.subType), m.params)
}
func threeDecimalPlaces(val float64) float64 {
	ratio := math.Pow(10, float64(3))
	return math.Floor(val*ratio) / ratio
}
