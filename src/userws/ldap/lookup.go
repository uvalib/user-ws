package ldap

import (
      "fmt"
      "github.com/nmcclain/ldap"
      "time"
      "userws/api"
      "userws/logger"
)

var Attributes []string = []string {
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

func LookupUser(endpoint string, timeout int, baseDn string, userId string) (*api.User, error) {

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
            fmt.Sprintf("(userId=%s)", userId),
            Attributes,
            nil)

      sr, err := l.Search(search)
      if err != nil {
            logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
            return nil, err
      }

      if len(sr.Entries) == 1 {
            logger.Log(fmt.Sprintf("Lookup %s OK, time %s", userId, time.Since(start)))
            return &api.User{
                  UserId:      userId,
                  DisplayName: sr.Entries[0].GetAttributeValue( Attributes[ 0 ] ),
                  FirstName:   sr.Entries[0].GetAttributeValue( Attributes[ 1 ] ),
                  Initials:    sr.Entries[0].GetAttributeValue( Attributes[ 2 ] ),
                  LastName:    sr.Entries[0].GetAttributeValue( Attributes[ 3 ] ),
                  Description: sr.Entries[0].GetAttributeValue( Attributes[ 4 ] ),
                  Department:  sr.Entries[0].GetAttributeValue( Attributes[ 5 ] ),
                  Title:       sr.Entries[0].GetAttributeValue( Attributes[ 6 ] ),
                  Office:      sr.Entries[0].GetAttributeValue( Attributes[ 7 ] ),
                  Phone:       sr.Entries[0].GetAttributeValue( Attributes[ 8 ] ),
                  Email:       sr.Entries[0].GetAttributeValue( Attributes[ 9 ] ),
            }, nil
      }

      logger.Log(fmt.Sprintf("Lookup %s NOT FOUND, time %s", userId, time.Since(start)))

      // return empty user if not found
      return nil, nil
}
