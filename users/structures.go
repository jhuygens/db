package users

// User db info
type User struct {
	ID               int64
	Email            string
	Password         string
	AppName          string
	ClientID         string
	CurrentSecretKey string
	Token            string
	RefreshToken     string
	RedirectURLs     []RedirectURL
	CreatedAt        string
}

// RedirectURL of the user app
type RedirectURL struct {
	ID          string
	URL         string
	Description string
	CreatedAt   string
}
