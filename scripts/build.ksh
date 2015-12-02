export GOPATH=$(pwd)

#env GOOS=linux go build -o bin/ldap.linux
env GOOS=darwin go build -o bin/ldap.osx
