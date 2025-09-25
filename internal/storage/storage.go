package storage

// Contact est notre structure de données centrale
type Contact struct {
	ID    int
	Name  string
	Email string
}

// Storer est un CONTRAT de stockage
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie par du comment c'est fait
// (en mémoire, fichier, BDD...) seulement de ce qui peut être fait
type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetByID(id int) (*Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}
