# Bug reproduction

## Bug
LifecyclePool.Run 的 admission 无条件放行，lifecycleJobContext 用 Background 切断父取消，job 返回后不把 ctx.Err 写回结果，commitLifecycle 也忽略取消且继续处理后续队列。运行中任务、排队任务和最终提交分别绕过了同一取消生命周期。

## Trigger
运行回归场景 `TestWorkerCancellationStopsCommitAndLaterAdmission`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`direct canceled commit was admitted`
