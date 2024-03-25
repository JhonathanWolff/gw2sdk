package achivements

import Connection "gw2sdk/connection"

type AchivementDailyTomorrow struct {
	Gw2sdk Connection.GW2sdk
}

// https://wiki.guildwars2.com/wiki/API:2/achievements/daily/tomorrow
func (a *AchivementDailyTomorrow) Get(parameters map[string]string) interface{} {

	response := AchivementResponseDetails{}
	a.Gw2sdk.Retrieve("achievements/daily/tomorrow", parameters, &response)
	return response

}
