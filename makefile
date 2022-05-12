generate-pb:
	 protoc \
		 --plugin=$(GOPATH)/bin/protoc-gen-go \
		 -I=./api \
		 --go_out=./pkg/pb ./api/api.proto; \
	 sed -i "" -e "s/,omitempty//g" ./pkg/pb/gen/api/*.go;

run:
	air