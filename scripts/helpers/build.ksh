export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  GOOS=darwin go build -o bin/user-ws.darwin userws
  res=$?
fi

if [ $res -eq 0 ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/user-ws.linux userws
  res=$?
fi

exit $res
