package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type JSONStore struct {
	file     string
	mu       sync.Mutex
	contacts map[int]*Contact
	nextID   int
}

func NewJSONStore(path string) (*JSONStore, error) {
	j := &JSONStore{file: path, contacts: make(map[int]*Contact), nextID: 1}
	if _, err := os.Stat(path); err == nil {
		if err := j.load(); err != nil {
			return nil, err
		}
	}
	return j, nil
}

func (j *JSONStore) load() error {
	j.mu.Lock()
	defer j.mu.Unlock()
	f, err := os.Open(j.file)
	if err != nil {
		return err
	}
	defer f.Close()
	var list []Contact
	if err := json.NewDecoder(f).Decode(&list); err != nil {
		return err
	}
	j.contacts = make(map[int]*Contact)
	max := 0
	for i := range list {
		c := list[i]
		j.contacts[c.ID] = &c
		if c.ID > max {
			max = c.ID
		}
	}
	j.nextID = max + 1
	return nil
}

func (j *JSONStore) save() error {
	j.mu.Lock()
	defer j.mu.Unlock()
	list := make([]Contact, 0, len(j.contacts))
	for _, c := range j.contacts {
		list = append(list, *c)
	}
	f, err := os.Create(j.file)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(list)
}

func (j *JSONStore) Add(c *Contact) error {
	if c == nil {
		return errors.New("contact nil")
	}
	j.mu.Lock()
	defer j.mu.Unlock()
	if c.ID == 0 {
		c.ID = j.nextID
		j.nextID++
	} else {
		if _, ok := j.contacts[c.ID]; ok {
			return errors.New("id already used")
		}
		if c.ID >= j.nextID {
			j.nextID = c.ID + 1
		}
	}
	j.contacts[c.ID] = c
	return j.save()
}

func (j *JSONStore) GetAll() ([]*Contact, error) {
	j.mu.Lock()
	defer j.mu.Unlock()
	out := make([]*Contact, 0, len(j.contacts))
	for _, c := range j.contacts {
		out = append(out, c)
	}
	return out, nil
}

func (j *JSONStore) GetByID(id int) (*Contact, error) {
	j.mu.Lock()
	defer j.mu.Unlock()
	if c, ok := j.contacts[id]; ok {
		return c, nil
	}
	return nil, errors.New("contact not found")
}

func (j *JSONStore) Update(id int, newName, newEmail string) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	c, ok := j.contacts[id]
	if !ok {
		return errors.New("contact not found")
	}
	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}
	return j.save()
}

func (j *JSONStore) Delete(id int) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	if _, ok := j.contacts[id]; !ok {
		return errors.New("contact not found")
	}
	delete(j.contacts, id)
	return j.save()
}

func (j *JSONStore) Close() error { return nil }
