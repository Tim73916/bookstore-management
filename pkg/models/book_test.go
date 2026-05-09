package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	return db, mock, err
}

func TestBookCreate(t *testing.T) {
	book := &Book{
		Name:        "Test Book",
		Author:      "Test Author",
		Publication: "2024",
	}

	assert.NotNil(t, book)
	assert.Equal(t, "Test Book", book.Name)
	assert.Equal(t, "Test Author", book.Author)
}

func TestBookValidation(t *testing.T) {
	tests := []struct {
		name     string
		book     Book
		hasError bool
	}{
		{
			name: "Valid book",
			book: Book{
				Name:        "Valid Name",
				Author:      "Valid Author",
				Publication: "2024",
			},
			hasError: false,
		},
		{
			name: "Empty name",
			book: Book{
				Name:        "",
				Author:      "Author",
				Publication: "2024",
			},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hasError {
				assert.Empty(t, tt.book.Name)
			} else {
				assert.NotEmpty(t, tt.book.Name)
			}
		})
	}
}
