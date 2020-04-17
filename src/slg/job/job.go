package Job

import (
	"slg/entity"
)

//事件参数类型 以首次注册时为准
var _jobDoneCallBack = make(map[int32]func(job *Entity.Job))

func RegJobDone(jobType int32, call func(job *Entity.Job)) {
	_jobDoneCallBack[jobType] = call
}
