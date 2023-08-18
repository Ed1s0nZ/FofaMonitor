package main

type FakeWebInfo struct {
	Results [][]string `json:"results"`
	Size    int        `json:"size"`
}

type FakeWebResults struct {
	IP              string `json:"ip" db:"ip"` //ip,port,domain,country,title,protocol,country_name,region,city,longitude,latitude,as_number,as_organization,host,os,server,icp,jarm,header,banner,cert,base_protocol,body,cname
	Port            string `json:"port" db:"port"`
	Domain          string `json:"domain" db:"domain"`
	Country         string `json:"country" db:"country"`
	Title           string `json:"title" db:"title"`
	Protocol        string `json:"protocol" db:"protocol"`
	Country_name    string `json:"country_name" db:"country_name"`
	Region          string `json:"region" db:"region"`
	City            string `json:"city" db:"city"`
	Longitude       string `json:"longitude" db:"longitude"`
	Latitude        string `json:"latitude" db:"latitude"`
	As_number       string `json:"as_number" db:"as_number"`
	As_organization string `json:"as_organization" db:"as_organization"`
	Host            string `json:"host" db:"host"`
	Os              string `json:"os" db:"os"`
	Server          string `json:"server" db:"server"`
	Icp             string `json:"icp" db:"icp"`
	Jarm            string `json:"jarm" db:"jarm"`
	Header          string `json:"header" db:"header"`
	Banner          string `json:"banner" db:"banner"`
	Cert            string `json:"cert" db:"cert"`
	Base_protocol   string `json:"base_protocol" db:"base_protocol"`
	Link            string `json:"link" db:"link"`
	Cname           string `json:"cname" db:"base_protocol"`
	HashValue       string //md5(ip+port+domain+host)
}

type MediaInfo struct {
	ErrMsg  string `json:"errmsg"`
	Type    string `json:"type"`
	MediaID string `json:"media_id"`
}

type Message struct {
	MsgType string `json:"msgtype"`
	File    struct {
		MediaID string `json:"media_id"`
	} `json:"file"`
}

type UrlsAlive struct {
	Shields []string
	Alive   []string
}
