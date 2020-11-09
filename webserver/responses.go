package main

type ResponseTask struct {
	UserID   string `json:"user_id"`
	UserTask string `json:"user_task"`
}

type ResponseID struct {
	UserID string `json:"user_id"`
}
