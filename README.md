# JWT
> This is a simple tool to sign, verify and show JSON Web Tokens ([JWT](http://jwt.io/)) from the command line, base [jwt-go](https://github.com/dgrijalva/jwt-go).

# Install

```
➜ go get -u github.com/ehlxr/jwt
```

# Usage

```
➜ jwt
JWT(Json Web Token) 工具
用于生成、验证、查看 JWT

Usage:
  jwt [command]

Available Commands:
  help        Help about any command
  show        查看 JWT Token
  sign        JWT 签名
  verify      JWT token 验证

Flags:
      --config string   config file (default is $HOME/.jwt.yaml)
  -h, --help            help for jwt
  -v, --version         show version of the jwt.

Use "jwt [command] --help" for more information about a command.
```