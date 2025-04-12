FROM golang:1.24.1-alpine

WORKDIR /app

# Install PostgreSQL client
RUN apk add --no-cache postgresql-client

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Create entrypoint script
RUN echo '#!/bin/sh' > /app/entrypoint.sh && \
    echo 'until pg_isready -h gondola.proxy.rlwy.net -p 25045 -U postgres -d railway; do' >> /app/entrypoint.sh && \
    echo '  echo "Waiting for PostgreSQL..."' >> /app/entrypoint.sh && \
    echo '  sleep 2' >> /app/entrypoint.sh && \
    echo 'done' >> /app/entrypoint.sh && \
    echo 'echo "PostgreSQL is ready!"' >> /app/entrypoint.sh && \
    echo 'exec "$@"' >> /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["./main"]