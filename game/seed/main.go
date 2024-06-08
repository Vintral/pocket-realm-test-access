package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	models "github.com/Vintral/pocket-realm-test-access/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Loading Environment")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	//==============================//
	//	Setup Telemetry							//
	//==============================//
	fmt.Println("Setting up telemetry")
	otelShutdown, tp, err := setupOTelSDK(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		fmt.Println("In shutdown")
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	fmt.Println("Setting Trace Provider")
	models.SetTracerProvider(tp)

	fmt.Println("Setting up database")
	db, err := models.Database(false)
	if err != nil {
		time.Sleep(3 * time.Second)
		panic(err)
	}

	dropTables(db)
	runMigrations(db)

	createRules(db)
	createNews(db)
	createUnits(db)
	createBuildings(db)
	createResources(db)
	createItems(db)
	createOverrides(db)
	round := createRounds(db)
	createUsers(db, round)
	createUserTables(db, round)
	createShouts(db)
	createConversations(db)
}

func createUserTables(db *gorm.DB, round *models.Round) {
	//================================//
	// Users Units										//
	//================================//
	fmt.Println("Seeding User's Units")
	// db.Create(&models.UserUnit{
	// 	UserID:   1,
	// 	UnitID:   1,
	// 	RoundID:  1,
	// 	Quantity: 15,
	// })
	// db.Create(&models.UserUnit{
	// 	UserID:   1,
	// 	UnitID:   2,
	// 	RoundID:  1,
	// 	Quantity: 20,
	// })
	// db.Create(&models.UserUnit{
	// 	UserID:   1,
	// 	UnitID:   3,
	// 	RoundID:  1,
	// 	Quantity: 30,
	// })
	// db.Create(&models.UserUnit{
	// 	UserID:   1,
	// 	UnitID:   4,
	// 	RoundID:  1,
	// 	Quantity: 40,
	// })
	// db.Create(&models.UserUnit{
	// 	UserID:   1,
	// 	UnitID:   5,
	// 	RoundID:  2,
	// 	Quantity: 45,
	// })

	//================================//
	// User Buildings									//
	//================================//
	fmt.Println("Seeding User's Buildings")
	// db.Create(&models.UserBuilding{
	// 	UserID:     1,
	// 	BuildingID: 1,
	// 	RoundID:    1,
	// 	Quantity:   10,
	// })
	// db.Create(&models.UserBuilding{
	// 	UserID:     1,
	// 	BuildingID: 2,
	// 	RoundID:    1,
	// 	Quantity:   10,
	// })
	// db.Create(&models.UserBuilding{
	// 	UserID:     1,
	// 	BuildingID: 3,
	// 	RoundID:    1,
	// 	Quantity:   10,
	// })
	// db.Create(&models.UserBuilding{
	// 	UserID:     1,
	// 	BuildingID: 4,
	// 	RoundID:    1,
	// 	Quantity:   10,
	// })
	// db.Create(&models.UserBuilding{
	// 	UserID:     1,
	// 	BuildingID: 5,
	// 	RoundID:    1,
	// 	Quantity:   10,
	// })
	// db.Create(&models.UserBuilding{
	// 	UserID:     1,
	// 	BuildingID: 6,
	// 	RoundID:    1,
	// 	Quantity:   10,
	// })

	//================================//
	// User Items											//
	//================================//
	fmt.Println("Seeding User's Items")
	// db.Create(&models.UserItem{
	// 	UserID: 1,
	// 	ItemID: 1,
	// })

	//================================//
	// Users Rounds										//
	//================================//
	fmt.Println("Seeding User's Round")
	db.Create(&models.UserRound{
		UserID:         1,
		RoundID:        1,
		CharacterClass: "mage",
		Energy:         int(round.EnergyMax),
		Gold:           200,
		Food:           200,
		Wood:           200,
		Metal:          200,
		Faith:          200,
		Stone:          200,
		Mana:           200,
		Land:           200,
		FreeLand:       200,
		BuildPower:     1,
		RecruitPower:   1,
	})
	db.Create(&models.UserRound{
		UserID:         1,
		RoundID:        2,
		CharacterClass: "mage",
		Energy:         int(round.EnergyMax),
		Gold:           200,
		TickGold:       5,
		Food:           200,
		TickFood:       5,
		Wood:           200,
		TickWood:       5,
		Metal:          200,
		TickMetal:      5,
		Faith:          200,
		TickFaith:      5,
		Stone:          200,
		TickStone:      5,
		Mana:           200,
		TickMana:       5,
		Land:           200,
		FreeLand:       200,
		BuildPower:     25,
		RecruitPower:   25,
	})
	db.Create(&models.UserRound{
		UserID:         3,
		RoundID:        1,
		CharacterClass: "priest",
	})
	db.Create(&models.UserRound{
		UserID:         4,
		RoundID:        1,
		CharacterClass: "warlord",
	})
	db.Create(&models.UserRound{
		UserID:         5,
		RoundID:        1,
		CharacterClass: "necromancer",
	})
	db.Create(&models.UserRound{
		UserID:         6,
		RoundID:        1,
		CharacterClass: "merchant",
	})
	db.Create(&models.UserRound{
		UserID:         7,
		RoundID:        1,
		CharacterClass: "druid",
	})
}

