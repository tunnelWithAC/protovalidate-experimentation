#!/bin/zsh

# copy proto from main folder
# cp ../../proto .
cp -r src/proto/ src/golang/lightweight-example/proto

# switch to go directory
cd src/golang/lightweight-example

# generate artefacts
buf generate
