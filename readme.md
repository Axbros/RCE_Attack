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
🔐 Security Considerations (Built-in)
Execution timeout: 3 seconds to prevent long-running commands

Logs all requests: IP, timestamp, command, output, error

Recommended additions:

- IP rate limiting middleware

- Token-based authentication

- Command whitelist filtering

- Execution sandbox (e.g. Docker)

🛠 Running Locally
1. Build the Project
```shell
go build -o executor main.go
```
2. Start the Server
```shell
./executor
# Listens on localhost:8080 by default
```
3. Test with Postman or curl
```shell
curl -X POST http://localhost:8080/run \
  -H "Content-Type: application/json" \
  -d '{"cmd": "date"}'

```
**📋 License**

This project is released under the MIT License.

⚠️ Use at your own risk. This service executes system-level commands. DO NOT deploy it to production or the public internet without strict access controls.
