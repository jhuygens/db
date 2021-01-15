package users

// Save one user
func Save(user User) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return usrs.Save(user)
}

// Get one user by user id
func Get(userID int64) (*User, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return nil, err
	}
	return usrs.Get(userID)
}

// GetAll users
func GetAll() ([]User, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return nil, err
	}
	return usrs.GetAll()
}

// Delete user by user id
func Delete(userID int64) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return usrs.Delete(userID)
}

// GetRedirectURLs of user app
func GetRedirectURLs(userID int64) ([]RedirectURL, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return nil, err
	}
	return usrs.GetRedirectURLs(userID)
}

// UpdateCurrentSecretKey of user
func UpdateCurrentSecretKey(userID int64, secretKey string) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return usrs.UpdateCurrentSecretKey(userID, secretKey)
}

// GetByEmail get user by email
func GetByEmail(email string) (*User, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return nil, err
	}
	return usrs.GetByEmail(email)
}

// UpdateToken update user access token
func UpdateToken(clientID, token, refreshToken string) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return usrs.UpdateToken(clientID, token, refreshToken)
}

// GetByClientID get user by client Id
func GetByClientID(clientID string) (*User, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return nil, err
	}
	return usrs.GetByClientID(clientID)
}
