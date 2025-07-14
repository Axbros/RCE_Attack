package handler

import (
	"awesomeProject/types"
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func Run(w http.ResponseWriter, r *http.Request) {
	var cmdReq types.CommandRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &cmdReq)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 防止 top, ping, yes 等长时间阻塞命令。
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, cmdReq.Cmd, cmdReq.Args...)
	out, err := cmd.CombinedOutput()

	// 构造返回
	resp := types.CommandResponse{Output: string(out)}
	if err != nil {
		resp.Error = err.Error()
	}

	// 记录日志
	log.Printf(
		"[COMMAND LOG] Time: %s | IP: %s | CMD: %s %v | ERROR: %v | OUTPUT: %s",
		time.Now().Format(time.RFC3339),
		getIP(r),
		cmdReq.Cmd,
		cmdReq.Args,
		err,
		sanitizeOutput(string(out)),
	)

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getIP(r *http.Request) string {
	// 支持反向代理场景
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func sanitizeOutput(out string) string {
	const maxLen = 500
	// 可选：移除控制字符等敏感内容
	clean := strings.ReplaceAll(out, "\n", "\\n")
	if len(clean) > maxLen {
		return clean[:maxLen] + "...[truncated]"
	}
	return clean
}
