package lasagna

const OvenTimeInMins int = 40
const LayerPrepTimeInMins int = 2

func OvenTime() int {
	return OvenTimeInMins
}

func RemainingOvenTime(minsPassed int) int {
	return OvenTimeInMins - minsPassed
}

func PreparationTime(layers int) int {
	return layers * LayerPrepTimeInMins
}

func ElapsedTime(layers int, minsPassed int) int {
	return PreparationTime(layers) + minsPassed
}