func createRounds(db *gorm.DB) *models.Round {
	//================================//
	// Rounds													//
	//================================//
	fmt.Println("Seeding Round")
	round := &models.Round{
		EnergyMax:   250,
		EnergyRegen: 10,
		Ends:        time.Now().Add(14 * 24 * time.Hour),
	}
	db.Create(round)

	return round
}

func createUsers(db *gorm.DB, round *models.Round) {
	//================================//
	// Users													//
	//================================//
	fmt.Println("Seeding users")
	db.Create(&models.User{
		Email:        "jeffrey.heater@gmail.com",
		Admin:        true,
		Username:     "Vintral",
		Avatar:       "1",
		RoundID:      1,
		RoundPlaying: round.GUID,
	})
	db.Create(&models.User{
		Email:    "jeffrey.heater0@gmail.com",
		Admin:    true,
		Username: "Trilanni",
		Avatar:   "2",
	})
	db.Create(&models.User{
		Email:        "jeffrey.heater1@gmail.com",
		Admin:        true,
		Username:     "Vintral1",
		Avatar:       "3",
		RoundID:      1,
		RoundPlaying: round.GUID,
	})
	db.Create(&models.User{
		Email:        "jeffrey.heater2@gmail.com",
		Admin:        true,
		Username:     "Vintral2",
		Avatar:       "4",
		RoundID:      1,
		RoundPlaying: round.GUID,
	})
	db.Create(&models.User{
		Email:        "jeffrey.heater3@gmail.com",
		Admin:        true,
		Username:     "Vintral3",
		Avatar:       "5",
		RoundID:      1,
		RoundPlaying: round.GUID,
	})
	db.Create(&models.User{
		Email:        "jeffrey.heater4@gmail.com",
		Admin:        true,
		Username:     "Vintral4",
		Avatar:       "6",
		RoundID:      1,
		RoundPlaying: round.GUID,
	})
	db.Create(&models.User{
		Email:        "jeffrey.heater5@gmail.com",
		Admin:        true,
		Username:     "Vintral5",
		Avatar:       "1",
		RoundID:      1,
		RoundPlaying: round.GUID,
	})
}

func createShouts(db *gorm.DB) {
	//================================//
	// Shouts													//
	//================================//
	fmt.Println("Seeding shouts")
	db.Create(&models.Shout{
		UserID: 1,
		Shout:  "Mage Shout",
	})
	db.Create(&models.Shout{
		UserID: 2,
		Shout:  "Not Playing Shout",
	})
	db.Create(&models.Shout{
		UserID: 3,
		Shout:  "Priest Shout",
	})
	db.Create(&models.Shout{
		UserID: 4,
		Shout:  "Warlord Shout",
	})
	db.Create(&models.Shout{
		UserID: 5,
		Shout:  "Necromacer Shout",
	})
	db.Create(&models.Shout{
		UserID: 6,
		Shout:  "Merchant shout",
	})
	db.Create(&models.Shout{
		UserID: 7,
		Shout:  "Druid shout",
	})
	db.Create(&models.Shout{
		UserID: 1,
		Shout:  "Mage shout",
	})
	db.Create(&models.Shout{
		UserID: 2,
		Shout:  "Not Playing shout",
	})
	db.Create(&models.Shout{
		UserID: 3,
		Shout:  "Priest shout",
	})
}

func dropTables(db *gorm.DB) {
	db.Exec("DROP TABLE user_units")
	db.Exec("DROP TABLE user_rounds")
	db.Exec("DROP TABLE user_buildings")
	db.Exec("DROP TABLE user_items")
	db.Exec("DROP TABLE round_resources")
	db.Exec("DROP TABLE round_buildings")
	db.Exec("DROP TABLE round_units")
	db.Exec("DROP TABLE units")
	db.Exec("DROP TABLE users")
	db.Exec("DROP TABLE buildings")
	db.Exec("DROP TABLE items")
	db.Exec("DROP TABLE rounds")
	db.Exec("DROP TABLE resources")
	db.Exec("DROP TABLE news_items")
	db.Exec("DROP TABLE rules")
	db.Exec("DROP TABLE shouts")
	db.Exec("DROP TABLE user_logs")
	db.Exec("DROP TABLE conversations")
	db.Exec("DROP TABLE messages")
}

