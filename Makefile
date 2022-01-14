pwd := $(shell pwd)

migrate:
	migrate create -ext sql -dir  $(pwd)/migrations -seq $(name)

migrate-up:
	docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://pavel:pass@tcp(localhost:3306)/todo" up
migrate-down:
	docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://todo:pass@tcp(localhost:3306)/todo" down
migrate-force:
	docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://todo:pass@tcp(localhost:3306)/todo" force $(version)

start:
	docker-compose up todo
