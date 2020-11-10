package main

type User struct {
	ID    string `json:"id"`
	Tasks []Task `json:"tasks"`
}

var users []User

func FindUser(id string) *User {
	for i, us := range users {
		if us.ID == id {
			return &users[i]
		}
	}
	users = append(users, User{
		ID:    id,
		Tasks: nil,
	})
	return &users[len(users)-1]
}
