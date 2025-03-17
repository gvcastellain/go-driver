package folders

import (
	"errors"
	"time"
)

var (
	ErrNameRequired = errors.New("name is required")
)

func New(name string, parentID int64) (*Folder, error) {
	f := &Folder{
		ParentID: parentID,
		Name:     name,
	}

	err := f.Validade()
	if err != nil {
		return nil, err
	}

	return f, nil
}

type Folder struct {
	ID         int64     `json:"id"`
	ParentID   int64     `json:"parent_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

type FolderContent struct {
	Folder  Folder           `json:"folder"`
	Content []FolderResource `json:"content"`
}

type FolderResource struct {
	ID         int64     `json:"id"`
	Type       string    `json:"Type"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (f *Folder) Validade() error {
	if f.Name == "" {
		return ErrNameRequired
	}

	return nil
}
