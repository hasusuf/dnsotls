# dnsotls
[![license](https://img.shields.io/github/license/hasusuf/dnsotls.svg?maxAge=2592000)](https://github.com/hasusuf/dnsotls/blob/master/LICENSE)

Very basic TCP DNS-over-TLS proxy for Cloudflare

## Available options
- `--bind`: Binding IP address, default: 127.0.0.1
- `--port`: Binding Port, default: 53
- `--debug`: Enable debug mode, default: false
- `--help`

## Usage
```
$ go get github.com/hasusuf/dnsotls
$ dnsotls --debug
```

## Docker
```
$ docker build --rm -t hasusuf/dnsotls:latest ./docker
$ docker run -d --rm --name dnsotls hasusuf/dnsotls
$ docker exec -it dnsotls bash
$ dig @127.0.0.1 -p 53 cloudflare.com +short
$ nslookup cloudflare.com 127.0.0.1
```