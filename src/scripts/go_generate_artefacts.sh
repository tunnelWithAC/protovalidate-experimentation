# copy proto from main folder
cp -r src/proto/* src/golang/lightweight-example/proto

# switch to go directory
cd src/golang/lightweight-example

# generate artefacts
buf generate

echo 'Generated Go files'
ls gen