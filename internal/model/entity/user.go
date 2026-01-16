// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id             uint64 `json:"id"             orm:"id"               description:"用户ID"`                         // 用户ID
	Provider       string `json:"provider"       orm:"provider"         description:"登录第三方(gmail/facebook/github)"` // 登录第三方(gmail/facebook/github)
	ProviderUserId string `json:"providerUserId" orm:"provider_user_id" description:"第三方用户ID"`                      // 第三方用户ID
	Name           string `json:"name"           orm:"name"             description:"用户名"`                          // 用户名
	Email          string `json:"email"          orm:"email"            description:"邮箱"`                           // 邮箱
	CreatedAt      int    `json:"createdAt"      orm:"created_at"       description:"创建时间"`                         // 创建时间
	UpdatedAt      int    `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`                         // 更新时间
}
