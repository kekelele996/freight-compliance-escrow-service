# Bug reproduction

## Bug
跨午夜窗口仍用查询时刻的 weekday 判归属，ContainsAt 只支持开门分钟小于关门分钟；候选开放时间固定在 UTC 构造，NextOpenAt 还会返回当天已经过去的候选。窗口归属日、区间状态和本地未来候选三套规则不一致。

## Trigger
运行回归场景 `TestScheduleStateUsesOwnerDayAndFutureLocalCandidate`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`contains=false/false/false`
