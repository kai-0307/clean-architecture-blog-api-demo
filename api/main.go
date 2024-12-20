package main

import (
	"api/infrastructure"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	// 共通のデータベース接続ロジックを使用
	db := infrastructure.NewDB()
	defer db.Close()

	mux := http.NewServeMux()

	// /blogs エンドポイントを定義
	mux.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// ブログ一覧を取得
			rows, err := db.Query("SELECT id, title, content, author, created_at FROM blogs")
			if err != nil {
				log.Printf("Database query failed: %v", err)
				http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var blogs []Blog
			for rows.Next() {
				var blog Blog
				err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt)
				if err != nil {
					log.Printf("Failed to parse row: %v", err)
					http.Error(w, "Failed to parse blogs", http.StatusInternalServerError)
					return
				}
				blogs = append(blogs, blog)
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(blogs); err != nil {
				log.Printf("Failed to encode JSON: %v", err)
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}

		case http.MethodPost:
			var blog Blog
			if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
			}

			// INSERTクエリに RETURNING id を追加して新しいIDを取得
			query := "INSERT INTO blogs (title, content, author, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id"
			err := db.QueryRow(query, blog.Title, blog.Content, blog.Author).Scan(&blog.ID)
			if err != nil {
					log.Printf("Failed to insert blog: %v", err)
					http.Error(w, "Failed to create blog", http.StatusInternalServerError)
					return
			}

			// 作成日時をセット
			blog.CreatedAt = time.Now().Format(time.RFC3339)

			// JSONレスポンスとして返す
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blog)

		case http.MethodDelete:
			// ブログ削除
			id := r.URL.Query().Get("id")
			if id == "" {
				http.Error(w, "Missing blog ID", http.StatusBadRequest)
				return
			}

			blogID, err := strconv.Atoi(id)
			if err != nil || blogID <= 0 {
				http.Error(w, "Invalid blog ID", http.StatusBadRequest)
				return
			}

			_, err = db.Exec("DELETE FROM blogs WHERE id = $1", blogID)
			if err != nil {
				log.Printf("Failed to delete blog: %v", err)
				http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Blog deleted successfully"))

		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// CORS設定
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c.Handler(mux))
}
