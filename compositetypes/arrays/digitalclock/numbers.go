package main

const (
	ZERO = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
)

type placeholder [5]string

var (
	zero = placeholder{
		"█████",
		"█   █",
		"█   █",
		"█   █",
		"█████",
	}

	one = placeholder{
		"███  ",
		"  █  ",
		"  █  ",
		"  █  ",
		"█████",
	}

	two = placeholder{
		"█████",
		"    █",
		"█████",
		"█    ",
		"█████",
	}
	three = placeholder{
		"█████",
		"    █",
		" ████",
		"    █",
		"█████",
	}

	four = placeholder{
		"█   █",
		"█   █",
		"█████",
		"    █",
		"    █",
	}

	five = placeholder{
		"█████",
		"█    ",
		"█████",
		"    █",
		"█████",
	}

	six = placeholder{
		"█████",
		"█    ",
		"█████",
		"█   █",
		"█████",
	}

	seven = placeholder{
		"█████",
		"    █",
		"    █",
		"    █",
		"    █",
	}

	eight = placeholder{
		"█████",
		"█   █",
		"█████",
		"█   █",
		"█████",
	}

	nine = placeholder{
		"█████",
		"█   █",
		"█████",
		"    █",
		"    █",
	}

	colon = placeholder{
		"     ",
		"  ░  ",
		"     ",
		"  ░  ",
		"     ",
	}

	alarm = [...]placeholder{
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"█ █",
		},
		{
			"█  ",
			"█  ",
			"█  ",
			"█  ",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"█ █",
		},
		{
			"██ ",
			"█ █",
			"██ ",
			"█ █",
			"█ █",
		},
		{
			"█ █",
			"███",
			"█ █",
			"█ █",
			"█ █",
		},
		{
			" █ ",
			" █ ",
			" █ ",
			"   ",
			" █ ",
		},
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}

	dot = placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		" ░ ",
	}
)

var digits = [...]placeholder{
	ZERO:  zero,
	ONE:   one,
	TWO:   two,
	THREE: three,
	FOUR:  four,
	FIVE:  five,
	SIX:   six,
	SEVEN: seven,
	EIGHT: eight,
	NINE:  nine,
}
