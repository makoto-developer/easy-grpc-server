# Learning gRPC

gRPCの簡単なサーバーとクライアントの実装サンプル

## Requirement

- asdf

## Setup

asdfにプラグインが入っていない場合は`plugin add`する

```shell
asdf plugin add golang
asdf plugin add protoc-gen-go-grpc
asdf plugin add protoc-gen-go
asdf plugin add protoc
```

Golangとgrpcの実装に必要なライブラリをインストールする

```shell
asdf install
```

## Start gRPC Server and Client

init

```shell
go mod tidy
```

starting server

```shell
make server
```

```shell
make client
```

## protobufの定義を更新してgoファイルを生成する

```shell
make generate
```

# References

- https://qiita.com/Ian_C/items/523912ad9b0169cacb3c
