package kairosdb

import (
    "errors"
    "encoding/json"
)


func (s *StringList) UnmarshalJSON(b []byte) error {
	x := []string{}
	err := json.Unmarshal(b, &x)
	if err == nil {
		s.Values = x
		return nil
	}

	r := make(map[string]interface{})
	err = json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	if v, ok := r["values"]; ok {
		if z, ok := v.([]string); ok {
			s.Values = z
			return nil
		}
	}
	return errors.New("stringlist must be of type []string or map key, value pair \"values\":[]string{...}")
}
