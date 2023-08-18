package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库结构体
var Db *sqlx.DB

func init() {
	//连接数据库，"用户名:密码@协议(地址:端口)/数据库名称"
	database, err := sqlx.Open("mysql", sqlstring)
	if err != nil {
		fmt.Println("打开数据库失败,err:", err)
		return
	}
	//初始化数据库
	Db = database
}

// 写sql
func sql_insert(fakeWeb FakeWebResults) {
	r, err := Db.Exec("insert into fake_web (ip,port,domain,country,title,protocol,country_name,region,city,longitude,latitude,as_number,as_organization,host,os,server,icp,jarm,header,banner,cert,base_protocol,cname,hashvalue) value(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", fakeWeb.IP, fakeWeb.Port, fakeWeb.Domain, fakeWeb.Country, fakeWeb.Title, fakeWeb.Protocol, fakeWeb.Country_name, fakeWeb.Region, fakeWeb.City, fakeWeb.Longitude, fakeWeb.Latitude, fakeWeb.As_number, fakeWeb.As_organization, fakeWeb.Host, fakeWeb.Os, fakeWeb.Server, fakeWeb.Icp, fakeWeb.Jarm, fakeWeb.Header, fakeWeb.Banner, fakeWeb.Cert, fakeWeb.Base_protocol, fakeWeb.Cname, fakeWeb.HashValue)
	if err != nil {
		fmt.Println("修改失败,err:", err)
		return
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("修改失败,err:", err)
		return
	}
}

// 查询是否存在
func queryRow(fakeWeb FakeWebResults) {
	var f FakeWebResults
	sqlStr := "select hashvalue from fake_web where hashvalue=?"
	err := Db.QueryRow(sqlStr, fakeWeb.HashValue).Scan(&f.HashValue)
	if err == sql.ErrNoRows {
		vxRobot(false, fakeWeb.IP, fakeWeb.Port, fakeWeb.Host, fakeWeb.Domain, fakeWeb.Country)
		log.Println("[新增]", fakeWeb.HashValue)
		sql_insert(fakeWeb)
		return
	} else if err == nil {
		log.Println("[重复]", fakeWeb.HashValue)
		return
	}
}
func writeSql(results []FakeWebResults) {
	for _, v := range results {
		queryRow(v)
	}
}

func selectSqlAliveUrls() []string {
	var urls []string
	rows, err := Db.Query("SELECT host FROM fake_web")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 创建一个通道用于从goroutine接收有效的URL
	validURLsCh := make(chan string)
	wg := sync.WaitGroup{}
	// 启动固定数量的goroutine以并发处理URL
	numWorkers := 50 // 根据需要进行调整
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rows.Next() {
				var host string
				if err := rows.Scan(&host); err != nil {
					log.Println(err)
					continue
				}
				if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
					host = "http://" + host
				}
				log.Println(host)
				check := checkURL(host)
				if check != "" {
					log.Println(check + host)
					validURLsCh <- check + host
				}
			}
		}()
	}

	// 在所有goroutine完成后关闭通道
	go func() {
		wg.Wait()
		close(validURLsCh)
	}()

	// 从通道中读取有效的URL
	for validURL := range validURLsCh {
		urls = append(urls, validURL)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return urls
}
