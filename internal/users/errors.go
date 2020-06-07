package users

// WrongUsernameOrPasswordError - Error used in authentication
type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
