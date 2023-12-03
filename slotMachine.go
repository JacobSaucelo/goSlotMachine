package main

var wearRatingType = []wearRating{
	{
		wearName: "Factory New",
		rating:   0.07,
	},
	{
		wearName: "Minimal Wear",
		rating:   0.15,
	},
	{
		wearName: "Field-Tested",
		rating:   0.37,
	},
	{
		wearName: "Well-Worn",
		rating:   0.44,
	},
	{
		wearName: "Battle-Scarred",
		rating:   1.00,
	},
}

type item struct {
	name       string
	rarity     string
	wearRating float64
	isStatTrak bool
	gunType    string
}

type wearRating struct {
	wearName string
	rating   float64
}
