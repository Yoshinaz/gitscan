db-up:
	docker-compose -f docker-compose.yaml up -d

migrate:
	go run cmd/cli/main.go db migrate

mock:
	go install github.com/golang/mock/mockgen@v1.6.0
	go mod download github.com/golang/mock
	mockgen -destination=./mocks/mock_repo.go -package=mocks github.com/gitscan/internal/service/repo Interface
	mockgen -destination=./mocks/mock_db.go -package=mocks github.com/gitscan/internal/database DB