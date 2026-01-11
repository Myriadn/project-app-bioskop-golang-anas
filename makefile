run:
	@echo running app...
	go run cmd/api/main.go

sync:
	@echo syncinc...
	go mod tidy
	make run

cover-test:
	@echo coverage test...
	go test ./internal/repository/... -cover
	go test ./internal/service/... -cover
