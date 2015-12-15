package main

import (
   "fmt"
   "log"
   "time"
   "github.com/nmcclain/ldap"
)

var	Attributes []string = []string{"displayName", "givenName", "initials", "sn", "description", "uvaDisplayDepartment", "title", "physicalDeliveryOfficeName", "mail", "telephoneNumber"}

func LookupUser( userId string ) ( User, error ) {

	start := time.Now( )

	l, err := ldap.DialTimeout("tcp", config.LdapUrl, time.Second * 10 )
	if err != nil {
		log.Printf( "ERROR: %s\n", err.Error( ) )
		return User{ }, err
	}

	defer l.Close()
	// l.Debug = true

	//err = l.Bind(user, passwd)
	//if err != nil {
	//   log.Printf("ERROR: Cannot bind: %s\n", err.Error())
	//   return
	//}

	search := ldap.NewSearchRequest(
		config.LdapBaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf( "(userId=%s)", userId ),
		Attributes,
		nil )

	sr, err := l.Search(search)
	if err != nil {
		log.Printf( "ERROR: %s\n", err.Error( ) )
		return User{ }, err
	}

	if len( sr.Entries ) == 1 {
		log.Printf( "Lookup %s OK\t%s", userId, time.Since( start ) )
        return User {
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

   log.Printf( "Lookup %s NOT FOUND\t%s", userId, time.Since( start ) )

   // return empty user if not found
   return User{ }, nil
}