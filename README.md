# go/gin

go + gin の小さなサンプル

# 実行

```zsh
go run .
```

# curl

get

```zsh
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

post

```zsh
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

