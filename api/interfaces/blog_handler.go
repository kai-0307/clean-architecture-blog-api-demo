package handler

import (
	"api/entity"
	"api/usecase"
	"encoding/json"
	"net/http"
	"time"
)

type BlogHandler struct {
	Usecase *usecase.BlogUsecase
}

type CreateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var req CreateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 現在のタイムスタンプを設定
	blog := entity.Blog{
		Title:     req.Title,
		Content:   req.Content,
		Author:    req.Author,
		CreatedAt: time.Now(),
	}

	err := h.Usecase.CreateBlog(blog)
	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Blog created successfully"))
}
