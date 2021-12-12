# Movie-Service

# Structure

```
├── action
├── builder
├── entity
├── repo
│ ├── curl
│ ├── mocks
│ ├── mysql
│ │ └── migrations
│ └── redis
├── transport
│ └── grpc
│     └── proto
│         └── movie
└── util
```

movie-gateway
https://github.com/luqman-v1/movie-gateway

# Testing 
- go test -cover -coverprofile=coverage.out ./...
- go tool cover -html=coverage.out
