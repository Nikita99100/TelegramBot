package main

type ResponseTask struct {
	UserID   string `json:"user_id"`
	UserTask string `json:"user_task"`
}
type ResponseTaskIndex struct {
	UserID    string `json:"user_id"`
	TaskIndex int    `json:"task_index"`
}
type ResponseID struct {
	UserID string `json:"user_id"`
}
