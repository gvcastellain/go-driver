package files

import (
	"errors"
	"time"
)

var (
	ErrOwnerRequired = errors.New("owner is required")

	ErrNameRequired = errors.New("name is required")

	ErrTypeRequired = errors.New("type is required")

	ErrPathRequired = errors.New("path is required")
)

func New(ownerID int64, name, fileType, path string) (*File, error) {
	f := File{
		OwnerID:    ownerID,
		Name:       name,
		Type:       fileType,
		Path:       path,
		ModifiedAt: time.Now(),
	}

	err := f.Validade()
	if err != nil {
		return nil, err
	}

	return &f, nil
}

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validade() error {
	if f.OwnerID == 0 {
		return ErrOwnerRequired
	}

	if f.Name == "" {
		return ErrNameRequired
	}

	if f.Type == "" {
		return ErrTypeRequired
	}

	if f.Path == "" {
		return ErrPathRequired
	}

	return nil
}
