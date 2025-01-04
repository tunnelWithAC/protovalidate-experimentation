# protovalidate-demo

## Prerequisites

Install `buf` CLI tool by following the [buf installation steps](https://buf.build/docs/installation/).

On Mac, you just need to run `brew install buf`

`buf generate proto`

`go run main.go`

Open your browser or use curl to test the endpoint:
Default: http://localhost:8080/hello
With query param: http://localhost:8080/hello?name=John

----- 

https://buf.build/docs/generate/tutorial/#3-configure-your-bufgenyaml-file

Based on [protovalidate](https://github.com/bufbuild/protovalidate)のデモリポジトリ