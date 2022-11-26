package main

// 一份setu的结构体
type setu struct {
	Pid    int      `json:"pid"`
	P      int      `json:"p"`
	Uid    int      `json:"uid"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	R18    bool     `json:"r18"`
	Tags   []string `json:"tags"`
	Ext    string   `json:"ext"`
	Urls   string   `json:"urls"`
}

// setuAPI返回的结构体
type setuList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []setu `json:"data"`
	Tips    string `json:"tips"`
}

// 返回一个setu的结构体
func newSetu(pid int, p int, uid int, title string, author string, r18 bool, tags []string, ext string, urls string) setu {
	return setu{
		Pid:    pid,
		P:      p,
		Uid:    uid,
		Title:  title,
		Author: author,
		R18:    r18,
		Tags:   tags,
		Ext:    ext,
		Urls:   urls,
	}
}

// setuAPI返回值的结构体
func newSetuList(code int, message string, data []setu) setuList {
	return setuList{
		Code:    code,
		Message: message,
		Data:    data,
		Tips:    "后可接请求参数: tag num r18, example: ?tag=loli&num=10&r18=true",
	}
}
