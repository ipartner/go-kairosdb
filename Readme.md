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
              "datapoints": [[1359788400000, 123], [1359788300000, 13.2], [1359788410000, 23.1]],
              "tags": {
                  "host": "server1",
                "data_center": "DC1"
              },
              "ttl": 300
          },
        ```
    - Unsupported
        * Format wherein there is no datapoints array like:
        ```
            // dont do this
            {
                  "name": "archive_file_search",
                  "timestamp": 1359786400000,
                  "value": 321,
                  "tags": {
                      "host": "server2"
                  }
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
