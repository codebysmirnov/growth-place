package handlers

// UserHandler handle user requests
type UserHandler struct {
	userService IUserService
}

// NewUserHandlers returns new UserHandler instance
func NewUserHandlers(userService IUserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
