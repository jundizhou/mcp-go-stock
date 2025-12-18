#!/bin/bash
# jd-go-stock 一键编译部署脚本
# 用法: ./deploy.sh

set -e

# 配置
SERVER_HOST="root@122.51.133.114"
SERVER_PATH="/opt/jd-go-stock"
LOCAL_PATH="$(cd "$(dirname "$0")" && pwd)"

echo "=========================================="
echo "jd-go-stock 一键编译部署脚本"
echo "=========================================="
echo "本地路径: $LOCAL_PATH"
echo "服务器: $SERVER_HOST:$SERVER_PATH"
echo ""

# 1. 编译 Linux 版本
echo "[1/5] 编译 Linux 版本..."
cd "$LOCAL_PATH"
GOOS=linux GOARCH=amd64 go build -o jd-go-stock-linux .
GOOS=linux GOARCH=amd64 go build -o init_db-linux ./cmd/init_db/main.go
echo "✅ 编译完成"

# 2. 停止远程服务
echo ""
echo "[2/5] 停止远程服务..."
ssh $SERVER_HOST "pkill -f 'jd-go-stock-linux' 2>/dev/null || true"
echo "✅ 远程服务已停止"

# 3. 上传文件
echo ""
echo "[3/5] 上传文件到服务器..."
ssh $SERVER_HOST "mkdir -p $SERVER_PATH"
scp jd-go-stock-linux init_db-linux backend/data/stock_basic.json $SERVER_HOST:$SERVER_PATH/
ssh $SERVER_HOST "chmod +x $SERVER_PATH/jd-go-stock-linux $SERVER_PATH/init_db-linux"
echo "✅ 文件上传完成"

# 4. 初始化数据库（如果需要）
echo ""
echo "[4/5] 检查数据库..."
DB_EXISTS=$(ssh $SERVER_HOST "[ -f $SERVER_PATH/data/stock.db ] && echo 'yes' || echo 'no'")
if [ "$DB_EXISTS" = "no" ]; then
    echo "数据库不存在，正在初始化..."
    ssh $SERVER_HOST "cd $SERVER_PATH && ./init_db-linux"
    echo "✅ 数据库初始化完成"
else
    echo "✅ 数据库已存在，跳过初始化"
fi

# 5. 启动服务
echo ""
echo "[5/5] 启动 MCP 服务器..."
ssh $SERVER_HOST "cd $SERVER_PATH && nohup ./jd-go-stock-linux --mode=http --port=3000 > server.log 2>&1 &"
sleep 2

# 验证服务
echo ""
echo "验证服务状态..."
RESULT=$(ssh $SERVER_HOST "ps aux | grep 'jd-go-stock-linux' | grep -v grep | wc -l")
if [ "$RESULT" -gt 0 ]; then
    echo "✅ 服务已启动"
    ssh $SERVER_HOST "ps aux | grep 'jd-go-stock-linux' | grep -v grep"
else
    echo "❌ 服务启动失败，请检查日志"
    ssh $SERVER_HOST "tail -20 $SERVER_PATH/server.log"
    exit 1
fi

echo ""
echo "=========================================="
echo "✅ 部署完成!"
echo "=========================================="
echo ""
echo "服务地址: http://122.51.133.114:3000/mcp"
echo ""
echo "MCP 客户端配置:"
echo '{'
echo '  "mcpServers": {'
echo '    "jd-go-stock": {'
echo '      "type": "streamableHttp",'
echo '      "url": "http://122.51.133.114:3000/mcp"'
echo '    }'
echo '  }'
echo '}'
echo ""
echo "查看日志: ssh $SERVER_HOST \"tail -f $SERVER_PATH/server.log\""
echo "停止服务: ssh $SERVER_HOST \"pkill -f jd-go-stock-linux\""
