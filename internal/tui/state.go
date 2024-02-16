package tui

type state struct {
	Details details
}

type details struct {
	TotalTests  int
	TotalPassed int
	TotalFailed int
}

func NewState() state {
	return state{
		Details: details{
			TotalTests:  0,
			TotalPassed: 0,
			TotalFailed: 0,
		},
	}
}
