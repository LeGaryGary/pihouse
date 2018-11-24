package data

type Action int

const (
	LivingRoomLightsOff    Action = 0
	LivingRoomLightsOn     Action = 1
	LivingRoomLightsToggle Action = 2
	HeatingOn              Action = 3
	HeatingOff             Action = 4
	HeatingToggle          Action = 5
	Swearing               Action = 6
	LightsHereOn           Action = 7
	LightsHereOff          Action = 8
	LightsHereToggle       Action = 9
)

func (action Action) String() string {
	names := [...]string{
		"LivingRoomLightsOff",
		"LivingRoomLightsOn",
		"LivingRoomLightsToggle",
		"HeatingOn",
		"HeatingOff",
		"HeatingToggle",
		"Swearing",
		"LightsHereOn",
		"LightsHereOff",
		"LightsHereToggle",
	}

	if action < LivingRoomLightsOff || action > LightsHereToggle {
		return "Unknown value: " + action.String()
	}

	return names[action]
}
