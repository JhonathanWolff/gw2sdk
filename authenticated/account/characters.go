package authenticated

import (
	Connection "gw2sdk/connection"
	"time"
)

type Characters struct {
	Gw2sdk Connection.GW2sdk
}
type CharacterCrafting struct {
	Discipline string `json:"discipline"`
	Rating     int    `json:"rating"`
	Active     bool   `json:"active"`
}

type CharacterWvwAbilities struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}
type CharacterStatsAttributes struct {
	Power      int `json:"Power"`
	Precision  int `json:"Precision"`
	CritDamage int `json:"CritDamage"`
}
type CharacterEquipmentStats struct {
	ID         int                      `json:"id"`
	Attributes CharacterStatsAttributes `json:"attributes"`
}

type CharacterEquipmentEquipment struct {
	ID        int                     `json:"id"`
	Slot      string                  `json:"slot"`
	Binding   string                  `json:"binding,omitempty"`
	BoundTo   string                  `json:"bound_to,omitempty"`
	Infusions []int                   `json:"infusions,omitempty"`
	Stats     CharacterEquipmentStats `json:"stats,omitempty"`
	Dyes      []any                   `json:"dyes,omitempty"`
	Upgrades  []int                   `json:"upgrades,omitempty"`
	Skin      int                     `json:"skin,omitempty"`
}

type CharacterBagsInventory struct {
	ID      int    `json:"id"`
	Count   int    `json:"count"`
	Binding string `json:"binding,omitempty"`
}
type CharacterBags struct {
	ID        int                      `json:"id"`
	Size      int                      `json:"size"`
	Inventory []CharacterBagsInventory `json:"inventory"`
}
type CharacterEquipmentPvp struct {
	Amulet int   `json:"amulet"`
	Rune   int   `json:"rune"`
	Sigils []int `json:"sigils"`
}
type CharacterSpecialization struct {
	ID     int   `json:"id"`
	Traits []int `json:"traits"`
}

type CharacterSkills struct {
	Heal      int      `json:"heal"`
	Utilities []int    `json:"utilities"`
	Elite     int      `json:"elite"`
	Legends   []string `json:"legends"`
}

type CharactersResponse []struct {
	Name            string                        `json:"name"`
	Race            string                        `json:"race"`
	Gender          string                        `json:"gender"`
	Flags           []any                         `json:"flags"`
	Profession      string                        `json:"profession"`
	Level           int                           `json:"level"`
	Guild           string                        `json:"guild"`
	Age             int                           `json:"age"`
	Created         time.Time                     `json:"created"`
	Deaths          int                           `json:"deaths"`
	Crafting        []CharacterCrafting           `json:"crafting"`
	Title           int                           `json:"title,omitempty"`
	Backstory       []string                      `json:"backstory"`
	WvwAbilities    []CharacterWvwAbilities       `json:"wvw_abilities"`
	Equipment       []CharacterEquipmentEquipment `json:"equipment"`
	Recipes         []int                         `json:"recipes"`
	Training        []any                         `json:"training"`
	Bags            []CharacterBags               `json:"bags"`
	EquipmentPvp    CharacterEquipmentPvp         `json:"equipment_pvp"`
	Specializations struct {
		Pve []CharacterSpecialization `json:"pve"`
		Pvp []CharacterSpecialization `json:"pvp"`
		Wvw []CharacterSpecialization `json:"wvw"`
	} `json:"specializations"`
	Skills struct {
		Pve CharacterSkills `json:"pve"`
		Pvp CharacterSkills `json:"pvp"`
		Wvw CharacterSkills `json:"wvw"`
	} `json:"skills"`
}

func (a *Characters) Get() CharactersResponse {

	response := CharactersResponse{}
	a.Gw2sdk.Retrieve("characters?ids=all", nil, &response)

	return response

}
