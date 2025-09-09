package conf

const (
	LocalUseridUint  = "user_id_uint"
	LocalUseridInt64 = "user_id_int64"
	LocalToken       = "token"
	LocalAuthority   = "authority"

	AdminUseridInt64 = "admin_id_int64"   //管理员id
	AdminUsername    = "admin_user_name"  //管理员用户名
	ManageRole       = "manage_user_role" //管理系统角色
	ManageUser       = "manage_user"      //管理系统用户
)

const MsgSuccess = 0
const MessageFail = -1
const TokenFail = -2
