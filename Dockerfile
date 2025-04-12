FROM golang:1.24.1-alpine

WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache gcc musl-dev postgresql-client

# Create migrations directory
RUN mkdir -p /app/migrations

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy migrations
COPY migrations/*.sql /app/migrations/

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Make sure migrations are executable
RUN chmod +x /app/migrations/*.sql

# Expose port 8080
EXPOSE 8080

# Create entrypoint script
RUN echo '#!/bin/sh' > /app/entrypoint.sh && \
    echo 'while ! pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do' >> /app/entrypoint.sh && \
    echo '  echo "Waiting for PostgreSQL..."' >> /app/entrypoint.sh && \
    echo '  sleep 2' >> /app/entrypoint.sh && \
    echo 'done' >> /app/entrypoint.sh && \
    echo 'echo "PostgreSQL is ready!"' >> /app/entrypoint.sh && \
    echo 'exec "$@"' >> /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["./main"]