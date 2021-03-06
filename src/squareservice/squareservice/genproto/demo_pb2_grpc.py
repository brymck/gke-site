# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import demo_pb2 as demo__pb2


class HelloServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetGreeting = channel.unary_unary(
        '/gkesite.HelloService/GetGreeting',
        request_serializer=demo__pb2.GreetingRequest.SerializeToString,
        response_deserializer=demo__pb2.Greeting.FromString,
        )


class HelloServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetGreeting(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_HelloServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetGreeting': grpc.unary_unary_rpc_method_handler(
          servicer.GetGreeting,
          request_deserializer=demo__pb2.GreetingRequest.FromString,
          response_serializer=demo__pb2.Greeting.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'gkesite.HelloService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class SquareServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetSquare = channel.unary_unary(
        '/gkesite.SquareService/GetSquare',
        request_serializer=demo__pb2.SquareRequest.SerializeToString,
        response_deserializer=demo__pb2.SquareResponse.FromString,
        )


class SquareServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetSquare(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SquareServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetSquare': grpc.unary_unary_rpc_method_handler(
          servicer.GetSquare,
          request_deserializer=demo__pb2.SquareRequest.FromString,
          response_serializer=demo__pb2.SquareResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'gkesite.SquareService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class CountServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetCount = channel.unary_unary(
        '/gkesite.CountService/GetCount',
        request_serializer=demo__pb2.CountRequest.SerializeToString,
        response_deserializer=demo__pb2.CountResponse.FromString,
        )


class CountServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetCount(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CountServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetCount': grpc.unary_unary_rpc_method_handler(
          servicer.GetCount,
          request_deserializer=demo__pb2.CountRequest.FromString,
          response_serializer=demo__pb2.CountResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'gkesite.CountService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
