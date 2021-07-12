package model

import "time"

type PublicRequest struct {
	Askorgcode		string			`json:"askorgcode"`
	Askorgname		string			`json:"askorgname"`
	Targetorgcode	string			`json:"targetorgcode"`
	Asktime			time.Time		`json:"asktime"`
	Version			string			`json:"version"`
	Signmsg			string			`json:"signmsg"`
}

type IntegratedQueryRequest struct {
	Publicrequest		PublicRequest	`json:"publicrequest"`
	Sql					string			`json:"sql"`
	Dbname				string			`json:"dbname"`
	Parameter			string			`json:"parameter"`
}

func RequestInstance(targetOrgCode string) PublicRequest {
	var pub PublicRequest
	pub.Askorgcode = "GO"
	pub.Askorgname = "GO"
	pub.Targetorgcode = targetOrgCode
	pub.Asktime = time.Now()
	pub.Version = "1.0.0"
	pub.Signmsg = "GO"
	return pub
}
