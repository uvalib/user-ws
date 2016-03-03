package authtoken

func Validate( token string ) bool {
    return token == "secret"
}
