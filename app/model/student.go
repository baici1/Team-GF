package model

// StuApiSignUpReq 用于学生注册请求参数。
type StuApiSignUpReq struct {
	Stuid     string `v:"required|length:10,12#学号不能为空|学号长度应该在:min到:max之间" `
	Password  string `v:"required|password2#密码不能为空|密码强度不够，长度在6~18之间,必须包含大小写字母和数字！"`
	Password2 string `v:"required|password2|same:Password#第二次密码不能为空|密码强度不够，长度在6~18之间,必须包含大小写字母和数字！|两次密码输入不相等！"`
}

// StuServiceSignUpReq 学生业务函数输入参数
type StuServiceSignUpReq struct {
	Stuid    string
	Password string
	ID       string
}

// StuApiSignInReq 学生登录的请求参数
type StuApiSignInReq struct {
	Stuid    string `json:"stuid" v:"required#学号不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

// StuApiSubmitDataReq 学生提交信息
type StuApiSubmitDataReq struct {
	Name      string `json:"name"`
	Gender    int    `json:"gender" v:"in:1,0,...#性别请选择男或者女"`
	Email     string `json:"email" v:"email#请输入正确的邮箱"`
	Introduce string `json:"introduce"`
}

// StuApiGetDetailRes 返回学生相关信息
type StuApiGetDetailRes struct {
	Stuid     string `json:"stuid,omitempty"`
	Name      string `json:"name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Email     string `json:"email,omitempty"`
	Introduce string `json:"introduce,omitempty"`
}
