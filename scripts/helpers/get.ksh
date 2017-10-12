if [ -z "$GOPATH" ]; then
   echo "ERROR: $GOPATH not defined"
   exit 1
fi

cd $GOPATH/src
rm -fr vendor

go get -u github.com/golang/lint/golint
go get -u github.com/FiloSottile/gvt

gvt fetch github.com/nmcclain/ldap
gvt fetch github.com/gorilla/mux
gvt fetch github.com/parnurzeal/gorequest

# for tests
gvt fetch gopkg.in/yaml.v2
