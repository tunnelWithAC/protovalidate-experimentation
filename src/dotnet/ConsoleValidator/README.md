```commandline
protoc -I=./proto \
    --python_out=./gen \
    --grpc_python_out=./gen $(find ./proto -name "*.proto")
```

dotnet sln ../todo.sln add todo-app.csproj 
