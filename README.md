# Mini-CRM CLI — Guide rapide

Petit outil CLI pour gérer des contacts (CRUD). Ce README va droit au but : build, exécution et tests.

Prérequis
- Go 1.18+ installé

1) Installer dépendances et formater

cd d:\M2_DEV\Go
go mod tidy
gofmt -w .


2) Compiler


go build -tags sqlite_omit_load_extension -o minicrm.exe


Note : le tag `sqlite_omit_load_extension` est recommandé quand on utilise le driver
`github.com/glebarez/sqlite` (pure-Go) pour éviter les problèmes liés aux extensions.

3) Fichier de configuration
- `config.yaml` (exemple existant) :
	- `type: "memory"` -> stockage en mémoire (non persistant)
	- `type: "json"` -> stockage dans `json.path` (persistant)
	- `type: "gorm"` -> stockage via GORM/SQLite (fichier `gorm.path`)

4) Commandes (exemples)

- Ajouter :

.\minicrm.exe add --nom "Jean Dupont" --email "jean@efrei.fr"


- Lister :

.\minicrm.exe list


- Mettre à jour :

.\minicrm.exe update --id 1 --name "Jean M. Dupont" --email "jm@ex.com"


- Supprimer :

.\minicrm.exe delete --id 1


-- Fin --
