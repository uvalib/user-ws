package main

import (
   "encoding/json"
   //"fmt"
   //"io"
   //"io/ioutil"
   "net/http"
   //"strconv"
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
      if err := json.NewEncoder( w ).Encode( user ); err != nil {
         panic( err )
      }
	   return
   }

   // If we didn't find it, 404
   w.WriteHeader( http.StatusNotFound )
   //if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
   //   panic(err)
   //}

}
