#
#
#

# ensure we have an endpoint
if [ -z "$USERINFO_URL" ]; then
   echo "ERROR: USERINFO_URL is not defined"
   exit 1
fi

# issue the command
echo "$USERINFO_URL"
curl $USERINFO_URL/version

exit 0

#
# end of file
#
