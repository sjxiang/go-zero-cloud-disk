// Code generated by goctl. DO NOT EDIT.
package types

type UserLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type UserDetailReq struct {
	Identity string `path:"identity"`
}

type UserDetailResp struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type VerifyCodeSendReq struct {
	Email string `json:"email"`
}

type VerifyCodeSendResp struct {
	Code string `json:"code"`
}

type UserRegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterResp struct {
}
