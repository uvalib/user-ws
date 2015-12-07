# set the definitions
INSTANCE=user-ws
NAMESPACE=uvadave

# pull the current runable image
docker pull $NAMESPACE/$INSTANCE:latest

# return status
exit $?
