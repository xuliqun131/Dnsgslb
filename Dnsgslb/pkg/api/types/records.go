package types


type Name struct {
	Name 		string		`json:"name"`
}

type NoCheck struct {
	Type 		string 		`json:"type"`
}

type Domains struct {
	Id 			int			`json:"id"`
	Domain 		string		`json:"domain"`
}

type ContentType struct {
	Id 			int			`json:"id"`
	Content		string		`json:"content"`
}

type Views struct {
	Id 			int 		`json:"id"`
	View		string		`json:"view"`
	Rule 		string		`json:"rule"`
}

type DnsType struct {
	Value 		int			`json:"value"`
	Type 		string		`json:"type"`
	Description string		`json:"description"`
}


type Monitors struct {
	Id 				int					`json:"id"`
	Monitor			string				`json:"monitor"`
	Type 			string				`json:"type"`
	Content 		string				`json:"content"`
	Port 			string				`json:"port"`
	Interval 		int					`json:"interval"`
	Timeout 		int 				`json:"timeout"`

}


type Records struct {
	Id   		int64 		`json:"id"`
	Domain 		string 		`json:"domain"`
	Name  		string		`json:"name"`
	Type 		string		`json:"type"`
	Content  	string		`json:"content"`
	TTL 		int			`json:"ttl"`
	Disable		int			`json:"disable"`
	Weight		int			`json:"weight"`
	Fallback	int 		`json:"fallback"`
	Monitors 	string		`json:"monitors"`
	Persistence int			`json:"persistence"`
	View		string		`json:"view"`
	Status 		bool		`json:"status"`
}

type ListRecords struct {
	Id 				int 				`json:"id"`
	ListRecord 		Records			`json:"list_record"`
}

type Listdomain struct {
	Id 			int			`json:"id"`
	Domain 		string		`json:"domain"`
}
