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

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-multierror"
	"github.com/ipartner/go-kairosdb/kairosdb"
	log "github.com/opsee/logrus"
)

const (
	KdbListMetricNames  = "api/v1/metricnames"
	KdbListTagNames     = "api/v1/tagnames"
	KdbListTagValues    = "api/v1/tagvalues"
	KdbQuery            = "api/v1/datapoints/query"
	KdbHealthCheck      = "api/v1/health/check"
	KdbHealthStatus     = "api/v1/health/status"
	KdbDeleteDatapoints = "api/v1/datapoints/delete"
	KdbAddDatapoints    = "api/v1/datapoints"
)

// This is the type that implements the Client interface.
type KDBClient struct {
	kdbAddress string
}

// Returns a new http client
func New(addr string) Client {
	return &KDBClient{
		kdbAddress: addr,
	}
}

// Get kairosdb health status data
// https://kairosdb.github.io/docs/build/html/restapi/Health.html
func (s *KDBClient) HealthStatus() (*kairosdb.HealthStatusResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", s.kdbAddress, KdbHealthStatus))
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("received list tag values response: %v", string(body))

	gr := &kairosdb.HealthStatusResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}
	return gr, nil
}

// Get kairosdb health
// server returns 204 if all checks pass
// https://kairosdb.github.io/docs/build/html/restapi/Health.html
func (s *KDBClient) HealthCheck() (*kairosdb.HealthCheckResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", s.kdbAddress, KdbHealthCheck))
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}

	healthy := false
	if resp.StatusCode == 204 {
		healthy = true
	}

	gr := &kairosdb.HealthCheckResponse{
		Healthy: healthy,
	}
	return gr, nil
}

// Queries the kdb server via the QueryMetricsRequest protobuf generate type
func (s *KDBClient) Query(in *kairosdb.QueryMetricsRequest) (*kairosdb.QueryMetricsResponse, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	hreq, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", s.kdbAddress, KdbQuery), bytes.NewBuffer(b))
	hreq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(hreq)
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("received query response: %v", string(body))

	gr := &kairosdb.QueryMetricsResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}

	// handle errors field
	if len(gr.Errors) > 0 {
		var errs error
		for _, e := range gr.Errors {
			errs = multierror.Append(errs, errors.New(e))
		}
		return nil, errs
	}

	return gr, nil
}

// List Metric Names
func (s *KDBClient) ListMetricNames() (*kairosdb.ListMetricNamesResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", s.kdbAddress, KdbListMetricNames))
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("received list metric names response: %v", string(body))

	gr := &kairosdb.ListMetricNamesResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}

	return gr, nil
}

// List Tag Names
func (s *KDBClient) ListTagNames() (*kairosdb.ListTagNamesResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", s.kdbAddress, KdbListTagNames))
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("received list tag names response: %v", string(body))

	gr := &kairosdb.ListTagNamesResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}
	return gr, nil
}

// List Tag Values
func (s *KDBClient) ListTagValues() (*kairosdb.ListTagValuesResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", s.kdbAddress, KdbListTagValues))
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("received list tag values response: %v", string(body))

	gr := &kairosdb.ListTagValuesResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}
	return gr, nil
}

// Queries the kdb server via the QueryMetricsRequest protobuf generate type
func (s *KDBClient) DeleteDatapoints(in *kairosdb.DeleteDatapointsRequest) (*kairosdb.DeleteDatapointsResponse, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	hreq, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", s.kdbAddress, KdbDeleteDatapoints), bytes.NewBuffer(b))
	hreq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(hreq)
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 204 {
		return &kairosdb.DeleteDatapointsResponse{}, nil
	}

	log.Debugf("received query response: %v", string(body))

	gr := &kairosdb.DeleteDatapointsResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}

	// handle errors field
	if len(gr.Errors) > 0 {
		var errs error
		for _, e := range gr.Errors {
			errs = multierror.Append(errs, errors.New(e))
		}
		return nil, errs
	}

	return gr, nil
}

// Queries the kdb server via the QueryMetricsRequest protobuf generate type
func (s *KDBClient) AddDatapoints(in *kairosdb.AddDatapointsRequest) (*kairosdb.AddDatapointsResponse, error) {
	b, err := json.Marshal(in.Metrics)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", s.kdbAddress, KdbAddDatapoints), bytes.NewBuffer(b))
	hreq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(hreq)
	if err != nil {
		log.WithError(err).Fatalf("couldn't query kairosdb server")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		return &kairosdb.AddDatapointsResponse{}, nil
	}

	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("received query response: %v", string(body))

	gr := &kairosdb.AddDatapointsResponse{}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return nil, err
	}

	// handle errors field
	if len(gr.Errors) > 0 {
		var errs error
		for _, e := range gr.Errors {
			errs = multierror.Append(errs, errors.New(e))
		}
		return nil, errs
	}

	return gr, nil
}
