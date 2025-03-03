package services

import result "brunaoliveira/pathfinder/models"

func CalculateDegrees(dc int, modifier int) map[string]int {
	var result result.Result

	result.CriticalFailures = max(0, min(20, dc-10-modifier))
	result.Failures = max(0, min(20, dc-modifier-1)) - result.CriticalFailures
	result.Successes = max(0, min(20, dc+10-modifier-1)) - result.Failures - result.CriticalFailures
	result.CriticalSuccesses = 20 - result.CriticalFailures - result.Failures - result.Successes

	result = AjustNaturalOne(modifier, dc, result)
	result = AdjustNaturalTwenty(modifier, dc, result)

	return map[string]int{
		"critical_failures":  result.CriticalFailures,
		"failures":           result.Failures,
		"successes":          result.Successes,
		"critical_successes": result.CriticalSuccesses,
	}
}

func AjustNaturalOne(modifier int, dc int, result result.Result) result.Result {
	if modifier+1 >= dc+10 { // critical success -> success
		result.CriticalSuccesses--
		result.Successes++
	} else if modifier+1 >= dc { // success -> failure
		result.Successes--
		result.Failures++
	} else if modifier+1 >= dc-10 { // failure -> critical failure
		result.Failures--
		result.CriticalFailures++
	}

	return result
}

func AdjustNaturalTwenty(modifier int, dc int, result result.Result) result.Result {

	if modifier+20 <= dc-10 { // critical failures -> failures
		result.CriticalFailures--
		result.Failures++
	} else if modifier+20 < dc { // failures -> successes
		result.Failures--
		result.Successes++
	} else if modifier+20 < dc+10 { // successes -> critical successes
		result.Successes--
		result.CriticalSuccesses++
	}

	return result
}
