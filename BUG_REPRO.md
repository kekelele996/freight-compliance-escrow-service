# Bug reproduction

## Bug
claim 直接把 pending 记录标成 published，而不是 inflight；send 出错时 PublishOne 提前返回并跳过 finish，finish 自身又无论错误都写 published，invokeDeliverySend 对 nil sender 和非法状态也无保护。发送、错误传播、最终状态和重试可见性顺序互相矛盾，失败事件因此变成不可重试。

## Trigger
运行回归场景 `TestDeliveryFailureIsFinalizedAsRetryable`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`claim={ID:e1 State:published Attempts:1 LastError:}`
