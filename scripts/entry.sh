# set blank options variables
LDAPURL_OPT=""
LDAPBINDACCT_OPT=""
LDAPBINDPWD_OPT=""
#LDAPBASEDN_OPT=""
TIMEOUT_OPT=""
TOKENURL_OPT=""
HEALTHCHECK_OPT=""
DEBUG_OPT=""

# LDAP endpoint
if [ -n "$LDAP_URL" ]; then
   LDAPURL_OPT="--ldapendpoint $LDAP_URL"
fi

# LDAP bind account
if [ -n "$LDAP_ACCOUNT" ]; then
   LDAPBINDACCT_OPT="--ldapbindacct $LDAP_ACCOUNT"
fi

# LDAP bind account password
if [ -n "$LDAP_ACCOUNT_PWD" ]; then
   LDAPBINDPWD_OPT="--ldapbindpwd $LDAP_ACCOUNT_PWD"
fi

# LDAP base DN
#if [ -n "$LDAP_BASEDN" ]; then
#   LDAPBASEDN_OPT="--ldapbasedn '$LDAP_BASEDN'"
#fi

# connection timeout
if [ -n "$SERVICE_TIMEOUT" ]; then
   TIMEOUT_OPT="--timeout $SERVICE_TIMEOUT"
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

bin/user-ws $LDAPURL_OPT $LDAPBINDACCT_OPT $LDAPBINDPWD_OPT $TIMEOUT_OPT $HEALTHCHECK_OPT $TOKENURL_OPT $DEBUG_OPT

#
# end of file
#
