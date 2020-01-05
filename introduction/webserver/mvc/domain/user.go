package domain

import (
	"bytes"
	"fmt"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u *User) Bytes() []byte {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("%#v", u))
	return buff.Bytes()
}
