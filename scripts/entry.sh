# set blank options variables
LDAPURL_OPT=""
LDAPTIMEOUT_OPT=""
TOKENURL_OPT=""
HEALTHCHECK_OPT=""
DEBUG_OPT=""

# public LDAP hostname:port
if [ -n "$LDAP_URL" ]; then
   LDAPURL_OPT="--url $LDAP_URL"
fi

# connection timeout
if [ -n "$LDAP_TIMEOUT" ]; then
   LDAPTIMEOUT_OPT="--timeout $LDAP_TIMEOUT"
fi

# token authentication service URL
if [ -n "$TOKENAUTH_URL" ]; then
   TOKENURL_OPT="--tokenauth $TOKENAUTH_URL"
fi

# healthcheck username
if [ -n "$HEALTHCHECK_USER" ]; then
   HEALTHCHECK_OPT="--hcuser $HEALTHCHECK_USER"
fi

# service debugging
if [ -n "$USERINFO_DEBUG" ]; then
   DEBUG_OPT="--debug"
fi

bin/user-ws $LDAPURL_OPT $LDAPTIMEOUT_OPT $HEALTHCHECK_OPT $TOKENURL_OPT $DEBUG_OPT

#
# end of file
#
