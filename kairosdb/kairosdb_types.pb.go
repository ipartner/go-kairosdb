// Code generated by protoc-gen-go.
// source: kairosdb_types.proto
// DO NOT EDIT!

package kairosdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a kairosdb datapoint
// https://github.com/kairosdb/kairosdb-client/blob/master/src/main/java/org/kairosdb/client/builder/DataPoint.java
type Datapoint struct {
	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *Datapoint) Reset()                    { *m = Datapoint{} }
func (m *Datapoint) String() string            { return proto.CompactTextString(m) }
func (*Datapoint) ProtoMessage()               {}
func (*Datapoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Represents a sampling in the QueryMetrics request body
type Sampling struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Unit  string `protobuf:"bytes,2,opt,name=unit" json:"unit,omitempty"`
}

func (m *Sampling) Reset()                    { *m = Sampling{} }
func (m *Sampling) String() string            { return proto.CompactTextString(m) }
func (*Sampling) ProtoMessage()               {}
func (*Sampling) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Represents a kairosdb aggregator
// https://kairosdb.github.io/docs/build/html/restapi/Aggregators.html
type Aggregator struct {
	Name           string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	AlignSampling  bool      `protobuf:"varint,2,opt,name=align_sampling,json=alignSampling" json:"align_sampling,omitempty"`
	Sampling       *Sampling `protobuf:"bytes,3,opt,name=sampling" json:"sampling,omitempty"`
	AlignStartTime bool      `protobuf:"varint,4,opt,name=align_start_time,json=alignStartTime" json:"align_start_time,omitempty"`
	Percentile     *string   `protobuf:"bytes,5,opt,name=percentile" json:"percentile,omitempty"`
}

func (m *Aggregator) Reset()                    { *m = Aggregator{} }
func (m *Aggregator) String() string            { return proto.CompactTextString(m) }
func (*Aggregator) ProtoMessage()               {}
func (*Aggregator) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Aggregator) GetSampling() *Sampling {
	if m != nil {
		return m.Sampling
	}
	return nil
}

// Represents the group_property in the request
// I've decided that it's not bad
type GroupBy struct {
	Name  string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags  []string          `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
	Group map[string]string `protobuf:"bytes,3,rep,name=group" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GroupBy) Reset()                    { *m = GroupBy{} }
func (m *GroupBy) String() string            { return proto.CompactTextString(m) }
func (*GroupBy) ProtoMessage()               {}
func (*GroupBy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GroupBy) GetGroup() map[string]string {
	if m != nil {
		return m.Group
	}
	return nil
}

// Reprents a kairosdb QueryMetric
// https://kairosdb.github.io/docs/build/html/restapi/QueryMetrics.html#id3
type QueryMetric struct {
	Name        string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags        map[string]*StringList `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GroupBy     []*GroupBy             `protobuf:"bytes,3,rep,name=group_by,json=groupBy" json:"group_by,omitempty"`
	Aggregators []*Aggregator          `protobuf:"bytes,4,rep,name=aggregators" json:"aggregators,omitempty"`
	Limit       int64                  `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	TimeZone    string                 `protobuf:"bytes,6,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
}

func (m *QueryMetric) Reset()                    { *m = QueryMetric{} }
func (m *QueryMetric) String() string            { return proto.CompactTextString(m) }
func (*QueryMetric) ProtoMessage()               {}
func (*QueryMetric) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *QueryMetric) GetTags() map[string]*StringList {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *QueryMetric) GetGroupBy() []*GroupBy {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *QueryMetric) GetAggregators() []*Aggregator {
	if m != nil {
		return m.Aggregators
	}
	return nil
}

type Result struct {
	Name    string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	GroupBy []*GroupBy             `protobuf:"bytes,2,rep,name=group_by,json=groupBy" json:"group_by,omitempty"`
	Tags    map[string]*StringList `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values  []*Datapoint           `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Result) GetGroupBy() []*GroupBy {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *Result) GetTags() map[string]*StringList {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Result) GetValues() []*Datapoint {
	if m != nil {
		return m.Values
	}
	return nil
}

// Represents a single set of results for a given Query
type Query struct {
	Results []*Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *Query) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type Metric struct {
	Name       string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Datapoints []*Datapoint      `protobuf:"bytes,2,rep,name=datapoints" json:"datapoints,omitempty"`
	Tags       map[string]string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ttl        int64             `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
	Type       string            `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *Metric) GetDatapoints() []*Datapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

func (m *Metric) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Request to Add data points to metric or metrics
type AddDatapointsRequest struct {
	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *AddDatapointsRequest) Reset()                    { *m = AddDatapointsRequest{} }
func (m *AddDatapointsRequest) String() string            { return proto.CompactTextString(m) }
func (*AddDatapointsRequest) ProtoMessage()               {}
func (*AddDatapointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *AddDatapointsRequest) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type AddDatapointsResponse struct {
	Errors []string `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
}

