package authenticated

import (
	Connection "gw2sdk/connection"
)

type WizzardsVaultDaily struct {
	Gw2sdk Connection.GW2sdk
}

type Objectives struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Track            string `json:"track"`
	Acclaim          int    `json:"acclaim"`
	ProgressCurrent  int    `json:"progress_current"`
	ProgressComplete int    `json:"progress_complete"`
	Claimed          bool   `json:"claimed"`
}

type WizzardsVaultDailyResponse struct {
	MetaProgressCurrent  int          `json:"meta_progress_current"`
	MetaProgressComplete int          `json:"meta_progress_complete"`
	MetaRewardItemID     int          `json:"meta_reward_item_id"`
	MetaRewardAstral     int          `json:"meta_reward_astral"`
	MetaRewardClaimed    bool         `json:"meta_reward_claimed"`
	Objectives           []Objectives `json:"objectives"`
}

func (a *WizzardsVaultDaily) Get() WizzardsVaultDailyResponse {

	response := WizzardsVaultDailyResponse{}
	a.Gw2sdk.Retrieve("account/wizardsvault/daily", nil, &response)

	return response

}
