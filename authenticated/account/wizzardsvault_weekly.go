package authenticated

import (
	Connection "gw2sdk/connection"
)

type WizzardsVaultWeekly struct {
	Gw2sdk Connection.GW2sdk
}

type ObjectivesWeekly struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Track            string `json:"track"`
	Acclaim          int    `json:"acclaim"`
	ProgressCurrent  int    `json:"progress_current"`
	ProgressComplete int    `json:"progress_complete"`
	Claimed          bool   `json:"claimed"`
}

type WizzardsVaultWeeklyResponse struct {
	MetaProgressCurrent  int                `json:"meta_progress_current"`
	MetaProgressComplete int                `json:"meta_progress_complete"`
	MetaRewardItemID     int                `json:"meta_reward_item_id"`
	MetaRewardAstral     int                `json:"meta_reward_astral"`
	MetaRewardClaimed    bool               `json:"meta_reward_claimed"`
	ObjectivesWeekly     []ObjectivesWeekly `json:"objectives"`
}

func (a *WizzardsVaultWeekly) Get() WizzardsVaultWeeklyResponse {

	response := WizzardsVaultWeeklyResponse{}
	a.Gw2sdk.Retrieve("account/wizardsvault/weekly", nil, &response)

	return response

}
