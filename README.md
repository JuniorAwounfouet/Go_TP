Ce programme Go implémente un gestionnaire de contacts en ligne de commande.
Fonctionnalités principales :

- Ajouter, lister, supprimer et mettre à jour des contacts via un menu interactif ou des flags CLI.
- Les contacts sont stockés en mémoire dans une map, chaque contact ayant un ID unique, un nom et un email.
- Utilisation des flags `-ajouter`, `-id`, `-nom`, `-email` pour ajouter un contact directement depuis la ligne de commande.
- Toutes les opérations sont réalisées en mémoire (pas de persistance sur disque).

Aspects techniques :

- Utilisation du package standard `flag` pour la gestion des arguments CLI.
- Utilisation de la map Go (`map[int]Contact`) pour un accès rapide aux contacts par ID.
- Gestion des erreurs et validation des entrées utilisateur.
