db-up:
	docker-compose -f docker-compose.yaml up -d

migrate:
	go run cmd/cli/main.go db migrate

mock:
	go install github.com/golang/mock/mockgen@v1.6.0
	go mod download github.com/golang/mock
	mockgen -destination=./mocks/repo/mock_repo.go -package=repoMocks github.com/gitscan/internal/service/repo Interface
	mockgen -destination=./mocks/db/mock_db.go -package=dbMocks github.com/gitscan/internal/database DB
	mockgen -destination=./mocks/db/mock_db_info.go -package=dbMocks github.com/gitscan/internal/database InfoInterface
	mockgen -destination=./mocks/db/mock_db_finding.go -package=dbMocks github.com/gitscan/internal/database FindingInterface
	mockgen -destination=./mocks/db/mock_db_location.go -package=dbMocks github.com/gitscan/internal/database LocationInterface
	mockgen -destination=./mocks/rule/mock_rule.go -package=ruleMocks github.com/gitscan/rules Interface

test_and_cover:
	go test -tags integration_test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func=coverage.txt