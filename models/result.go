package result

type Result struct {
	CriticalFailures  int `json:"critical_failures"`
	Failures          int `json:"failures"`
	Successes         int `json:"successes"`
	CriticalSuccesses int `json:"critical_successes"`
}
