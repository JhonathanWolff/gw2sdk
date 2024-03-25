# gw2sdk

**Unofficial Guild Wars 2 Wiki Data Extraction SDK (Go - V2 API Only)**

This project provides an unofficial Go library for extracting information from the Guild Wars 2 Wiki using **version 2 of their official REST API**.

**Important Note:**

This SDK currently only supports requests made against version 2 of the Guild Wars 2 Wiki REST API. Please refer to the official API documentation for details on V2 and any potential compatibility considerations.

**Installation**

**Usage**

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

**Available Resources**

**Achievements**
* **achievements**: Returns information about achievements.
* **achievements/daily**: Returns information about daily achievements.
* **achievements/daily/tomorrow**: Returns information about the next daily achievements.
* **achievements/groups**: Returns information about achievement groups.
* **achievements/categories**: Returns information about achievement categories.

**Authenticated**
These endpoints access account data and require an API key to view.

* **account**: Returns information about an account associated with an API key.
* **account/achievements**: Returns information about an account's achievement progress.
* **account/bank**: Returns information about a bank associated with an API key.
* **account/dailycrafting**: Returns information about (daily) crafted materials (since the most recent reset) associated with an API key.
* **account/dungeons**: Returns information about the current daily cleared dungeons associated with an API key.
* **account/dyes**: Returns information about unlocked dyes associated with an API key.
* **account/finishers**: Returns information about unlocked finishers associated with an API key.
* **account/gliders**: Returns information about unlocked gliders associated with an API key.
* **account/home/cats**: Returns information about unlocked cats in the home instance associated with an API key.
* **account/home/nodes**: Returns information about unlocked nodes in the home instance associated with an API key.
* **account/inventory**: Returns information about the shared inventory slots associated with an API key.
* **account/jadebots**: Returns information about unlocked jade bot skins associated with an API key.
* **account/luck**: Returns information about acquired luck associated with an API key.
* **account/legendaryarmory**: Returns information about the Legendary Armory associated with an API key.
* **account/mailcarriers**: Returns information about the mail carriers associated with an API key.
* **account/mapchests**: Returns information about (daily) map chest rewards received (since the most recent reset) associated with an API key.
* **account/masteries**: Returns information about unlocked masteries associated with an API key.
* **account/mastery/points**: Returns information about the total amount of mastery points associated with an API key.
* **account/materials**: Returns information about a material storage associated with an API key.
* **account/minis**: Returns information about unlocked miniatures associated with an API key.
* **account/mounts/skins**: Returns information about unlocked mount skins associated with an API key.
* **account/mounts/types**: Returns information about unlocked mounts associated with an API key.
* **account/novelties**: Returns information about unlocked novelties associated with an API key.
* **account/outfits**: Returns information about unlocked outfits associated with an API key.
* **account/progression**: Returns information about unlocked fractal account augmentation and acquired luck associated with an API key.
* **account/pvp/heroes**: Returns information about unlocked PvP heroes associated with an API key.
* **account/raids**: Returns information about completed raid events between weekly resets associated with an API key.
* **account/recipes**: Returns information about unlocked crafting recipes associated with an API key.
* **account/skiffs**: Returns information about unlocked skiff skins associated with an API key.
* **account/skins**: Returns information about unlocked skins associated with an API key.
* **account/titles**: Returns information about unlocked titles associated with an API key.
* **account/wallet**: Returns information about wealth associated with an API key.
* **account/wizardsvault/daily**: Returns information about the Daily Wizard's Vault objectives and progress associated with an API Key.
* **account/wizardsvault/listings**: Returns information about the Astral Rewards associated with an API Key.
* **account/wizardsvault/special**: Returns information about the Special Wizard's Vault objectives and progress associated with an API Key.
* **account/wizardsvault/weekly**: Returns information about the Weekly Wizard's Vault objectives and progress associated with an API Key.
* **account/worldbosses**: Returns information about (daily) Core Tyria world boss clears (since the most recent reset) associated with an API key.
* **characters**: Returns information on an account's characters.
* **commerce/transactions**: Returns information on an account's past and current trading post transactions.
* **createsubtoken**: Creates a subtoken that can be used in place of API Keys. The subtoken will have an expiry date and can be limited to a subset of the permissions granted to the API Key used during its creation as well as to a subset of accessible endpoints.
* **pvp/stats**: Returns general information on a player's performance in sPvP.
* **pvp/games**: Returns more detailed information on the player's most recent sPvP matches.
* **pvp/standings**: Returns the best and current standing of a player in sPvP leagues.
* **tokeninfo**: Returns information about the supplied API Key.

**Daily Rewards**
* **dailycrafting**: Returns information about daily craftable items.
* **mapchests**: Returns information about daily acquirable map chests.
* **worldbosses**: Returns information about daily (Core Tyria) world boss clears.

