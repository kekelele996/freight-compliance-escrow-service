# Bug reproduction

## Bug
租户作用域在 leaseStorageKey 前被折叠为 shared，ReserveLease 又把客户端 key 直接当作 LeaseID；CommitLease 按 ClientKey 扫描并忽略租户、LeaseID 与 reserved 状态，AbortLease 则直接按裸 key 删除。四处对同一租约身份和状态转换采用了不同规则，造成跨租户冲突、外部租约可提交以及取消删除错记录。

## Trigger
运行回归场景 `TestRequestLeaseLifecycleIsTenantScopedAndAtomic`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`bad lease identity: {Scope: ClientKey:same Fingerprint:fa LeaseID:same Response: State:reserved}`
