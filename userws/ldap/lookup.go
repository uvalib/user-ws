package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/nmcclain/ldap"
	"github.com/uvalib/user-ws/userws/api"
	"github.com/uvalib/user-ws/userws/config"
	"github.com/uvalib/user-ws/userws/logger"
	"net"
	"regexp"
	"sort"
	"time"
)

var attributes = []string{
	"displayName",
	"givenName",
	"initials",
	"sn",
	"description",                // multi-field
	"uvaDisplayDepartment",       // multi-field
	"title",                      // multi-field
	"physicalDeliveryOfficeName", // multi-field
	"telephoneNumber",            // multi-field
	"mail",
	"uvaPersonUniversityID",
	"uvRestrict",
	"uvaPersonIAMAffiliation", // multi-field
}

//
// openConnection -- open the connection to the LDAP server
//
func openConnection() (*ldap.Conn, error) {

	// are we using TLS for our connection
	if config.Configuration.LdapUseTls == true {
		dialer := &net.Dialer{
			Timeout: time.Second * time.Duration(config.Configuration.ServiceTimeout),
		}
		tlsConf := &tls.Config{
			InsecureSkipVerify: config.Configuration.LdapSkipTlsVerify,
		}
		connection, err := ldap.DialTLSDialer("tcp", config.Configuration.LdapEndpoint, tlsConf, dialer)
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
		err = connection.Bind(config.Configuration.LdapBindAccount, config.Configuration.LdapBindPassword)
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
		logger.Log(fmt.Sprintf("INFO: lookup %s OK, time %s", userID, time.Since(start)))

		if config.Configuration.Debug == true {
			logger.Log(fmt.Sprintf("DEBUG: %#v", sr.Entries[0].Attributes))
			sr.PrettyPrint(0)
		}

		// a special case
		private := "false"
		if len(sr.Entries[0].GetAttributeValue(attributes[11])) != 0 {
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
			Affiliation: makeOrderedField(sr.Entries[0].GetAttributeValues(attributes[12])),
			Email:       sr.Entries[0].GetAttributeValue(attributes[9]),
			UvaID:       sr.Entries[0].GetAttributeValue(attributes[10]),
			Private:     private,
		}, nil
	}

	logger.Log(fmt.Sprintf("WARNING: lookup %s NOT FOUND, time %s", userID, time.Since(start)))

	// return empty user if not found
	return nil, nil
}

//
// makeOrderedField -- Convert the multi-field into an ordered array stripping out the cruft
//
func makeOrderedField(fields []string) []string {

	mf := make([]string, len(fields))
	rx := regexp.MustCompile("^[EUSW]\\d:")

	// if these are to be ordered, they are tagged in a lexically appropriate manner
	sort.Strings(fields)

	// remove the ordering tags
	for f_index, field := range fields {
		mf[f_index] = rx.ReplaceAllString(field, "")
	}

	return mf
}

//
// end of file
//
