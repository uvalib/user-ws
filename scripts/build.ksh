export GOPATH=$(pwd)

#env GOOS=linux go build -o bin/user-ws userws
env GOOS=darwin go build -o bin/user-ws userws
