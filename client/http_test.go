package client

import (
	"testing"

	"github.com/ipartner/go-kairosdb/kairosdb"
)

var testHttpClient = New("http://localhost:8080")

func TestAddDatapoints(t *testing.T) {
	req := &kairosdb.AddDatapointsRequest{
		Metrics: []*kairosdb.Metric{
			// 13, 14, 15 august 2013
			&kairosdb.Metric{
				Name: "test_data",
				Datapoints: []*kairosdb.Datapoint{
					&kairosdb.Datapoint{
						Timestamp: 1359786400000,
						Value:     1,
					},
					&kairosdb.Datapoint{
						Timestamp: 1471218571000,
						Value:     2,
					},
					&kairosdb.Datapoint{
						Timestamp: 1471304971000,
						Value:     3,
					},
				},
				Tags: map[string]string{
					"atag": "is here",
				},
			},
		},
	}

	_, err := testHttpClient.AddDatapoints(req)
	if err != nil {
		t.Fatalf("client returned err: %s", err.Error())
	}
}

func TestListMetricNames(t *testing.T) {
	resp, err := testHttpClient.ListMetricNames()
	if err != nil {
		t.Fatalf("client return err: %s", err.Error())
	}
	if len(resp.Results) == 0 {
		t.Fatal("kairosdb should have at least default metrics like kairosdb.jvm.free_memory, got none.")
	}
}

func TestListTagNames(t *testing.T) {
	_, err := testHttpClient.ListTagNames()
	if err != nil {
		t.Fatalf("client returned err: %s", err.Error())
	}
}

func TestListTagValues(t *testing.T) {
	_, err := testHttpClient.ListTagValues()
	if err != nil {
		t.Fatalf("client returned err: %s", err.Error())
	}
}

func TestHealthStatus(t *testing.T) {
	resp, err := testHttpClient.HealthStatus()
	if err != nil {
		t.Fatalf("client return err: %s", err.Error())
	}
	if len(resp.Results) == 0 {
		t.Fatal("kairosdb at least some status information, got none.")
	}
}

func TestHealthChecks(t *testing.T) {
	resp, err := testHttpClient.HealthCheck()
	if err != nil {
		t.Fatalf("client returned err: %s", err.Error())
	}
	if resp.Healthy != true {
		t.Fatalf("kairosdb not healthy")
	}
}

// Test Query for built-in kairosdb.jvm.free_memory
func TestQueryMetrics(t *testing.T) {
	req := &kairosdb.QueryMetricsRequest{
		Metrics: []*kairosdb.QueryMetric{
			&kairosdb.QueryMetric{
				Name: "go-kairosdb.test_data",
			},
		},
		CacheTime:     0,
		StartAbsolute: 1339786400000, // before 13 august 2013
		EndAbsolute:   1481304971000, // after 15 august 2013
	}

	resp, err := testHttpClient.Query(req)
	if err != nil {
		t.Fatalf("client returned err: %s", err.Error())
	}
	if len(resp.Errors) > 0 {
		t.Fatalf("kairosdb return err: %v", resp.Errors)
	}
}

func TestDeleteDatapoints(t *testing.T) {
	req := &kairosdb.DeleteDatapointsRequest{
		Metrics: []*kairosdb.QueryMetric{
			&kairosdb.QueryMetric{
				Name: "go-kairosdb.test_data",
			},
		},
		CacheTime:     0,
		StartAbsolute: 1339786400000, // before 13 august 2013
		EndAbsolute:   1481304971000, // after 15 august 2013
	}

	resp, err := testHttpClient.DeleteDatapoints(req)
	if err != nil {
		t.Fatalf("client returned err: %s", err.Error())
	}
	if len(resp.Errors) > 0 {
		t.Fatalf("kairosdb return err: %v", resp.Errors)
	}
}
