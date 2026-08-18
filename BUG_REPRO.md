# Bug reproduction

## Bug
运营日 key 按 UTC 日期生成且不包含站点时区，claim 在工作开始前直接标 done，也不拒绝 running/done；finish 在失败时保留错误的完成状态，service 仍继续执行 work。DST 切换日的同一本地运营日可能被错误分组，失败重试和重复执行状态也无法正确转换。

## Trigger
运行回归场景 `TestOperationalDayRetriesWithoutDuplicateAcrossDST`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`key=2026-03-08`
