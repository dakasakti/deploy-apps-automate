package models

type Book struct {
	ID       uint   `json:"id"`
	Title    string `json:"title,omitempty"`
	Penulis  string `json:"penulis,omitempty"`
	Penerbit string `json:"penerbit,omitempty"`
	UserID   uint   `json:"-"`
}

type BookRequest struct {
	Title    string `json:"title" validate:"required,min=5,max=255"`
	Penulis  string `json:"penulis" validate:"required,min=5,max=255"`
	Penerbit string `json:"penerbit" validate:"required,min=5,max=255"`
}

type BookUpdateRequest struct {
	Title    string `json:"title,omitempty" validate:"omitempty,min=5,max=255"`
	Penulis  string `json:"penulis,omitempty" validate:"omitempty,min=5,max=255"`
	Penerbit string `json:"penerbit,omitempty" validate:"omitempty,min=5,max=255"`
}
