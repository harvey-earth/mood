package models

import "errors"

// ErrNoRecord is the record returned when no models match
var ErrNoRecord = errors.New("models: no matching record found")
