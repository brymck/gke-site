#!/bin/bash -eu
package=squareservice
python -c 'import grpc_tools' >/dev/null 2>&1 || pip install grpcio-tools
protodir=../../pb
outdir="$package"/genproto
python -m grpc_tools.protoc -I "$protodir" --python_out="$outdir" --grpc_python_out="$outdir" "$protodir/demo.proto"
find "$outdir" -name '*_pb2*.py' -print0 | xargs -0 -I {} sed -i.bak 's/^\(import.*_pb2\)/from . \1/' {}
find "$outdir" -name '*_pb2*.py.bak' -print0 | xargs -0 rm
