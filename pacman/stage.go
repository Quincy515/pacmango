package pacman

type stage struct {
	matrix []string
}

var stage1 = stage{
	[]string{
		"3888888884888888885",
		"gqrrrrrrrgrrrrrrrqg",
		"gr02r789ror789r02rg",
		"gracrrrrrrrrrrracrg",
		"grikr6r78489r6rikrg",
		"grrrrgrrrgrrrgrrrrg",
		"gr02rd89ror78fr02rg",
		"grikrgrrrrrrrgrikrg",
		"grrrror39v75rorrrrg",
		"gr39rrrguwtgrrr75rg",
		"grgrr6rl888nr6rrgrg",
		"gror7nrrrxrrrl9rorg",
		"grrrrrr6r6r6rrrrrrg",
		"gr012r3nrorl5r012rg",
		"grijkrorrprrorijkrg",
		"gqrrrrrr385rrrrrrqg",
		"l8888888m8m8888888n",
	},
}

var defaultStage = &stage1
