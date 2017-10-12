package handlers

import (
   "github.com/gorilla/mux"
   "net/http"
   "userws/authtoken"
   "userws/config"
   "userws/ldap"
)

//
// UserLookup -- do the user lookup
//
func UserLookup(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   userID := vars["userId"]
   token := r.URL.Query().Get("auth")

   // parameters OK ?
   if parametersOk(userID, token) == false {
      encodeStandardResponse(w, http.StatusBadRequest, nil)
      return
   }

   // validate the token
   if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
      encodeStandardResponse(w, http.StatusForbidden, nil)
      return
   }

   // do the lookup
   user, err := ldap.LookupUser(config.Configuration.EndpointURL,
      config.Configuration.Timeout,
      config.Configuration.LdapBaseDn,
      userID)

   // lookup error?
   if err != nil {
      encodeStandardResponse(w, http.StatusInternalServerError, nil)
      return
   }

   // user not found (probably an error)?
   if user == nil {
      encodeStandardResponse(w, http.StatusNotFound, nil)
      return
   }

   // all good...
   encodeStandardResponse(w, http.StatusOK, user)
}

//
// end of file
//

