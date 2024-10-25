package generator

import (
	"github.com/loyality7/dockerize-it/internal/detector"
)

func GenerateDockerCompose(stack string, structure detector.ProjectStructure) string {
	switch stack {
	case "node":
		return `version: '3'
services:
  app:
    build: .
    ports:
      - "3000:3000"
    volumes:
      - .:/app
    environment:
      - NODE_ENV=development
`
	case "go":
		return `version: '3'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - .:/app
`
	default:
		return "# Unable to generate docker-compose.yml for unknown stack"
	}
}
