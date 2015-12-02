package main

import (
   "fmt"
   //"log"
   "github.com/nmcclain/ldap"
)

var (
	ldapServer string   = "ldap.virginia.edu"
	ldapPort   uint16   = 389
	baseDN     string   = "o=University of Virginia,c=US"
	Attributes []string = []string{"displayName", "givenName", "initials", "sn", "description", "uvaDisplayDepartment", "title", "physicalDeliveryOfficeName", "mail", "telephoneNumber"}
)

func LookupUser( userId string ) ( User, error ) {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		//log.Fatalf("ERROR: %s\n", err.Error())
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
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf( "(userId=%s)", userId ),
		Attributes,
		nil )

	sr, err := l.Search(search)
	if err != nil {
		// log.Fatalf("ERROR: %s\n", err.Error())
		return User{ }, err
	}

	if len( sr.Entries ) == 1 {
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

   // return empty user if not found
   return User{ }, nil
}