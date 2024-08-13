package models

import (
	"database/sql"
	"errors"
	"time"
)

// TeamModelInterface is interface for teams model
type TeamModelInterface interface {
	Insert(name string) (int, error)
	Get(id int) (*Team, error)
	Update(id int, score int) error
}

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

// Insert creates a new team in the database
func (m *TeamModel) Insert(name string) (int, error) {
	stmt := "INSERT INTO teams (name, score, created) VALUES(?, ?, CURRENT_TIMESTAMP)"
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

// Get returns a row from the database that matches the id
func (m *TeamModel) Get(id int) (*Team, error) {
	stmt := "SELECT id, name, score, created FROM teams WHERE id = ?"
	row := m.DB.QueryRow(stmt, id)
	s := &Team{}
	err := row.Scan(&s.ID, &s.Name, &s.Score, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return s, nil
}

// Update takes a team id and new score and updates the score in the database
func (m *TeamModel) Update(id int, score int) error {
	stmt := "UPDATE teams SET score = ? WHERE id = ?"
	_, err := m.DB.Exec(stmt, score, id)
	if err != nil {
		return err
	}
	return nil
}
