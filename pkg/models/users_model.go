package models

//users_model.go

type Users struct {
	Id_user           int    `json:"id_user"`
	Surname           string `json:"surname"`
	Username          string `json:"username"`
	Patronymic        string `json:"patronymic"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Login             string `json:"login"`
	PasswordHash      string `json:"-"`
	Salt              string `json:"-"`
	PasswordAlgorithm string `json:"-"`
	Admin             bool   `json:"admin"`
	Manager           bool   `json:"manager"`
	Roles             []Role `json:"roles"`

	//Notification
	Id_notification_settings      int    `json:"id_notification"`
	Variant_notification_settings string `json:"name_notification"`

	//Role
	Id_role   int    `json:"id_role"`
	Name_role string `json:"name_role"`
}

type Role struct {
	Id_role   int    `json:"id_role"`
	Name_role string `json:"name_role"`
}
