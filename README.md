```
  __  __            _      _       
 |  \/  | ___   ___| | __ (_) ___  
 | |\/| |/ _ \ / __| |/ / | |/ _ \ 
 | |  | | (_) | (__|   < _| | (_) |
 |_|  |_|\___/ \___|_|\_(_)_|\___/ 
                                   
```

A very simple mock server for quick API emulation.

### Features

- simple json format
- response latency
- reverse proxy mode
- requests and responses logs

## How to use

### Create a mock file

The json format is as follows:

- **key** - route
- **value** - response

The **key** is a colon-separated string and consists of three parts:

| # | Part   | Required | Default |
|---|--------|:--------:|:-------:|
| 1 | Method |          |   GET   |
| 2 | Code   |          |   200   |
| 3 | Path   |     *    |         |

Some examples of key mapping:

| Key              | Method | Path      | Code |
|------------------|:------:|-----------|:----:|
| /orders          |   GET  | /orders   |  200 |
| 404:/orders/1    |   GET  | /orders/1 |  404 |
| GET:/orders/2    |   GET  | /orders/2 |  200 |
| POST:201:/orders |  POST  | /orders   |  201 |

#### Example

```json
{
  "/v1/orders/1": {
    "id": 1,
    "status": "confirmed"
  },
  "POST:201:/v1/orders": {
    "success": true
  }
}
```

### Go

```bash
mockio -i routes.json
```

```bash
Options:
  -c    color output (default true)
  -i string
        path to mock file
  -l int
        response latency (ms)
  -p int
        server port (default 8080)
  -v    verbose output
  -vv
        very verbose output (print response headers)
```

### Testing

```bash
curl -i localhost:8080/v1/orders/1
```

```bash
curl -i -X POST localhost:8080/v1/orders
```

## Reverse proxy mode

Redirect all non-defined routes to the specified host.

```json
{
  "*": "http://localhost:8080"
}
```

Override a specific route:

```json
{
  "*": "http://localhost:8080",
  "/v1/orders/1": {
    "id": 1,
    "status": "confirmed"
  }
}
```