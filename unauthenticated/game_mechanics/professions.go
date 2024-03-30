package unauthenticated

import (
	Connection "gw2sdk/connection"
)

type Professions struct {
	Gw2sdk Connection.GW2sdk
}

type ProfessionsResponse []string

type WeaponDetails struct {
	Specialization int `json:"specialization"`
	Skills         []struct {
		ID   int    `json:"id"`
		Slot string `json:"slot"`
	} `json:"skills"`
}

type ProfessionsResponseDetails []struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	IconBig         string `json:"icon_big"`
	Specializations []int  `json:"specializations"`
	Weapons         struct {
		Hammer     WeaponDetails `json:"Hammer"`
		Mace       WeaponDetails `json:"Mace"`
		Pistol     WeaponDetails `json:"Pistol"`
		Sword      WeaponDetails `json:"Sword"`
		Scepter    WeaponDetails `json:"Scepter"`
		Focus      WeaponDetails `json:"Focus"`
		Shield     WeaponDetails `json:"Shield"`
		Torch      WeaponDetails `json:"Torch"`
		Warhorn    WeaponDetails `json:"Warhorn"`
		Greatsword WeaponDetails `json:"Greatsword"`
		Axe        WeaponDetails `json:"Axe"`
		Longbow    WeaponDetails `json:"Longbow"`
		Rifle      WeaponDetails `json:"Rifle"`
		Shortbow   WeaponDetails `json:"Shortbow"`
		Staff      WeaponDetails `json:"Staff"`
		Speargun   WeaponDetails `json:"Speargun"`
		Spear      WeaponDetails `json:"Spear"`
		Trident    WeaponDetails `json:"Trident"`
		Dagger     WeaponDetails `json:"Dagger"`
	} `json:"weapons"`
	Training []struct {
		ID       int    `json:"id"`
		Category string `json:"category"`
		Name     string `json:"name"`
		Track    []struct {
			Cost    int    `json:"cost"`
			Type    string `json:"type"`
			SkillID int    `json:"skill_id"`
		} `json:"track"`
	} `json:"training"`
	Flags  []string `json:"flags"`
	Skills []struct {
		ID   int    `json:"id"`
		Slot string `json:"slot"`
		Type string `json:"type"`
	} `json:"skills"`
}

func (a *Professions) Get() ProfessionsResponse {

	response := ProfessionsResponse{}
	a.Gw2sdk.Retrieve("professions", nil, &response)

	return response

}

func (a *Professions) GetDetails(parameters map[string]string) ProfessionsResponseDetails {

	response := ProfessionsResponseDetails{}
	a.Gw2sdk.Retrieve("professions", parameters, &response)

	return response

}
