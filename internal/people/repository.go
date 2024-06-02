package people

import (
	"database/sql"
	"luizilha/rinha-ilha/internal/types"

	"github.com/google/uuid"
)

type Repository struct {
	*sql.DB
}

func NewRepository(dbValue *sql.DB) *Repository {
	return &Repository{DB: dbValue}
}

func (p *Repository) Insert(people *types.People) (string, error) {
	var id string
	row := p.DB.QueryRow("INSERT INTO people(id, name, nickname, birthdate, stack, search) values($1, $2, $3, $4, $5, $6) RETURNING id",
		uuid.New().String(), people.Name, people.Nickname, people.Birthdate, people.Stack, people.SearchValue()).Scan(&id)

	if row != nil {
		return "", row
	}
	return id, nil
}

func (p *Repository) Count() (int, error) {
	row := p.DB.QueryRow("SELECT count(name) FROM people")

	var count int
	nError := row.Scan(&count)
	if nError != nil {
		return 0, nError
	}

	return count, nil
}

func (p *Repository) FindById(id string) (*types.People, error) {

	rows, error := p.DB.Query("SELECT id, name, nickname, birthdate, stack from people p WHERE p.id = $1", id)

	if error != nil {
		return nil, error
	}

	var people []types.People
	for rows.Next() {
		p := types.People{}
		nError := rows.Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthdate, &p.Stack)
		if nError != nil {
			return nil, nError
		}
		people = append(people, p)
	}
	if len(people) > 0 {
		return &people[0], nil
	}
	return nil, nil
}
