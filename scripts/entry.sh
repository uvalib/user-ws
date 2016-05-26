# set blank options variables
LDAPURL_OPT=""
TOKENURL_OPT=""

# public LDAP hostname:port
if [ -n "$LDAP_URL" ]; then
   LDAPURL_OPT="--url $LDAP_URL"
fi

# token authentication service URL
if [ -n "$TOKENAUTH_URL" ]; then
   TOKENURL_OPT="--tokenauth $TOKENAUTH_URL"
fi

bin/user-ws $LDAPURL_OPT $TOKENURL_OPT

#
# end of file
#
