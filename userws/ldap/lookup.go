package ldap

import (
	"fmt"
	"github.com/nmcclain/ldap"
	"github.com/uvalib/user-ws/userws/api"
	"github.com/uvalib/user-ws/userws/logger"
	"time"
)

var attributes = []string{
	"displayName",
	"givenName",
	"initials",
	"sn",
	"description",
	"uvaDisplayDepartment",
	"title",
	"physicalDeliveryOfficeName",
	"telephoneNumber",
	"mail",
}

//
// LookupUser -- the user lookup handler
//
func LookupUser( ldapEndpoint string, timeout int, ldapBindAccount string, ldapBindPasswd string, ldapBaseDn string, userID string) (*api.User, error) {

	start := time.Now()

	l, err := ldap.DialTimeout("tcp", ldapEndpoint, time.Second*time.Duration(timeout))
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return nil, err
	}

	defer l.Close()

    // if we have credentials then attempt to use them
	if len( ldapBindAccount ) != 0 && len( ldapBindPasswd ) != 0 {
		err = l.Bind(ldapBindAccount, ldapBindPasswd)
		if err != nil {
			logger.Log(fmt.Sprintf("ERROR: Cannot bind: %s\n", err.Error()))
			return nil, err
		}
	}

	search := ldap.NewSearchRequest(
		ldapBaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(userID=%s)", userID),
		attributes,
		nil)

	sr, err := l.Search(search)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return nil, err
	}

	if len(sr.Entries) == 1 {
		logger.Log(fmt.Sprintf("Lookup %s OK, time %s", userID, time.Since(start)))
		return &api.User{
			UserID:      userID,
			DisplayName: sr.Entries[0].GetAttributeValue(attributes[0]),
			FirstName:   sr.Entries[0].GetAttributeValue(attributes[1]),
			Initials:    sr.Entries[0].GetAttributeValue(attributes[2]),
			LastName:    sr.Entries[0].GetAttributeValue(attributes[3]),
			Description: sr.Entries[0].GetAttributeValue(attributes[4]),
			Department:  sr.Entries[0].GetAttributeValue(attributes[5]),
			Title:       sr.Entries[0].GetAttributeValue(attributes[6]),
			Office:      sr.Entries[0].GetAttributeValue(attributes[7]),
			Phone:       sr.Entries[0].GetAttributeValue(attributes[8]),
			Email:       sr.Entries[0].GetAttributeValue(attributes[9]),
		}, nil
	}

	logger.Log(fmt.Sprintf("Lookup %s NOT FOUND, time %s", userID, time.Since(start)))

	// return empty user if not found
	return nil, nil
}

//
// end of file
//
