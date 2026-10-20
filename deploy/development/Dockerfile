FROM node:24-alpine3.21 AS npm-deps

ENV PROJECT_DIR=/app

WORKDIR $PROJECT_DIR

COPY package*.json ./

RUN npm install

FROM golang:1.26.0 AS development

# Set environment variables for the project directory and Go modules
ENV PROJECT_DIR=/app \
    GOPROXY=https://goproxy.io,direct \
    GO111MODULE=on\
    CGO_ENABLED=0

# Set the working directory to PROJECT_DIR
WORKDIR $PROJECT_DIR

# Copy Go module files to the working directory
COPY go.mod go.sum ./

# Download project dependencies
RUN go mod download

# Copy all project files to the working directory
COPY . .

# Install the Air tool for live reloading during development
RUN go install github.com/air-verse/air@latest

# RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# USER appuser

# Copy the built project files from the builder stage & npm deps stage to the development stage
COPY --from=npm-deps /app/node_modules /app/node_modules

# Expose ports to the outside world
EXPOSE 8000 50000

# Specify the command to run when the container starts (execute Air with the specified config)
CMD ["air", "-c", ".air.toml"]