func createConversations(db *gorm.DB) {
	fmt.Println("Seeding conversations")

	conversation := &models.Conversation{
		User1ID:       1,
		User2ID:       2,
		User2LastRead: time.Now(),
	}
	db.Create(conversation)

	for i := 0; i < 15; i++ {
		db.Create(&models.Message{
			Conversation: conversation.ID,
			UserID:       1 + uint(i%2),
			Text:         "Message should show",
		})
	}

	conversation = &models.Conversation{
		User1ID:       2,
		User2ID:       3,
		User2LastRead: time.Now(),
	}
	db.Create(conversation)

	for i := 0; i < 15; i++ {
		db.Create(&models.Message{
			Conversation: conversation.ID,
			UserID:       1 + uint(i%2),
			Text:         "Message should not show",
		})
	}
}

func runMigrations(db *gorm.DB) {
	models.RunMigrations(db)
}

func createOverrides(db *gorm.DB) {
	db.Create(&models.RoundUnit{
		RoundID:     1,
		UnitID:      1,
		Attack:      1.00,
		Defense:     1.00,
		Power:       1.00,
		Health:      5,
		Ranged:      false,
		CostGold:    1,
		CostPoints:  1,
		CostFood:    1,
		CostWood:    1,
		CostMetal:   1,
		CostStone:   1,
		CostFaith:   1,
		CostMana:    1,
		UpkeepGold:  1,
		UpkeepFood:  1,
		UpkeepWood:  1,
		UpkeepStone: 1,
		UpkeepMetal: 1,
		UpkeepFaith: 1,
		UpkeepMana:  1,
		Available:   true,
		Recruitable: true,
	})

	db.Create(&models.RoundResource{RoundID: 1, ResourceID: 6, CanGather: false, CanMarket: false})
	db.Create(&models.RoundResource{RoundID: 1, ResourceID: 7, CanGather: false, CanMarket: false})

	db.Create(&models.RoundBuilding{
		BuildingID:  1,
		RoundID:     1,
		CostPoints:  1,
		CostWood:    1,
		CostStone:   1,
		CostGold:    1,
		CostFood:    1,
		CostMetal:   1,
		CostFaith:   1,
		CostMana:    1,
		BonusValue:  1,
		UpkeepGold:  1,
		UpkeepFood:  1,
		UpkeepWood:  1,
		UpkeepStone: 1,
		UpkeepMetal: 1,
		UpkeepFaith: 1,
		UpkeepMana:  1,
		Buildable:   true,
		Available:   true,
	})
}

func createBuildings(db *gorm.DB) {
	//================================//
	// Buildings											//
	//================================//
	fmt.Println("Seeding Buildings")
	db.Create(&models.Building{Name: "farm", BonusField: "food_tick"})
	db.Create(&models.Building{Name: "barracks", BonusField: "recruit_power"})
	db.Create(&models.Building{Name: "lumber-mill", BonusField: "wood_tick"})
	db.Create(&models.Building{Name: "quarry", BonusField: "stone_tick"})
	db.Create(&models.Building{Name: "wall", BonusField: "defense"})
	db.Create(&models.Building{Name: "workshop", BonusField: "build_power"})
	db.Create(&models.Building{Name: "mine", BonusField: "metal_tick"})

	//================================//
	// Building Defaults							//
	//================================//
	fmt.Println("Seeding Building Defaults")
	db.Create(&models.RoundBuilding{
		BuildingID:      1,
		RoundID:         0,
		CostPoints:      1,
		CostWood:        1,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundBuilding{
		BuildingID:      2,
		RoundID:         0,
		CostWood:        100,
		CostStone:       100,
		CostPoints:      10,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundBuilding{
		BuildingID:      3,
		RoundID:         0,
		CostWood:        15,
		CostStone:       0,
		CostPoints:      2,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundBuilding{
		BuildingID:      4,
		RoundID:         0,
		CostWood:        5,
		CostStone:       10,
		CostPoints:      2,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundBuilding{
		BuildingID:      5,
		RoundID:         0,
		CostWood:        0,
		CostStone:       25,
		CostPoints:      2,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundBuilding{
		BuildingID:      6,
		RoundID:         0,
		CostWood:        20,
		CostStone:       25,
		CostPoints:      2,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundBuilding{
		BuildingID:      7,
		RoundID:         0,
		CostWood:        20,
		CostStone:       25,
		CostPoints:      2,
		BonusValue:      1,
		Available:       true,
		Buildable:       true,
		SupportsPartial: false,
	})
}

func createResources(db *gorm.DB) {
	//================================//
	// Resources											//
	//================================//
	fmt.Println("Seeding Resources")
	db.Create(&models.Resource{Name: "gold"})
	db.Create(&models.Resource{Name: "wood"})
	db.Create(&models.Resource{Name: "food"})
	db.Create(&models.Resource{Name: "stone"})
	db.Create(&models.Resource{Name: "metal"})
	db.Create(&models.Resource{Name: "faith"})
	db.Create(&models.Resource{Name: "mana"})

	//================================//
	// Resource Defaults							//
	//================================//
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 1, CanGather: true, CanMarket: false})
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 2, CanGather: true, CanMarket: true})
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 3, CanGather: true, CanMarket: true})
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 4, CanGather: true, CanMarket: true})
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 5, CanGather: true, CanMarket: true})
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 6, CanGather: true, CanMarket: false})
	db.Create(&models.RoundResource{RoundID: 0, ResourceID: 7, CanGather: true, CanMarket: false})
}

