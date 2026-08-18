# Bug reproduction

## Bug
routingContext 用 Background 替换调用方 context，worker 返回后不复查取消；RoutingResultCache.put 在加锁前后都忽略取消，RunRouting 也无入口检查和 commit 前回滚。取消信号因此在请求、worker、缓存和最终提交四个边界连续丢失，晚到结果被缓存并计为已提交。

## Trigger
运行回归场景 `TestRoutingCancellationBoundaries`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`request cancellation was not preserved`
