package viewmodel

import "github.com/kataras/iris"

type DataTablesResult struct {
	//总条数
	recordsTotal int64
	//总条数
	recordsFiltered int64
	// 数据
	data interface{}
}

func (this DataTablesResult) ToJson() iris.Map {
	return iris.Map{"recordsTotal":this.recordsTotal,"recordsFiltered":this.recordsFiltered,"data":this.data}
}

func NewDataTablesResult(total int64,data interface{})DataTablesResult  {
	return DataTablesResult{
		recordsTotal:total,
		recordsFiltered:total,
		data:data,
	}
}