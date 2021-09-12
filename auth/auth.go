//handles web app authentication and authorization
package auth

type AuthManager interface {
	//accepts username and password and transforms to custom user details type
	NewCredentials(username, password string) Credentials
	//accepts user details and given password, compares equality and returns auth type
	Authenticate(u Credentials, password string) (Auth, error)
	//serves as a filter to verify authentication
	CheckAuth() error
}
type Auth interface {
	//retrieves string details about authentication jwt, session etc
	Get() string
}

//interface to represent credentials
type Credentials interface {
	Password
	UserDetails
}

type UserDetails interface {
	//retrieves the principal associated with the credentials eg. username, id etc
	Principal() string
}

type Password interface {
	//returns the given password
	Password() string
	//returns a string representation of the hashed password
	Hash() (string, error)
}
