package models

import (
	"database/sql"
	"errors"
	"time"
)

// Team table model
type Team struct {
	ID      int
	Name    string
	Score   int
	Created time.Time
}

// TeamModel is a Struct for database model
type TeamModel struct {
	DB *sql.DB
}

// Creates a new team
func (m *TeamModel) Insert(name string) (int, error) {
	stmt := "INSERT INTO teams (name, score, created) VALUES(?, ?, UTC_TIMESTAMP())"
	result, err := m.DB.Exec(stmt, name, 50)
	if err != nil {
		return 0, nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m *TeamModel) Get(id int) (*Team, error) {
	stmt := "SELECT id, name, score, created FROM teams WHERE id = ?"
	row := m.DB.QueryRow(stmt, id)
	s := &Team{}
	err := row.Scan(&s.ID, &s.Name, &s.Score, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *TeamModel) Update(id int, score int) error {
	stmt := "UPDATE teams SET score = ? WHERE id = ?"
	_, err := m.DB.Exec(stmt, score, id)
	if err != nil {
		return err
	}
	return nil
}
