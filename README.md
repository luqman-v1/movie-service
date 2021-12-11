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

# Testing 
- go test -cover -coverprofile=coverage.out ./...
- go tool cover -html=coverage.out