# sync-image

sync-image是一个轻量级工具，通过使用代理能力，比如clash。将DockerHub镜像同步到私有harbor仓库，简化了镜像管理流程，自动从源拉取镜像并同步到目标仓库。未来版本将支持符合 Docker repository v2 协议的其他仓库。

## 安装

```bash
git clone https://github.com/Shuimo03/sync-image
cd sync-image
make build
```

## 使用方法

### 配置文件

配置文件分为两种: auth.yaml和image.yaml，auth是关于harbor或者说私有仓库授权，image是需要同步的镜像。

auth.yaml

```yaml
auths:
  registry: ""
  username: ""
  password: ""
```

image.yaml

```yaml
source:
  registries:
    - name: monitoring #保证和target中repositories一致
      images:
        - registry.k8s.io/kube-state-metrics/kube-state-metrics:v2.13.0

target:
  registry: 192.168.0.101:8080 #目标镜像地址
  repositories:
    - name: monitoring #需要确保存在
```

以上配置好之后，按照以下方法运行:

```bash
sync-image --config=cnf/image.yaml --auth=cnf/auth.yaml
```

## 许可证

本项目使用 MIT 许可证 - 详见 LICENSE 文件。