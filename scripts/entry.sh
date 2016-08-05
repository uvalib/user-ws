# set blank options variables
LDAPURL_OPT=""
TOKENURL_OPT=""
DEBUG_OPT=""

# public LDAP hostname:port
if [ -n "$LDAP_URL" ]; then
   LDAPURL_OPT="--url $LDAP_URL"
fi

# token authentication service URL
if [ -n "$TOKENAUTH_URL" ]; then
   TOKENURL_OPT="--tokenauth $TOKENAUTH_URL"
fi

# service debugging
if [ -n "$USERINFO_DEBUG" ]; then
   DEBUG_OPT="--debug"
fi

bin/user-ws $LDAPURL_OPT $TOKENURL_OPT $DEBUG_OPT

#
# end of file
#
