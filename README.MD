protoc \
--proto_path=deps \
--proto_path=deps/include \
--proto_path=deps/github.com/gogo/protobuf@v1.3.0/protobuf \
--proto_path=. \
--gogo_out=Mdescriptor.proto=:. \
validator.proto


protoc \
-I. \
-I $GOPATH/src/ \
-I $GOPATH/src/github.com/google/protobuf/src/ \
--go_out=plugins=grpc:./ \
--govalidators_out=./ \
test.proto