# privatemode-public-demo

## Testing

```pwsh
Invoke-RestMethod -Uri "http://localhost:8080/v1/chat/completions" `
  -Method Post `
  -ContentType "application/json" `
  -Body '{
    "model": "ibnzterrell/Meta-Llama-3.3-70B-Instruct-AWQ-INT4",
    "messages": [
      {
        "role": "system",
        "content": "Hello Privatemode!"
      }
    ]
  }' | Tee-Object -Variable respons
```  

For wsl user: `ip route show | grep -i default | awk '{ print $3}'`

```bash
curl localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "ibnzterrell/Meta-Llama-3.3-70B-Instruct-AWQ-INT4",
    "messages": [
      {
        "role": "system",
        "content": "Hello Privatemode!"
      }
    ]
  }' | jq
```