**Game Mechanics**
* **jadebots**: Returns information about jade bot skins.
* **legendaryarmory**: Returns information about the Legendary Armory
* **masteries**: Returns information about masteries.
* **mounts**: Returns information about the v2/mounts endpoints.
* **mounts/skins**: Returns information about mount skins.
* **mounts/types**: Returns information about mount types.
* **outfits**: Returns information about outfits.
* **pets**: Returns information about pets.
* **professions**: Returns information about professions.
* **races**: Returns information about particular racial skills.
* **specializations**: Returns information about specializations.
* **skiffs**: Returns information about skiff skins.
* **skills**: Returns information about skills.
* **traits**: Returns information about traits.
* **legends**: Returns information about revenant legends.

**Guild**
* **guild/**:id: Returns core details about a given guild.
* **emblem**: Returns image resources needed to render guild emblems.
* **guild/permissions**: Returns information about guild rank permissions.
* **guild/search**: Returns information on guild ids to be used for other API queries.
* **guild/upgrades**: Returns information about guild upgrades and scribe decorations.

**Guild Authenticated**

These endpoints access guild-specific data and require an API key from the guild owner to view.

* **guild/**:id/log: Returns information about a guild's event log.
* **guild/**:id/members: Returns information about members of a guild.
* **guild/**:id/ranks: Returns information about the permission ranks of a guild.
* **guild/**:id/stash: Returns information about the contents of a guild's stash.
* **guild/**:id/storage: Returns information about the contents of a guild's storage.
* **guild/**:id/treasury: Returns information about a guild's treasury contents.
* **guild/**:id/teams: Returns information about a guild's teams.
* **guild/**:id/upgrades: Returns information about a guild's upgrades.

**Home Instance**
* **home/cats**: Returns information about cats that can be unlocked in the home instance.
* **home/nodes**: Returns information about home instance upgrades.

**Items**
* **finishers**: Returns information about finishers.
* **items**: Returns information about items.
* **itemstats**: Returns information about item stats.
* **materials**: Returns information about materials.
* **pvp/amulets**: Returns information about pvp amulets.
* **recipes**: Returns information about recipes.
* **recipes/search**: A search interface for recipes.
* **skins**: Returns information about skins.

**Map information**
* **continents**: Returns a list of continents and information about each continent.
* **maps**: Returns information about maps in the game.
* Miscellaneous
* **build**: Returns the current build id.
* **colors**: Returns information about dye colors.
* **currencies**: Returns information about wallet currencies.
* **dungeons**: Returns information about each dungeons and its associated paths.
* **files**: Returns commonly requested assets.
* **quaggans**: Returns quaggan icons.
* **minis**: Returns minis.
* **novelties**: Returns novelties.
* **raids**: Returns information about each raid and its associated events.
* **titles**: Returns the list of titles.
* **worlds**: Returns world names.

**Story**
* **backstory/answers**: Returns information about Biography answers.
* **backstory/questions**: Returns information about Biography questions presented during character creation.
* **stories**: Returns information about stories in the Story Journal.
* **stories/seasons**: Returns information about the story seasons in the Story Journal.
* **quests**: Returns information about the quests constituting the stories found in the Story Journal.

**Structured PvP**
* **pvp**: Returns information about the v2/pvp endpoints.
* **pvp/ranks**: Returns information about PvP Rank.
* **pvp/seasons**: Returns information about League seasons.
* **pvp/seasons/**:id/leaderboards: Returns information about League leaderboards.

**Trading post**
* **commerce/delivery**: Returns the items and coin available for pickup for a given account.
* **commerce/exchange**: Returns a list of accepted resources for the gem exchange.
* **commerce/exchange/coins**: Returns the current coins to gems exchange rate.
* **commerce/exchange/gems**: Returns the current gems to coins exchange rate.
* **commerce/listings**: Returns trading post listings.
* **commerce/prices**: Returns buy and sell listing information.
* **commerce/transactions**: Returns the current and historical buy and sell transactions for a given account.

**Wizard's Vault**
* **wizardsvault/listings**: Returns information about listings in the Wizard's Vault.
* **wizardsvault/objectives**: Returns information about objectives in the Wizard's Vault.

**World vs. World**
* **wvw**: Returns information about the v2/wvw endpoints.
* **wvw/abilities**: Returns information about WvW abilities.
* **wvw/matches**: Returns information about current WvW matchups.
* **wvw/objectives**: Returns information about WvW objectives.
* **wvw/ranks**: Returns information about current WvW Rank.
* **wvw/upgrades**: Returns information about available upgrades for objectives.


**Contributing**

We welcome contributions to this project! 

**License**

This project is licensed under the MIT License (see `LICENSE.md` for details).

**Disclaimer**

This is an unofficial SDK and is not affiliated with Guild Wars 2 or the Guild Wars 2 Wiki. Please refer to the official API documentation for the latest information and usage guidelines, especially regarding V2: [https://wiki.guildwars2.com/wiki/API:Main](https://wiki.guildwars2.com/wiki/API:Main)
