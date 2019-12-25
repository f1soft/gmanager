// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package sys_role_menu

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table sys_role_menu.
type Entity struct {
	Id     int64 `orm:"id,primary" json:"id"`      // 主键
	RoleId int64 `orm:"role_id"    json:"role_id"` // 角色id
	MenuId int64 `orm:"menu_id"    json:"menu_id"` // 菜单id
}

// SysRoleMenu is alias of Entity, which some developers say they just want.
type SysRoleMenu = Entity

// New creates and returns an *Entity.
func New() *Entity {
	return new(Entity)
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
