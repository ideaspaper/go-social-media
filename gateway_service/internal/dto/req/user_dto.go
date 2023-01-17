package req

type UserDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (ud UserDto) ErrorMessages(field, tag string) string {
	switch field {
	case "Email":
		switch tag {
		case "required":
			return "email is required"
		case "email":
			return "email format is wrong"
		}
	case "Password":
		switch tag {
		case "required":
			return "password is required"
		case "min":
			return "password minimum length is 8"
		}
	case "FirstName":
		switch tag {
		case "required":
			return "first_name is required"
		}
	case "LastName":
		switch tag {
		case "required":
			return "last_name is required"
		}
	}
	return ""
}
