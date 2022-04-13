generate-pb:
	 protoc \
		 --plugin=$(GOPATH)/bin/protoc-gen-go \
		 -I=./api \
		 --go_out=./pkg/pb ./api/api.proto

run:
	air