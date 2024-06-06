
# Create a new migration
create:
	migrate create -ext sql -dir migrations -seq create_sql_tables


mig-up:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/vote?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/vote?sslmode=disable' -path migrations down


# gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
# 	protoc --go_out=. --go-grpc_out=. candidate.proto
# 	protoc --go_out=. --go-grpc_out=. election.proto
# 	protoc --go_out=. --go-grpc_out=. public_vote.proto
# 	protoc --go_out=. --go-grpc_out=. vote.proto


gen-proto:
# proto fileni bittalab ruchnoy generatsiya qilish 
  protoc --go_out=services/ \
       --go-grpc_out=services/ \
       services/protos/cart.proto


gen-protoAll:
#proto fileni hammasini bittada generatsiya qilish 
  protoc --go_out=./ \
    --go-grpc_out=./ \
    protos/*.proto



gen-proto_curr:
# bu proto fayllarimiz turgan joyga generatsiya qilib beradi bittalab
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto


CURRENT_DIR=$(shell pwd)


proto-gen:
#script file orqali proto fileni hammasini bittada generatsiya qilish 
	./scripts/gen-proto.sh "${CURRENT_DIR}"