//handles web app authentication and authorization
package auth

type AuthManager interface {
	//accepts username and password and transforms to custom credentials type
	NewCredentials(username, password string) Credentials
	//accepts two credential instances, compares equality and returns auth type
	Authenticate(a, b Credentials) (Auth, error)
	CheckAuth() error
}
type Auth interface {
	//retrieves string details about authentication jwt, session etc
	Get() string
}

//interface to represent credentials
type Credentials interface {
	//retrieves the principal associated with the credentials eg. username, id etc
	Principal() string
}
