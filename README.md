# HyacineGo


## 功能
- 登录
- 进入游戏
- 其它功能开发中

## 初始化
下载 Go语言环境（https://go.dev/dl/go1.25.5.windows-amd64.msi）
在 `Base/HyacineGo` 目录运行：

```powershell
go run .\cmd\hyacine-server -config .\configs\config.json
```

首次启动会使用/创建：
- `configs/config.json`：主配置
- `configs/accounts.json`：账号数据
- `configs/players/`：玩家存档目录（`<uid>.json`）

## 配置

配置文件：`configs/config.json`

常用字段：
- `dispatch.addr`：Dispatch 监听地址（HTTP）
- `dispatch.account_path`：账号文件路径（默认 `./configs/accounts.json`）
- `gameserver.addr`：GameServer 监听地址（TCP）
- `gameserver.resource_root`：资源根目录（默认 `./resources`）
- `logging.level` / `logging.format` / `logging.file`：日志输出
- `debug.trace_file`：包 trace（可选）
