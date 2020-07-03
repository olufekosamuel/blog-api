package helpers

import "github.com/olufekosamuel/blog-api/models"

func Validate(u models.User) (string, error) {
	if u.Email == "" {
		return "Email cannot be empty", nil
	}
	if u.Firstname == "" {
		return "Firstname cannot be empty", nil
	}
	if u.Lastname == "" {
		return "Lastname cannot be empty", nil
	}
	if u.Password == "" {
		return "Password cannot be empty", nil
	}

	return "", nil
}
