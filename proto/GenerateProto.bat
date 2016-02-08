set protopath=..\..\..\google\protobuf\src
.\protoc.exe %protopath%\google\protobuf\descriptor.proto --plugin=protoc-gen-go=..\..\..\..\..\bin\protoc-gen-go.exe --go_out . --proto_path=%protopath%
copy .\google\protobuf\descriptor.pb.go *.go

mkdir compiler
.\protoc.exe %protopath%\google\protobuf\compiler\plugin.proto --plugin=protoc-gen-go=..\..\..\..\..\bin\protoc-gen-go.exe --go_out . --proto_path=%protopath%
copy .\google\protobuf\compiler\plugin.pb.go .\compiler\*.go

rmdir /s/q google