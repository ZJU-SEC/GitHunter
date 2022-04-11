FROM golang:1.18 as builder

# Set Environment Variables
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
COPY . .
RUN go mod download

# Build app
RUN go build -a -installsuffix cgo -o GitHunter .

FROM alpine:latest

WORKDIR /

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/GitHunter .

ENTRYPOINT [ "/GitHunter" ]
