// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta         `orm:"table:user, do:true"`
	Id             any // 用户ID
	Provider       any // 登录第三方(gmail/facebook/github)
	ProviderUserId any // 第三方用户ID
	Name           any // 用户名
	Email          any // 邮箱
	CreatedAt      any // 创建时间
	UpdatedAt      any // 更新时间
}
