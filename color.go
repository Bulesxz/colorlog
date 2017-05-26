package main

type ColorType int
type ColorCode string

const (
	ColorOff ColorType = iota

	Black
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White

	BoldBlack
	BoldRed
	BoldGreen
	BoldYellow
	BoldBlue
	BoldPurple
	BoldCyan
	BoldWhite

	UnderlineBlack
	UnderlineRed
	UnderlineGreen
	UnderlineYellow
	UnderlineBlue
	UnderlinePurple
	UnderlineCyan
	UnderlineWhite

	BackgroundBlack
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundPurple
	BackgroundCyan
	BackgroundWhite

	HighIntensityBlack
	HighIntensityRed
	HighIntensityGreen
	HighIntensityYellow
	HighIntensityBlue
	HighIntensityPurple
	HighIntensityCyan
	HighIntensityWhite

	BoldHighIntensityBlack
	BoldHighIntensityRed
	BoldHighIntensityGreen
	BoldHighIntensityYellow
	BoldHighIntensityBlue
	BoldHighIntensityPurple
	BoldHighIntensityCyan
	BoldHighIntensityWhite

	HighIntensityBackgroundBlack
	HighIntensityBackgroundRed
	HighIntensityBackgroundGreen
	HighIntensityBackgroundYellow
	HighIntensityBackgroundBlue
	HighIntensityBackgroundPurple
	HighIntensityBackgroundCyan
	HighIntensityBackgroundWhite
)

var colors = map[ColorType]ColorCode{
	ColorOff: "\033[0m",
	Black:    "\033[0;30m",
	Red:      "\033[0;31m",
	Green:    "\033[0;32m",
	Yellow:   "\033[0;33m",
	Blue:     "\033[0;34m",
	Purple:   "\033[0;35m",
	Cyan:     "\033[0;36m",
	White:    "\033[0;37m",

	BoldBlack:  "\033[1;30m",
	BoldRed:    "\033[1;31m",
	BoldGreen:  "\033[1;32m",
	BoldYellow: "\033[1;33m",
	BoldBlue:   "\033[1;34m",
	BoldPurple: "\033[1;35m",
	BoldCyan:   "\033[1;36m",
	BoldWhite:  "\033[1;37m",

	UnderlineBlack:  "\033[4;30m",
	UnderlineRed:    "\033[4;31m",
	UnderlineGreen:  "\033[4;32m",
	UnderlineYellow: "\033[4;33m",
	UnderlineBlue:   "\033[4;34m",
	UnderlinePurple: "\033[4;35m",
	UnderlineCyan:   "\033[4;36m",
	UnderlineWhite:  "\033[4;37m",

	BackgroundBlack:  "\033[40m",
	BackgroundRed:    "\033[41m",
	BackgroundGreen:  "\033[42m",
	BackgroundYellow: "\033[43m",
	BackgroundBlue:   "\033[44m",
	BackgroundPurple: "\033[45m",
	BackgroundCyan:   "\033[46m",
	BackgroundWhite:  "\033[47m",

	HighIntensityBlack:  "\033[0;90m",
	HighIntensityRed:    "\033[0;91m",
	HighIntensityGreen:  "\033[0;92m",
	HighIntensityYellow: "\033[0;93m",
	HighIntensityBlue:   "\033[0;94m",
	HighIntensityPurple: "\033[0;95m",
	HighIntensityCyan:   "\033[0;96m",
	HighIntensityWhite:  "\033[0;97m",

	BoldHighIntensityBlack:  "\033[1;90m",
	BoldHighIntensityRed:    "\033[1;91m",
	BoldHighIntensityGreen:  "\033[1;92m",
	BoldHighIntensityYellow: "\033[1;93m",
	BoldHighIntensityBlue:   "\033[1;94m",
	BoldHighIntensityPurple: "\033[1;95m",
	BoldHighIntensityCyan:   "\033[1;96m",
	BoldHighIntensityWhite:  "\033[1;97m",

	HighIntensityBackgroundBlack:  "\033[0;100m",
	HighIntensityBackgroundRed:    "\033[0;101m",
	HighIntensityBackgroundGreen:  "\033[0;102m",
	HighIntensityBackgroundYellow: "\033[0;103m",
	HighIntensityBackgroundBlue:   "\033[0;104m",
	HighIntensityBackgroundPurple: "\033[0;105m",
	HighIntensityBackgroundCyan:   "\033[0;106m",
	HighIntensityBackgroundWhite:  "\033[0;107m",
}
