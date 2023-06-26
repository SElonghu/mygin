package common

type PageInfo struct {
	PageSize int `form:"pagesize"`
	PageNum  int `form:"pagenum"`
}
