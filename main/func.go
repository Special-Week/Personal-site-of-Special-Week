package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

// 从slice中随机获取一个元素
func randomChoiceItem(slice []interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 随机脏话
func selectDirtyLanguage() map[string]any {
	db, err := sql.Open("sqlite3", "./database/zhuan.db")                                 // 打开数据库
	checkErr(err)                                                                         // 检查错误
	defer db.Close()                                                                      // 关闭数据库
	sqlCommand := "select id,text from main where level='max'  order by random() limit 1" // 随机获取一条脏话
	rows, err := db.Query(sqlCommand)                                                     // 执行sql语句
	checkErr(err)                                                                         // 检查错误
	var id int                                                                            // 定义id
	var text string                                                                       // 定义脏话
	for rows.Next() {                                                                     // 遍历结果
		err = rows.Scan(&id, &text) // 获取结果
		checkErr(err)               // 检查错误
	}
	return map[string]any{ // 返回结果
		"id":   id,
		"text": text,
	}
}

func selectSetu(keyword string, argum int, argr18 bool) setuList {
	db, err := sql.Open("sqlite3", "./database/setu.db") // 打开数据库
	checkErr(err)                                        // 检查错误
	defer db.Close()                                     // 关闭数据库
	r18String := "False"                                 // 定义r18String
	numString := fmt.Sprintf("%d", argum)                // 将int转换为string
	if argr18 {
		r18String = "True" // 如果argr18为true,则r18String为True
	} else {
		r18String = "False" // 如果argr18为false,则r18String为False
	}
	keyword = strings.ReplaceAll(keyword, "'", "") // 去除keyword中的单引号, 可能会导致sql注入
	sqlCommand := "select * from main where (tags like '%" + keyword + "%' or title like '%" + keyword + "%' or author like '%" + keyword + "%')and r18 = '" + r18String + "' order by random() limit " + numString
	rows, err := db.Query(sqlCommand) // 执行sql语句
	checkErr(err)                     // 检查错误
	data := make([]setu, 0)           // 定义data
	for rows.Next() {                 // 遍历结果
		var pid, p, uid int                                                        // 定义pid,p,uid
		var title, author, r18, tags, ext, urls string                             // 定义title,author,r18,tags,ext,urls
		err = rows.Scan(&pid, &p, &uid, &title, &author, &r18, &tags, &ext, &urls) // 获取结果
		checkErr(err)                                                              // 检查错误
		tagsArray := strings.Split(tags, ",")                                      // 将tags转换为数组
		newTagsArray := make([]string, 0)                                          // 定义newTagsArray
		for _, tag := range tagsArray {                                            // 遍历tagsArray
			tag = strings.Trim(tag, " ")             // 去除tag两边的空格
			tag = strings.Trim(tag, "\"")            // 去除tag两边的双引号
			newTagsArray = append(newTagsArray, tag) // 将tag添加到newTagsArray中
		}
		r18Bool := false // 定义r18Bool
		if r18 == "True" {
			r18Bool = true // 如果r18为True,则r18Bool为true
		} else {
			r18Bool = false // 如果r18为False,则r18Bool为false
		}
		setu := newSetu(pid, p, uid, title, author, r18Bool, newTagsArray, ext, urls) // 创建setu
		data = append(data, setu)                                                     // 将setu添加到data中
	}
	return newSetuList(200, "success", data) // 返回结果
}
