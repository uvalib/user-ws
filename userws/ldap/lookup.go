package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/nmcclain/ldap"
	"github.com/uvalib/user-ws/userws/config"
	"github.com/uvalib/user-ws/userws/api"
	"github.com/uvalib/user-ws/userws/logger"
	"regexp"
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
	"uvRestrict",
}

//
// openConnection -- open the connection to the LDAP server
//
func openConnection( ) ( *ldap.Conn, error ){

	// are we using TLS for our connection
	if config.Configuration.LdapUseTls == true {
		tlsConf := &tls.Config{
			InsecureSkipVerify: config.Configuration.LdapSkipTlsVerify,
		}
		connection, err := ldap.DialTLS("tcp", config.Configuration.LdapEndpoint, tlsConf )
		return connection, err
	} else {
		connection, err := ldap.DialTimeout("tcp", config.Configuration.LdapEndpoint,
			time.Second * time.Duration( config.Configuration.ServiceTimeout ))
		return connection, err
	}
}

//
// LookupUser -- the user lookup handler
//
func LookupUser( userID string) (*api.User, error) {

	start := time.Now()

	connection, err := openConnection( )
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		return nil, err
	}

	defer connection.Close()

    // if we have credentials then attempt to use them
	if len( config.Configuration.LdapBindAccount ) != 0 && len( config.Configuration.LdapBindPassword ) != 0 {
		err := connection.Bind( config.Configuration.LdapBindAccount, config.Configuration.LdapBindPassword )
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
		//logger.Log(fmt.Sprintf( "RES: %#v", sr.Entries[0].Attributes ))
		//sr.PrettyPrint(0)

		// a special case
		private := "false"
		if len( sr.Entries[0].GetAttributeValue(attributes[10]) ) != 0 {
		   private = "true"
		}
		return &api.User{
			UserID:      userID,
			DisplayName: sr.Entries[0].GetAttributeValue(attributes[0]),
			FirstName:   sr.Entries[0].GetAttributeValue(attributes[1]),
			Initials:    sr.Entries[0].GetAttributeValue(attributes[2]),
			LastName:    sr.Entries[0].GetAttributeValue(attributes[3]),
			Description: stripLDAPPrefix( sr.Entries[0].GetAttributeValue(attributes[4])),
			Department:  stripLDAPPrefix( sr.Entries[0].GetAttributeValue(attributes[5])),
			Title:       stripLDAPPrefix( sr.Entries[0].GetAttributeValue(attributes[6])),
			Office:      stripLDAPPrefix( sr.Entries[0].GetAttributeValue(attributes[7])),
			Phone:       stripLDAPPrefix( sr.Entries[0].GetAttributeValue(attributes[8])),
			Email:       sr.Entries[0].GetAttributeValue(attributes[9]),
			Private:     private,
		}, nil
	}

	logger.Log(fmt.Sprintf("Lookup %s NOT FOUND, time %s", userID, time.Since(start)))

	// return empty user if not found
	return nil, nil
}

//
// stripLDAPPrefix -- Strip the prefix that * might * appear at the start of the field
//
func stripLDAPPrefix( field string ) string {

	r := regexp.MustCompile("^[EUSW]\\d:")
	return r.ReplaceAllString( field, "" )
}
//
// end of file
//
