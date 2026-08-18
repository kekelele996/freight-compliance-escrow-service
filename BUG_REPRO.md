# Bug reproduction

## Bug
decode、normalize、merge、persist 都按值复制 ManifestDraft 结构体，却继续共享 Stops、嵌套 Seals 切片和 Labels map；normalize 的 append/写 map、merge 的追加和 overlay 合并、persist 的返回值都没有建立阶段所有权。外部输入或任一中间快照被修改时，其他阶段快照随共享底层数据一起变化。

## Trigger
运行回归场景 `TestManifestPipelineKeepsStageOwnership`，构造题面描述的业务状态并观察跨层状态变化。

## Actual error
`decoded={Stops:[{Code:A Seals:[input-change normalized]}] Labels:map[carrier:x normalized:yes source:input-change]}`
