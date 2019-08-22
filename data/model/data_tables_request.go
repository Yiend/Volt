package model

type DataTablesRequest struct {
	//分页大小
	  PageSize int
	//页码
	 PageIndex int
	//排序字段
	 OrderBy string
	//是否倒序
	 Isdesc bool
	//记录数
	 TotalItemCount int
	//带升降序的排序字段 例如： Name desc
	 OrderByWithAcsOrDesc string
}

