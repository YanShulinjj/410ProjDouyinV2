package xerr

type ErrCodeType uint32

const (
	// 成功返回
	OK ErrCodeType = 200
	/**(前3位代表业务,后三位代表具体功能)**/
	/**全局错误码*/
	// 服务器开小差
	ServerCommonErr ErrCodeType = 100000 + iota
	// 请求参数错误
	ReuqestParamErr
	// 生成token失败
	TokenGenerateErr
	// Token验证失败
	TokenNotMatchErr
	// Token过期
	TokenExpiredErr
	// Token还未激活
	TokenNotActiveErr
	// 数据库繁忙,请稍后再试
	DbErr
	// 更新数据影响行数为0
	DbUpdateAffectedZeroErr
	// 数据不存在
	DataNoExistErr

	// 用户服务
	UserNotExistErr ErrCodeType = 200000 + iota
	UserExistedErr
	UserPwdNotMatchErr

	// 视频服务
	FileUploadErr ErrCodeType = 300000 + iota

	// 商品服务

	// 支付服务

)
