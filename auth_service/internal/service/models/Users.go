package models

type (
	User struct {
		Username    string
		Password    string
		Permissions map[string]string
	}
)
