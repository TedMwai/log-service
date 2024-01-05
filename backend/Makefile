include .env
export
COMMIT := $(shell git log -1 --format="%H")
DOCKER_IMAGE := "log-management:latest"

up:
	docker compose up -d
down:
	docker compose down
run:
	go run ./cmd/server/.

# go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2 
# make migration name=initial
# use underscores in the name instead of spaces
migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down 1

drop:
	migrate -database ${POSTGRESQL_URL} -path migrations drop

# go-build
guild:
	go build -ldflags "-X main.commit=${COMMIT}" -o bin/server cmd/server/main.go

start:
	./server

build:
	docker build -t ${DOCKER_IMAGE} --network host --build-arg COMMIT .

test:
	go test -cover ./... -count=2