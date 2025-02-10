# TTAuto
甜糖心愿自动签到

功能：每天 3:00 自动签到

## 使用

### 正常

1. 从 [Releases](https://github.com/ygxbnet/ttauto/releases) 页面中按照电脑系统及架构下载最新的压缩包到本地
2. 解压，运行可执行文件
3. 按照提示分别输入手机号，4位数图形验证码，短信验证码
4. 稳定运行

### 推荐

使用 **Windows** 系统电脑进行配置，用带 **Docker** 的 **Linux** 服务器 7*24 运行签到

> 别问为什么推荐这种，因为我比较菜，只会这样用🤣

1. 从 [Releases](https://github.com/ygxbnet/ttauto/releases) 页面中下载最新的 `ttauto_Windows_x86_64.zip` 到本地 Windows 电脑
2. 解压，并运行 `ttauto.exe`
3. 初次使用请按照提示分别输入手机号，4位数图形验证码，短信验证码
4. 登陆成功后出现以下提示信息，这时就已经可以使用了（后面还有第5步）

```
登陆成功！
===================================
用户名：xxx
手机号：xxxxxxx
等级：xx
union_id：xxx
token：xxx
===================================
已将 union_id 保存至以下文件中
C:\xxx\xxx\xxx\xxx/union_id
登陆成功，开始执行后续程序...

已有 union_id，正在验证是否可用...
union_id 有效，token 刷新成功！
【定时任务】已开启定时任务，每天 3:00 定时签到
```

不过如果你跟作者一样，电脑不会 7*24 小时开机，或者不想看见这个黑不溜秋的框框在桌面上，那就请接着往下看

提示：以下操作默认你已经有一个带 **Docker** 的 **Linux** 服务器

5. 选择一个你喜欢的目录（例如 `/data/ttauto/`）在当前目录下创建 `docker-compose.yaml` 并写入以下内容

```yaml
services:
  ttauto:
    image: ygxb/ttauto
    container_name: ttauto
    restart: always
    network_mode: host
    volumes:
      - ./union_id:/data/union_id
```

6. 将第4步提示信息中给到的 `C:\xxx\xxx\xxx\xxx/union_id` 文件复制到服务器目录下（例如 `/data/ttauto/`）
7. 现在如果按照我给的例如来做，你服务器的 `/data/ttauto/` 目录下应该有 `docker-compose.yaml` `union_id` 两个文件
8. 运行 `docker compose up -d`，等待运行完成
9. 然后运行 `docker compose logs -f` 查看日志，如果出现以下信息，那恭喜你，你已经拥有了一个甜糖自动打卡机器人了，达到60级也指日可待！

```
欢迎使用 TTAuto

已有 union_id，正在验证是否可用...
union_id 有效，token 刷新成功！
【定时任务】已开启定时任务，每天 3:00 定时签到
```

## 要饭🙏

（哭）各位看官求求你们了，孩子已近几天几夜没吃饭了，如果觉得工具还不错，能不能在甜糖中填一下邀请码：`774172` 

（跪下）谢谢了！

## 特别感谢

[boris1993/tiantang-auto-harvest](https://github.com/boris1993/tiantang-auto-harvest)（提供甜糖 API 及调用）
