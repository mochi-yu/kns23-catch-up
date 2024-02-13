package model

type RegisterPostParam struct {
	UserID      string `json:"userID"`
	DisplayName string `json:"displayName"`
	UserName    string `json:"userName"`
	ClassID     string `json:"ClassID"`
	MailAddress string `json:"mailAddress"`
}

type RegisterPostResponse struct {
	UserID      string `json:"userID"`
	DisplayName string `json:"displayName"`
	ClassID     string `json:"ClassID"`
	MailAddress string `json:"mailAddress"`
}
