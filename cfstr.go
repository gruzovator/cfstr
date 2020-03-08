package cfstr

import (
	"encoding/json"
	"strings"
)

type UpperCase struct {
	value string
}

func NewUpperCase(s string) UpperCase {
	return UpperCase{value: strings.ToUpper(s)}
}

func (s UpperCase) String() string {
	return s.value
}

func (s *UpperCase) UnmarshalJSON(bb []byte) error {
	if err := json.Unmarshal(bb, &s.value); err != nil {
		return err
	}
	s.value = strings.ToUpper(s.value)
	return nil
}

func (s UpperCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

type LowerCase struct {
	value string
}

func (s LowerCase) String() string {
	return s.value
}

func NewLowerCase(s string) LowerCase {
	return LowerCase{value: strings.ToLower(s)}
}

func (s *LowerCase) UnmarshalJSON(bb []byte) error {
	if err := json.Unmarshal(bb, &s.value); err != nil {
		return err
	}
	s.value = strings.ToLower(s.value)
	return nil
}

func (s LowerCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}
