# gw2sdk

**Unofficial Guild Wars 2 Wiki Data Extraction SDK (Go - V2 API Only)**

This project provides an unofficial Go library for extracting information from the Guild Wars 2 Wiki using **version 2 of their official REST API**.

**Important Note:**

This SDK currently only supports requests made against version 2 of the Guild Wars 2 Wiki REST API. Please refer to the official API documentation for details on V2 and any potential compatibility considerations.


## Contributing

We welcome contributions to this project! 

## License

This project is licensed under the MIT License (see `LICENSE.md` for details).

## Disclaimer

This repository was created for practicing Golang. It implements the methods of the Guild Wars 2 API /V2:

* https://wiki.guildwars2.com/wiki/API:Main

However, the following sub-endpoints were not implemented:

* /v2/emblem/backgrounds
* /v2/emblem/foregrounds

And other sub-endpoints.

These sub-endpoints were not mentioned on the main page, and since this project is for study purposes, I did not implement them.
Please note that **contributions are welcome to update this repository with new resources or to fix unforeseen bugs**.


This is an unofficial SDK and is not affiliated with Guild Wars 2 or the Guild Wars 2 Wiki. Please refer to the official API documentation for the latest information and usage guidelines, especially regarding V2: [https://wiki.guildwars2.com/wiki/API:Main](https://wiki.guildwars2.com/wiki/API:Main)


## Installation and Usage


1. Import the package:

   ```go
   import "github.com/JhonathanWolff/gw2sdk-sdk"
   ```

2. (Optional) Set your API key (for authenticated requests):

   ```go
	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	var achivement = achivements.Category{Gw2sdk: gw2sdk}
   ```

3. Use the provided functions to interact with the **version 2** Guild Wars 2 Wiki API:

   ```go
   // Get Achivements
	var achivement = achivements.Category{Gw2sdk: gw2sdk}

   ```

## Available methods

each struct for each resource has one of those 3 methods available to use


| Endpoint Methods | Args              | description                                                                                                                                                                                             |
|------------------|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Get              | null              | simple request without any paramter                                                                                                                                                                     |
| GetDetails       | string \| int     | updates the path based on the arg                                                                                                                                                                       |
| GetDetails       | map[string]string | a list of parameters to add to the request, note : every method that use this arg must have the paramter key "ids" and  the value must be a list of values splited with "," exmaple "ids" : "1,2,3,4,5" |


## Parallalel request
If you want to run a parallel request 
its possible to run using the parallel methods inside the sdk

```go


//example for simple request

import 	"gw2sdk/parallel"

func parallel_request() {

	var gw2sdk connection.GW2sdk = connection.GW2sdk{Token: os.Getenv("GW2_TOKEN")}
	var array_request []parallel.GetRequest
	for i := 0; i < 10; i++ {
		client := unauthenticated.Achivement{Gw2sdk: gw2sdk}
		array_request = append(array_request, &client)
	}
	array_response := parallel.ParallelGet(array_request)

	for _, resp := range array_response {
		fmt.Println(string(resp))
	}

}

its possible to run this method for `GetDetails (value string)` and `GetDetails(parameters map[string]string)`



* `GetDetails (value string)` --> `ParallelGetDetailsPath`
* `GetDetails(parameters map[string]string)` --> `ParallelGetDetailsParameters`


```

# Available Resources

## Achievements
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [achievements](https://wiki.guildwars2.com/wiki/API:2/achievements) | Returns information about achievements. | Achievements | :heavy_check_mark: |
| [achievements/daily](https://wiki.guildwars2.com/wiki/API:2/achievements/daily) | Returns information about daily achievements. | AchievementsDaily | :heavy_check_mark: |
| [achievements/daily/tomorrow](https://wiki.guildwars2.com/wiki/API:2/achievements/daily/tomorrow) | Returns information about the next daily achievements. | AchievementsDailyTomorrow | :heavy_check_mark: |
| [achievements/groups](https://wiki.guildwars2.com/wiki/API:2/achievements/groups) | Returns information about achievement groups. | AchievementsGroups | :heavy_check_mark: |
| [achievements/categories](https://wiki.guildwars2.com/wiki/API:2/achievements/categories) | Returns information about achievement categories. | AchievementsCategories | :heavy_check_mark: |
## Authenticated
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [account](https://wiki.guildwars2.com/wiki/API:2/account) | Returns information about an account associated with an API key. | Account | :heavy_check_mark: |
| [account/achievements](https://wiki.guildwars2.com/wiki/API:2/account/achievements) | Returns information about an account's achievement progress. | AccountAchievements | :heavy_check_mark: |
| [account/bank](https://wiki.guildwars2.com/wiki/API:2/account/bank) | Returns information about a bank associated with an API key. | AccountBank | :heavy_check_mark: |
| [account/dailycrafting](https://wiki.guildwars2.com/wiki/API:2/account/dailycrafting) | Returns information about (daily) crafted materials (since the most recent reset) associated with an API key. | AccountDailycrafting | :heavy_check_mark: |
| [account/dungeons](https://wiki.guildwars2.com/wiki/API:2/account/dungeons) | Returns information about the current daily cleared dungeons associated with an API key. | AccountDungeons | :heavy_check_mark: |
| [account/dyes](https://wiki.guildwars2.com/wiki/API:2/account/dyes) | Returns information about unlocked dyes associated with an API key. | AccountDyes | :heavy_check_mark: |
| [account/finishers](https://wiki.guildwars2.com/wiki/API:2/account/finishers) | Returns information about unlocked finishers associated with an API key. | AccountFinishers | :heavy_check_mark: |
| [account/gliders](https://wiki.guildwars2.com/wiki/API:2/account/gliders) | Returns information about unlocked gliders associated with an API key. | AccountGliders | :heavy_check_mark: |
| [account/home/cats](https://wiki.guildwars2.com/wiki/API:2/account/home/cats) | Returns information about unlocked cats in the home instance associated with an API key. | AccountHomeCats | :heavy_check_mark: |
| [account/home/nodes](https://wiki.guildwars2.com/wiki/API:2/account/home/nodes) | Returns information about unlocked nodes in the home instance associated with an API key. | AccountHomeNodes | :heavy_check_mark: |
| [account/inventory](https://wiki.guildwars2.com/wiki/API:2/account/inventory) | Returns information about the shared inventory slots associated with an API key. | AccountInventory | :heavy_check_mark: |
| [account/jadebots](https://wiki.guildwars2.com/wiki/API:2/account/jadebots) | Returns information about unlocked jade bot skins associated with an API key. | AccountJadebots | :heavy_check_mark: |
| [account/luck](https://wiki.guildwars2.com/wiki/API:2/account/luck) | Returns information about acquired luck associated with an API key. | AccountLuck | :heavy_check_mark: |
| [account/legendaryarmory](https://wiki.guildwars2.com/wiki/API:2/account/legendaryarmory) | Returns information about the Legendary Armory associated with an API key. | AccountLegendaryarmory | :heavy_check_mark: |
| [account/mailcarriers](https://wiki.guildwars2.com/wiki/API:2/account/mailcarriers) | Returns information about the mail carriers associated with an API key. | AccountMailcarriers | :heavy_check_mark: |
| [account/mapchests](https://wiki.guildwars2.com/wiki/API:2/account/mapchests) | Returns information about (daily) map chest rewards received (since the most recent reset) associated with an API key. | AccountMapchests | :heavy_check_mark: |
| [account/masteries](https://wiki.guildwars2.com/wiki/API:2/account/masteries) | Returns information about unlocked masteries associated with an API key. | AccountMasteries | :heavy_check_mark: |
| [account/mastery/points](https://wiki.guildwars2.com/wiki/API:2/account/mastery/points) | Returns information about the total amount of mastery points associated with an API key. | AccountMasteryPoints | :heavy_check_mark: |
| [account/materials](https://wiki.guildwars2.com/wiki/API:2/account/materials) | Returns information about a material storage associated with an API key. | AccountMaterials | :heavy_check_mark: |
| [account/minis](https://wiki.guildwars2.com/wiki/API:2/account/minis) | Returns information about unlocked miniatures associated with an API key. | AccountMinis | :heavy_check_mark: |
| [account/mounts/skins](https://wiki.guildwars2.com/wiki/API:2/account/mounts/skins) | Returns information about unlocked mount skins associated with an API key. | AccountMountsSkins | :heavy_check_mark: |
| [account/mounts/types](https://wiki.guildwars2.com/wiki/API:2/account/mounts/types) | Returns information about unlocked mounts associated with an API key. | AccountMountsTypes | :heavy_check_mark: |
| [account/novelties](https://wiki.guildwars2.com/wiki/API:2/account/novelties) | Returns information about unlocked novelties associated with an API key. | AccountNovelties | :heavy_check_mark: |
| [account/outfits](https://wiki.guildwars2.com/wiki/API:2/account/outfits) | Returns information about unlocked outfits associated with an API key. | AccountOutfits | :heavy_check_mark: |
| [account/progression](https://wiki.guildwars2.com/wiki/API:2/account/progression) | Returns information about unlocked fractal account augmentation and acquired luck associated with an API key. | AccountProgression | :heavy_check_mark: |
| [account/pvp/heroes](https://wiki.guildwars2.com/wiki/API:2/account/pvp/heroes) | Returns information about unlocked PvP heroes associated with an API key. | AccountPvpHeroes | :heavy_check_mark: |
| [account/raids](https://wiki.guildwars2.com/wiki/API:2/account/raids) | Returns information about completed raid events between weekly resets associated with an API key. | AccountRaids | :heavy_check_mark: |
| [account/recipes](https://wiki.guildwars2.com/wiki/API:2/account/recipes) | Returns information about unlocked crafting recipes associated with an API key. | AccountRecipes | :heavy_check_mark: |
| [account/skiffs](https://wiki.guildwars2.com/wiki/API:2/account/skiffs) | Returns information about unlocked skiff skins associated with an API key. | AccountSkiffs | :heavy_check_mark: |
| [account/skins](https://wiki.guildwars2.com/wiki/API:2/account/skins) | Returns information about unlocked skins associated with an API key. | AccountSkins | :heavy_check_mark: |
| [account/titles](https://wiki.guildwars2.com/wiki/API:2/account/titles) | Returns information about unlocked titles associated with an API key. | AccountTitles | :heavy_check_mark: |
| [account/wallet](https://wiki.guildwars2.com/wiki/API:2/account/wallet) | Returns information about wealth associated with an API key. | AccountWallet | :heavy_check_mark: |
| [account/wizardsvault/daily](https://wiki.guildwars2.com/wiki/API:2/account/wizardsvault/daily) | Returns information about the Daily Wizard's Vault objectives and progress associated with an API Key. | AccountWizardsvaultDaily | :heavy_check_mark: |
| [account/wizardsvault/listings](https://wiki.guildwars2.com/wiki/API:2/account/wizardsvault/listings) | Returns information about the Astral Rewards associated with an API Key. | AccountWizardsvaultListings | :heavy_check_mark: |
| [account/wizardsvault/special](https://wiki.guildwars2.com/wiki/API:2/account/wizardsvault/special) | Returns information about the Special Wizard's Vault objectives and progress associated with an API Key. | AccountWizardsvaultSpecial | :heavy_check_mark: |
| [account/wizardsvault/weekly](https://wiki.guildwars2.com/wiki/API:2/account/wizardsvault/weekly) | Returns information about the Weekly Wizard's Vault objectives and progress associated with an API Key. | AccountWizardsvaultWeekly | :heavy_check_mark: |
| [account/worldbosses](https://wiki.guildwars2.com/wiki/API:2/account/worldbosses) | Returns information about (daily) Core Tyria world boss clears (since the most recent reset) associated with an API key. | AccountWorldbosses | :heavy_check_mark: |
| [characters](https://wiki.guildwars2.com/wiki/API:2/characters) | Returns information on an account's characters. | Characters | :heavy_check_mark: |
| [commerce/transactions](https://wiki.guildwars2.com/wiki/API:2/commerce/transactions) | Returns information on an account's past and current trading post transactions. | CommerceTransactions | :heavy_check_mark: |
| [createsubtoken](https://wiki.guildwars2.com/wiki/API:2/createsubtoken) | Creates a subtoken that can be used in place of API Keys. The subtoken will have an expiry date and can be limited to a subset of the permissions granted to the API Key used during its creation as well as to a subset of accessible endpoints. | Createsubtoken | :heavy_check_mark: |
| [pvp/stats](https://wiki.guildwars2.com/wiki/API:2/pvp/stats) | Returns general information on a player's performance in sPvP. | PvpStats | :heavy_check_mark: |
| [pvp/games](https://wiki.guildwars2.com/wiki/API:2/pvp/games) | Returns more detailed information on the player's most recent sPvP matches. | PvpGames | :heavy_check_mark: |
| [pvp/standings](https://wiki.guildwars2.com/wiki/API:2/pvp/standings) | Returns the best and current standing of a player in sPvP leagues. | PvpStandings | :heavy_check_mark: |
| [tokeninfo](https://wiki.guildwars2.com/wiki/API:2/tokeninfo) | Returns information about the supplied API Key. | Tokeninfo | :heavy_check_mark: |
## Daily Rewards
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [dailycrafting](https://wiki.guildwars2.com/wiki/API:2/dailycrafting) | Returns information about daily craftable items. | Dailycrafting | :heavy_check_mark: |
| [mapchests](https://wiki.guildwars2.com/wiki/API:2/mapchests) | Returns information about daily acquirable map chests. | Mapchests | :heavy_check_mark: |
| [worldbosses](https://wiki.guildwars2.com/wiki/API:2/worldbosses) | Returns information about daily (Core Tyria) world boss clears. | Worldbosses | :heavy_check_mark: |
## Game Mechanics
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [jadebots](https://wiki.guildwars2.com/wiki/API:2/jadebots) | Returns information about jade bot skins. | Jadebots | :heavy_check_mark: |
| [legendaryarmory](https://wiki.guildwars2.com/wiki/API:2/legendaryarmory) | Returns information about the Legendary Armory | Legendaryarmory | :heavy_check_mark: |
| [masteries](https://wiki.guildwars2.com/wiki/API:2/masteries) | Returns information about masteries. | Masteries | :heavy_check_mark: |
| [mounts](https://wiki.guildwars2.com/wiki/API:2/mounts) | Returns information about the v2/mounts endpoints. | Mounts | :heavy_check_mark: |
| [mounts/skins](https://wiki.guildwars2.com/wiki/API:2/mounts/skins) | Returns information about mount skins. | MountsSkins | :heavy_check_mark: |
| [mounts/types](https://wiki.guildwars2.com/wiki/API:2/mounts/types) | Returns information about mount types. | MountsTypes | :heavy_check_mark: |
| [outfits](https://wiki.guildwars2.com/wiki/API:2/outfits) | Returns information about outfits. | Outfits | :heavy_check_mark: |
| [pets](https://wiki.guildwars2.com/wiki/API:2/pets) | Returns information about pets. | Pets | :heavy_check_mark: |
| [professions](https://wiki.guildwars2.com/wiki/API:2/professions) | Returns information about professions. | Professions | :heavy_check_mark: |
| [races](https://wiki.guildwars2.com/wiki/API:2/races) | Returns information about particular racial skills. | Races | :heavy_check_mark: |
| [specializations](https://wiki.guildwars2.com/wiki/API:2/specializations) | Returns information about specializations. | Specializations | :heavy_check_mark: |
| [skiffs](https://wiki.guildwars2.com/wiki/API:2/skiffs) | Returns information about skiff skins. | Skiffs | :heavy_check_mark: |
| [skills](https://wiki.guildwars2.com/wiki/API:2/skills) | Returns information about skills. | Skills | :heavy_check_mark: |
| [traits](https://wiki.guildwars2.com/wiki/API:2/traits) | Returns information about traits. | Traits | :heavy_check_mark: |
| [legends](https://wiki.guildwars2.com/wiki/API:2/legends) | Returns information about revenant legends. | Legends | :heavy_check_mark: |
## Guild
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [guild/:id](https://wiki.guildwars2.com/wiki/API:2/guild/:id) | Returns core details about a given guild. | Guild:id | :heavy_check_mark: |
| [emblem](https://wiki.guildwars2.com/wiki/API:2/emblem) | Returns image resources needed to render guild emblems. | Emblem | :heavy_check_mark: |
| [guild/permissions](https://wiki.guildwars2.com/wiki/API:2/guild/permissions) | Returns information about guild rank permissions. | GuildPermissions | :heavy_check_mark: |
| [guild/search](https://wiki.guildwars2.com/wiki/API:2/guild/search) | Returns information on guild ids to be used for other API queries. | GuildSearch | :heavy_check_mark: |
| [guild/upgrades](https://wiki.guildwars2.com/wiki/API:2/guild/upgrades) | Returns information about guild upgrades and scribe decorations. | GuildUpgrades | :heavy_check_mark: |
## Guild Authenticated
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [guild/:id/log](https://wiki.guildwars2.com/wiki/API:2/guild/:id/log) | Returns information about a guild's event log. | Guild:idLog | :heavy_check_mark: |
| [guild/:id/members](https://wiki.guildwars2.com/wiki/API:2/guild/:id/members) | Returns information about members of a guild. | Guild:idMembers | :heavy_check_mark: |
| [guild/:id/ranks](https://wiki.guildwars2.com/wiki/API:2/guild/:id/ranks) | Returns information about the permission ranks of a guild. | Guild:idRanks | :heavy_check_mark: |
| [guild/:id/stash](https://wiki.guildwars2.com/wiki/API:2/guild/:id/stash) | Returns information about the contents of a guild's stash. | Guild:idStash | :heavy_check_mark: |
| [guild/:id/storage](https://wiki.guildwars2.com/wiki/API:2/guild/:id/storage) | Returns information about the contents of a guild's storage. | Guild:idStorage | :heavy_check_mark: |
| [guild/:id/treasury](https://wiki.guildwars2.com/wiki/API:2/guild/:id/treasury) | Returns information about a guild's treasury contents. | Guild:idTreasury | :heavy_check_mark: |
| [guild/:id/teams](https://wiki.guildwars2.com/wiki/API:2/guild/:id/teams) | Returns information about a guild's teams. | Guild:idTeams | :heavy_check_mark: |
| [guild/:id/upgrades](https://wiki.guildwars2.com/wiki/API:2/guild/:id/upgrades) | Returns information about a guild's upgrades. | Guild:idUpgrades | :heavy_check_mark: |
## Home Instance
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [home/cats](https://wiki.guildwars2.com/wiki/API:2/home/cats) | Returns information about cats that can be unlocked in the home instance. | HomeCats | :x: |
| [home/nodes](https://wiki.guildwars2.com/wiki/API:2/home/nodes) | Returns information about home instance upgrades. | HomeNodes | :x: |
## Items
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [finishers](https://wiki.guildwars2.com/wiki/API:2/finishers) | Returns information about finishers. | Finishers | :x: |
| [items](https://wiki.guildwars2.com/wiki/API:2/items) | Returns information about items. | Items | :x: |
| [itemstats](https://wiki.guildwars2.com/wiki/API:2/itemstats) | Returns information about item stats. | Itemstats | :x: |
| [materials](https://wiki.guildwars2.com/wiki/API:2/materials) | Returns information about materials. | Materials | :x: |
| [pvp/amulets](https://wiki.guildwars2.com/wiki/API:2/pvp/amulets) | Returns information about pvp amulets. | PvpAmulets | :x: |
| [recipes](https://wiki.guildwars2.com/wiki/API:2/recipes) | Returns information about recipes. | Recipes | :x: |
| [recipes/search](https://wiki.guildwars2.com/wiki/API:2/recipes/search) | A search interface for recipes. | RecipesSearch | :x: |
| [skins](https://wiki.guildwars2.com/wiki/API:2/skins) | Returns information about skins. | Skins | :x: |
## Map information
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [continents](https://wiki.guildwars2.com/wiki/API:2/continents) | Returns a list of continents and information about each continent. | Continents | :x: |
| [maps](https://wiki.guildwars2.com/wiki/API:2/maps) | Returns information about maps in the game. | Maps | :x: |
## Miscellaneous
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [build](https://wiki.guildwars2.com/wiki/API:2/build) | Returns the current build id. | Build | :x: |
| [colors](https://wiki.guildwars2.com/wiki/API:2/colors) | Returns information about dye colors. | Colors | :x: |
| [currencies](https://wiki.guildwars2.com/wiki/API:2/currencies) | Returns information about wallet currencies. | Currencies | :x: |
| [dungeons](https://wiki.guildwars2.com/wiki/API:2/dungeons) | Returns information about each dungeons and its associated paths. | Dungeons | :x: |
| [files](https://wiki.guildwars2.com/wiki/API:2/files) | Returns commonly requested assets. | Files | :x: |
| [quaggans](https://wiki.guildwars2.com/wiki/API:2/quaggans) | Returns quaggan icons. | Quaggans | :x: |
| [minis](https://wiki.guildwars2.com/wiki/API:2/minis) | Returns minis. | Minis | :x: |
| [novelties](https://wiki.guildwars2.com/wiki/API:2/novelties) | Returns novelties. | Novelties | :x: |
| [raids](https://wiki.guildwars2.com/wiki/API:2/raids) | Returns information about each raid and its associated events. | Raids | :x: |
| [titles](https://wiki.guildwars2.com/wiki/API:2/titles) | Returns the list of titles. | Titles | :x: |
| [worlds](https://wiki.guildwars2.com/wiki/API:2/worlds) | Returns world names. | Worlds | :x: |
## Story
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [backstory/answers](https://wiki.guildwars2.com/wiki/API:2/backstory/answers) | Returns information about Biography answers. | BackstoryAnswers | :x: |
| [backstory/questions](https://wiki.guildwars2.com/wiki/API:2/backstory/questions) | Returns information about Biography questions presented during character creation. | BackstoryQuestions | :x: |
| [stories](https://wiki.guildwars2.com/wiki/API:2/stories) | Returns information about stories in the Story Journal. | Stories | :x: |
| [stories/seasons](https://wiki.guildwars2.com/wiki/API:2/stories/seasons) | Returns information about the story seasons in the Story Journal. | StoriesSeasons | :x: |
| [quests](https://wiki.guildwars2.com/wiki/API:2/quests) | Returns information about the quests constituting the stories found in the Story Journal. | Quests | :x: |
## Structured PvP
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [pvp](https://wiki.guildwars2.com/wiki/API:2/pvp) | Returns information about the v2/pvp endpoints. | Pvp | :x: |
| [pvp/ranks](https://wiki.guildwars2.com/wiki/API:2/pvp/ranks) | Returns information about PvP Rank. | PvpRanks | :x: |
| [pvp/seasons](https://wiki.guildwars2.com/wiki/API:2/pvp/seasons) | Returns information about League seasons. | PvpSeasons | :x: |
| [pvp/seasons/:id/leaderboards](https://wiki.guildwars2.com/wiki/API:2/pvp/seasons/:id/leaderboards) | Returns information about League leaderboards. | PvpSeasons:idLeaderboards | :x: |
## Trading post
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [commerce/delivery](https://wiki.guildwars2.com/wiki/API:2/commerce/delivery) | Returns the items and coin available for pickup for a given account. | CommerceDelivery | :x: |
| [commerce/exchange](https://wiki.guildwars2.com/wiki/API:2/commerce/exchange) | Returns a list of accepted resources for the gem exchange. | CommerceExchange | :x: |
| [commerce/exchange/coins](https://wiki.guildwars2.com/wiki/API:2/commerce/exchange/coins) | Returns the current coins to gems exchange rate. | CommerceExchangeCoins | :x: |
| [commerce/exchange/gems](https://wiki.guildwars2.com/wiki/API:2/commerce/exchange/gems) | Returns the current gems to coins exchange rate. | CommerceExchangeGems | :x: |
| [commerce/listings](https://wiki.guildwars2.com/wiki/API:2/commerce/listings) | Returns trading post listings. | CommerceListings | :x: |
| [commerce/prices](https://wiki.guildwars2.com/wiki/API:2/commerce/prices) | Returns buy and sell listing information. | CommercePrices | :x: |
| [commerce/transactions](https://wiki.guildwars2.com/wiki/API:2/commerce/transactions) | Returns the current and historical buy and sell transactions for a given account. | CommerceTransactions | :x: |
## Wizard's Vault
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [wizardsvault/listings](https://wiki.guildwars2.com/wiki/API:2/wizardsvault/listings) | Returns information about listings in the Wizard's Vault. | WizardsvaultListings | :x: |
| [wizardsvault/objectives](https://wiki.guildwars2.com/wiki/API:2/wizardsvault/objectives) | Returns information about objectives in the Wizard's Vault. | WizardsvaultObjectives | :x: |
## World vs. World
| endpoint | description | Api equivalent | implemented |
| --- | --- | --- | --- |
| [wvw](https://wiki.guildwars2.com/wiki/API:2/wvw) | Returns information about the v2/wvw endpoints. | Wvw | :x: |
| [wvw/abilities](https://wiki.guildwars2.com/wiki/API:2/wvw/abilities) | Returns information about WvW abilities. | WvwAbilities | :x: |
| [wvw/matches](https://wiki.guildwars2.com/wiki/API:2/wvw/matches) | Returns information about current WvW matchups. | WvwMatches | :x: |
| [wvw/objectives](https://wiki.guildwars2.com/wiki/API:2/wvw/objectives) | Returns information about WvW objectives. | WvwObjectives | :x: |
| [wvw/ranks](https://wiki.guildwars2.com/wiki/API:2/wvw/ranks) | Returns information about current WvW Rank. | WvwRanks | :x: |
| [wvw/upgrades](https://wiki.guildwars2.com/wiki/API:2/wvw/upgrades) | Returns information about available upgrades for objectives. | WvwUpgrades | :x: |

