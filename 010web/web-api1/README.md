
良い記事。

[Go言語で基本的なCRUD操作を行うREST APIを作成](https://dev.classmethod.jp/articles/go-sample-rest-api/)
[gitソース](https://github.com/koga456/sample-api)

```
# 実行
docker-compose -f build/docker-compose.yml up -d
# ビルドして実行
docker-compose -f build/docker-compose.yml up --build -d

# docker確認
cd build
docker-compose ps
docker-compose down

docker-compose exec api-server bash


#TODO一覧取得
curl -i localhost:8080/todos/
#登録
curl -i -X POST -H "Content-Type: application/json" -d '{"title":"test", "content":"テストです。"}' localhost:8080/todos/
#更新
curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"test", "content":"変更テスト"}' localhost:8080/todos/4
#削除
curl -i -X DELETE localhost:8080/todos/4
```

