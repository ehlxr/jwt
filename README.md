# JWT
> This is a simple tool to sign, verify and show JSON Web Tokens ([JWT](http://jwt.io/)) from the command line, base [jwt-go](https://github.com/dgrijalva/jwt-go).

[![jwt](https://asciinema.org/a/P0O3XBCslMNam0UduazwPhB6o.png)](https://asciinema.org/a/P0O3XBCslMNam0UduazwPhB6o?autoplay=1)

# Install

build with go get

```
➜ go get -u github.com/ehlxr/jwt
```

build with make

```
➜ git clone https://github.com/ehlxr/jwt.git

➜ cd jwt && make install
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
  show        JWT Token 查看
  sign        JWT 签名
  verify      JWT token 验证
  version     Print version

Flags:
      --config string   config file (default is $HOME/.jwt.yaml)
  -h, --help            help for jwt

Use "jwt [command] --help" for more information about a command.

```

## JWT version

```
➜ jwt version

   __     __     __     ______
  /\ \   /\ \  _ \ \   /\__  _\
 _\_\ \  \ \ \/ ".\ \  \/_/\ \/
/\_____\  \ \__/".~\_\    \ \_\
\/_____/   \/_/   \/_/     \/_/



Name: jwt
Version: v1.0.2
BuildTime: 2019-10-02 16:56:20
GitCommit: c546aaaee1b6a6b03eabf396f9cab01718e22104
GoVersion: go version go1.13.1 darwin/amd64
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

## Thanks to the following organizations for providing open source licenses

[<img src="https://cdn.ehlxr.top/jetbrains.png" width = "200" height = "217" alt="图片名称" align=center />](https://www.jetbrains.com/?from=jwt)


