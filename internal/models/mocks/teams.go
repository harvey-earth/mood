package mocks

import (
	"time"

	"github.com/harvey-earth/mood/internal/models"
)

var mockTeam = &models.Team{
	ID:      1,
	Name:    "Example Team",
	Score:   50,
	Created: time.Now(),
}

// TeamModel is a struct to test the database model
type TeamModel struct{}

// Insert mocks a call to insert a new team to the database
func (m *TeamModel) Insert(name string) (int, error) {
	return 2, nil
}

// Get mocks a call to retrieve a team by ID
func (m *TeamModel) Get(id int) (*models.Team, error) {
	switch id {
	case 1:
		return mockTeam, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *TeamModel) Update(id int, score int) error {
	return nil
}
