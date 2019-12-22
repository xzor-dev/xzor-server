package team

// ID is a unique string assigned to individual teams.
type ID string

// Team contains data for individual teams.
type Team struct {
	ID      ID        `json:"id"`
	Name    string    `json:"name"`
	Tag     string    `json:"tag"`
	Members []*Member `json:"members;omitempty"`
}

// CreateTeamParams are parameters required for creating new teams.
type CreateTeamParams struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// Validate determines if all parameters are set and returns
// a map of empty or invalid keys.  Returns nil if all parameters are valid.
func (p *CreateTeamParams) Validate() map[string]string {
	invalidParams := make(map[string]string)
	isValid := true
	if p.Name == "" {
		isValid = false
		invalidParams["name"] = "a team name is required"
	}

	if isValid {
		return nil
	}
	return invalidParams
}
