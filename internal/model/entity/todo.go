// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Todo is the golang structure for table todo.
type Todo struct {
	Id        uint64 `json:"id"        orm:"id"         description:"TODO ID"`     // TODO ID
	UserId    uint64 `json:"userId"    orm:"user_id"    description:"用户ID"`        // 用户ID
	Content   string `json:"content"   orm:"content"    description:"TODO内容"`      // TODO内容
	Status    int    `json:"status"    orm:"status"     description:"是否完成，0否、1是)"` // 是否完成，0否、1是)
	CreatedAt int    `json:"createdAt" orm:"created_at" description:"创建时间"`        // 创建时间
	UpdatedAt int    `json:"updatedAt" orm:"updated_at" description:"更新时间"`        // 更新时间
}
