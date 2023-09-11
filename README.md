# Gin

go/ginサンプルコード

## Stack
- go v1.21.1
- gin

## 実行

```zsh
go run .
```

全てのアルバムリスト取得


```zsh
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

アルバムを追加

```zsh
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "only my railgun","artist": "FripSide","price": 30.2, "tax": 0.1}'
```

アルバムIDからアルバムの詳細を取得

```zsh
curl http://localhost:8080/albums/1 \
    --header "Content-Type: application/json" \
    --request "GET"
```