func (m *AddDatapointsResponse) Reset()                    { *m = AddDatapointsResponse{} }
func (m *AddDatapointsResponse) String() string            { return proto.CompactTextString(m) }
func (*AddDatapointsResponse) ProtoMessage()               {}
func (*AddDatapointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

// Request to delete data points
type DeleteDatapointsRequest struct {
	Metrics       []*QueryMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	CacheTime     int64          `protobuf:"varint,2,opt,name=cache_time,json=cacheTime" json:"cache_time,omitempty"`
	StartAbsolute int64          `protobuf:"varint,3,opt,name=start_absolute,json=startAbsolute" json:"start_absolute,omitempty"`
	EndAbsolute   int64          `protobuf:"varint,4,opt,name=end_absolute,json=endAbsolute" json:"end_absolute,omitempty"`
}

func (m *DeleteDatapointsRequest) Reset()                    { *m = DeleteDatapointsRequest{} }
func (m *DeleteDatapointsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDatapointsRequest) ProtoMessage()               {}
func (*DeleteDatapointsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *DeleteDatapointsRequest) GetMetrics() []*QueryMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type DeleteDatapointsResponse struct {
	Errors []string `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
}

func (m *DeleteDatapointsResponse) Reset()                    { *m = DeleteDatapointsResponse{} }
func (m *DeleteDatapointsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDatapointsResponse) ProtoMessage()               {}
func (*DeleteDatapointsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

// Reprents a QueryMetrics request to kairosdb
// https://kairosdb.github.io/docs/build/html/restapi/QueryMetrics.html#query-properties
type QueryMetricsRequest struct {
	Metrics       []*QueryMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	CacheTime     int64          `protobuf:"varint,2,opt,name=cache_time,json=cacheTime" json:"cache_time,omitempty"`
	StartAbsolute int64          `protobuf:"varint,3,opt,name=start_absolute,json=startAbsolute" json:"start_absolute,omitempty"`
	EndAbsolute   int64          `protobuf:"varint,4,opt,name=end_absolute,json=endAbsolute" json:"end_absolute,omitempty"`
}

func (m *QueryMetricsRequest) Reset()                    { *m = QueryMetricsRequest{} }
func (m *QueryMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsRequest) ProtoMessage()               {}
func (*QueryMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *QueryMetricsRequest) GetMetrics() []*QueryMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// Represents response to kairosdb QueryMetrics request
type QueryMetricsResponse struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
	Errors  []string `protobuf:"bytes,5,rep,name=errors" json:"errors,omitempty"`
}

func (m *QueryMetricsResponse) Reset()                    { *m = QueryMetricsResponse{} }
func (m *QueryMetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryMetricsResponse) ProtoMessage()               {}
func (*QueryMetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *QueryMetricsResponse) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ListMetricNamesRequest struct {
}

func (m *ListMetricNamesRequest) Reset()                    { *m = ListMetricNamesRequest{} }
func (m *ListMetricNamesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListMetricNamesRequest) ProtoMessage()               {}
func (*ListMetricNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

type ListMetricNamesResponse struct {
	Results []string `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ListMetricNamesResponse) Reset()                    { *m = ListMetricNamesResponse{} }
func (m *ListMetricNamesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListMetricNamesResponse) ProtoMessage()               {}
func (*ListMetricNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

type ListTagNamesRequest struct {
}

func (m *ListTagNamesRequest) Reset()                    { *m = ListTagNamesRequest{} }
func (m *ListTagNamesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTagNamesRequest) ProtoMessage()               {}
func (*ListTagNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

type ListTagNamesResponse struct {
	Results []string `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ListTagNamesResponse) Reset()                    { *m = ListTagNamesResponse{} }
func (m *ListTagNamesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTagNamesResponse) ProtoMessage()               {}
func (*ListTagNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

type ListTagValuesRequest struct {
}

func (m *ListTagValuesRequest) Reset()                    { *m = ListTagValuesRequest{} }
func (m *ListTagValuesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTagValuesRequest) ProtoMessage()               {}
func (*ListTagValuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

type ListTagValuesResponse struct {
	Results []string `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ListTagValuesResponse) Reset()                    { *m = ListTagValuesResponse{} }
func (m *ListTagValuesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTagValuesResponse) ProtoMessage()               {}
func (*ListTagValuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

type HealthCheckRequest struct {
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{20} }

type HealthCheckResponse struct {
	Healthy bool `protobuf:"varint,1,opt,name=healthy" json:"healthy,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{21} }

type HealthStatusRequest struct {
}

func (m *HealthStatusRequest) Reset()                    { *m = HealthStatusRequest{} }
func (m *HealthStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthStatusRequest) ProtoMessage()               {}
func (*HealthStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{22} }

type HealthStatusResponse struct {
	Results []string `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *HealthStatusResponse) Reset()                    { *m = HealthStatusResponse{} }
func (m *HealthStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthStatusResponse) ProtoMessage()               {}
func (*HealthStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{23} }

func init() {
	proto.RegisterType((*Datapoint)(nil), "kairosdb.Datapoint")
	proto.RegisterType((*Sampling)(nil), "kairosdb.Sampling")
	proto.RegisterType((*Aggregator)(nil), "kairosdb.Aggregator")
	proto.RegisterType((*GroupBy)(nil), "kairosdb.GroupBy")
	proto.RegisterType((*QueryMetric)(nil), "kairosdb.QueryMetric")
	proto.RegisterType((*Result)(nil), "kairosdb.Result")
	proto.RegisterType((*Query)(nil), "kairosdb.Query")
	proto.RegisterType((*Metric)(nil), "kairosdb.Metric")
	proto.RegisterType((*AddDatapointsRequest)(nil), "kairosdb.AddDatapointsRequest")
	proto.RegisterType((*AddDatapointsResponse)(nil), "kairosdb.AddDatapointsResponse")
	proto.RegisterType((*DeleteDatapointsRequest)(nil), "kairosdb.DeleteDatapointsRequest")
	proto.RegisterType((*DeleteDatapointsResponse)(nil), "kairosdb.DeleteDatapointsResponse")
	proto.RegisterType((*QueryMetricsRequest)(nil), "kairosdb.QueryMetricsRequest")
	proto.RegisterType((*QueryMetricsResponse)(nil), "kairosdb.QueryMetricsResponse")
	proto.RegisterType((*ListMetricNamesRequest)(nil), "kairosdb.ListMetricNamesRequest")
	proto.RegisterType((*ListMetricNamesResponse)(nil), "kairosdb.ListMetricNamesResponse")
	proto.RegisterType((*ListTagNamesRequest)(nil), "kairosdb.ListTagNamesRequest")
	proto.RegisterType((*ListTagNamesResponse)(nil), "kairosdb.ListTagNamesResponse")
	proto.RegisterType((*ListTagValuesRequest)(nil), "kairosdb.ListTagValuesRequest")
	proto.RegisterType((*ListTagValuesResponse)(nil), "kairosdb.ListTagValuesResponse")
	proto.RegisterType((*HealthCheckRequest)(nil), "kairosdb.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "kairosdb.HealthCheckResponse")
	proto.RegisterType((*HealthStatusRequest)(nil), "kairosdb.HealthStatusRequest")
	proto.RegisterType((*HealthStatusResponse)(nil), "kairosdb.HealthStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for KairosDB service

type KairosDBClient interface {
	ListMetricNames(ctx context.Context, in *ListMetricNamesRequest, opts ...grpc.CallOption) (*ListMetricNamesResponse, error)
	ListTagNames(ctx context.Context, in *ListTagNamesRequest, opts ...grpc.CallOption) (*ListTagNamesResponse, error)
	ListTagValues(ctx context.Context, in *ListTagValuesRequest, opts ...grpc.CallOption) (*ListTagValuesResponse, error)
	QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error)
	HealthStatus(ctx context.Context, in *HealthStatusRequest, opts ...grpc.CallOption) (*HealthStatusResponse, error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	AddDatapoints(ctx context.Context, in *AddDatapointsRequest, opts ...grpc.CallOption) (*AddDatapointsResponse, error)
	DeleteDatapoints(ctx context.Context, in *DeleteDatapointsRequest, opts ...grpc.CallOption) (*DeleteDatapointsResponse, error)
}

type kairosDBClient struct {
	cc *grpc.ClientConn
}

func NewKairosDBClient(cc *grpc.ClientConn) KairosDBClient {
	return &kairosDBClient{cc}
}

func (c *kairosDBClient) ListMetricNames(ctx context.Context, in *ListMetricNamesRequest, opts ...grpc.CallOption) (*ListMetricNamesResponse, error) {
	out := new(ListMetricNamesResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/ListMetricNames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) ListTagNames(ctx context.Context, in *ListTagNamesRequest, opts ...grpc.CallOption) (*ListTagNamesResponse, error) {
	out := new(ListTagNamesResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/ListTagNames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) ListTagValues(ctx context.Context, in *ListTagValuesRequest, opts ...grpc.CallOption) (*ListTagValuesResponse, error) {
	out := new(ListTagValuesResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/ListTagValues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error) {
	out := new(QueryMetricsResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/QueryMetrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) HealthStatus(ctx context.Context, in *HealthStatusRequest, opts ...grpc.CallOption) (*HealthStatusResponse, error) {
	out := new(HealthStatusResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/HealthStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) AddDatapoints(ctx context.Context, in *AddDatapointsRequest, opts ...grpc.CallOption) (*AddDatapointsResponse, error) {
	out := new(AddDatapointsResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/AddDatapoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kairosDBClient) DeleteDatapoints(ctx context.Context, in *DeleteDatapointsRequest, opts ...grpc.CallOption) (*DeleteDatapointsResponse, error) {
	out := new(DeleteDatapointsResponse)
	err := grpc.Invoke(ctx, "/kairosdb.KairosDB/DeleteDatapoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KairosDB service

type KairosDBServer interface {
	ListMetricNames(context.Context, *ListMetricNamesRequest) (*ListMetricNamesResponse, error)
	ListTagNames(context.Context, *ListTagNamesRequest) (*ListTagNamesResponse, error)
	ListTagValues(context.Context, *ListTagValuesRequest) (*ListTagValuesResponse, error)
	QueryMetrics(context.Context, *QueryMetricsRequest) (*QueryMetricsResponse, error)
	HealthStatus(context.Context, *HealthStatusRequest) (*HealthStatusResponse, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	AddDatapoints(context.Context, *AddDatapointsRequest) (*AddDatapointsResponse, error)
	DeleteDatapoints(context.Context, *DeleteDatapointsRequest) (*DeleteDatapointsResponse, error)
}

func RegisterKairosDBServer(s *grpc.Server, srv KairosDBServer) {
	s.RegisterService(&_KairosDB_serviceDesc, srv)
}

func _KairosDB_ListMetricNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).ListMetricNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/ListMetricNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).ListMetricNames(ctx, req.(*ListMetricNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_ListTagNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).ListTagNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/ListTagNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).ListTagNames(ctx, req.(*ListTagNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_ListTagValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).ListTagValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/ListTagValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).ListTagValues(ctx, req.(*ListTagValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_QueryMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).QueryMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/QueryMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).QueryMetrics(ctx, req.(*QueryMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_HealthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).HealthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/HealthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).HealthStatus(ctx, req.(*HealthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_AddDatapoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDatapointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).AddDatapoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/AddDatapoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).AddDatapoints(ctx, req.(*AddDatapointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KairosDB_DeleteDatapoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatapointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KairosDBServer).DeleteDatapoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kairosdb.KairosDB/DeleteDatapoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KairosDBServer).DeleteDatapoints(ctx, req.(*DeleteDatapointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KairosDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kairosdb.KairosDB",
	HandlerType: (*KairosDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMetricNames",
			Handler:    _KairosDB_ListMetricNames_Handler,
		},
		{
			MethodName: "ListTagNames",
			Handler:    _KairosDB_ListTagNames_Handler,
		},
		{
			MethodName: "ListTagValues",
			Handler:    _KairosDB_ListTagValues_Handler,
		},
		{
			MethodName: "QueryMetrics",
			Handler:    _KairosDB_QueryMetrics_Handler,
		},
		{
			MethodName: "HealthStatus",
			Handler:    _KairosDB_HealthStatus_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _KairosDB_HealthCheck_Handler,
		},
		{
			MethodName: "AddDatapoints",
			Handler:    _KairosDB_AddDatapoints_Handler,
		},
		{
			MethodName: "DeleteDatapoints",
			Handler:    _KairosDB_DeleteDatapoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("kairosdb_types.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x96, 0x93, 0x9b, 0xc4, 0x39, 0xb9, 0xe5, 0x86, 0x49, 0xd2, 0x1b, 0x99, 0x96, 0xb6, 0x96,
	0x90, 0xca, 0x8f, 0x52, 0x48, 0x11, 0x54, 0x6c, 0x50, 0x43, 0x11, 0x08, 0x28, 0x82, 0x69, 0x55,
	0xa9, 0xdd, 0x44, 0x4e, 0x32, 0x72, 0xac, 0x3a, 0x76, 0xb0, 0x27, 0x48, 0xe1, 0x51, 0xd8, 0x20,
	0x9e, 0x80, 0x05, 0x7b, 0x9e, 0x85, 0xa7, 0x60, 0xcd, 0xfc, 0x78, 0xec, 0xb1, 0xe3, 0x24, 0x1b,
	0x36, 0xac, 0x3a, 0x3e, 0xf3, 0x9d, 0x33, 0xdf, 0xf9, 0xce, 0x4f, 0x03, 0xdd, 0x67, 0xc7, 0x8b,
	0xc2, 0x78, 0x36, 0x19, 0xd3, 0xf5, 0x92, 0xc4, 0x83, 0x65, 0x14, 0xd2, 0x10, 0x99, 0xca, 0x6a,
	0xa1, 0x39, 0xf1, 0x97, 0x24, 0xd2, 0x6f, 0xed, 0xcf, 0xa1, 0x79, 0xe3, 0x50, 0x67, 0x19, 0x7a,
	0x01, 0x45, 0x47, 0xd0, 0xa4, 0xde, 0x82, 0xc4, 0xd4, 0x59, 0x2c, 0xfb, 0xc6, 0xa9, 0x71, 0x5e,
	0xc5, 0x99, 0x01, 0x75, 0xa1, 0xf6, 0xb3, 0xe3, 0xaf, 0x48, 0xbf, 0xc2, 0x6e, 0x0c, 0x2c, 0x3f,
	0xec, 0x8f, 0xc1, 0xbc, 0x63, 0xb7, 0xbe, 0x17, 0xb8, 0x19, 0x82, 0xfb, 0x36, 0x13, 0x04, 0x42,
	0xf0, 0x62, 0x15, 0x78, 0x54, 0xb8, 0x35, 0xb1, 0x38, 0xdb, 0xbf, 0x1b, 0x00, 0xd7, 0xae, 0x1b,
	0x11, 0xd7, 0xa1, 0x61, 0xc4, 0x21, 0x81, 0xb3, 0x50, 0x7e, 0xe2, 0x8c, 0xde, 0x81, 0x37, 0x1c,
	0xdf, 0x73, 0x83, 0x71, 0x9c, 0x84, 0x17, 0x01, 0x4c, 0x7c, 0x20, 0xac, 0xe9, 0x9b, 0x03, 0x30,
	0x53, 0x40, 0x95, 0x01, 0x5a, 0x43, 0x34, 0x50, 0x19, 0x0f, 0x14, 0x0a, 0xa7, 0x18, 0x74, 0x0e,
	0xed, 0x24, 0x2c, 0x75, 0x22, 0x3a, 0xe6, 0xe9, 0xf5, 0x5f, 0x88, 0xc0, 0xf2, 0xb9, 0x3b, 0x6e,
	0xbe, 0x67, 0x56, 0xfb, 0x37, 0x03, 0x1a, 0x5f, 0x45, 0xe1, 0x6a, 0x39, 0x5a, 0x97, 0x12, 0x64,
	0x36, 0xea, 0xb8, 0x31, 0xa3, 0x55, 0xe5, 0x36, 0x7e, 0x46, 0x43, 0xa8, 0xb9, 0xdc, 0x85, 0x51,
	0xa9, 0x32, 0x2a, 0x47, 0x19, 0x95, 0x24, 0x92, 0xfc, 0xfb, 0x65, 0x40, 0xa3, 0x35, 0x96, 0x50,
	0xeb, 0x0a, 0x20, 0x33, 0xa2, 0x36, 0x54, 0x9f, 0xc9, 0x3a, 0x79, 0x88, 0x1f, 0xf3, 0xba, 0x2b,
	0x55, 0x3f, 0xab, 0x5c, 0x19, 0xf6, 0x5f, 0x15, 0x68, 0xfd, 0xb8, 0x22, 0xd1, 0xfa, 0x96, 0xd0,
	0xc8, 0x9b, 0x96, 0xb2, 0xbc, 0xd4, 0x58, 0xb6, 0x86, 0x27, 0x19, 0x21, 0xcd, 0x71, 0x70, 0xcf,
	0x10, 0x92, 0x93, 0x4c, 0xe3, 0x03, 0x30, 0x05, 0xb7, 0xf1, 0x64, 0x9d, 0x64, 0xf2, 0xe6, 0x46,
	0x26, 0xb8, 0xe1, 0x26, 0xe2, 0x7c, 0x02, 0x2d, 0x27, 0xad, 0x65, 0xcc, 0xd4, 0xe4, 0x0e, 0xdd,
	0xcc, 0x21, 0x2b, 0x34, 0xd6, 0x81, 0x3c, 0x31, 0xdf, 0x5b, 0xb0, 0xce, 0xa8, 0x89, 0x56, 0x93,
	0x1f, 0xe8, 0x2d, 0xd9, 0x84, 0xe3, 0x5f, 0xc2, 0x80, 0xf4, 0xeb, 0x22, 0x13, 0x93, 0x1b, 0x9e,
	0xd8, 0xb7, 0x75, 0x0b, 0xcd, 0x94, 0x6b, 0x89, 0x54, 0xef, 0xe9, 0x52, 0xe5, 0x38, 0xdc, 0xb1,
	0x3c, 0x03, 0xf7, 0x3b, 0x2f, 0xa6, 0xba, 0x80, 0xff, 0x18, 0x50, 0xc7, 0x24, 0x5e, 0xf9, 0xb4,
	0x54, 0x3b, 0x5d, 0x86, 0xca, 0x5e, 0x19, 0x06, 0x89, 0xd2, 0x52, 0x30, 0x2b, 0x43, 0xca, 0x17,
	0x36, 0x44, 0x7e, 0x1f, 0xea, 0x82, 0x89, 0x52, 0xac, 0x93, 0x79, 0xa4, 0x23, 0x89, 0x13, 0xc8,
	0x7f, 0x9d, 0xf8, 0x25, 0xd4, 0x44, 0xfd, 0x99, 0x63, 0x23, 0x12, 0xf4, 0x62, 0x16, 0x8e, 0xb3,
	0x68, 0x17, 0x79, 0x63, 0x05, 0xb0, 0xff, 0x66, 0x6a, 0xed, 0xec, 0x34, 0x98, 0x29, 0xde, 0xaa,
	0xdf, 0x4a, 0x73, 0xd2, 0x60, 0xdb, 0x45, 0xdb, 0xd2, 0x99, 0x2c, 0x75, 0x4a, 0x7d, 0x31, 0xb1,
	0x55, 0xcc, 0x8f, 0x62, 0x0c, 0xd9, 0x42, 0x13, 0x4d, 0xc4, 0xc7, 0x90, 0x9d, 0xad, 0x4f, 0x77,
	0xab, 0xb5, 0x7d, 0xa2, 0x46, 0xd0, 0xbd, 0x9e, 0xcd, 0x52, 0xaa, 0x31, 0x26, 0x3f, 0x31, 0xf5,
	0x29, 0x97, 0x69, 0x21, 0x08, 0x95, 0xc8, 0x24, 0x99, 0x62, 0x05, 0xb0, 0x2f, 0xa0, 0x57, 0x88,
	0x11, 0x2f, 0xc3, 0x20, 0x26, 0xe8, 0x10, 0xea, 0x24, 0x8a, 0xf8, 0x88, 0x18, 0x62, 0x65, 0x24,
	0x5f, 0xf6, 0x9f, 0x06, 0xbc, 0xbe, 0x21, 0x3e, 0xa1, 0x64, 0xf3, 0xe1, 0x8b, 0xe2, 0xc3, 0xbd,
	0xd2, 0x09, 0x4e, 0x5f, 0x47, 0xc7, 0x00, 0x53, 0x67, 0x3a, 0x27, 0x72, 0xb3, 0x55, 0xe4, 0x12,
	0x17, 0x16, 0xbe, 0xd4, 0xf8, 0x56, 0x95, 0x8b, 0xcf, 0x99, 0xc4, 0xa1, 0xbf, 0xa2, 0x44, 0x2c,
	0xcd, 0x2a, 0x3e, 0x10, 0xd6, 0xeb, 0xc4, 0x88, 0xce, 0xe0, 0x25, 0x09, 0x66, 0x19, 0x48, 0xea,
	0xdd, 0x62, 0x36, 0x05, 0xb1, 0x87, 0xd0, 0xdf, 0x24, 0xbd, 0x27, 0xd3, 0x3f, 0x0c, 0xe8, 0x68,
	0xac, 0xff, 0x07, 0x59, 0x3e, 0x42, 0x37, 0x4f, 0x38, 0xc9, 0xf0, 0x5d, 0x68, 0x30, 0xea, 0x91,
	0x47, 0x14, 0xe3, 0x57, 0x05, 0xc6, 0x58, 0xdd, 0x6b, 0x62, 0xd4, 0x72, 0x62, 0xf4, 0xe1, 0x90,
	0x8f, 0xa5, 0x8c, 0xfc, 0x3d, 0x9b, 0x20, 0x25, 0x07, 0x9b, 0xce, 0xd7, 0x1b, 0x37, 0xc9, 0xbb,
	0xfd, 0xfc, 0xbc, 0x36, 0xb3, 0xe9, 0xec, 0x41, 0x87, 0x3b, 0xb1, 0xbe, 0xcf, 0xc5, 0xfa, 0x10,
	0xba, 0x79, 0xf3, 0xde, 0x40, 0x87, 0xa9, 0xc7, 0x83, 0xd8, 0x3d, 0x2a, 0xd2, 0x47, 0xd0, 0x2b,
	0xd8, 0xf7, 0x86, 0xea, 0x02, 0xfa, 0x9a, 0x38, 0x3e, 0x9d, 0x7f, 0x31, 0x27, 0xd3, 0x67, 0x15,
	0xe8, 0x02, 0x3a, 0x39, 0x6b, 0x16, 0x66, 0x2e, 0xcc, 0x72, 0x56, 0x4d, 0xac, 0x3e, 0x79, 0x6a,
	0xd2, 0x81, 0xfd, 0x73, 0xa6, 0x2b, 0x3d, 0xb5, 0xbc, 0x79, 0x1f, 0x9f, 0xe1, 0xaf, 0x35, 0x30,
	0xbf, 0x15, 0x65, 0xba, 0x19, 0xa1, 0x07, 0x78, 0x55, 0x50, 0x19, 0x9d, 0x66, 0x45, 0x2c, 0x2f,
	0x8d, 0x75, 0xb6, 0x03, 0x91, 0x3c, 0x7f, 0x0b, 0x2f, 0x75, 0xc5, 0xd1, 0x71, 0xde, 0xa5, 0x50,
	0x20, 0xeb, 0xed, 0x6d, 0xd7, 0x49, 0xb8, 0x1f, 0xe0, 0x20, 0x27, 0x3b, 0xda, 0x74, 0xc8, 0xd5,
	0xc9, 0x3a, 0xd9, 0x7a, 0x9f, 0x11, 0xd4, 0x7b, 0x5a, 0x27, 0x58, 0x32, 0x9c, 0x3a, 0xc1, 0xd2,
	0x51, 0x60, 0xe1, 0xf4, 0x32, 0xe8, 0xe1, 0x4a, 0xaa, 0xa6, 0x87, 0x2b, 0xad, 0xde, 0x37, 0xd0,
	0xd2, 0xba, 0x03, 0x1d, 0x15, 0xe1, 0x7a, 0x2b, 0x59, 0xc7, 0x5b, 0x6e, 0x33, 0xed, 0x72, 0xab,
	0x58, 0xd7, 0xae, 0x6c, 0xcf, 0xeb, 0xda, 0x95, 0xef, 0xf0, 0x47, 0x68, 0x17, 0xb7, 0x1e, 0xd2,
	0x7a, 0x62, 0xcb, 0x1a, 0xb7, 0xec, 0x5d, 0x10, 0x19, 0x7a, 0x04, 0x4f, 0xe9, 0x4f, 0xf5, 0x49,
	0x5d, 0xfc, 0x3a, 0xbf, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x62, 0xb1, 0x7c, 0xd3, 0x0b,
	0x00, 0x00,
}
