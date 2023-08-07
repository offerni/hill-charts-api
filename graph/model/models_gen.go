// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewScope struct {
	Colour  string `json:"colour"`
	Name    string `json:"name"`
	SquadID string `json:"squad_id"`
}

type NewSquad struct {
	CurrentCycleName *string `json:"current_cycle_name"`
	Name             string  `json:"name"`
}

type Scope struct {
	Colour     string  `json:"colour"`
	CreatedAt  *string `json:"created_at"`
	ID         string  `json:"id"`
	ModifiedAt *string `json:"modified_at"`
	Name       string  `json:"name"`
	Progress   string  `json:"progress"`
	SquadID    string  `json:"squad_id"`
}

type Squad struct {
	CreatedAt        *string  `json:"created_at"`
	CurrentCycleName string   `json:"current_cycle_name"`
	ID               string   `json:"id"`
	ModifiedAt       *string  `json:"modified_at"`
	Name             string   `json:"name"`
	Scope            []*Scope `json:"scope"`
}

type SquadList struct {
	Data       []*Squad `json:"data"`
	HasMore    bool     `json:"has_more"`
	TotalCount int      `json:"total_count"`
}

type UpdateScope struct {
	Colour   *string `json:"colour"`
	Name     *string `json:"name"`
	Progress *string `json:"progress"`
}

type UpdateSquad struct {
	CurrentCycleName *string `json:"current_cycle_name"`
	Name             *string `json:"name"`
}
