package model

type Lengths struct {
	Header              uint32
	File                uint32
	MetadataOffset      uint32
	Metadata            uint32
	PayloadHeaderOffset uint32
	PayloadHeader       uint32
	PayloadOffset       uint32
}

type Rofl struct {
	Magic         [6]byte
	Signature     [256]byte
	Lengths       Lengths
	Metadata      MetadataJson
	PayloadHeader struct {
		GameId              uint64
		GameLength          uint32
		KeyframeCount       uint32
		ChunkCount          uint32
		EndStartupChunkId   uint32
		StartGameChunkId    uint32
		KeyframeInterval    uint32
		EncryptionKeyLength uint16
		EncryptionKey       string
	}
	headerEnd int64
	Chunks    []Chunk
	Keyframes []Keyframe
}

type MetadataJson struct {
	GameLength      int    `json:"gameLength"`
	GameVersion     string `json:"gameVersion"`
	LastGameChunkID int    `json:"lastGameChunkId"`
	LastKeyFrameID  int    `json:"lastKeyFrameId"`
	// statsJSON is like a buffer, json string is loaded into it first during reads, and the contents are written during writes
	StatsJSON string                   `json:"statsJson"`
	Stats     []map[string]interface{} `json:"stats,omitempty"`
}

type PlayerStatsJson struct {
	Assists                                     string `json:"ASSISTS"`
	BaronKills                                  string `json:"BARON_KILLS"`
	BarracksKilled                              string `json:"BARRACKS_KILLED"`
	BarracksTakedowns                           string `json:"BARRACKS_TAKEDOWNS"`
	BountyLevel                                 string `json:"BOUNTY_LEVEL"`
	ChampionsKilled                             string `json:"CHAMPIONS_KILLED"`
	ChampionMissionStat0                        string `json:"CHAMPION_MISSION_STAT_0"`
	ChampionMissionStat1                        string `json:"CHAMPION_MISSION_STAT_1"`
	ChampionMissionStat2                        string `json:"CHAMPION_MISSION_STAT_2"`
	ChampionMissionStat3                        string `json:"CHAMPION_MISSION_STAT_3"`
	ChampionTransform                           string `json:"CHAMPION_TRANSFORM"`
	ConsumablesPurchased                        string `json:"CONSUMABLES_PURCHASED"`
	DoubleKills                                 string `json:"DOUBLE_KILLS"`
	DragonKills                                 string `json:"DRAGON_KILLS"`
	Exp                                         string `json:"EXP"`
	FriendlyDampenLost                          string `json:"FRIENDLY_DAMPEN_LOST"`
	FriendlyHqLost                              string `json:"FRIENDLY_HQ_LOST"`
	FriendlyTurretLost                          string `json:"FRIENDLY_TURRET_LOST"`
	GameEndedInEarlySurrender                   string `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
	GameEndedInSurrender                        string `json:"GAME_ENDED_IN_SURRENDER"`
	GoldEarned                                  string `json:"GOLD_EARNED"`
	GoldSpent                                   string `json:"GOLD_SPENT"`
	HqKilled                                    string `json:"HQ_KILLED"`
	HqTakedowns                                 string `json:"HQ_TAKEDOWNS"`
	ID                                          string `json:"ID"`
	IndividualPosition                          string `json:"INDIVIDUAL_POSITION"`
	Item0                                       string `json:"ITEM0"`
	Item1                                       string `json:"ITEM1"`
	Item2                                       string `json:"ITEM2"`
	Item3                                       string `json:"ITEM3"`
	Item4                                       string `json:"ITEM4"`
	Item5                                       string `json:"ITEM5"`
	Item6                                       string `json:"ITEM6"`
	ItemsPurchased                              string `json:"ITEMS_PURCHASED"`
	KeystoneID                                  string `json:"KEYSTONE_ID"`
	KillingSprees                               string `json:"KILLING_SPREES"`
	LargestCriticalStrike                       string `json:"LARGEST_CRITICAL_STRIKE"`
	LargestKillingSpree                         string `json:"LARGEST_KILLING_SPREE"`
	LargestMultiKill                            string `json:"LARGEST_MULTI_KILL"`
	Level                                       string `json:"LEVEL"`
	LongestTimeSpentLiving                      string `json:"LONGEST_TIME_SPENT_LIVING"`
	MagicDamageDealtPlayer                      string `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
	MagicDamageDealtToChampions                 string `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
	MagicDamageTaken                            string `json:"MAGIC_DAMAGE_TAKEN"`
	MinionsKilled                               string `json:"MINIONS_KILLED"`
	MutedAll                                    string `json:"MUTED_ALL"`
	Name                                        string `json:"NAME"`
	NeutralMinionsKilled                        string `json:"NEUTRAL_MINIONS_KILLED"`
	NeutralMinionsKilledEnemyJungle             string `json:"NEUTRAL_MINIONS_KILLED_ENEMY_JUNGLE"`
	NeutralMinionsKilledYourJungle              string `json:"NEUTRAL_MINIONS_KILLED_YOUR_JUNGLE"`
	NodeCapture                                 string `json:"NODE_CAPTURE"`
	NodeCaptureAssist                           string `json:"NODE_CAPTURE_ASSIST"`
	NodeNeutralize                              string `json:"NODE_NEUTRALIZE"`
	NodeNeutralizeAssist                        string `json:"NODE_NEUTRALIZE_ASSIST"`
	NumDeaths                                   string `json:"NUM_DEATHS"`
	ObjectivesStolen                            string `json:"OBJECTIVES_STOLEN"`
	ObjectivesStolenAssists                     string `json:"OBJECTIVES_STOLEN_ASSISTS"`
	PentaKills                                  string `json:"PENTA_KILLS"`
	Perk0                                       string `json:"PERK0"`
	Perk0Var1                                   string `json:"PERK0_VAR1"`
	Perk0Var2                                   string `json:"PERK0_VAR2"`
	Perk0Var3                                   string `json:"PERK0_VAR3"`
	Perk1                                       string `json:"PERK1"`
	Perk1Var1                                   string `json:"PERK1_VAR1"`
	Perk1Var2                                   string `json:"PERK1_VAR2"`
	Perk1Var3                                   string `json:"PERK1_VAR3"`
	Perk2                                       string `json:"PERK2"`
	Perk2Var1                                   string `json:"PERK2_VAR1"`
	Perk2Var2                                   string `json:"PERK2_VAR2"`
	Perk2Var3                                   string `json:"PERK2_VAR3"`
	Perk3                                       string `json:"PERK3"`
	Perk3Var1                                   string `json:"PERK3_VAR1"`
	Perk3Var2                                   string `json:"PERK3_VAR2"`
	Perk3Var3                                   string `json:"PERK3_VAR3"`
	Perk4                                       string `json:"PERK4"`
	Perk4Var1                                   string `json:"PERK4_VAR1"`
	Perk4Var2                                   string `json:"PERK4_VAR2"`
	Perk4Var3                                   string `json:"PERK4_VAR3"`
	Perk5                                       string `json:"PERK5"`
	Perk5Var1                                   string `json:"PERK5_VAR1"`
	Perk5Var2                                   string `json:"PERK5_VAR2"`
	Perk5Var3                                   string `json:"PERK5_VAR3"`
	PerkPrimaryStyle                            string `json:"PERK_PRIMARY_STYLE"`
	PerkSubStyle                                string `json:"PERK_SUB_STYLE"`
	PhysicalDamageDealtPlayer                   string `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
	PhysicalDamageDealtToChampions              string `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	PhysicalDamageTaken                         string `json:"PHYSICAL_DAMAGE_TAKEN"`
	Ping                                        string `json:"PING"`
	PlayersIMuted                               string `json:"PLAYERS_I_MUTED"`
	PlayersThatMutedMe                          string `json:"PLAYERS_THAT_MUTED_ME"`
	PlayerPosition                              string `json:"PLAYER_POSITION"`
	PlayerRole                                  string `json:"PLAYER_ROLE"`
	PlayerScore0                                string `json:"PLAYER_SCORE_0"`
	PlayerScore1                                string `json:"PLAYER_SCORE_1"`
	PlayerScore10                               string `json:"PLAYER_SCORE_10"`
	PlayerScore11                               string `json:"PLAYER_SCORE_11"`
	PlayerScore2                                string `json:"PLAYER_SCORE_2"`
	PlayerScore3                                string `json:"PLAYER_SCORE_3"`
	PlayerScore4                                string `json:"PLAYER_SCORE_4"`
	PlayerScore5                                string `json:"PLAYER_SCORE_5"`
	PlayerScore6                                string `json:"PLAYER_SCORE_6"`
	PlayerScore7                                string `json:"PLAYER_SCORE_7"`
	PlayerScore8                                string `json:"PLAYER_SCORE_8"`
	PlayerScore9                                string `json:"PLAYER_SCORE_9"`
	QuadraKills                                 string `json:"QUADRA_KILLS"`
	SightWardsBoughtInGame                      string `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
	Skin                                        string `json:"SKIN"`
	Spell1Cast                                  string `json:"SPELL1_CAST"`
	Spell2Cast                                  string `json:"SPELL2_CAST"`
	Spell3Cast                                  string `json:"SPELL3_CAST"`
	Spell4Cast                                  string `json:"SPELL4_CAST"`
	StatPerk0                                   string `json:"STAT_PERK_0"`
	StatPerk1                                   string `json:"STAT_PERK_1"`
	StatPerk2                                   string `json:"STAT_PERK_2"`
	SummonSpell1Cast                            string `json:"SUMMON_SPELL1_CAST"`
	SummonSpell2Cast                            string `json:"SUMMON_SPELL2_CAST"`
	Team                                        string `json:"TEAM"`
	TeamEarlySurrendered                        string `json:"TEAM_EARLY_SURRENDERED"`
	TeamObjective                               string `json:"TEAM_OBJECTIVE"`
	TeamPosition                                string `json:"TEAM_POSITION"`
	TimeCcingOthers                             string `json:"TIME_CCING_OTHERS"`
	TimeOfFromLastDisconnect                    string `json:"TIME_OF_FROM_LAST_DISCONNECT"`
	TimePlayed                                  string `json:"TIME_PLAYED"`
	TimeSpentDisconnected                       string `json:"TIME_SPENT_DISCONNECTED"`
	TotalDamageDealt                            string `json:"TOTAL_DAMAGE_DEALT"`
	TotalDamageDealtToBuildings                 string `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
	TotalDamageDealtToChampions                 string `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
	TotalDamageDealtToObjectives                string `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
	TotalDamageDealtToTurrets                   string `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
	TotalDamageSelfMitigated                    string `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
	TotalDamageShieldedOnTeammates              string `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
	TotalDamageTaken                            string `json:"TOTAL_DAMAGE_TAKEN"`
	TotalHeal                                   string `json:"TOTAL_HEAL"`
	TotalHealOnTeammates                        string `json:"TOTAL_HEAL_ON_TEAMMATES"`
	TotalTimeCrowdControlDealt                  string `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
	TotalTimeSpentDead                          string `json:"TOTAL_TIME_SPENT_DEAD"`
	TotalUnitsHealed                            string `json:"TOTAL_UNITS_HEALED"`
	TripleKills                                 string `json:"TRIPLE_KILLS"`
	TrueDamageDealtPlayer                       string `json:"TRUE_DAMAGE_DEALT_PLAYER"`
	TrueDamageDealtToChampions                  string `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
	TrueDamageTaken                             string `json:"TRUE_DAMAGE_TAKEN"`
	TurretsKilled                               string `json:"TURRETS_KILLED"`
	TurretTakedowns                             string `json:"TURRET_TAKEDOWNS"`
	UnrealKills                                 string `json:"UNREAL_KILLS"`
	VictoryPointTotal                           string `json:"VICTORY_POINT_TOTAL"`
	VisionScore                                 string `json:"VISION_SCORE"`
	VisionWardsBoughtInGame                     string `json:"VISION_WARDS_BOUGHT_IN_GAME"`
	WardKilled                                  string `json:"WARD_KILLED"`
	WardPlaced                                  string `json:"WARD_PLACED"`
	WardPlacedDetector                          string `json:"WARD_PLACED_DETECTOR"`
	WasAfk                                      string `json:"WAS_AFK"`
	WasAfkAfterFailedSurrender                  string `json:"WAS_AFK_AFTER_FAILED_SURRENDER"`
	WasEarlySurrenderAccomplice                 string `json:"WAS_EARLY_SURRENDER_ACCOMPLICE"`
	WasSurrenderDueToAfk                        string `json:"WAS_SURRENDER_DUE_TO_AFK"`
	Win                                         string `json:"WIN"`
	Missions_TakedownDragons                    string `json:"Missions_TakedownDragons"`
	Missions_TrueDamageToStructures             string `json:"Missions_TrueDamageToStructures"`
	PLAYER_AUGMENT_6                            string `json:"PLAYER_AUGMENT_6"`
	RIFT_HERALD_KILLS                           string `json:"RIFT_HERALD_KILLS"`
	Missions_CreepScoreBy10Minutes              string `json:"Missions_CreepScoreBy10Minutes"`
	WAS_LEAVER                                  string `json:"WAS_LEAVER"`
	Missions_TakedownsWithHelpFromMonsters      string `json:"Missions_TakedownsWithHelpFromMonsters"`
	Missions_VoidMitesSummoned                  string `json:"Missions_VoidMitesSummoned"`
	HOLD_PINGS                                  string `json:"HOLD_PINGS"`
	LARGEST_ABILITY_DAMAGE                      string `json:"LARGEST_ABILITY_DAMAGE"`
	PLAYER_SUBTEAM                              string `json:"PLAYER_SUBTEAM"`
	Missions_HealingFromLevelObjects            string `json:"Missions_HealingFromLevelObjects"`
	Missions_Crepe_DamageDealtSpeedZone         string `json:"Missions_Crepe_DamageDealtSpeedZone"`
	Missions_MinionsKilled                      string `json:"Missions_MinionsKilled"`
	SUMMONER_ID                                 string `json:"SUMMONER_ID"`
	Missions_ChampionsKilled                    string `json:"Missions_ChampionsKilled"`
	PLAYER_AUGMENT_2                            string `json:"PLAYER_AUGMENT_2"`
	PLAYER_AUGMENT_5                            string `json:"PLAYER_AUGMENT_5"`
	Missions_GoldFromStructuresDestroyed        string `json:"Missions_GoldFromStructuresDestroyed"`
	RIOT_ID_TAG_LINE                            string `json:"RIOT_ID_TAG_LINE"`
	Missions_LegendaryItems                     string `json:"Missions_LegendaryItems"`
	Missions_TakedownEpicMonsters               string `json:"Missions_TakedownEpicMonsters"`
	Missions_TakedownBaronsElderDragons         string `json:"Missions_TakedownBaronsElderDragons"`
	SeasonalMissions_TakedownAtakhan            string `json:"SeasonalMissions_TakedownAtakhan"`
	Missions_ChampionTakedownsWithIgnite        string `json:"Missions_ChampionTakedownsWithIgnite"`
	Missions_DamageToStructures                 string `json:"Missions_DamageToStructures"`
	TOTAL_TIME_CROWD_CONTROL_DEALT_TO_CHAMPIONS string `json:"TOTAL_TIME_CROWD_CONTROL_DEALT_TO_CHAMPIONS"`
	Missions_PlaceUsefulWards                   string `json:"Missions_PlaceUsefulWards"`
	Missions_ResolveRune                        string `json:"Missions_ResolveRune"`
	Missions_TakedownsAfterTeleporting          string `json:"Missions_TakedownsAfterTeleporting"`
	LAST_TAKEDOWN_TIME                          string `json:"LAST_TAKEDOWN_TIME"`
	SUMMONER_SPELL_1                            string `json:"SUMMONER_SPELL_1"`
	Missions_PrecisionRune                      string `json:"Missions_PrecisionRune"`
	BASIC_PINGS                                 string `json:"BASIC_PINGS"`
	ASSIST_ME_PINGS                             string `json:"ASSIST_ME_PINGS"`
	Missions_ImmobilizeChampions                string `json:"Missions_ImmobilizeChampions"`
	Missions_TakedownGold                       string `json:"Missions_TakedownGold"`
	Missions_TakedownsBefore15Min               string `json:"Missions_TakedownsBefore15Min"`
	PLAYER_AUGMENT_1                            string `json:"PLAYER_AUGMENT_1"`
	ALL_IN_PINGS                                string `json:"ALL_IN_PINGS"`
	Missions_DominationRune                     string `json:"Missions_DominationRune"`
	Missions_HexgatesUsed                       string `json:"Missions_HexgatesUsed"`
	DANGER_PINGS                                string `json:"DANGER_PINGS"`
	Missions_SnowballsHit                       string `json:"Missions_SnowballsHit"`
	Missions_TakedownsAfterExhausting           string `json:"Missions_TakedownsAfterExhausting"`
	Missions_TakedownsUnderTurret               string `json:"Missions_TakedownsUnderTurret"`
	ATAKHAN_KILLS                               string `json:"ATAKHAN_KILLS"`
	Missions_PeriodicDamage                     string `json:"Missions_PeriodicDamage"`
	HORDE_KILLS                                 string `json:"HORDE_KILLS"`
	Missions_TakedownStructures                 string `json:"Missions_TakedownStructures"`
	Missions_TotalGold                          string `json:"Missions_TotalGold"`
	Missions_PlaceUsefulControlWards            string `json:"Missions_PlaceUsefulControlWards"`
	Missions_CannonMinionsKilled                string `json:"Missions_CannonMinionsKilled"`
	PLAYER_SUBTEAM_PLACEMENT                    string `json:"PLAYER_SUBTEAM_PLACEMENT"`
	Missions_TurretPlatesDestroyed              string `json:"Missions_TurretPlatesDestroyed"`
	SUMMONER_SPELL_2                            string `json:"SUMMONER_SPELL_2"`
	VISION_CLEARED_PINGS                        string `json:"VISION_CLEARED_PINGS"`
	ON_MY_WAY_PINGS                             string `json:"ON_MY_WAY_PINGS"`
	PUUID                                       string `json:"PUUID"`
	GET_BACK_PINGS                              string `json:"GET_BACK_PINGS"`
	Missions_GoldFromTurretPlatesTaken          string `json:"Missions_GoldFromTurretPlatesTaken"`
	NEED_VISION_PINGS                           string `json:"NEED_VISION_PINGS"`
	ENEMY_MISSING_PINGS                         string `json:"ENEMY_MISSING_PINGS"`
	Event_2025LR_StructuresEpicMonsters         string `json:"Event_2025LR_StructuresEpicMonsters"`
	Missions_TwoChampsKilledWithSameAbility     string `json:"Missions_TwoChampsKilledWithSameAbility"`
	PUSH_PINGS                                  string `json:"PUSH_PINGS"`
	RIOT_ID_GAME_NAME                           string `json:"RIOT_ID_GAME_NAME"`
	Missions_Crepe_SnowballLanded               string `json:"Missions_Crepe_SnowballLanded"`
	Missions_DamageToChampsWithItems            string `json:"Missions_DamageToChampsWithItems"`
	RETREAT_PINGS                               string `json:"RETREAT_PINGS"`
	Missions_CreepScore                         string `json:"Missions_CreepScore"`
	Missions_TakedownEpicMonstersSingleGame     string `json:"Missions_TakedownEpicMonstersSingleGame"`
	LARGEST_ATTACK_DAMAGE                       string `json:"LARGEST_ATTACK_DAMAGE"`
	Missions_ChampionsHitWithAbilitiesEarlyGame string `json:"Missions_ChampionsHitWithAbilitiesEarlyGame"`
	Missions_DestroyPlants                      string `json:"Missions_DestroyPlants"`
	Missions_GoldPerMinute                      string `json:"Missions_GoldPerMinute"`
	Missions_TakedownWards                      string `json:"Missions_TakedownWards"`
	PLAYER_AUGMENT_3                            string `json:"PLAYER_AUGMENT_3"`
	ENEMY_VISION_PINGS                          string `json:"ENEMY_VISION_PINGS"`
	Missions_InspirationRune                    string `json:"Missions_InspirationRune"`
	Missions_Crepe_TakedownsWithInhibBuff       string `json:"Missions_Crepe_TakedownsWithInhibBuff"`
	Missions_PorosFed                           string `json:"Missions_PorosFed"`
	Missions_SorceryRune                        string `json:"Missions_SorceryRune"`
	PLAYER_AUGMENT_4                            string `json:"PLAYER_AUGMENT_4"`
	COMMAND_PINGS                               string `json:"COMMAND_PINGS"`
	Missions_ChampionTakedownsWhileGhosted      string `json:"Missions_ChampionTakedownsWhileGhosted"`
}

type Chunk struct {
	Id        uint32
	ChunkType byte
	Length    uint32
	NextId    uint32
	Offset    uint32
	Data      []byte
}

type Keyframe struct {
	Id           uint32
	KeyframeType byte
	Length       uint32
	NextId       uint32
	Offset       uint32
	Data         []byte
}
