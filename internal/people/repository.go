package people

import (
	"database/sql"
)

type repository struct {
	*sql.DB
}

func NewRepository(dbValue *sql.DB) *repository {
	return &repository{ DB: dbValue }
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
func (p *repository) Count() (int, error) {
  row := p.DB.QueryRow("SELECT count(name) FROM people")

  var count int
  nError := row.Scan(&count)
  if (nError != nil) {
    return 0, nError
  }

  return count, nil
}
