package main

type Dog struct {
	Uuid     string `json:"Uuid"`
	Title    string `json:"Title"`
	Subtitle string `json:"Subtitle"`
	Sex      string `json:"Sex"`
	Preview  string `json:"Preview"`
	Picture1 string `json:"Picture1"`
	Picture2 string `json:"Picture2"`
	Picture3 string `json:"Picture3"`
	Picture4 string `json:"Picture4"`
	Picture5 string `json:"Picture5"`
	Memo     string `json:"Memo"`
	Like     int    `json:"Like"`
}

type Pup struct {
	Uuid     string `json:"Uuid"`
	Title    string `json:"Title"`
	Subtitle string `json:"Subtitle"`
	Sex      string `json:"Sex"`
	Preview  string `json:"Preview"`
	Picture1 string `json:"Picture1"`
	Picture2 string `json:"Picture2"`
	Picture3 string `json:"Picture3"`
	Picture4 string `json:"Picture4"`
	Picture5 string `json:"Picture5"`
	Memo     string `json:"Memo"`
	Like     int    `json:"Like"`
}

type Comment struct {
	Uuid     string `json:"Uuid"`
	Title    string `json:"Title"`
	Content  string `json:"Content"`
	DateTime int64  `json:"DateTime"`
}

type Global struct {
	AdminName string `json:"AdminName"`
	AdminPass string `json:"AdminPass"`
	Introduce string `json:"Introduce"`
}
