# go-kairosdb, a protobuf based kairosdb client for golang

go-kairosdb provides protobuf message types that marshal and unmarshal to valid kairosdb json schema.


# What and what is not currently supported?

Going down the list in the Rest API section in the docs:

1. Add Data Points
    - Unsupported
2. Delete Data Points
    - Unsupported
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
