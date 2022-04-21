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
-I $GOPATH/pkg/mod/github.com/google/protobuf@v3.18.0+incompatible/src/ \
--proto_path=. \
--gogo_out=${GOPATH}/src \
--gogo_opt=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor \
secvalidator.proto

Test case 
protoc \
--proto_path=. \
--proto_path=${GOPATH}/pkg/mod \
--proto_path=${GOPATH}/src \
--proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf \
--gogo_out=. \
--gogo_out=${GOPATH}/src \
--gogo_opt=Mgogoproto/gogo.proto=github.com/gogo/protobuf@v1.3.2 \
--gogo_opt=Mgo-proto-validators/validator.proto=github.com/maanasasubrahmanyam-sd \
--govalidators_out=./ \
validator_proto3.proto

working command proto give complete path of secvalidate.proto
protoc \
--proto_path=. \
--proto_path=${GOPATH}/pkg/mod \
--proto_path=${GOPATH}/src \
--proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf \
--gogo_out=./gogo \
--gogo_out=${GOPATH}/src \
--gogo_opt=Mgogoproto/gogo.proto=github.com/gogo/protobuf@v1.3.2 \
--gosecvalidator_out=./gogo \
validator_proto3.proto


// working command in proto keep secvalidat.proto 
protoc \
--proto_path=. \
--proto_path=C:\\Users\\Rohit\\go\\src\\github.com\\maanasasubrahmanyam-sd\\test\\secvalidator \
--proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf \
-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
--gogo_out=./ \
--gogo_opt=Mgogoproto/gogo.proto=github.com/gogo/protobuf@v1.3.2 \
--gosecvalidator_out=gogoimport=true:./ \
--validate_out="lang=go:./" \
*.proto

$ go install protoc-gen-gosecvalidator.go