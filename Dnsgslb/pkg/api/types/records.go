package types


type Records struct {
	Domain 		string 		`json:"domain"`
	Name  		string		`json:"name"`
	Type 		string		`json:"type"`
	Content  	string		`json:"content"`
	TTL 		int			`json:"ttl"`
	Disable		string		`json:"disable"`
	Weight		string		`json:"weight"`
	Monitors 	string		`json:"monitors"`
	View		string		`json:"view"`
}
