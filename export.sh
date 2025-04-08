#! /bin/bash

# 需要执行：source export.sh，而不是 ./export.sh，直接执行是子进程中执行，环境变量不会生效！！！

# 获取当前脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 设置环境变量
export DUBBO_GO_ROOT_PATH="$SCRIPT_DIR"
export DUBBO_GO_CONFIG_PATH="$SCRIPT_DIR/config/dubbogo.yaml"

echo "DUBBO_GO_ROOT_PATH: $DUBBO_GO_ROOT_PATH"
echo "DUBBO_GO_CONFIG_PATH: $DUBBO_GO_CONFIG_PATH"
