
# Create a new migration
create:
	migrate create -ext sql -dir migrations -seq create_sql_tables


mig-up:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/vote?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/vote?sslmode=disable' -path migrations down


gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. candidate.proto
	protoc --go_out=. --go-grpc_out=. election.proto
	protoc --go_out=. --go-grpc_out=. public_vote.proto
	protoc --go_out=. --go-grpc_out=. vote.proto