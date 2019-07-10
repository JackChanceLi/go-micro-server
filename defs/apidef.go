package defs

type UserIdentity struct {
	userName string `json:"userName"`
	passwd string `json:"password"`
	email string `json: "email"`
	role int `json:"role"`
}