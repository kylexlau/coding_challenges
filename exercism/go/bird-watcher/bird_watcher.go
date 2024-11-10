package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, b := range birdsPerDay {
		sum = sum + b
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	week_start := (week - 1) * 7
	week_end := week_start + 7
	//	if len(birdsPerDay) < week_end {
	//		week_end = len(birdsPerDay)
	//	}
	for _, b := range birdsPerDay[week_start:week_end] {
		sum = sum + b
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}
	return birdsPerDay
}
