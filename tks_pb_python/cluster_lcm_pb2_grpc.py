# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import cluster_lcm_pb2 as cluster__lcm__pb2
import common_pb2 as common__pb2


class ClusterLcmServiceStub(object):
    """ClusterLcmService is a service to manage cluster's document.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateCluster = channel.unary_unary(
                '/tks_pb.ClusterLcmService/CreateCluster',
                request_serializer=cluster__lcm__pb2.CreateClusterRequest.SerializeToString,
                response_deserializer=common__pb2.IDResponse.FromString,
                )
        self.ScaleCluster = channel.unary_unary(
                '/tks_pb.ClusterLcmService/ScaleCluster',
                request_serializer=cluster__lcm__pb2.ScaleClusterRequest.SerializeToString,
                response_deserializer=common__pb2.SimpleResponse.FromString,
                )
        self.DeleteCluster = channel.unary_unary(
                '/tks_pb.ClusterLcmService/DeleteCluster',
                request_serializer=common__pb2.IDRequest.SerializeToString,
                response_deserializer=common__pb2.SimpleResponse.FromString,
                )
        self.InstallAppGroups = channel.unary_unary(
                '/tks_pb.ClusterLcmService/InstallAppGroups',
                request_serializer=cluster__lcm__pb2.InstallAppGroupsRequest.SerializeToString,
                response_deserializer=common__pb2.IDsResponse.FromString,
                )
        self.UninstallAppGroups = channel.unary_unary(
                '/tks_pb.ClusterLcmService/UninstallAppGroups',
                request_serializer=cluster__lcm__pb2.UninstallAppGroupsRequest.SerializeToString,
                response_deserializer=common__pb2.IDsResponse.FromString,
                )


class ClusterLcmServiceServicer(object):
    """ClusterLcmService is a service to manage cluster's document.
    """

    def CreateCluster(self, request, context):
        """CreateCluster creates a Kubernetes cluster and returns cluster id
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ScaleCluster(self, request, context):
        """ScaleCluster scales Kubernetes cluster
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteCluster(self, request, context):
        """DeleteCluster deletes Kubernetes cluster
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def InstallAppGroups(self, request, context):
        """InstallAppGroups installs app groups, returns an array of app group id
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UninstallAppGroups(self, request, context):
        """UninstallAppGroups uninstalls app groups.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ClusterLcmServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateCluster': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCluster,
                    request_deserializer=cluster__lcm__pb2.CreateClusterRequest.FromString,
                    response_serializer=common__pb2.IDResponse.SerializeToString,
            ),
            'ScaleCluster': grpc.unary_unary_rpc_method_handler(
                    servicer.ScaleCluster,
                    request_deserializer=cluster__lcm__pb2.ScaleClusterRequest.FromString,
                    response_serializer=common__pb2.SimpleResponse.SerializeToString,
            ),
            'DeleteCluster': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteCluster,
                    request_deserializer=common__pb2.IDRequest.FromString,
                    response_serializer=common__pb2.SimpleResponse.SerializeToString,
            ),
            'InstallAppGroups': grpc.unary_unary_rpc_method_handler(
                    servicer.InstallAppGroups,
                    request_deserializer=cluster__lcm__pb2.InstallAppGroupsRequest.FromString,
                    response_serializer=common__pb2.IDsResponse.SerializeToString,
            ),
            'UninstallAppGroups': grpc.unary_unary_rpc_method_handler(
                    servicer.UninstallAppGroups,
                    request_deserializer=cluster__lcm__pb2.UninstallAppGroupsRequest.FromString,
                    response_serializer=common__pb2.IDsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'tks_pb.ClusterLcmService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ClusterLcmService(object):
    """ClusterLcmService is a service to manage cluster's document.
    """

    @staticmethod
    def CreateCluster(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tks_pb.ClusterLcmService/CreateCluster',
            cluster__lcm__pb2.CreateClusterRequest.SerializeToString,
            common__pb2.IDResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ScaleCluster(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tks_pb.ClusterLcmService/ScaleCluster',
            cluster__lcm__pb2.ScaleClusterRequest.SerializeToString,
            common__pb2.SimpleResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteCluster(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tks_pb.ClusterLcmService/DeleteCluster',
            common__pb2.IDRequest.SerializeToString,
            common__pb2.SimpleResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def InstallAppGroups(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tks_pb.ClusterLcmService/InstallAppGroups',
            cluster__lcm__pb2.InstallAppGroupsRequest.SerializeToString,
            common__pb2.IDsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UninstallAppGroups(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/tks_pb.ClusterLcmService/UninstallAppGroups',
            cluster__lcm__pb2.UninstallAppGroupsRequest.SerializeToString,
            common__pb2.IDsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
