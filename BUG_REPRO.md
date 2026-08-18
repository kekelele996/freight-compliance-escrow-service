# Bug reproduction

## Bug
报价缓存键只使用报价 ID，失败 candidate 在校验完成前就以原指针写入缓存，rollback 为空操作，load 也直接返回缓存指针。一次金额校验错误会留下可跨租户命中的失败对象，后续调用和调用方对返回对象的修改都会继续污染缓存。

## Trigger
运行回归场景 `TestFailedQuoteDoesNotPoisonRetryOrOtherTenant`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`failed candidate cached: 1`
