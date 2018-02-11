// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package pheromone

const (
	ErrLocalSocketTimeout = 1001

	ErrRemoteSocketEmpty  = 2001
	ErrRemoteSocketExist  = 2002
	ErrRemoteSocketMisType  = 2003

	ErrUnKnownProtocal     = 3003
	ErrMismatchProtocalReq = 3101
	ErrMismatchProtocalConnectReq = 3102
	ErrMismatchProtocalResp = 3202

	ErrUnknuowPeer		= 4001
)

type Error int

func (err Error) Error() string {
	return errMap[err]
}

var errMap = map[Error]string{
	ErrLocalSocketTimeout : "链接超时",

	ErrRemoteSocketEmpty  : "链接为空",
	ErrRemoteSocketExist  : "链接已存在",
	ErrRemoteSocketMisType  : "链接类型错误",

	ErrUnKnownProtocal : "未知的协议类型",
	ErrMismatchProtocalReq   : "请求协议数据类型不匹配",
	ErrMismatchProtocalConnectReq : "连接请求不合法",
	ErrMismatchProtocalResp   : "返回协议数据类型不匹配",

	ErrUnknuowPeer : "未知peer",
}