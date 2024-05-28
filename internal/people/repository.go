package people

import (
	"database/sql"
	"luizilha/rinha-ilha/internal/types"
  "github.com/google/uuid"
)

type repository struct {
	*sql.DB
}

func NewRepository(dbValue *sql.DB) *repository {
	return &repository{DB: dbValue}
}

/*
func (p *repository) AllPeople() ([]types.People, error) {
	rows, error := p.DB.Query("SELECT id, nickname, name, birthdate, stack FROM people")
	if error != nil {
		return nil, error
	}

	var arrayPeople []types.People
	for rows.Next() {
		p := types.People{}
		nError := rows.Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthdate, &p.Stack)
		if nError != nil {
			return nil, nError
		}
		arrayPeople = append(arrayPeople, p)
	}
	return arrayPeople, nil
}
*/

func (p *repository) Insert(people *types.People) (string, error) {
  var id string
	row := p.DB.QueryRow("INSERT INTO people(id, name, nickname, birthdate, stack, search) values($1, $2, $3, $4, $5, $6) RETURNING id", 
    uuid.New().String(), people.Name, people.Nickname, people.Birthdate, people.Stack, people.SearchValue()).Scan(&id)

  if row != nil {
    return "", row
  }
  return id, nil
}

func (p *repository) Count() (int, error) {
	row := p.DB.QueryRow("SELECT count(name) FROM people")

	var count int
	nError := row.Scan(&count)
	if nError != nil {
		return 0, nError
	}

	return count, nil
}
