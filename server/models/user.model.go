/*
User Model

Defines database schema representation for users.
*/

package models

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Salt         string
}
