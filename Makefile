dev: 
	docker compose up
build-ui:
	docker compose run vue-ui npm run build
build-server:
	cd ./software/dccpi/cmd && go build ./main.go
	sudo chown root ./software/dccpi/cmd/main
	sudo chmod +s ./software/dccpi/cmd/main
bash:
	docker compose run vue-ui bash

lint-go:
	@golangci-lint run --fix --config .golangci.yml