func createUnits(db *gorm.DB) {
	//================================//
	// Units													//
	//================================//
	fmt.Println("Seeding Units")
	db.Create(&models.Unit{Name: "peasant"})
	db.Create(&models.Unit{Name: "footman"})
	db.Create(&models.Unit{Name: "archer"})
	db.Create(&models.Unit{Name: "crusader"})
	db.Create(&models.Unit{Name: "cavalry"})

	//================================//
	// Unit Defaults  								//
	//================================//
	fmt.Println("Seeding Unit Defaults")
	db.Create(&models.RoundUnit{
		RoundID:         0,
		UnitID:          1,
		Attack:          1.00,
		Defense:         1.00,
		Power:           1.00,
		Health:          5,
		Ranged:          false,
		CostGold:        1,
		CostPoints:      1,
		CostFood:        1,
		UpkeepFood:      1,
		Available:       true,
		Recruitable:     true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundUnit{
		RoundID:         0,
		UnitID:          2,
		Attack:          2.00,
		Defense:         2.00,
		Power:           2.00,
		Health:          15,
		Ranged:          false,
		CostGold:        2,
		CostPoints:      2,
		UpkeepGold:      1,
		UpkeepFood:      1,
		Available:       true,
		Recruitable:     true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundUnit{
		RoundID:         0,
		UnitID:          3,
		Attack:          3.00,
		Defense:         1.00,
		Power:           3.00,
		Health:          15,
		Ranged:          true,
		CostGold:        5,
		CostPoints:      5,
		UpkeepGold:      2,
		UpkeepFood:      1,
		Available:       true,
		Recruitable:     true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundUnit{
		RoundID:         0,
		UnitID:          4,
		Attack:          5.00,
		Defense:         5.00,
		Power:           10.00,
		Health:          30,
		Ranged:          false,
		CostGold:        10,
		CostPoints:      10,
		UpkeepGold:      3,
		UpkeepFood:      2,
		Available:       true,
		Recruitable:     true,
		SupportsPartial: false,
	})
	db.Create(&models.RoundUnit{
		RoundID:         0,
		UnitID:          5,
		Attack:          10.00,
		Defense:         5.00,
		Power:           20.00,
		Health:          50,
		Ranged:          false,
		CostGold:        25,
		CostPoints:      20,
		UpkeepGold:      5,
		UpkeepFood:      5,
		Available:       true,
		Recruitable:     true,
		SupportsPartial: false,
	})
}

func createItems(db *gorm.DB) {
	//================================//
	// Items													//
	//================================//
	fmt.Println("Seeding Items")
	db.Create(&models.Item{
		Name:   "Hourglass",
		Plural: "Hourglasses",
	})
}

func createRules(db *gorm.DB) {
	fmt.Println("Seeding Rules")

	db.Create(&models.Rule{
		Rule:   "rule-1",
		Active: true,
	})
	db.Create(&models.Rule{
		Rule:   "rule-2",
		Active: false,
	})
	db.Create(&models.Rule{
		Rule:   "rule-3",
		Active: true,
	})
}

func createNews(db *gorm.DB) {
	fmt.Println("Seeding news")

	db.Create(&models.NewsItem{
		Title:  "Test Title",
		Body:   "News body goes here",
		Active: true,
	})

	db.Create(&models.NewsItem{
		Title:  "Test Title 2",
		Body:   "News body goes here 2",
		Active: true,
	})
}
