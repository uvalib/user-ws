package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/nmcclain/ldap"
	"github.com/uvalib/user-ws/userws/api"
	"github.com/uvalib/user-ws/userws/config"
	"github.com/uvalib/user-ws/userws/logger"
	"regexp"
	"strconv"
	"time"
)

var attributes = []string{
	"displayName",
	"givenName",
	"initials",
	"sn",
	"description",                // this is a multi-field
	"uvaDisplayDepartment",       // this is a multi-field
	"title",                      // this is a multi-field
	"physicalDeliveryOfficeName", // this is a multi-field
	"telephoneNumber",            // this is a multi-field
	"mail",
	"uvRestrict",
}

//
// openConnection -- open the connection to the LDAP server
//
func openConnection() (*ldap.Conn, error) {

	// are we using TLS for our connection
	if config.Configuration.LdapUseTls == true {
		tlsConf := &tls.Config{
			InsecureSkipVerify: config.Configuration.LdapSkipTlsVerify,
		}
		connection, err := ldap.DialTLS("tcp", config.Configuration.LdapEndpoint, tlsConf)
		return connection, err
	} else {
		connection, err := ldap.DialTimeout("tcp", config.Configuration.LdapEndpoint,
			time.Second*time.Duration(config.Configuration.ServiceTimeout))
		return connection, err
	}
}

//
// LookupUser -- the user lookup handler
//
func LookupUser(userID string) (*api.User, error) {

	start := time.Now()

	connection, err := openConnection()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return nil, err
	}

	defer connection.Close()

	// if we have credentials then attempt to use them
	if len(config.Configuration.LdapBindAccount) != 0 && len(config.Configuration.LdapBindPassword) != 0 {
		err := connection.Bind(config.Configuration.LdapBindAccount, config.Configuration.LdapBindPassword)
		if err != nil {
			logger.Log(fmt.Sprintf("ERROR: Cannot bind: %s\n", err.Error()))
			return nil, err
		}
	}

	search := ldap.NewSearchRequest(
		config.Configuration.LdapBaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(userID=%s)", userID),
		attributes,
		nil)

	sr, err := connection.Search(search)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return nil, err
	}

	if len(sr.Entries) == 1 {
		logger.Log(fmt.Sprintf("Lookup %s OK, time %s", userID, time.Since(start)))

		if config.Configuration.Debug == true {
			logger.Log(fmt.Sprintf("RES: %#v", sr.Entries[0].Attributes))
			sr.PrettyPrint(0)
		}

		// a special case
		private := "false"
		if len(sr.Entries[0].GetAttributeValue(attributes[10])) != 0 {
			private = "true"
		}
		return &api.User{
			UserID:      userID,
			DisplayName: sr.Entries[0].GetAttributeValue(attributes[0]),
			FirstName:   sr.Entries[0].GetAttributeValue(attributes[1]),
			Initials:    sr.Entries[0].GetAttributeValue(attributes[2]),
			LastName:    sr.Entries[0].GetAttributeValue(attributes[3]),
			Description: makeOrderedField(sr.Entries[0].GetAttributeValues(attributes[4])),
			Department:  makeOrderedField(sr.Entries[0].GetAttributeValues(attributes[5])),
			Title:       makeOrderedField(sr.Entries[0].GetAttributeValues(attributes[6])),
			Office:      makeOrderedField(sr.Entries[0].GetAttributeValues(attributes[7])),
			Phone:       makeOrderedField(sr.Entries[0].GetAttributeValues(attributes[8])),
			Email:       sr.Entries[0].GetAttributeValue(attributes[9]),
			Private:     private,
		}, nil
	}

	logger.Log(fmt.Sprintf("Lookup %s NOT FOUND, time %s", userID, time.Since(start)))

	// return empty user if not found
	return nil, nil
}

//
// makeOrderedField -- Convert the multi-field into an ordered array stripping out the cruft
//
func makeOrderedField(fields []string) []string {

	mf := make([]string, len(fields))
	rx := regexp.MustCompile("^[EUSW]\\d:")

	for f_index, field := range fields {

		match, err := regexp.MatchString("^[EUSW]\\d:", field)
		if err == nil {

			// if we match then we process as a set of ordered fields
			if match == true {

				ix, err := strconv.Atoi(string(field[1]))
				if err == nil {
					mf[ix] = rx.ReplaceAllString(field, "")
				}
			} else {
				mf[f_index] = field
			}
		}
	}

	return mf
}

//
// end of file
//
