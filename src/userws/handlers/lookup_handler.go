package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"userws/authtoken"
	"userws/config"
	"userws/ldap"
)

func UserLookup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if parametersOk(userId, token) == false {
		encodeStandardResponse(w, http.StatusBadRequest, nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
		encodeStandardResponse(w, http.StatusForbidden, nil)
		return
	}

	// do the lookup
	user, err := ldap.LookupUser(config.Configuration.EndpointUrl,
		config.Configuration.Timeout,
		config.Configuration.LdapBaseDn,
		userId)

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
