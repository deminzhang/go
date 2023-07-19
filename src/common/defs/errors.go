package defs

const (
	ErrOK                  = 0
	ErrBadRequest          = 400 // 客户端请求的语法错误，服务器无法理解
	ErrUnauthorized        = 401 // 请求要求用户的身份认证
	ErrNotAllowed          = 403 //	服务器理解请求客户端的请求，但是拒绝执行此请求
	ErrNotFound            = 404 //	服务器无法根据客户端的请求找到资源（网页）
	ErrConflict            = 409 // 服务器端目标资源的当前状态相冲突
	ErrUnsupportedMedia    = 415 // 服务器无法处理请求附带的媒体格式
	ErrInvalidData         = 422 // 请求的语义错误
	ErrLocked              = 423 // 当前资源被锁定
	ErrTooManyRequests     = 429 // 服务器请求过载
	ErrInternal            = 500 //	服务器内部错误，无法完成请求
	ErrNotImplemented      = 501 // 服务器不支持请求的功能，无法完成请求
	ErrNetworkAuthRequired = 511 // 需验证以许可连接

	ErrMailSendFail       = 601 // 邮件发送失败
	ErrMailQueueFull      = 602 // 邮箱已满
	ErrMailPrefabNotExist = 603 // 错误邮件预制不存在

	ErrDbsDbWrong        = 701 // db错误
	ErrDbsHostDbNotFound = 702 // db host 错误
	ErrDbsIdGenFail      = 703 // db id gen错误
	ErrDbsStrEscape      = 704 // db 字符串逃逸错误
	ErrDbsRows           = 705 // db返回数据记录条目数错误

	ErrImSessionWrong   = 801 // 会话错误
	ErrImPersonNotFound = 802 // 找不到人
	ErrImMsgSendFail    = 803 // 消息发送失败
	ErrImPersonBanned   = 804 // 被禁言

	ErrGptReachLimited = 851 // 达到请求限制了

	ErrTODO          = 1000 //开发中未开放
	ErrModuleNotOpen = 1002 //功能未开启

)
