.PHONY: mock test

# run: test-api test-service
run:
	./cmds/env .env go run main.go

test:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./... -v -cover
	go mod tidy

test-repo:
	go get -u github.com/kyoh86/richgo
	./cmds/env env-test richgo test -count=1 ./repositories/mysql -v -cover
	go mod tidy

migrate-test-up:
	migrate \
  	-source file://migrations \
  	-database "mysql://test:test@tcp(localhost:3316)/test" up

migrate-test-down:
	migrate \
  	-source file://migrations \
  	-database "mysql://test:test@tcp(localhost:3316)/test" down

migrate-dev-up:
	migrate \
  	-source file://migrations \
  	-database "mysql://user:password@tcp(localhost:3326)/ordering" up

migrate-dev-down:
	migrate \
  	-source file://migrations \
  	-database "mysql://user:password@tcp(localhost:3326)/ordering" down
 
mock:
	@mockery --dir models --all