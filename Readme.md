# go-kairosdb, a protobuf based kairosdb client for golang

go-kairosdb provides protobuf message types that marshal and unmarshal to valid kairosdb json schema.

See https://github.com/dan-compton/go-kairosdb/blob/master/kairosdb/kairosdb_types.proto for the message types
and https://github.com/dan-compton/go-kairosdb/blob/master/client/http_test.go if you dont know what the generated
go types will look like.


# What and what is not currently supported?

Going down the list in the Rest API section in the docs:

1. Add Data Points
    - Supported
        * Format containing datapoints array like following example:
        ```
            // do this
            {
                "name": "archive_file_tracked",
                "datapoints": [[1359788400000, 1], [1359788300000, 2], [1359788410000, 3]],
                "tags": {
                    "host": "server1",
                    "data_center": "DC1"
                 },
                 "ttl": 300
            },
       ```

        ```
            // like this:
            &kairosdb.AddDatapointsRequest{
                Metrics: []*kairosdb.Metric{
                    &kairosdb.Metric{
                        Name: "some",
                        Datapoints: []*kairosdb.Datapoint{
                            &kairosdb.Datapoint{
                                Timestamp: 1359786400000,
                                Value: 1,
                            },
                            &kairosdb.Datapoint{
                                Timestamp: 1471218571000,
                                Value: 2,
                            },
                            &kairosdb.Datapoint{
                                Timestamp: 1471304971000,
                                Value: 3,
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
            panic(err)  // or w/e
        }
        ```
2. Delete Data Points
    - Supported
3. Delete Metric
    - Unsupported
4. Health Checks
    - Supported
        * Health Check Status Data
        * Health Check
5. List Metric Names
    - Supported
6. List Tag Names
    - Supported
7. Query Metric
    - Supported
        * Grouping
        * Aggregators
        * Filtering by tag
    - Unsupported
        * Relative Start and End Time
8. Aggregators
    - Supported
        * Average
        * Standard Deviation
        * Count
        * First
        * Gaps
        * Histogram
        * Last
        * Least Squares
        * Max
        * Min
    - Unsupported
        * Divide
        * Rate
        * Sampler
        * Scale
        * Trim
        * Save As
9. Query Metric Tags
    - Unsupported
