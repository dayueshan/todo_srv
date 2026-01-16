// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Todo is the golang structure of table todo for DAO operations like Where/Data.
type Todo struct {
	g.Meta    `orm:"table:todo, do:true"`
	Id        any // TODO ID
	UserId    any // 用户ID
	Content   any // TODO内容
	Status    any // 是否完成，0否、1是)
	CreatedAt any // 创建时间
	UpdatedAt any // 更新时间
}
