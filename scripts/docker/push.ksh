# set the definitions
INSTANCE=user-ws
NAMESPACE=uvadave

# push the current image
docker push $NAMESPACE/$INSTANCE

# return status
exit $?
