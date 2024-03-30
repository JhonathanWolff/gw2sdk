package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Skills struct {
	Gw2sdk Connection.GW2sdk
}

type SkillsResponseDetails struct {
	Name  string `json:"name"`
	Facts []struct {
		Text        string `json:"text"`
		Type        string `json:"type"`
		Value       int    `json:"value,omitempty"`
		Icon        string `json:"icon,omitempty"`
		HitCount    int    `json:"hit_count,omitempty"`
		Duration    int    `json:"duration,omitempty"`
		Status      string `json:"status,omitempty"`
		Description string `json:"description,omitempty"`
		ApplyCount  int    `json:"apply_count,omitempty"`
	} `json:"facts"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	WeaponType   string   `json:"weapon_type"`
	Professions  []string `json:"professions"`
	Slot         string   `json:"slot"`
	Cost         int      `json:"cost"`
	FlipSkill    int      `json:"flip_skill"`
	Categories   []string `json:"categories"`
	TraitedFacts []struct {
		Text          string `json:"text"`
		Type          string `json:"type"`
		Icon          string `json:"icon"`
		RequiresTrait int    `json:"requires_trait"`
		HitCount      int    `json:"hit_count"`
		Overrides     int    `json:"overrides"`
	} `json:"traited_facts"`
	Icon     string `json:"icon"`
	ID       int    `json:"id"`
	ChatLink string `json:"chat_link"`
}

func (a *Skills) GetDetails(parameters map[string]string) SkillsResponseDetails {

	response := SkillsResponseDetails{}
	a.Gw2sdk.Retrieve("skills", parameters, &response)

	return response

}
