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


protoc \
--proto_path=deps \
--proto_path=deps/include \
--proto_path=deps/github.com/gogo/protobuf/protobuf \
--proto_path=. \
--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
validator.proto


protoc \
-I. \
-I $GOPATH/src/ \
-I $GOPATH/pkg/mod/github.com/google/protobufv3.18.0+incompatible/src/ \
--proto_path=. \
--gogo_out=${GOPATH}/src \
--gogo_opt=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor \
secvalidator.proto

