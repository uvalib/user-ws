package ldap

import (
	"fmt"
	"github.com/nmcclain/ldap"
	"time"
	"userws/api"
	"userws/logger"
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
func LookupUser(endpoint string, timeout int, baseDn string, userID string) (*api.User, error) {

	start := time.Now()

	l, err := ldap.DialTimeout("tcp", endpoint, time.Second*time.Duration(timeout))
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return nil, err
	}

	defer l.Close()
	// l.Debug = true

	//err = l.Bind(user, passwd)
	//if err != nil {
	//   logger.Log( fmt.Sprintf("ERROR: Cannot bind: %s\n", err.Error() ) )
	//   return
	//}

	search := ldap.NewSearchRequest(
		baseDn,
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
