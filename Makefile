run:
	go run cmd/igcgo/main.go

test:
	go test ./...

apitest:
	pytest

auth:
	curl --location 'http://localhost:9487/auth' --header 'Content-Type: application/json' --data-raw '{"id":12345,"email":"yale@gmail.com","password":"12345"}'

watch:
	cp main.go ./cmd/igcgo/main.go
	gowatch

#database
DOCKER_NETWORK="igcgo_igcgo"
psql:
	docker run -it --rm --network=$(DOCKER_NETWORK) postgres psql -h postgres -U postgres

install_migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create_migrate:
	migrate create -ext sql -dir db/migrations -seq create_igcgo_tables

mup:
	migrate -database postgres://postgres:12345@localhost:5432/igcgo?sslmode=disable -path db/migrations up

mdown:
	migrate -database postgres://postgres:12345@localhost:5432/igcgo?sslmode=disable -path db/migrations down

.PHONT: run test apitest auth psql install_migrate
.SILENT: run test apitest auth psql install_migrate