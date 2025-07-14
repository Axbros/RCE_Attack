package types

type CommandRequest struct {
	Cmd  string   `json:"cmd"`  // 要执行的命令
	Args []string `json:"args"` // 命令参数
}

type CommandResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}
