start:
	make stop
	docker compose build
	docker compose up -d
	
yolo:
ifndef ADDR
	@echo "address is missing default value (0.0.0.0) was set"
endif
	
ifndef PORT
	@echo "port is missing default value (8000) was set"
endif
	make stop
	docker compose build
	SERVER_ADDRESS=$(ADDR) SERVER_PORT=$(PORT) docker-compose up -d



stop:
	docker compose stop

rebuild:
ifndef service
	@echo "service parameter is missing"
	@exit 1
endif
	docker compose stop ${service}
	docker compose build ${service}
	docker compose up -d ${service}

generate-mock:
ifndef file
	@echo "file parameter is missing"
	@exit 1
endif
	make test-build
	@docker run --volume "$(PWD)/powerfactors-assignment":/app --workdir /app \
	assessment-test-build /bin/bash -c "mockgen -source=${file} -destination=mocks/${file}"

tests-unit:
	make test-build
	@docker run \
		--rm \
		--volume "$(PWD)/powerfactors-assignment":/app \
		--workdir /app \
		assessment-test-build go test -short -cover -count=1 ./...

tests-all:
	make test-build
	@docker run \
		--rm \
		--volume "$(PWD)/powerfactors-assignment":/app \
		--workdir /app \
		assessment-test-build godotenv -f .env go test ./... -cover -count=1

test-build:
	@docker build \
		--tag assessment-test-build \
		-f powerfactors-assignment/Dockerfile.test ./powerfactors-assignment
