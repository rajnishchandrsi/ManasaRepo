# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.

mkfile_dir = "$(dir $(abspath $(lastword $(MAKEFILE_LIST))))"

ifdef GOBIN
extra_path = "$(mkfile_dir)deps/bin:$(GOBIN)"
else
extra_path = "$(mkfile_dir)deps/bin:$(HOME)/go/bin"
endif

prepare_deps:
	@echo "--- Preparing dependencies."
	@bash scripts/prepare-deps.sh

regenerate: prepare_deps
	@echo "--- Regenerating validator.proto"
	export PATH=$(extra_path):$${PATH}; protoc \
		--proto_path=deps \
		--proto_path=deps/include \
		--proto_path=deps/github.com/gogo/protobuf/protobuf \
		--proto_path=. \
		--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
		validator.proto

clean:
	rm -rf "deps"