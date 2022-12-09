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

type UserRepositorySaveReq struct {
	Parentld           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveResp struct {
	Identity string `json:"identity"`
}

type UserFileListReq struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFileListResp struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Name               string `json:"name"`
	Size               string `json:"size"`
	Path               string `json:"path"`
	Ext                string `json:"ext"`
}

type UserFileNameUpdateReq struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateResp struct {
}

type UserFolderCreateReq struct {
	Parentld int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderCreateResp struct {
	Identity string `json:"identity"`
}

type UserFileDeleteReq struct {
	Identity string `json:"identity"`
}

type UserFileDeleteResp struct {
}

type UserFileMoveReq struct {
	Identity       string `json:"identity"`        // 文件 uuid
	Parentldentity string `json:"parent_identity"` // 文件夹、目录 uuid
}

type UserFileMoveResp struct {
}

type FileUploadReq struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadResp struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type FileUploadPrepareReq struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

type FileUploadPrepareResp struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
}

type FileUploadChunkReq struct {
}

type FileUploadChunkResp struct {
	Etag string `json:"etag"` // md5
}

type FileUploadChunkCompleteReq struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type FileUploadChunkCompleteResp struct {
}

type CosObject struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}
