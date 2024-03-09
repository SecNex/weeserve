# WeeServe

WeeServe is a development server. It is designed to be used in development environments and is not suitable for production use.

## Installation

```bash
curl -sSL https://serverify.de/scripts/weeserve-install.sh | bash
```

## Usage

```bash
# Start simple server
weeserve -host 127.0.0.1 -port 8080

# Start simple server with reverse proxy
weeserve -host 127.0.0.1 -port 8080 -reverse
```