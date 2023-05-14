#!/usr/bin/env bash

FuzzFUNC="Fuzz" #"FuzzReverse"
CGO_ENABLED=1
export CGO_ENABLED

if ! command -v golangci-lint  &> /dev/null
then
  echo "golangci-lint required but it's not installed. Skipping."
else
  echo "Let's Lint, first.."
  golangci-lint run #./...
fi

echo "Let's Test"
go test -v ./... -coverprofile=coverage.out

echo "Let's Test (race detector)"
go test -race ./...
# https://go.dev/doc/articles/race_detector

echo "Let's Fuzz" #cannot use -fuzz flag with multiple packages
go test ./internal -fuzz ${FuzzFUNC} -fuzztime 15s

echo "Let's Bench"
go test -v ./... -run=^$ -bench . -benchmem -benchtime=3s ./

echo "Finally, the security..."
if ! command -v gosec &> /dev/null
then
  echo "gosec required but it's not installed. Skipping."
  exit
else
  echo "Let's Gosec (Scan code vulnerabilities)"
  gosec ./...
  #gosec -no-fail -fmt sarif -out results.sarif ./...
fi

if ! command -v govulncheck &> /dev/null
then # https://go.dev/blog/vuln
  echo "govulncheck required but it's not installed. Skipping."
  echo "Can be install: go install golang.org/x/vuln/cmd/govulncheck@latest"
  exit
else
  echo "Let's Govulncheck (Scan code/pkg for known vulnerabilities)"
  govulncheck ./...
fi