package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]

	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekEnd := week * 7
	weekStart := weekEnd - 7
	return TotalBirdCount(birdsPerDay[weekStart:weekEnd])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	newBirds := birdsPerDay
	for i := 0; i < len(newBirds); i++ {
		if i%2 == 0 {
			newBirds[i] += 1
		}
	}
	return newBirds
}
