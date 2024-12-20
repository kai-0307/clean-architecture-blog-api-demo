package usecase

import (
	"api/entity"
	"errors"
)

type BlogRepository interface {
	GetAll() ([]entity.Blog, error)
	GetByID(id int) (*entity.Blog, error)
	Create(blog entity.Blog) error
	Delete(id int) error
}

type BlogUsecase struct {
	Repo BlogRepository // リポジトリを依存として注入
}

// 記事一覧を取得
func (u *BlogUsecase) GetBlogs() ([]entity.Blog, error) {
	return u.Repo.GetAll()
}

// 記事詳細を取得
func (u *BlogUsecase) GetBlogByID(id int) (*entity.Blog, error) {
	if id <= 0 {
		return nil, errors.New("invalid blog ID")
	}
	return u.Repo.GetByID(id)
}

// 記事を作成
func (u *BlogUsecase) CreateBlog(blog entity.Blog) error {
	return u.Repo.Create(blog)
}

// 記事を削除
func (u *BlogUsecase) DeleteBlog(id int) error {
	if id <= 0 {
		return errors.New("invalid blog ID")
	}
	return u.Repo.Delete(id)
}
