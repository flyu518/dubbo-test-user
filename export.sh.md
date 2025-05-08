#! /bin/bash

# 需要执行：source export.sh，而不是 ./export.sh，直接执行是子进程中执行，环境变量不会生效！！！

LM_CONFIG_CENTER_CONFIG="{\"address\":\"118.25.146.1:31238\",\"username\":\"\",\"password\":\"\",\"namespace\":\"cui_test\",\"group\":\"lingmou\",\"data_id\":\"user-server.yaml\",\"service_data_id\":\"user-server-config.yaml\"}"

# 设置环境变量
export LM_CONFIG_CENTER_CONFIG="$LM_CONFIG_CENTER_CONFIG"

echo "LM_CONFIG_CENTER_CONFIG: $LM_CONFIG_CENTER_CONFIG"
