#!/bin/bash

rm -rf ./mocks

mockgen -source=./src/core/domains/planet/repositories/planet_repository.go -destination=./mocks/planet_repository_mock.go -package=mocks
mockgen -source=./src/shared/providers/logger/logger_provider.go -destination=./mocks/logger_provider_mock.go -package=mocks
