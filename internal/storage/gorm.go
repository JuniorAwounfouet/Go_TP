package storage

import (
	"errors"

	"minicrm/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(path string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.Contact{}); err != nil {
		return nil, err
	}
	return &GORMStore{db: db}, nil
}

func (g *GORMStore) Add(c *Contact) error {
	if c == nil {
		return errors.New("contact nil")
	}
	m := model.Contact{ID: c.ID, Name: c.Name, Email: c.Email}
	if err := g.db.Create(&m).Error; err != nil {
		return err
	}
	c.ID = m.ID
	return nil
}

func (g *GORMStore) GetAll() ([]*Contact, error) {
	var list []model.Contact
	if err := g.db.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*Contact, 0, len(list))
	for i := range list {
		m := list[i]
		out = append(out, &Contact{ID: m.ID, Name: m.Name, Email: m.Email})
	}
	return out, nil
}

func (g *GORMStore) GetByID(id int) (*Contact, error) {
	var m model.Contact
	if err := g.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &Contact{ID: m.ID, Name: m.Name, Email: m.Email}, nil
}

func (g *GORMStore) Update(id int, newName, newEmail string) error {
	var m model.Contact
	if err := g.db.First(&m, id).Error; err != nil {
		return err
	}
	if newName != "" {
		m.Name = newName
	}
	if newEmail != "" {
		m.Email = newEmail
	}
	return g.db.Save(&m).Error
}

func (g *GORMStore) Delete(id int) error {
	return g.db.Delete(&model.Contact{}, id).Error
}

func (g *GORMStore) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
