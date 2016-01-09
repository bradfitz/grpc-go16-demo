// Code generated by protoc-gen-go.
// source: github.com/bradfitz/grpc-go16-demo/proto/cluster_service_proto/bigtable_cluster_service.proto
// DO NOT EDIT!

package google_bigtable_admin_cluster_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_cluster_v11 "github.com/bradfitz/grpc-go16-demo/proto/cluster_data_proto"
import google_protobuf "github.com/bradfitz/grpc-go16-demo/proto/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for BigtableClusterService service

type BigtableClusterServiceClient interface {
	// Lists the supported zones for the given project.
	ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error)
	// Gets information about a particular cluster.
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Lists all clusters in the given project, along with any zones for which
	// cluster information could not be retrieved.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates a cluster and begins preparing it to begin serving. The returned
	// cluster embeds as its "current_operation" a long-running operation which
	// can be used to track the progress of turning up the new cluster.
	// Immediately upon completion of this request:
	//  * The cluster will be readable via the API, with all requested attributes
	//    but no allocated resources.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will render the cluster immediately unreadable
	//    via the API.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// Upon completion of the embedded operation:
	//  * Billing for all successfully-allocated resources will begin (some types
	//    may have lower than the requested levels).
	//  * New tables can be created in the cluster.
	//  * The cluster's allocated resource levels will be readable via the API.
	// The embedded operation's "metadata" field type is
	// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Updates a cluster, and begins allocating or releasing resources as
	// requested. The returned cluster embeds as its "current_operation" a
	// long-running operation which can be used to track the progress of updating
	// the cluster.
	// Immediately upon completion of this request:
	//  * For resource types where a decrease in the cluster's allocation has been
	//    requested, billing will be based on the newly-requested level.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will set its metadata's "cancelled_at_time",
	//    and begin restoring resources to their pre-request values. The operation
	//    is guaranteed to succeed at undoing all resource changes, after which
	//    point it will terminate with a CANCELLED status.
	//  * All other attempts to modify or delete the cluster will be rejected.
	//  * Reading the cluster via the API will continue to give the pre-request
	//    resource levels.
	// Upon completion of the embedded operation:
	//  * Billing will begin for all successfully-allocated resources (some types
	//    may have lower than the requested levels).
	//  * All newly-reserved resources will be available for serving the cluster's
	//    tables.
	//  * The cluster's new resource levels will be readable via the API.
	// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UpdateCluster(ctx context.Context, in *google_bigtable_admin_cluster_v11.Cluster, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Marks a cluster and all of its tables for permanent deletion in 7 days.
	// Immediately upon completion of the request:
	//  * Billing will cease for all of the cluster's reserved resources.
	//  * The cluster's "delete_time" field will be set 7 days in the future.
	// Soon afterward:
	//  * All tables within the cluster will become unavailable.
	// Prior to the cluster's "delete_time":
	//  * The cluster can be recovered with a call to UndeleteCluster.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// At the cluster's "delete_time":
	//  * The cluster and *all of its tables* will immediately and irrevocably
	//    disappear from the API, and their data will be permanently deleted.
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type bigtableClusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableClusterServiceClient(cc *grpc.ClientConn) BigtableClusterServiceClient {
	return &bigtableClusterServiceClient{cc}
}

func (c *bigtableClusterServiceClient) ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error) {
	out := new(ListZonesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListZones", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error) {
	out := new(google_bigtable_admin_cluster_v11.Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/GetCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error) {
	out := new(google_bigtable_admin_cluster_v11.Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) UpdateCluster(ctx context.Context, in *google_bigtable_admin_cluster_v11.Cluster, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error) {
	out := new(google_bigtable_admin_cluster_v11.Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/DeleteCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableClusterService service

type BigtableClusterServiceServer interface {
	// Lists the supported zones for the given project.
	ListZones(context.Context, *ListZonesRequest) (*ListZonesResponse, error)
	// Gets information about a particular cluster.
	GetCluster(context.Context, *GetClusterRequest) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Lists all clusters in the given project, along with any zones for which
	// cluster information could not be retrieved.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates a cluster and begins preparing it to begin serving. The returned
	// cluster embeds as its "current_operation" a long-running operation which
	// can be used to track the progress of turning up the new cluster.
	// Immediately upon completion of this request:
	//  * The cluster will be readable via the API, with all requested attributes
	//    but no allocated resources.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will render the cluster immediately unreadable
	//    via the API.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// Upon completion of the embedded operation:
	//  * Billing for all successfully-allocated resources will begin (some types
	//    may have lower than the requested levels).
	//  * New tables can be created in the cluster.
	//  * The cluster's allocated resource levels will be readable via the API.
	// The embedded operation's "metadata" field type is
	// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	CreateCluster(context.Context, *CreateClusterRequest) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Updates a cluster, and begins allocating or releasing resources as
	// requested. The returned cluster embeds as its "current_operation" a
	// long-running operation which can be used to track the progress of updating
	// the cluster.
	// Immediately upon completion of this request:
	//  * For resource types where a decrease in the cluster's allocation has been
	//    requested, billing will be based on the newly-requested level.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will set its metadata's "cancelled_at_time",
	//    and begin restoring resources to their pre-request values. The operation
	//    is guaranteed to succeed at undoing all resource changes, after which
	//    point it will terminate with a CANCELLED status.
	//  * All other attempts to modify or delete the cluster will be rejected.
	//  * Reading the cluster via the API will continue to give the pre-request
	//    resource levels.
	// Upon completion of the embedded operation:
	//  * Billing will begin for all successfully-allocated resources (some types
	//    may have lower than the requested levels).
	//  * All newly-reserved resources will be available for serving the cluster's
	//    tables.
	//  * The cluster's new resource levels will be readable via the API.
	// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UpdateCluster(context.Context, *google_bigtable_admin_cluster_v11.Cluster) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Marks a cluster and all of its tables for permanent deletion in 7 days.
	// Immediately upon completion of the request:
	//  * Billing will cease for all of the cluster's reserved resources.
	//  * The cluster's "delete_time" field will be set 7 days in the future.
	// Soon afterward:
	//  * All tables within the cluster will become unavailable.
	// Prior to the cluster's "delete_time":
	//  * The cluster can be recovered with a call to UndeleteCluster.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// At the cluster's "delete_time":
	//  * The cluster and *all of its tables* will immediately and irrevocably
	//    disappear from the API, and their data will be permanently deleted.
	DeleteCluster(context.Context, *DeleteClusterRequest) (*google_protobuf.Empty, error)
}

func RegisterBigtableClusterServiceServer(s *grpc.Server, srv BigtableClusterServiceServer) {
	s.RegisterService(&_BigtableClusterService_serviceDesc, srv)
}

func _BigtableClusterService_ListZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).ListZones(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).GetCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).ListClusters(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).CreateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_bigtable_admin_cluster_v11.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).UpdateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).DeleteCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BigtableClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.cluster.v1.BigtableClusterService",
	HandlerType: (*BigtableClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListZones",
			Handler:    _BigtableClusterService_ListZones_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _BigtableClusterService_GetCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _BigtableClusterService_ListClusters_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _BigtableClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _BigtableClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _BigtableClusterService_DeleteCluster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
