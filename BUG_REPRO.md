# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

批量登记十五个站点时，中间只要有一个坐标非法，接口虽然失败，之前写入的站点仍留在库里，连接也长时间处于事务状态。请修复失败后的清理，让整批写入完整回滚并及时释放资源。事务回滚测试文件必须保持不变，不得跳过非法坐标场景。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-16
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-16.git
- parent SHA：982b3681d658adc885b1b86858a71fc5d2293588

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-16.git bug-repro
cd bug-repro
git checkout --detach 982b3681d658adc885b1b86858a71fc5d2293588
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.01s)
    task_behavior_test.go:14: status=400 body={"active":0,"stored":2}
FAIL
FAIL	chargeguard/internal/httpapi	0.067s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:14: status=400 body={"active":0,"stored":2}
FAIL
FAIL	chargeguard/internal/httpapi	0.004s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，在题面描述的触发条件下应得到预期业务结果且不再出现原始症状；定向验证命令修复前必须失败、应用修复后必须通过，相关回归和仓库全量测试必须通过；不得新增、删除或修改测试文件，不得跳过测试、降低断言或绕过目标逻辑。
