# copy proto from main folder
cp -r src/proto/* src/dotnet/ConsoleValidator/proto/

# switch to go directory
cd src/dotnet/ConsoleValidator

# generate artefacts
protoc \
  -I ./protovalidate/proto/protovalidate \
  --csharp_out=../Models \
  --proto_path=. \
  ./user.proto

#protoc -I=./proto --python_out=./gen --grpc_python_out=./gen ./proto/*.proto

echo 'Generated Go files'
ls ../Models