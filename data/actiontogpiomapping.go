package data

type ActionToGpioMapping struct {
	Action  Action `json:"action"`
	GpioPin int    `json:"gpioPin"`
}
