package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func RuneMetricsGetUser(username string) (*RuneScapeUser, error) {
	var Endpoint = fmt.Sprintf("https://apps.runescape.com/runemetrics/profile/profile?user=%s&activities=20", username)

	resp, err := http.Get(Endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var returnUser RuneScapeUser
	err = json.Unmarshal(body, &returnUser)
	if err != nil {
		return nil, err
	}
	return &returnUser, nil
}

type RuneScapeUser struct {
	Magic            int         `json:"magic"`
	QuestsStarted    int         `json:"questsstarted"`
	TotalSkill       int         `json:"totalskill"`
	QuestsCompleted  int         `json:"questscomplete"`
	QuestsNotStarted int         `json:"questsnotstarted"`
	TotalXP          int         `json:"totalxp"`
	Ranged           int         `json:"ranged"`
	Activities       *[]Activity `json:"activities"`
	SkillValues      []Skill     `json:"skillvalues"`
	Name             string      `json:"name"`
	Rank             string      `json:"rank"`
	Melee            int         `json:"melee"`
	CombatLevel      int         `json:"combatlevel"`
	LoggedIn         string      `json:"loggedin"`
}

type Activity struct {
	Date    time.Time `json:"date"`
	Details string    `json:"details"`
	Text    string    `json:"text"`
}

func (a *Activity) UnmarshalJSON(data []byte) error {
	// Define a struct that matches the JSON structure
	type Alias Activity
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	// Unmarshal the rest of the fields
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the date using a custom format
	layout := "02-Jan-2006 15:04"
	parsedDate, err := time.Parse(layout, aux.Date)
	if err != nil {
		return err
	}

	// Set the parsed date
	a.Date = parsedDate
	return nil
}

type Skill struct {
	Level int `json:"level"`
	XP    int `json:"xp"`
	Rank  int `json:"rank"`
	ID    int `json:"id"`
}

func (u *RuneScapeUser) GetSkillByID(id int) (*Skill, error) {
	// Iterate through the slice of SkillValues
	for _, skill := range u.SkillValues {
		if skill.ID == id {
			// Return the skill if the ID matches
			return &skill, nil
		}
	}
	// Return an error if no matching ID is found
	return nil, fmt.Errorf("skill with ID %d not found", id)
}
