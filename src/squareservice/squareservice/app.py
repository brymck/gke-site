import os
import time
from concurrent.futures.thread import ThreadPoolExecutor

import grpc

from .genproto.demo_pb2 import SquareRequest, SquareResponse
from .genproto.demo_pb2_grpc import SquareServiceServicer, add_SquareServiceServicer_to_server
from .logger import get_json_logger

logger = get_json_logger('squareservice-server')


class SquareService(SquareServiceServicer):
    def GetSquare(self, request: SquareRequest, context) -> SquareResponse:
        number = request.number
        square = number * number
        response = SquareResponse(number=square)
        return response


def start():
    server = grpc.server(ThreadPoolExecutor(max_workers=10))
    service = SquareService()

    add_SquareServiceServicer_to_server(service, server)

    port = os.environ.get('PORT', '8080')
    logger.info(f'listening on port {port}')
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    try:
        while True:
            time.sleep(3600)
    except KeyboardInterrupt:
        server.stop(0)
