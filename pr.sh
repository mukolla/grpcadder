protoc --go_out=./pkg --go-grpc_out=./pkg api/proto/adder.proto


#evans api/proto/adder.proto -p 8081


#protoc --go_out=paths=source_relative:. api/proto/adder.proto
#protoc --go_out=./pkg api/proto/adder.proto
#protoc -I api/proto --go-grpc_out=pkg/api api/proto/adder.proto
#protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
#protoc -I api/proto --go-grpc_out=generated --go_out=plugins=grpc:pkg/api api/proto/adder.proto
#protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto
#protoc -I /usr/local/include -I $PWD/api/dummy-proto --go_out=generated --go-grpc_out=generated --go_opt=paths=source_relative proto/v1/foo.proto
#protoc --go_out=. --go_opt=paths=source_relative \
#    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#    api/proto/adder.proto
#protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
#protoc --go_out=paths=./pkg/api api/proto/adder.proto
#
#sudo ln -s /home/mukolla/go/bin/protoc-gen-go-grpc /usr/bin/protoc-gen-go-grpc
#
#
#docker run --rm -v "$(pwd):/mount:ro" \
#    ghcr.io/ktr0731/evans:latest \
#      --path ./api/proto \
#      --proto adder.proto \
#      --host http://localhost \
#      --port 8081 \
#      repl
