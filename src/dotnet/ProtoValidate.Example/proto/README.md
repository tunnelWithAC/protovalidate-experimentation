Generate artifacts from proto file

protoc \
  -I ./vendor/protovalidate/proto/protovalidate \
  --csharp_out=./gen  --proto_path=. ./user.proto
