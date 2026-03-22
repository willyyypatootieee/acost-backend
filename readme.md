# Acost Backend

This project uses Gin and Swaggo to auto-generate Swagger docs and serve Swagger UI.

## Swagger Quick Start

1. Install Swagger generator tool (one-time):
	go install github.com/swaggo/swag/cmd/swag@latest

2. Install project dependencies (if needed):
	go mod tidy

3. Generate Swagger docs from annotations:
	go build ./...
	go generate ./...

4. Run the server:
	go run ./cmd/server

5. Open Swagger UI in browser:
	http://localhost:8080/swagger/index.html

## How Auto Generation Works

The server entry file includes a go:generate directive that runs Swag and writes generated files into the docs folder.

Current generation target:
- source entrypoint: cmd/server/main.go
- output folder: docs

Generated files:
- docs/docs.go
- docs/swagger.json
- docs/swagger.yaml

## When To Re-Generate

Run go generate ./... after:
- adding or changing endpoint annotations
- adding or changing request/response structs used in Swagger comments
- changing routes or API base path

## Common Commands

Regenerate docs only:
go run github.com/swaggo/swag/cmd/swag@latest init -g ./cmd/server/main.go -o ./docs --parseInternal

Run app from root:
go run ./cmd/server

Run app from cmd/server folder:
go run .\main.go

## Troubleshooting

1. Error about unknown fields in swag.Spec
- Update dependency:
  go get github.com/swaggo/swag@latest
- Then run:
  go mod tidy
  go generate ./...

2. Swagger page opens but no routes appear
- Make sure handler functions have Swagger annotations.
- Re-run go generate ./....

3. Build fails with missing acost/docs package
- Generate docs first with go generate ./....

## Notes

- Swagger UI route is exposed at /swagger/*any.
- API BasePath is /api/v1.
