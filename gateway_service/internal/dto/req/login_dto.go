package req

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (lv LoginDto) ErrorMessages(field, tag string) string {
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
	}
	return ""
}
