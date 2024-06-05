
# Create a new migration
create:
	migrate create -ext sql -dir migrations -seq create_sql_tables

mig-up:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/public?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/public?sslmode=disable' -path migrations down

gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. party.proto
	protoc --go_out=. --go-grpc_out=. public.proto
	
gen-protoAll:
#proto fileni hammasini bittada generatsiya qilish 
	protoc --go_out=./ \
	--go-grpc_out=./ \
	protos/*.proto