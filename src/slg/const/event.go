package Const

//事件
const (
	OnTest        = iota //空置
	OnDBConnect          //DB已连接 可初始化/增量更新数据
	OnLoadConfig         //配置加载完成 可重组关联项/预处理权重计算项
	OnServerStart        //服务器已经开启
	OnUserNew            //角色创建初始化
	OnUserLogin          //角色登陆 可处理离线结算 新功能数据初始化 异常数据修复
	OnUserOffline        //角色离线 可处理离线结算 社交关系通知 私有缓存释放
	OnUserGetData        //角色收集数据 给前端或跨服集
	OnUserLevelUp        //角色升级 更新任务/解锁/成就
)
