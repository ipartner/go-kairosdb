// Copyright 2016
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package kairosdb

import (
    "encoding/json"
)

// Make QueryMetric marshal into valid kairosdb json
func (q *QueryMetric) MarshalJSON() ([]byte, error) {
    tags := map[string][]string{}
    for k, v := range q.Tags {
        var z []string
        if v == nil {
            continue
        }
        for _, r := range v.Values {
            z = append(z, r)
        }
        tags[k] = z
    }

    s := struct {
        Name        string              `json:"name"`
        GroupBy     []*GroupBy          `json:"group_by,omitempty"`
        Tags        map[string][]string `json:"tags,omitempty"`
        Aggregators []*Aggregator       `json:"aggregators,omitempty"`
        Limit       int64               `json:"limit,omitempty"`
    }{
        Name:        q.Name,
        GroupBy:     q.GroupBy,
        Tags:        tags,
        Aggregators: q.Aggregators,
        Limit:       q.Limit,
    }

    return json.Marshal(s)
}

// Unmarshal kairosdb json QueryMetric into go type
func (q *QueryMetric) UnmarshalJSON(b []byte) error {
    type s struct {
        Name        string              `json:"name"`
        GroupBy     []*GroupBy          `json:"group_by"`
        Tags        map[string][]string `json:"tags"`
        Aggregators []*Aggregator       `json:"aggregators,omitempty"`
        Limit       int64               `json:"limit"`
    }
    ss := &s{}
    err := json.Unmarshal(b, &ss)
    if err != nil {
        return err
    }

    q.Name = ss.Name
    q.GroupBy = ss.GroupBy
    q.Aggregators = ss.Aggregators
    q.Limit = ss.Limit
    q.Tags = map[string]*StringList{}
    for k, v := range ss.Tags {
        sl := &StringList{Values: v}
        q.Tags[k] = sl
    }

    return nil
}

func (d *Datapoint) MarshalJSON() ([]byte,error) {
    return json.Marshal([]interface{}{d.Timestamp, d.Value})
}

//Unmarshal kairosdb json datapoint into go type
func (d *Datapoint) UnmarshalJSON(b []byte) error {
    var l []interface{}
    err := json.Unmarshal(b, &l)
    if err != nil {
        return err
    }

    if len(l) == 2 {
        if t, ok := l[0].(float64); ok {
            d.Timestamp = int64(t)
        }

        switch value := l[1].(type) {
        case int:
            d.Value = float64(value)
        case int8:
            d.Value = float64(value)
        case int16:
            d.Value = float64(value)
        case int32:
            d.Value = float64(value)
        case int64:
            d.Value = float64(value)
        case uint:
            d.Value = float64(value)
        case uint8:
            d.Value = float64(value)
        case uint16:
            d.Value = float64(value)
        case uint32:
            d.Value = float64(value)
        case uint64:
            d.Value = float64(value)
        case float32:
            d.Value = float64(value)
        case float64:
            d.Value = value
        }
    }
    return nil
}

func (h *HealthStatusResponse) UnmarshalJSON(b []byte) error {
    s := []string{}
    err := json.Unmarshal(b, &s)
    if err != nil {
        return err
    }
    h.Results = s
    return nil
}
