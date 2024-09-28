package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"strings"
	"time"
)

// here we define all the metrics we want to export
var (
	TotalXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_totalxp",
		Help: "The total xp accumulated by the account",
	})
	QuestsStarted = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_quests_started",
		Help: "The total number of in-progress quests",
	})
	CombatLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_combat_level",
		Help: "The combat level of the user",
	})
	QuestsCompleted = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_quests_completed",
		Help: "The total number of completed quests",
	})
	QuestsNotStarted = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_quests_not_started",
		Help: "The total number of quests not yet started",
	})
	TotalLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_total_level",
		Help: "The total number of levels",
	})
	Rank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_rank",
		Help: "The overall xp rank of the user",
	})
	Online = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_user_online",
		Help: "The total xp online",
	})
	// skills :ugh:
	AttackLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_attack_level",
		Help: "The user's attack level",
	})
	AttackXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_attack_xp",
		Help: "The user's attack XP",
	})
	AttackRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_attack_rank",
		Help: "The user's attack xp ranking",
	})

	DefenceLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_defence_level",
		Help: "The user's defence level",
	})
	DefenceXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_defence_xp",
		Help: "The user's defence XP",
	})
	DefenceRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_defence_rank",
		Help: "The user's defence xp ranking",
	})

	StrengthLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_strength_level",
		Help: "The user's strength level",
	})
	StrengthXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_strength_xp",
		Help: "The user's strength XP",
	})
	StrengthRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_strength_rank",
		Help: "The user's strength xp ranking",
	})

	ConstitutionLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_constitution_level",
		Help: "The user's constitution level",
	})
	ConstitutionXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_constitution_xp",
		Help: "The user's constitution XP",
	})
	ConstitutionRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_constitution_rank",
		Help: "The user's constitution xp ranking",
	})

	RangedLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_ranged_level",
		Help: "The user's ranged level",
	})
	RangedXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_ranged_xp",
		Help: "The user's ranged XP",
	})
	RangedRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_ranged_rank",
		Help: "The user's ranged xp ranking",
	})

	PrayerLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_prayer_level",
		Help: "The user's prayer level",
	})
	PrayerXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_prayer_xp",
		Help: "The user's prayer XP",
	})
	PrayerRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_prayer_rank",
		Help: "The user's prayer xp ranking",
	})

	MagicLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_magic_level",
		Help: "The user's magic level",
	})
	MagicXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_magic_xp",
		Help: "The user's magic XP",
	})
	MagicRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_magic_rank",
		Help: "The user's magic xp ranking",
	})

	CookingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_cooking_level",
		Help: "The user's cooking level",
	})
	CookingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_cooking_xp",
		Help: "The user's cooking XP",
	})
	CookingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_cooking_rank",
		Help: "The user's cooking xp ranking",
	})

	WoodcuttingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_woodcutting_level",
		Help: "The user's woodcutting level",
	})
	WoodcuttingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_woodcutting_xp",
		Help: "The user's woodcutting XP",
	})
	WoodcuttingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_woodcutting_rank",
		Help: "The user's woodcutting xp ranking",
	})

	FletchingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_fletching_level",
		Help: "The user's fletching level",
	})
	FletchingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_fletching_xp",
		Help: "The user's fletching XP",
	})
	FletchingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_fletching_rank",
		Help: "The user's fletching xp ranking",
	})

	FishingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_fishing_level",
		Help: "The user's fishing level",
	})
	FishingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_fishing_xp",
		Help: "The user's fishing XP",
	})
	FishingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_fishing_rank",
		Help: "The user's fishing xp ranking",
	})

	FiremakingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_firemaking_level",
		Help: "The user's firemaking level",
	})
	FiremakingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_firemaking_xp",
		Help: "The user's firemaking XP",
	})
	FiremakingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_firemaking_rank",
		Help: "The user's firemaking xp ranking",
	})

	CraftingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_crafting_level",
		Help: "The user's crafting level",
	})
	CraftingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_crafting_xp",
		Help: "The user's crafting XP",
	})
	CraftingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_crafting_rank",
		Help: "The user's crafting xp ranking",
	})

	SmithingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_smithing_level",
		Help: "The user's smithing level",
	})
	SmithingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_smithing_xp",
		Help: "The user's smithing XP",
	})
	SmithingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_smithing_rank",
		Help: "The user's smithing xp ranking",
	})

	MiningLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_mining_level",
		Help: "The user's mining level",
	})
	MiningXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_mining_xp",
		Help: "The user's mining XP",
	})
	MiningRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_mining_rank",
		Help: "The user's mining xp ranking",
	})

	HerbloreLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_herblore_level",
		Help: "The user's herblore level",
	})
	HerbloreXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_herblore_xp",
		Help: "The user's herblore XP",
	})
	HerbloreRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_herblore_rank",
		Help: "The user's herblore xp ranking",
	})

	AgilityLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_agility_level",
		Help: "The user's agility level",
	})
	AgilityXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_agility_xp",
		Help: "The user's agility XP",
	})
	AgilityRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_agility_rank",
		Help: "The user's agility xp ranking",
	})

	ThievingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_thieving_level",
		Help: "The user's thieving level",
	})
	ThievingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_thieving_xp",
		Help: "The user's thieving XP",
	})
	ThievingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_thieving_rank",
		Help: "The user's thieving xp ranking",
	})

	SlayerLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_slayer_level",
		Help: "The user's slayer level",
	})
	SlayerXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_slayer_xp",
		Help: "The user's slayer XP",
	})
	SlayerRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_slayer_rank",
		Help: "The user's slayer xp ranking",
	})

	FarmingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_farming_level",
		Help: "The user's farming level",
	})
	FarmingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_farming_xp",
		Help: "The user's farming XP",
	})
	FarmingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_farming_rank",
		Help: "The user's farming xp ranking",
	})

	RunecraftingLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_runecrafting_level",
		Help: "The user's runecrafting level",
	})
	RunecraftingXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_runecrafting_xp",
		Help: "The user's runecrafting XP",
	})
	RunecraftingRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_runecrafting_rank",
		Help: "The user's runecrafting xp ranking",
	})
	HunterLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_hunter_level",
		Help: "The user's hunter level",
	})
	HunterXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_hunter_xp",
		Help: "The user's hunter XP",
	})
	HunterRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_hunter_rank",
		Help: "The user's hunter xp ranking",
	})

	ConstructionLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_construction_level",
		Help: "The user's construction level",
	})
	ConstructionXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_construction_xp",
		Help: "The user's construction XP",
	})
	ConstructionRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_construction_rank",
		Help: "The user's construction xp ranking",
	})

	SummoningLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_summoning_level",
		Help: "The user's summoning level",
	})
	SummoningXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_summoning_xp",
		Help: "The user's summoning XP",
	})
	SummoningRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_summoning_rank",
		Help: "The user's summoning xp ranking",
	})

	DungeoneeringLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_dungeoneering_level",
		Help: "The user's dungeoneering level",
	})
	DungeoneeringXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_dungeoneering_xp",
		Help: "The user's dungeoneering XP",
	})
	DungeoneeringRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_dungeoneering_rank",
		Help: "The user's dungeoneering xp ranking",
	})

	DivinationLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_divination_level",
		Help: "The user's divination level",
	})
	DivinationXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_divination_xp",
		Help: "The user's divination XP",
	})
	DivinationRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_divination_rank",
		Help: "The user's divination xp ranking",
	})

	InventionLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_invention_level",
		Help: "The user's invention level",
	})
	InventionXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_invention_xp",
		Help: "The user's invention XP",
	})
	InventionRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_invention_rank",
		Help: "The user's invention xp ranking",
	})

	ArchaeologyLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_archaeology_level",
		Help: "The user's archaeology level",
	})
	ArchaeologyXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_archaeology_xp",
		Help: "The user's archaeology XP",
	})
	ArchaeologyRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_archaeology_rank",
		Help: "The user's archaeology xp ranking",
	})

	NecromancyLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_necromancy_level",
		Help: "The user's necromancy level",
	})
	NecromancyXP = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_necromancy_xp",
		Help: "The user's necromancy XP",
	})
	NecromancyRank = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "rs3_skills_necromancy_rank",
		Help: "The user's necromancy xp ranking",
	})
)

