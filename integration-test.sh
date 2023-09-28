#!/bin/bash

go test -v ./... 2>&1 | go-junit-report -set-exit-code -package-name github.com/joesantos418/integration-test-example-docker-go > test-results/api/results.xml
