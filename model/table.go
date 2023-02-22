package model

type Tables struct {
	Schemaname string
	Tablename  string
}

// TableName 定义表名
func (Tables) TableName() string {
	return "pg_table"
}
