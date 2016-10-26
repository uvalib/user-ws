package ldap

import (
   "fmt"
   "time"
   "github.com/nmcclain/ldap"
    "userws/api"
    "userws/logger"
)

var	Attributes []string = []string{"displayName", "givenName", "initials", "sn", "description", "uvaDisplayDepartment", "title", "physicalDeliveryOfficeName", "mail", "telephoneNumber"}

func LookupUser( endpoint string, baseDn string, userId string ) ( * api.User, error ) {

	start := time.Now( )

	l, err := ldap.DialTimeout("tcp", endpoint, time.Second * 5 )
	if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: %s\n", err.Error( ) ) )
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
		fmt.Sprintf( "(userId=%s)", userId ),
		Attributes,
		nil )

	sr, err := l.Search(search)
	if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: %s\n", err.Error( ) ) )
		return nil, err
	}

	if len( sr.Entries ) == 1 {
        logger.Log( fmt.Sprintf( "Lookup %s OK, time %s", userId, time.Since( start ) ) )
        return &api.User {
		    UserId:       userId,
			DisplayName:  sr.Entries[ 0 ].GetAttributeValue( "displayName" ),
			FirstName:    sr.Entries[ 0 ].GetAttributeValue( "givenName" ),
			Initials:     sr.Entries[ 0 ].GetAttributeValue( "initials" ),
			LastName:     sr.Entries[ 0 ].GetAttributeValue( "sn" ),
			Description:  sr.Entries[ 0 ].GetAttributeValue( "description" ),
			Department:   sr.Entries[ 0 ].GetAttributeValue( "uvaDisplayDepartment" ),
			Title:        sr.Entries[ 0 ].GetAttributeValue( "title" ),
			Office:       sr.Entries[ 0 ].GetAttributeValue( "physicalDeliveryOfficeName" ),
			Phone:        sr.Entries[ 0 ].GetAttributeValue( "telephoneNumber" ),
			Email:        sr.Entries[ 0 ].GetAttributeValue( "mail" ),
		}, nil
	}

    logger.Log( fmt.Sprintf( "Lookup %s NOT FOUND, time %s", userId, time.Since( start ) ) )

   // return empty user if not found
   return nil, nil
}