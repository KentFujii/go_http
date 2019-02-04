https://github.com/oreilly-japan/real-world-http

## 環境構築

```
docker-compose up --build
docker-compose exec 試したいプロトコル bash
./bin/試したい通信処理
```

## gist

### 1991年 – HTTP/0.9

GETメソッドのみ。ヘッダもレスポンスコードの規定も存在しない簡素な仕様

### 1996年 – HTTP/1.0

RFC1945として公開。ステータスコードを含むレスポンスヘッダが付加されるようになる。GET以外にPOSTメソッド等の新たなメソッドも追加される

### 1999年 – HTTP/1.1

RFC2068として公開。Keep-Aliveやパイプライン化をサポート。1.0から大幅に機能が追加される

### 2015年 – HTTP/2

RFC7540として公開。HTTP/1.1との互換性を保ちつつもWEBを効率化するための様々な機能をサポート
