export GOPATH=$(pwd)

env GOOS=linux go build -o bin/user-ws
#env GOOS=darwin go build -o bin/user-ws
