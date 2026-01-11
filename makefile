run:
	@echo running app...
	go run cmd/api/main.go

sync:
	@echo syncinc...
	go mod tidy
	make run
