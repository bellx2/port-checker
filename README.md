# Multi Port Checker for Test

Open multiple port checker.

## Parameters

| Option | Description       | Default   |
| ------ | ----------------- | --------- |
| -h     | host              | localhost |
| -p     | Start Port Number | 58000     |
| -n     | Count             | 1         |
| -json  | Output JSON       | false     |

## Build

```bash
go build
GOOS=windows GOARCH=amd64 go build # for windows
```

## Usage

```bash
➜ portcheker -h www.yahoo.co.jp -p 80 -n 10
Host: www.yahoo.co.jp
80 : OPEN
81 : CLOSE -  dial tcp 183.79.219.252:81: i/o timeout
82 : CLOSE -  dial tcp 183.79.219.252:82: i/o timeout
83 : CLOSE -  dial tcp 183.79.219.252:83: i/o timeout
84 : CLOSE -  dial tcp 183.79.219.252:84: i/o timeout
...

➜ portcheker -h www.yahoo.co.jp -p 80 -n 10 --json | jq
{
  "host": "www.yahoo.co.jp",
  "ports": [
    {
      "port": "80",
      "status": "OPEN",
      "open": true
    },
    {
      "port": "81",
      "status": "CLOSED : dial tcp 182.22.25.252:81: i/o timeout",
      "open": false
    },
    ...
  ]
}

```
