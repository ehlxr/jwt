# JWT
> This is a simple tool to sign, verify and show JSON Web Tokens ([JWT](http://jwt.io/)) from the command line, base [jwt-go](https://github.com/dgrijalva/jwt-go).

![image](https://wx4.sinaimg.cn/large/687148dbly1fprjoxvd2gg213g0od7eq.gif)

[more...](https://wx4.sinaimg.cn/large/687148dbly1fprk7dnqvsg213g0odx6s.gif)

# Install

build with go get

```
➜ go get -u github.com/ehlxr/jwt
```

build with go [dep](https://github.com/golang/dep)

```
➜ git clone https://github.com/ehlxr/jwt.git

➜ cd jwt && dep ensure
```

or download [releases](https://github.com/ehlxr/jwt/releases) binary package.

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

## sign JWT

```
➜ jwt sign -h

签名 JWT token 并复制到剪切板
标记 * 号的 flag 为必须项

Usage:
  jwt sign [flags]

Flags:
  -c, --claims argList   add additional claims. may be used more than once (default {})
  -d, --data string      * path or json to claims object to sign, '-' to read from clipboard, or '+' to use only -claim args
  -H, --header argList   add additional header params. may be used more than once (default {})
  -h, --help             help for sign
  -k, --key string       * path of keyfile or key argument

Global Flags:
      --config string   config file (default is $HOME/.jwt.yaml)
```

## show JWT

```
➜ jwt show -h

查看 JWT Token 内容
标记 * 号的 flag 为必须项

Usage:
  jwt show [flags]

Flags:
  -h, --help           help for show
  -t, --token string   * path or arg of JWT token to verify, '-' to read from clipboard

Global Flags:
      --config string   config file (default is $HOME/.jwt.yaml)
```

## verify JWT

```
➜ jwt verify -h

验证 JWT token 是否有效
标记 * 号的 flag 为必须项

Usage:
  jwt verify [flags]

Flags:
  -h, --help           help for verify
  -k, --key string     * path of keyfile or key argument
  -t, --token string   * path or arg of JWT token to verify, '-' to read from clipboard

Global Flags:
      --config string   config file (default is $HOME/.jwt.yaml)
```