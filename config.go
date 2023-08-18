package main

var (
	fofaEmail  string   = ""
	fofaKey    string   = ""
	size       int      = 1000
	fields     string   = "ip,port,domain,country,title,protocol,country_name,region,city,longitude,latitude,as_number,as_organization,host,os,server,icp,jarm,header,banner,cert,base_protocol,cname"
	qbase64s   []string = []string{"dGl0bGU9ImJhaWR1Ig%3D%3D", "dGl0bGU9ImFsaSI%3D"} //查询语法的base64,可添加多个
	page       int      = 1                                                          //勿动
	i                   = 0                                                          //勿动
	vxRobotKey string   = "123123123"                                                //企业微信机器key
	sqlstring  string   = "root:123456@tcp(1.1.1.1:3306)/test"
)
