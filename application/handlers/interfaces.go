package handlers

// IUserService presents user services methods
type IUserService interface {
	Create(login string, name, email, phone *string) error
}
