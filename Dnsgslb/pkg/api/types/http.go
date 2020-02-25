package types


type HttpCheck struct {
	Type 		string	`json:"type"`
	Url			string	`json:"url"`
	// interval between checks
	Interval	int		`json:"interval"`
	// check timeout
	Timeout		int		`json:"timeout"`
	// number of failed checks to disable record
	Fall		int		`json:"fall"`
	// number of successful checks to enable record
	Rise		int		`json:"rise"`
}