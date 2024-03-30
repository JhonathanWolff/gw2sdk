package authenticated

import (
	"fmt"
	Connection "gw2sdk/connection"
	"time"
)

type GuildLog struct {
	Gw2sdk Connection.GW2sdk
}

type GuildLogResponseDetails []struct {
	ID        int       `json:"id"`
	Time      time.Time `json:"time"`
	Type      string    `json:"type"`
	User      string    `json:"user"`
	ItemID    int       `json:"item_id,omitempty"`
	Count     int       `json:"count,omitempty"`
	Motd      string    `json:"motd,omitempty"`
	InvitedBy string    `json:"invited_by,omitempty"`
	Operation string    `json:"operation,omitempty"`
	Coins     int       `json:"coins,omitempty"`
	UpgradeID int       `json:"upgrade_id,omitempty"`
	Action    string    `json:"action,omitempty"`
}

func (a *GuildLog) GetDetails(guild_id string) GuildLogResponseDetails {

	path := fmt.Sprintf("guild/%s/log", guild_id)

	response := GuildLogResponseDetails{}
	a.Gw2sdk.Retrieve(path, nil, &response)

	return response

}
