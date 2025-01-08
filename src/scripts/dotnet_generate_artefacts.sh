# copy proto from main folder
cp -r src/proto/* src/golang/lightweight-example/proto/

# switch to go directory
cd src/golang/lightweight-example

# generate artefacts
buf generate

echo 'Generated Go files'
ls gen

# src/dotnet/ConsoleValidator/proto

# copy proto from main folder
cp -r src/proto/* src/dotnet/ConsoleValidator/proto/

# switch to go directory
cd src/dotnet/ConsoleValidator

# generate artefacts
protoc \
  -I ./protovalidate/proto/protovalidate \
  --csharp_out=../Models \
  --proto_path=. \
  ./*.proto

#protoc -I=./proto --python_out=./gen --grpc_python_out=./gen ./proto/*.proto

echo 'Generated Go files'
ls ../Models