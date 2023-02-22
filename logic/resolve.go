package logic

import (
	"github.com/labstack/echo/v4"

	"dbAdmin/frame"
	"dbAdmin/types"
)

type resolve struct {
	context echo.Context
}

func NewResolveLogic(context echo.Context) *resolve {
	return &resolve{context: context}
}

func (*resolve) TableList() []types.ResolveTables {

	type tableResolve struct {
		Comment      string
		TableComment string
		TableName    string
		Name         string
	}

	var (
		tables          []string
		tableResolves   []tableResolve
		tableMap        = make(map[string][]types.ResolveColumns)
		tableCommentMap = make(map[string]string)
		resp            []types.ResolveTables
	)

	frame.UserDB.Table("pg_tables").Where("schemaname = 'public'").
		Joins("").
		Pluck("tablename", &tables)

	frame.UserDB.Table("pg_class").
		Joins("join pg_attribute on pg_class.oid = pg_attribute.attrelid").
		Where("pg_attribute.attnum > 0 and pg_class.relname in ?", tables).
		Select(
			"col_description(pg_attribute.attrelid, pg_attribute.attnum) as comment",
			"pg_class.relname as table_name",
			"pg_attribute.attname as name",
			"cast(obj_description(relfilenode,'pg_class') as varchar) as table_comment",
		).
		Find(&tableResolves)

	for _, table := range tableResolves {
		tableMap[table.TableName] = append(tableMap[table.TableName], types.ResolveColumns{
			Name:    table.Name,
			Comment: table.Comment,
		})
		tableCommentMap[table.TableName] = table.TableComment
	}

	for tableName, resolveColumns := range tableMap {
		resp = append(resp, types.ResolveTables{
			Name:    tableName,
			Comment: tableCommentMap[tableName],
			Columns: resolveColumns,
		})
	}

	return resp
}
