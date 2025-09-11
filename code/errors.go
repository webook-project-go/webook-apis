package code

const (
	//user
	// 用户输入错误
	ErrCodeInvalidEmailOrPassword = 10101 // 邮箱或密码错误

	// 业务逻辑错误
	ErrCodeEmailDuplicate = 10201 // 邮箱已存在
	ErrCodeUserNotExist   = 10202 // 用户不存在
)
const (
	// code
	// 业务逻辑错误
	ErrCodeSendTooFrequent      = 11201 // 发送过于频繁
	ErrCodeTooManyVerifications = 11202 // 验证次数过多

	// 客户端输入错误
	ErrCodeInvalidCode = 11101 // 验证码格式无效
	ErrCodeWrongCode   = 11102 // 验证码错误

)
