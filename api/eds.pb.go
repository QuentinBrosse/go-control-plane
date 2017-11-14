// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/eds.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LbEndpoint struct {
	Endpoint            *Endpoint                    `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	HealthStatus        HealthStatus                 `protobuf:"varint,2,opt,name=health_status,json=healthStatus,enum=envoy.api.v2.HealthStatus" json:"health_status,omitempty"`
	Metadata            *Metadata                    `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	LoadBalancingWeight *google_protobuf.UInt32Value `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight" json:"load_balancing_weight,omitempty"`
}

func (m *LbEndpoint) Reset()                    { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string            { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()               {}
func (*LbEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetHealthStatus() HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *google_protobuf.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

type LocalityLbEndpoints struct {
	Locality            *Locality                    `protobuf:"bytes,1,opt,name=locality" json:"locality,omitempty"`
	LbEndpoints         []*LbEndpoint                `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints" json:"lb_endpoints,omitempty"`
	LoadBalancingWeight *google_protobuf.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight" json:"load_balancing_weight,omitempty"`
}

func (m *LocalityLbEndpoints) Reset()                    { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string            { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()               {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *LocalityLbEndpoints) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *google_protobuf.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

type EndpointLoadMetricStats struct {
	MetricName                    string  `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
	NumRequestsFinishedWithMetric uint64  `protobuf:"varint,2,opt,name=num_requests_finished_with_metric,json=numRequestsFinishedWithMetric" json:"num_requests_finished_with_metric,omitempty"`
	TotalMetricValue              float64 `protobuf:"fixed64,3,opt,name=total_metric_value,json=totalMetricValue" json:"total_metric_value,omitempty"`
}

func (m *EndpointLoadMetricStats) Reset()                    { *m = EndpointLoadMetricStats{} }
func (m *EndpointLoadMetricStats) String() string            { return proto.CompactTextString(m) }
func (*EndpointLoadMetricStats) ProtoMessage()               {}
func (*EndpointLoadMetricStats) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *EndpointLoadMetricStats) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *EndpointLoadMetricStats) GetNumRequestsFinishedWithMetric() uint64 {
	if m != nil {
		return m.NumRequestsFinishedWithMetric
	}
	return 0
}

func (m *EndpointLoadMetricStats) GetTotalMetricValue() float64 {
	if m != nil {
		return m.TotalMetricValue
	}
	return 0
}

type UpstreamLocalityStats struct {
	Locality                *Locality                  `protobuf:"bytes,1,opt,name=locality" json:"locality,omitempty"`
	TotalSuccessfulRequests uint64                     `protobuf:"varint,2,opt,name=total_successful_requests,json=totalSuccessfulRequests" json:"total_successful_requests,omitempty"`
	TotalRequestsInProgress uint64                     `protobuf:"varint,3,opt,name=total_requests_in_progress,json=totalRequestsInProgress" json:"total_requests_in_progress,omitempty"`
	TotalErrorRequests      uint64                     `protobuf:"varint,4,opt,name=total_error_requests,json=totalErrorRequests" json:"total_error_requests,omitempty"`
	LoadMetricStats         []*EndpointLoadMetricStats `protobuf:"bytes,5,rep,name=load_metric_stats,json=loadMetricStats" json:"load_metric_stats,omitempty"`
}

func (m *UpstreamLocalityStats) Reset()                    { *m = UpstreamLocalityStats{} }
func (m *UpstreamLocalityStats) String() string            { return proto.CompactTextString(m) }
func (*UpstreamLocalityStats) ProtoMessage()               {}
func (*UpstreamLocalityStats) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *UpstreamLocalityStats) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *UpstreamLocalityStats) GetTotalSuccessfulRequests() uint64 {
	if m != nil {
		return m.TotalSuccessfulRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalRequestsInProgress() uint64 {
	if m != nil {
		return m.TotalRequestsInProgress
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalErrorRequests() uint64 {
	if m != nil {
		return m.TotalErrorRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetLoadMetricStats() []*EndpointLoadMetricStats {
	if m != nil {
		return m.LoadMetricStats
	}
	return nil
}

type ClusterStats struct {
	ClusterName           string                   `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	UpstreamLocalityStats []*UpstreamLocalityStats `protobuf:"bytes,2,rep,name=upstream_locality_stats,json=upstreamLocalityStats" json:"upstream_locality_stats,omitempty"`
	TotalDroppedRequests  uint64                   `protobuf:"varint,3,opt,name=total_dropped_requests,json=totalDroppedRequests" json:"total_dropped_requests,omitempty"`
}

func (m *ClusterStats) Reset()                    { *m = ClusterStats{} }
func (m *ClusterStats) String() string            { return proto.CompactTextString(m) }
func (*ClusterStats) ProtoMessage()               {}
func (*ClusterStats) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *ClusterStats) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterStats) GetUpstreamLocalityStats() []*UpstreamLocalityStats {
	if m != nil {
		return m.UpstreamLocalityStats
	}
	return nil
}

func (m *ClusterStats) GetTotalDroppedRequests() uint64 {
	if m != nil {
		return m.TotalDroppedRequests
	}
	return 0
}

type LoadStatsRequest struct {
	Node         *Node           `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	ClusterStats []*ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats" json:"cluster_stats,omitempty"`
}

func (m *LoadStatsRequest) Reset()                    { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()               {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *LoadStatsRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

type ClusterLoadAssignment struct {
	ClusterName       string                        `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Endpoints         []*LocalityLbEndpoints        `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
	FailoverEndpoints []*LocalityLbEndpoints        `protobuf:"bytes,3,rep,name=failover_endpoints,json=failoverEndpoints" json:"failover_endpoints,omitempty"`
	Policy            *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy" json:"policy,omitempty"`
}

func (m *ClusterLoadAssignment) Reset()                    { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string            { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()               {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetFailoverEndpoints() []*LocalityLbEndpoints {
	if m != nil {
		return m.FailoverEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverload float64 `protobuf:"fixed64,1,opt,name=drop_overload,json=dropOverload" json:"drop_overload,omitempty"`
}

func (m *ClusterLoadAssignment_Policy) Reset()                    { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string            { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()               {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6, 0} }

func (m *ClusterLoadAssignment_Policy) GetDropOverload() float64 {
	if m != nil {
		return m.DropOverload
	}
	return 0
}

type LoadStatsResponse struct {
	Clusters              []string                   `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	LoadReportingInterval *google_protobuf2.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval" json:"load_reporting_interval,omitempty"`
}

func (m *LoadStatsResponse) Reset()                    { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string            { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()               {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.LocalityLbEndpoints")
	proto.RegisterType((*EndpointLoadMetricStats)(nil), "envoy.api.v2.EndpointLoadMetricStats")
	proto.RegisterType((*UpstreamLocalityStats)(nil), "envoy.api.v2.UpstreamLocalityStats")
	proto.RegisterType((*ClusterStats)(nil), "envoy.api.v2.ClusterStats")
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.api.v2.LoadStatsRequest")
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.api.v2.LoadStatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EndpointDiscoveryService service

type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamLoadStatsClient, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointDiscoveryServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamLoadStatsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], c.cc, "/envoy.api.v2.EndpointDiscoveryService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamLoadStatsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EndpointDiscoveryService service

type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
	StreamLoadStats(EndpointDiscoveryService_StreamLoadStatsServer) error
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointDiscoveryService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamLoadStats(&endpointDiscoveryServiceStreamLoadStatsServer{stream})
}

type EndpointDiscoveryService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamLoadStats",
			Handler:       _EndpointDiscoveryService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/eds.proto",
}

func init() { proto.RegisterFile("api/eds.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xee, 0xca, 0xae, 0x11, 0x8f, 0x64, 0x27, 0xa6, 0x6b, 0x4b, 0x11, 0xdc, 0xc4, 0x56, 0xd0,
	0x42, 0x08, 0x5a, 0x39, 0x50, 0x7a, 0x72, 0x0e, 0x41, 0x53, 0x27, 0x88, 0x01, 0x27, 0x75, 0x68,
	0xa4, 0x29, 0xd0, 0xc3, 0x82, 0xda, 0xa5, 0xb5, 0x44, 0x57, 0xe4, 0x96, 0xe4, 0xca, 0x30, 0x7a,
	0xcb, 0xa5, 0x0f, 0xd0, 0x37, 0xe9, 0x13, 0xf4, 0xd0, 0x1e, 0x7a, 0xee, 0xa5, 0x0f, 0xd0, 0x07,
	0x29, 0xf8, 0xb7, 0x5a, 0x29, 0x32, 0x10, 0x04, 0xbd, 0x89, 0x33, 0xf3, 0x71, 0xe6, 0xfb, 0x66,
	0x38, 0x2b, 0xd8, 0x20, 0x05, 0x3b, 0xa4, 0xa9, 0x1a, 0x14, 0x52, 0x68, 0x81, 0x5a, 0x94, 0x4f,
	0xc5, 0xd5, 0x80, 0x14, 0x6c, 0x30, 0x1d, 0x76, 0x37, 0x8d, 0x73, 0x44, 0x14, 0x75, 0xde, 0xee,
	0xb6, 0x39, 0xa7, 0x4c, 0x25, 0x62, 0x4a, 0xe5, 0x95, 0x37, 0xee, 0x1a, 0x63, 0x46, 0x49, 0xae,
	0xb3, 0x38, 0xc9, 0x68, 0xf2, 0xa3, 0xb7, 0xef, 0x8d, 0x85, 0x18, 0xe7, 0xf4, 0xd0, 0xb8, 0x09,
	0xe7, 0x42, 0x13, 0xcd, 0x04, 0xf7, 0x89, 0xba, 0x77, 0xbc, 0xd7, 0x9e, 0x46, 0xe5, 0xc5, 0x61,
	0x5a, 0x4a, 0x1b, 0x70, 0x9d, 0xff, 0x52, 0x92, 0xa2, 0xa0, 0xd2, 0xe3, 0x7b, 0xbf, 0x34, 0x00,
	0x4e, 0x47, 0x4f, 0x79, 0x5a, 0x08, 0xc6, 0x35, 0x1a, 0xc2, 0x0d, 0xea, 0x7f, 0x77, 0xa2, 0xfd,
	0xa8, 0xdf, 0x1c, 0xee, 0x0e, 0xea, 0x54, 0x06, 0x21, 0x12, 0x57, 0x71, 0xe8, 0x31, 0x6c, 0xf8,
	0xb2, 0x95, 0x26, 0xba, 0x54, 0x9d, 0xc6, 0x7e, 0xd4, 0xdf, 0x1c, 0x76, 0xe7, 0x81, 0xcf, 0x6d,
	0xc8, 0xb9, 0x8d, 0xc0, 0xad, 0xac, 0x76, 0x32, 0x49, 0x27, 0x54, 0x93, 0x94, 0x68, 0xd2, 0x59,
	0x59, 0x96, 0xf4, 0x85, 0xf7, 0xe2, 0x2a, 0x0e, 0x9d, 0xc1, 0x4e, 0x2e, 0x48, 0x1a, 0x8f, 0x48,
	0x4e, 0x78, 0xc2, 0xf8, 0x38, 0xbe, 0xa4, 0x6c, 0x9c, 0xe9, 0xce, 0xaa, 0xbd, 0x60, 0x6f, 0xe0,
	0x78, 0x0f, 0x02, 0xef, 0xc1, 0xeb, 0x13, 0xae, 0x1f, 0x0e, 0xbf, 0x23, 0x79, 0x49, 0xf1, 0xb6,
	0x81, 0x3e, 0x09, 0xc8, 0x37, 0x16, 0xd8, 0xfb, 0x27, 0x82, 0xed, 0x53, 0x91, 0x90, 0x9c, 0xe9,
	0xab, 0x99, 0x22, 0xb6, 0xba, 0xdc, 0x9b, 0x97, 0x4b, 0x12, 0x40, 0xb8, 0x8a, 0x43, 0x8f, 0xa0,
	0x95, 0x8f, 0xe2, 0xa0, 0x90, 0x51, 0x64, 0xa5, 0xdf, 0x1c, 0x76, 0x16, 0x70, 0x55, 0x12, 0xdc,
	0xcc, 0x6b, 0x09, 0xaf, 0xa5, 0xb6, 0xf2, 0xa1, 0xd4, 0x7e, 0x8b, 0xa0, 0x1d, 0xee, 0x3f, 0x15,
	0x24, 0x7d, 0x41, 0xb5, 0x64, 0x89, 0x51, 0x5f, 0xa1, 0xbb, 0xd0, 0x9c, 0xd8, 0x63, 0xcc, 0xc9,
	0x84, 0x5a, 0x86, 0xeb, 0x18, 0x9c, 0xe9, 0x25, 0x99, 0x50, 0xf4, 0x1c, 0x0e, 0x78, 0x39, 0x89,
	0x25, 0xfd, 0xa9, 0xa4, 0x4a, 0xab, 0xf8, 0x82, 0x71, 0xa6, 0x32, 0x9a, 0xc6, 0x97, 0x4c, 0x67,
	0xb1, 0x0b, 0xb4, 0x2d, 0x5f, 0xc5, 0x9f, 0xf2, 0x72, 0x82, 0x7d, 0xdc, 0x33, 0x1f, 0xf6, 0x86,
	0xe9, 0xcc, 0xe5, 0x43, 0x5f, 0x00, 0xd2, 0x42, 0x93, 0xdc, 0x83, 0xe2, 0xa9, 0xa9, 0xd8, 0xb2,
	0x8a, 0xf0, 0x2d, 0xeb, 0x71, 0x81, 0x96, 0x49, 0xef, 0xaf, 0x06, 0xec, 0xbc, 0x2e, 0x94, 0x96,
	0x94, 0x4c, 0x82, 0xc4, 0xae, 0xe4, 0x0f, 0xe9, 0xc8, 0x11, 0xdc, 0x76, 0xb9, 0x55, 0x99, 0x24,
	0x54, 0xa9, 0x8b, 0x32, 0xaf, 0x28, 0xf9, 0xea, 0xdb, 0x36, 0xe0, 0xbc, 0xf2, 0x07, 0x26, 0xe8,
	0x11, 0x74, 0x1d, 0xb6, 0xd2, 0x80, 0xf1, 0xb8, 0x90, 0x62, 0x2c, 0xa9, 0x52, 0xb6, 0xfe, 0x00,
	0x0e, 0x90, 0x13, 0x7e, 0xe6, 0xdd, 0xe8, 0x01, 0x7c, 0xe2, 0xc0, 0x54, 0x4a, 0x21, 0x67, 0x39,
	0x57, 0x2d, 0xcc, 0x09, 0xf2, 0xd4, 0xb8, 0xaa, 0x74, 0xaf, 0x60, 0xcb, 0xf6, 0xdf, 0xab, 0x64,
	0x1e, 0x95, 0xea, 0x7c, 0x6c, 0x27, 0xe8, 0xb3, 0xe5, 0x8f, 0x71, 0xa1, 0xa7, 0xf8, 0x66, 0x3e,
	0x6f, 0xe8, 0xfd, 0x11, 0x41, 0xeb, 0x9b, 0xbc, 0x54, 0x9a, 0x4a, 0x27, 0xe1, 0x01, 0xb4, 0x12,
	0x77, 0xae, 0xb7, 0xbd, 0xe9, 0x6d, 0xb6, 0xef, 0x3f, 0x40, 0xbb, 0xf4, 0xf2, 0xc7, 0x41, 0x46,
	0x5f, 0x8c, 0x1b, 0xe7, 0x7b, 0xf3, 0xc5, 0x2c, 0xed, 0x15, 0xde, 0x29, 0x97, 0xb6, 0xf0, 0x2b,
	0xd8, 0x75, 0xaa, 0xa4, 0x52, 0x14, 0x05, 0x4d, 0x67, 0xba, 0x38, 0x39, 0x9d, 0x66, 0xc7, 0xce,
	0x19, 0x94, 0xe9, 0xfd, 0x0c, 0xb7, 0x0c, 0x55, 0x77, 0xb3, 0x33, 0xa2, 0xcf, 0x61, 0x95, 0x8b,
	0x94, 0xfa, 0x41, 0x40, 0xf3, 0x35, 0xbd, 0x14, 0x29, 0xc5, 0xd6, 0x6f, 0xb6, 0x54, 0x60, 0x5c,
	0x27, 0xb1, 0xb0, 0xa5, 0xea, 0x22, 0xe1, 0x20, 0x91, 0xd3, 0xf0, 0xf7, 0x06, 0xec, 0x78, 0xb7,
	0x29, 0xe2, 0x6b, 0xa5, 0xd8, 0x98, 0x4f, 0x28, 0xd7, 0xef, 0x23, 0xe6, 0x63, 0x58, 0x5f, 0xdc,
	0x06, 0x07, 0xcb, 0x67, 0xb6, 0xb6, 0x7a, 0xf0, 0x0c, 0x83, 0xce, 0x00, 0x5d, 0x10, 0x96, 0x9b,
	0x0f, 0x46, 0x6d, 0xaf, 0xac, 0xbc, 0xef, 0x4d, 0x5b, 0x01, 0x3c, 0x5b, 0x33, 0x4f, 0x60, 0xad,
	0x10, 0x39, 0x4b, 0xae, 0xfc, 0xca, 0xbc, 0xbf, 0x54, 0x89, 0x79, 0xaa, 0x83, 0x33, 0x8b, 0xc0,
	0x1e, 0xd9, 0xfd, 0x12, 0xd6, 0x9c, 0x05, 0xdd, 0x83, 0x0d, 0xd3, 0xca, 0xd8, 0xe4, 0x30, 0xd3,
	0x67, 0x45, 0x88, 0x70, 0xcb, 0x18, 0xbf, 0xf5, 0xb6, 0xde, 0xdb, 0x08, 0xb6, 0x6a, 0x0d, 0x54,
	0x85, 0xe0, 0x8a, 0xa2, 0x2e, 0xdc, 0xf0, 0x52, 0xa9, 0x4e, 0xb4, 0xbf, 0xd2, 0x5f, 0xc7, 0xd5,
	0x19, 0xbd, 0x82, 0xb6, 0x7d, 0x0b, 0x92, 0x16, 0x42, 0x6a, 0xb3, 0x0b, 0x19, 0xd7, 0x54, 0x4e,
	0x49, 0x6e, 0x1f, 0x6d, 0x73, 0x78, 0xfb, 0x9d, 0x6d, 0x78, 0xec, 0x3f, 0x80, 0xd8, 0x6e, 0x51,
	0x1c, 0x80, 0x27, 0x1e, 0x37, 0xfc, 0xb3, 0x01, 0x9d, 0xa0, 0xc2, 0x71, 0xf8, 0x06, 0x9f, 0x53,
	0x39, 0x65, 0x09, 0x45, 0xdf, 0xc3, 0xcd, 0x73, 0x3b, 0xae, 0x33, 0x9d, 0xee, 0xcc, 0xeb, 0x52,
	0x41, 0xfc, 0x00, 0x76, 0xef, 0x5e, 0xeb, 0x77, 0xfc, 0x7a, 0x1f, 0xf5, 0xa3, 0x07, 0x11, 0x2a,
	0x61, 0xf3, 0x19, 0xd5, 0x49, 0xf6, 0x3f, 0x5e, 0xdc, 0x7b, 0xfb, 0xf7, 0xbf, 0xbf, 0x36, 0xf6,
	0x7a, 0xed, 0xc3, 0xe9, 0x70, 0xf6, 0x77, 0xe2, 0xa8, 0x1a, 0x8f, 0xa3, 0xe8, 0xfe, 0x8c, 0x50,
	0xa5, 0xfb, 0x62, 0xde, 0xc5, 0x17, 0xb5, 0x98, 0xf7, 0x9d, 0x86, 0x39, 0x42, 0xa3, 0x35, 0xab,
	0xf8, 0xc3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0x9c, 0xbb, 0xf0, 0xfa, 0x08, 0x00, 0x00,
}