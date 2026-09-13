# 承诺行动状态服务

该服务把承诺、证据和核查动作串成带版本的后端状态链，适配 PostgreSQL 与 NATS。`fixtures/` 记录参与方和阶段规则样例，敏感字段只在受权接口返回。

```bash
docker compose up --build
```
