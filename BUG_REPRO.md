# Bug reproduction

## Bug
decodeScopedCursor 解码后清空租户和过滤条件，authorizeScopedCursor 无条件放行；checkpoint staging 与授权/提交边界脱节，service 对 staged 结果不做显式 commit，rollback 也为空。非法或外租户 cursor 可被接受并影响会话，而合法 cursor 的版本和位置又不能可靠提交。

## Trigger
运行回归场景 `TestPageCursorCommitRequiresParseScopeAndTransaction`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`foreign cursor accepted`
