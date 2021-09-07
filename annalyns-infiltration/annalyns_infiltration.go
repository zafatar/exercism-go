package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	canSpy := false

	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		canSpy = true
	}

	return canSpy
}

// CanSignalPrisoner can be executed if prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	canSignal := false

	if !archerIsAwake && prisonerIsAwake {
		canSignal = true
	}

	return canSignal
}

// CanFreePrisoner can be executed if prisoner is awake and the other 3 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	canFree := false

	if (!knightIsAwake && !archerIsAwake && prisonerIsAwake) || (!archerIsAwake && petDogIsPresent) {
		canFree = true
	}

	return canFree
}
