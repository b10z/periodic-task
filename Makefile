#APPLICATION OPTIONS--------------------------------------------------------------------------
app.start:
ifndef ADDR
	@echo "address is missing default value (0.0.0.0) was set"
endif
ifndef PORT
	@echo "port is missing default value (8000) was set"
endif
	make app.stop
	docker compose build
	SERVER_ADDRESS=$(ADDR) SERVER_PORT=$(PORT) docker-compose up -d

app.stop:
	docker compose stop

#TESTS OPTIONS--------------------------------------------------------------------------------
tests.generate-mock:
ifndef file
	@echo "file parameter is missing"
	@exit 1
endif
	make tests.test-build
	@docker run --volume "$(PWD)/powerfactors-assignment":/app --workdir /app \
	assessment-test-build /bin/bash -c "mockgen -source=${file} -destination=mocks/${file}"


tests.tests-all:
	make tests.test-build
	@docker run \
		--rm \
		--volume "$(PWD)/powerfactors-assignment":/app \
		--workdir /app \
		assessment-test-build godotenv -f .env go test ./... -cover -count=1

tests.test-build:
	@docker build \
		--tag assessment-test-build \
		-f powerfactors-assignment/deploy/test/Dockerfile.test ./powerfactors-assignment
