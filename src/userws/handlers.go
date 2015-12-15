package main

import (
   "encoding/json"
   "net/http"
   "github.com/gorilla/mux"
)

func UserShow( w http.ResponseWriter, r *http.Request ) {
   vars := mux.Vars(r)
   userId := vars["userId"]

   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   user, err := LookupUser( userId )

   if err != nil {
      w.WriteHeader( http.StatusInternalServerError )
      return
   }
   
   if user.UserId == userId {
      w.WriteHeader( http.StatusOK )
      if err := json.NewEncoder( w ).Encode( Response{ User: user, Status: http.StatusOK, Message: http.StatusText( http.StatusOK ) } ); err != nil {
         panic( err )
      }
      return
   }

   // If we didn't find it, 404
   w.WriteHeader( http.StatusNotFound )
   if err := json.NewEncoder(w).Encode( Response{ Status: http.StatusNotFound, Message: http.StatusText( http.StatusNotFound ) } ); err != nil {
      panic(err)
   }

}

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

	healthy := true
	message := ""

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader( http.StatusOK )

	user, err := LookupUser( config.HealthCheckUser )
	if err != nil || user.UserId != config.HealthCheckUser {
		healthy = false
		message = err.Error( )
	}

	if err := json.NewEncoder(w).Encode( HealthCheckResponse { CheckType: HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
		panic(err)
	}
}