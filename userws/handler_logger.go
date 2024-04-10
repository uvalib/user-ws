package main

import (
	"fmt"
	"github.com/uvalib/user-ws/userws/logger"
	"net/http"
	"regexp"
	"time"
)

// HandlerLogger -- middleware handler
func HandlerLogger(inner http.Handler, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		matchExpr := fmt.Sprintf("\\?auth=.*$")
		re := regexp.MustCompile(matchExpr)
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Log(fmt.Sprintf(
			"%s %s (%s) -> method %s, time %s",
			r.Method,
			re.ReplaceAllString(r.RequestURI, "?auth=<secret>"),
			r.RemoteAddr,
			name,
			time.Since(start),
		))
	})
}

//
// end of file
//
