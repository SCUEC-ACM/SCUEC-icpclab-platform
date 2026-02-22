package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response 结构体用于返回 JSON
type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	
	// 准备返回数据
	res := Response{
		Message: "来自 Go 后端的连接成功！",
	}

	// 编码并发送
	json.NewEncoder(w).Encode(res)
	fmt.Println("接收到一次前端请求")
}

func main() {
	// 路由映射
	// 注意：这里的路径要和前端 fetch 的路径一致
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("后端服务启动在 :3000 端口...")
	// 启动服务器
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("启动失败: %s\n", err)
	}
}