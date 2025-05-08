package main

type MenuItem struct {
	Name        string
	Cost        int64 // Score cost
	Count       int64 // How many times does player have of this
	UnlockScore int64 // Score needed to unlock this upgrade
	IsUnlocked  bool

	ScorePerSecond int64
	Instability    float64
}

func InitalizeItems() []*MenuItem {
	items := []*MenuItem{
		{
			Name:           "Rustilng of Leaves",
			Cost:           5,
			Count:          0,
			UnlockScore:    0,
			IsUnlocked:     true,
			ScorePerSecond: 1,
			Instability:    0,
		},
		{
			Name:           "Waterstream babbling",
			Cost:           20,
			Count:          0,
			UnlockScore:    10,
			IsUnlocked:     false,
			ScorePerSecond: 5,
			Instability:    0,
		},
		{
			Name:           "Birds Chirping",
			Cost:           100,
			Count:          0,
			UnlockScore:    20,
			IsUnlocked:     false,
			ScorePerSecond: 10,
			Instability:    0,
		},
		{
			Name:           "Children Talking",
			Cost:           500,
			Count:          0,
			UnlockScore:    100,
			IsUnlocked:     false,
			ScorePerSecond: 20,
			Instability:    0,
		},
		{
			Name:           "Cars Driving By",
			Cost:           1000,
			Count:          0,
			UnlockScore:    500,
			IsUnlocked:     false,
			ScorePerSecond: 50,
			Instability:    0,
		},
		{
			Name:           "Train Horn",
			Cost:           3000,
			Count:          0,
			UnlockScore:    1000,
			IsUnlocked:     false,
			ScorePerSecond: 100,
			Instability:    0.1,
		},
		{
			Name:           "Neighbor Playing Music",
			Cost:           7500,
			Count:          0,
			UnlockScore:    3000,
			IsUnlocked:     false,
			ScorePerSecond: 150,
			Instability:    0.05,
		},
		{
			Name:           "Children Playing",
			Cost:           10000,
			Count:          0,
			UnlockScore:    7500,
			IsUnlocked:     false,
			ScorePerSecond: 200,
			Instability:    0.01,
		},
		{
			Name:           "Amateur Garage Band",
			Cost:           30000,
			Count:          0,
			UnlockScore:    10000,
			IsUnlocked:     false,
			ScorePerSecond: 500,
			Instability:    0.1,
		},
		{
			Name:           "Neighbor Morning Drilling",
			Cost:           100000,
			Count:          0,
			UnlockScore:    30000,
			IsUnlocked:     false,
			ScorePerSecond: 800,
			Instability:    0.15,
		},
		{
			Name:           "Background TV Show",
			Cost:           500000,
			Count:          0,
			UnlockScore:    100000,
			IsUnlocked:     false,
			ScorePerSecond: 1000,
			Instability:    0.2,
		},
		{
			Name:           "Unexpected Advertisment",
			Cost:           850000,
			Count:          0,
			UnlockScore:    150000,
			IsUnlocked:     false,
			ScorePerSecond: 1200,
			Instability:    0.3,
		},
		{
			Name:           "Metal Pipe Sound",
			Cost:           1000000,
			Count:          0,
			UnlockScore:    200000,
			IsUnlocked:     false,
			ScorePerSecond: 1500,
			Instability:    0.5,
		},
	}

	return items
}