func updateMetrics() {
	go func() {
		for {
			// first pull the metrics from the RuneMetrics API
			RSUser, err := RuneMetricsGetUser(RuneScapeUsername)
			if err != nil {
				return
			}
			// update our metrics
			TotalXP.Set(float64(RSUser.TotalXP))
			QuestsStarted.Set(float64(RSUser.QuestsStarted))
			QuestsCompleted.Set(float64(RSUser.QuestsCompleted))
			QuestsNotStarted.Set(float64(RSUser.QuestsNotStarted))
			TotalLevel.Set(float64(RSUser.TotalSkill))
			CombatLevel.Set(float64(RSUser.CombatLevel))
			//try to update rank
			newRank, err := parseRank(RSUser.Rank)
			if err == nil {
				Rank.Set(newRank)
			}
			// we can actually track online/offline with a bool. since prom only supports float64, we need to encode it to 0/1 and send it
			Online.Set(parseOnline(RSUser.LoggedIn))

			//time for the annoying part, skills
			AttackSkill, err := RSUser.GetSkillByID(0)
			if err == nil {
				AttackLevel.Set(float64(AttackSkill.Level))
				AttackXP.Set(float64(AttackSkill.XP))
				AttackRank.Set(float64(AttackSkill.Rank))
			}

			DefenceSkill, err := RSUser.GetSkillByID(1)
			if err == nil {
				DefenceLevel.Set(float64(DefenceSkill.Level))
				DefenceXP.Set(float64(DefenceSkill.XP))
				DefenceRank.Set(float64(DefenceSkill.Rank))
			}

			StrengthSkill, err := RSUser.GetSkillByID(2)
			if err == nil {
				StrengthLevel.Set(float64(StrengthSkill.Level))
				StrengthXP.Set(float64(StrengthSkill.XP))
				StrengthRank.Set(float64(StrengthSkill.Rank))
			}

			ConstitutionSkill, err := RSUser.GetSkillByID(3)
			if err == nil {
				ConstitutionLevel.Set(float64(ConstitutionSkill.Level))
				ConstitutionXP.Set(float64(ConstitutionSkill.XP))
				ConstitutionRank.Set(float64(ConstitutionSkill.Rank))
			}

			RangedSkill, err := RSUser.GetSkillByID(4)
			if err == nil {
				RangedLevel.Set(float64(RangedSkill.Level))
				RangedXP.Set(float64(RangedSkill.XP))
				RangedRank.Set(float64(RangedSkill.Rank))
			}

			PrayerSkill, err := RSUser.GetSkillByID(5)
			if err == nil {
				PrayerLevel.Set(float64(PrayerSkill.Level))
				PrayerXP.Set(float64(PrayerSkill.XP))
				PrayerRank.Set(float64(PrayerSkill.Rank))
			}

			MagicSkill, err := RSUser.GetSkillByID(6)
			if err == nil {
				MagicLevel.Set(float64(MagicSkill.Level))
				MagicXP.Set(float64(MagicSkill.XP))
				MagicRank.Set(float64(MagicSkill.Rank))
			}

			CookingSkill, err := RSUser.GetSkillByID(7)
			if err == nil {
				CookingLevel.Set(float64(CookingSkill.Level))
				CookingXP.Set(float64(CookingSkill.XP))
				CookingRank.Set(float64(CookingSkill.Rank))
			}

			WoodcuttingSkill, err := RSUser.GetSkillByID(8)
			if err == nil {
				WoodcuttingLevel.Set(float64(WoodcuttingSkill.Level))
				WoodcuttingXP.Set(float64(WoodcuttingSkill.XP))
				WoodcuttingRank.Set(float64(WoodcuttingSkill.Rank))
			}

			FletchingSkill, err := RSUser.GetSkillByID(9)
			if err == nil {
				FletchingLevel.Set(float64(FletchingSkill.Level))
				FletchingXP.Set(float64(FletchingSkill.XP))
				FletchingRank.Set(float64(FletchingSkill.Rank))
			}

			FishingSkill, err := RSUser.GetSkillByID(10)
			if err == nil {
				FishingLevel.Set(float64(FishingSkill.Level))
				FishingXP.Set(float64(FishingSkill.XP))
				FishingRank.Set(float64(FishingSkill.Rank))
			}

			FiremakingSkill, err := RSUser.GetSkillByID(11)
			if err == nil {
				FiremakingLevel.Set(float64(FiremakingSkill.Level))
				FiremakingXP.Set(float64(FiremakingSkill.XP))
				FiremakingRank.Set(float64(FiremakingSkill.Rank))
			}

			CraftingSkill, err := RSUser.GetSkillByID(12)
			if err == nil {
				CraftingLevel.Set(float64(CraftingSkill.Level))
				CraftingXP.Set(float64(CraftingSkill.XP))
				CraftingRank.Set(float64(CraftingSkill.Rank))
			}

			SmithingSkill, err := RSUser.GetSkillByID(13)
			if err == nil {
				SmithingLevel.Set(float64(SmithingSkill.Level))
				SmithingXP.Set(float64(SmithingSkill.XP))
				SmithingRank.Set(float64(SmithingSkill.Rank))
			}

			MiningSkill, err := RSUser.GetSkillByID(14)
			if err == nil {
				MiningLevel.Set(float64(MiningSkill.Level))
				MiningXP.Set(float64(MiningSkill.XP))
				MiningRank.Set(float64(MiningSkill.Rank))
			}

			HerbloreSkill, err := RSUser.GetSkillByID(15)
			if err == nil {
				HerbloreLevel.Set(float64(HerbloreSkill.Level))
				HerbloreXP.Set(float64(HerbloreSkill.XP))
				HerbloreRank.Set(float64(HerbloreSkill.Rank))
			}

			AgilitySkill, err := RSUser.GetSkillByID(16)
			if err == nil {
				AgilityLevel.Set(float64(AgilitySkill.Level))
				AgilityXP.Set(float64(AgilitySkill.XP))
				AgilityRank.Set(float64(AgilitySkill.Rank))
			}

			ThievingSkill, err := RSUser.GetSkillByID(17)
			if err == nil {
				ThievingLevel.Set(float64(ThievingSkill.Level))
				ThievingXP.Set(float64(ThievingSkill.XP))
				ThievingRank.Set(float64(ThievingSkill.Rank))
			}

			SlayerSkill, err := RSUser.GetSkillByID(18)
			if err == nil {
				SlayerLevel.Set(float64(SlayerSkill.Level))
				SlayerXP.Set(float64(SlayerSkill.XP))
				SlayerRank.Set(float64(SlayerSkill.Rank))
			}

			FarmingSkill, err := RSUser.GetSkillByID(19)
			if err == nil {
				FarmingLevel.Set(float64(FarmingSkill.Level))
				FarmingXP.Set(float64(FarmingSkill.XP))
				FarmingRank.Set(float64(FarmingSkill.Rank))
			}

			RunecraftingSkill, err := RSUser.GetSkillByID(20)
			if err == nil {
				RunecraftingLevel.Set(float64(RunecraftingSkill.Level))
				RunecraftingXP.Set(float64(RunecraftingSkill.XP))
				RunecraftingRank.Set(float64(RunecraftingSkill.Rank))
			}

			HunterSkill, err := RSUser.GetSkillByID(21)
			if err == nil {
				HunterLevel.Set(float64(HunterSkill.Level))
				HunterXP.Set(float64(HunterSkill.XP))
				HunterRank.Set(float64(HunterSkill.Rank))
			}

			ConstructionSkill, err := RSUser.GetSkillByID(22)
			if err == nil {
				ConstructionLevel.Set(float64(ConstructionSkill.Level))
				ConstructionXP.Set(float64(ConstructionSkill.XP))
				ConstructionRank.Set(float64(ConstructionSkill.Rank))
			}

			SummoningSkill, err := RSUser.GetSkillByID(23)
			if err == nil {
				SummoningLevel.Set(float64(SummoningSkill.Level))
				SummoningXP.Set(float64(SummoningSkill.XP))
				SummoningRank.Set(float64(SummoningSkill.Rank))
			}

			DungeoneeringSkill, err := RSUser.GetSkillByID(24)
			if err == nil {
				DungeoneeringLevel.Set(float64(DungeoneeringSkill.Level))
				DungeoneeringXP.Set(float64(DungeoneeringSkill.XP))
				DungeoneeringRank.Set(float64(DungeoneeringSkill.Rank))
			}

			DivinationSkill, err := RSUser.GetSkillByID(25)
			if err == nil {
				DivinationLevel.Set(float64(DivinationSkill.Level))
				DivinationXP.Set(float64(DivinationSkill.XP))
				DivinationRank.Set(float64(DivinationSkill.Rank))
			}

			InventionSkill, err := RSUser.GetSkillByID(26)
			if err == nil {
				InventionLevel.Set(float64(InventionSkill.Level))
				InventionXP.Set(float64(InventionSkill.XP))
				InventionRank.Set(float64(InventionSkill.Rank))
			}

			ArchaeologySkill, err := RSUser.GetSkillByID(27)
			if err == nil {
				ArchaeologyLevel.Set(float64(ArchaeologySkill.Level))
				ArchaeologyXP.Set(float64(ArchaeologySkill.XP))
				ArchaeologyRank.Set(float64(ArchaeologySkill.Rank))
			}

			NecromancySkill, err := RSUser.GetSkillByID(28)
			if err == nil {
				NecromancyLevel.Set(float64(NecromancySkill.Level))
				NecromancyXP.Set(float64(NecromancySkill.XP))
				NecromancyRank.Set(float64(NecromancySkill.Rank))
			}

			// update data every x timeframe
			time.Sleep(*UpdatePeriod)
		}
	}()
}

func parseRank(rank string) (float64, error) {
	//for whatever reason the API returns rank as a comma seperated string. let's make it to float64 like we want.
	newInt, err := strconv.Atoi(strings.ReplaceAll(rank, ",", ""))
	if err != nil {
		return 0, err
	}
	return float64(newInt), nil
}

func parseOnline(online string) float64 {
	if online == "true" {
		return 1
	}
	if online == "false" {
		return 0
	}
	return -1
}
