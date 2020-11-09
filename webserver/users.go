package main

type User struct {
	ID    string `json:"id"`
	Tasks []Task `json:"tasks"`
}

var users []User

func FindUser(id string) User {
	for _, us := range users {
		if us.ID == id {
			return us
		}
	}
	users = append(users, User{
		ID:    id,
		Tasks: nil,
	})
	return users[len(users)-1]
}
func FindUserIndex(id string) int {
	for i, us := range users {
		if us.ID == id {
			return i
		}
	}
	users = append(users, User{
		ID:    id,
		Tasks: nil,
	})
	return len(users) - 1
}
