// Link - https://leetcode.com/problems/angle-between-hands-of-a-clock/

import "math"

func angleClock(hour int, minutes int) float64 {
	var hourAngle float64 = float64(hour%12)*30 + float64(minutes%60)*0.5
	var minAngle float64 = float64(minutes%60) * 6
	if hourAngle < minAngle {
		hourAngle, minAngle = minAngle, hourAngle
	}
	return math.Min(hourAngle-minAngle, 360-hourAngle+minAngle)
}

// Time - 0ms(beats 100%)
// Memory - 2MB(beats 50%)