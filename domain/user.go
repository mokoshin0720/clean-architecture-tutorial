// Entity（ビジネスルールのためのデータを定義する）
package domain

type User struct {
	ID int
	FirstName string
	LastName string
}

type Users []User