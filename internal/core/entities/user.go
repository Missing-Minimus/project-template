package entities

import "time"

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	UpdatedAt time.Time
}

// Fiz a entidade seguindo os campos padrões, mas podemos adicionar mais campos se necessário (adicionei alguns que acredito que podem ser úteis)
// Basicamente, no core a gnt n coloca os campos em JSON ainda
// Pq vamos colocar dentro daquelas request entities
