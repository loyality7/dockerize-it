package generator

import (
	"github.com/yourusername/dockerize-it/internal/detector"
)

func GenerateDockerfile(stack string, structure detector.ProjectStructure) string {
	switch stack {
	case "node":
		return `FROM node:14
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "start"]
`
	case "go":
		return `FROM golang:1.16
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
`
	default:
		return "# Unable to generate Dockerfile for unknown stack"
	}
}
