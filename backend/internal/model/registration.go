package model

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Registration struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	DOB       time.Time      `gorm:"not null" json:"dob"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// UnmarshalJSON for custom date parsing
func (r *Registration) UnmarshalJSON(data []byte) error {
	type Alias Registration
	aux := &struct {
		DOB string `json:"dob"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the date
	dob, err := time.Parse("2006-01-02", aux.DOB)
	if err != nil {
		return fmt.Errorf("invalid date format for dob: %v", err)
	}
	r.DOB = dob

	return nil
}
