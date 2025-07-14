# Go Remote Command Executor API

> ⚠️ **Warning: This project is intended for internal and controlled environments only. Do NOT expose it to the public internet without proper authentication and security controls.**

This is a lightweight HTTP service written in Go that allows clients to send a command and arguments via a POST request. The server executes the command locally and returns both the output and error (if any). The service includes a timeout mechanism and logs key request data.

---

## ✨ Features

- Accepts JSON-based command execution requests
- Executes shell commands with optional arguments
- Command execution timeout (default: 3 seconds)
- Logs client IP, timestamp, command, output, and error
- JSON request/response format

---

## 🚀 API Endpoint

### POST `/run`

Executes a system command on the server and returns the output.

#### ✅ Request Payload (Example)

```json
{
  "cmd": "ls",
  "args": ["-la", "/"]
}
```
🔁 Successful Response Example
```json
{
  "output": "total 60\ndrwxr-xr-x   1 root root 4096 Jul 9 10:00 bin\n...",
  "error": ""
}
```

🔁 Error Response Example
```json
{
  "output": "",
  "error": "exec: \"notarealcmd\": executable file not found in $PATH"
}
```