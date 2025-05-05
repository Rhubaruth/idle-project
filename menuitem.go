package main

type MenuItem struct {
	Name           string
	Cost           int64 // Score cost
	Count          int64 // How many times does player have of this
	UnlockScore    int64 // Score needed to unlock this upgrade
	IsUnlocked     bool
	ScorePerSecond int64
}

func InitalizeItems() []*MenuItem {
	items := []*MenuItem{
		{
			Name:           "Random Agent",
			Cost:           10,
			Count:          0,
			UnlockScore:    0,
			IsUnlocked:     true,
			ScorePerSecond: 1,
		},
		{
			Name:           "Monte Carlo",
			Cost:           100,
			Count:          0,
			UnlockScore:    50,
			IsUnlocked:     false,
			ScorePerSecond: 5,
		},
		{
			Name:           "Heuristics",
			Cost:           750,
			Count:          0,
			UnlockScore:    500,
			IsUnlocked:     false,
			ScorePerSecond: 20,
		},
		{
			Name:           "Clustering",
			Cost:           3000,
			Count:          0,
			UnlockScore:    1500,
			IsUnlocked:     false,
			ScorePerSecond: 100,
		},
	}

	return items
}
