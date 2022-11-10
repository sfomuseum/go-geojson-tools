cli:
	go build -mod vendor -o bin/bounds cmd/bounds/main.go
	go build -mod vendor -o bin/bbox2feature cmd/bbox2feature/main.go
