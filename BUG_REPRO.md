# Bug reproduction

## Bug
approvalCandidate 原地修改候选记录，repository.save 在报告冲突前先覆盖持久化记录；service 又在保存成功前推进 checkpoint 并准备事件，而 rollback 不恢复这些副作用。乐观并发冲突因此跨 repository、checkpoint 和 event 三层留下部分提交。

## Trigger
运行回归场景 `TestApprovalConflictLeavesNoPartialState`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`repository mutated on conflict: {ID:s1 State:approved Version:2}`
