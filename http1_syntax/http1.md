# HTTP/1.0のシンタックス

HTTPを構成する要素である

- メソッドとパス
- ヘッダー
- ボディ
- ステータスコード

にフォーカスして解説する

```
docker-compose run app go run http1_syntax/server.go
```

## HTTP/0.9でできることを試す

リクエストを投げてレスポンスを得る

```
curl --http1.0 http://localhost:18888/greeting
```

リクエストにクエリを付与する。

```
curl --http1.0 --get --data-urlencode "search word" http://localhost:18888
```

下のコマンドは`http://localhost:18888/?search+word`というリクエストを投げているのと同じ

## HTTP/0.9から1.0への道のり

HTTP/0.9では

- 一つのドキュメントを送る機能しかなかった
- 通信される全ての内容はHTML文書であるという想定だったため、ダウンロードするコンテンツのフォーマットをサーバーから伝える手段がなかった
- クライアント側から検索のリクエストを送る以外のリクエストを送信できなかった
- 新しい文章を送信したり、更新したり、削除することはできなかった
- リクエストが正しかったか、もしくはサーバーが正しく応答することができたかといった情報をしる方法がなかった

HTTP/1.0のリクエストを試す

```
curl -v --http1.0 http://localhost:18888/greeting
```

HTTP/1.0では

- リクエスト時にメソッドが追加された(GET)
- リクエスト時にHTTPバージョンが追加された(HTTP/1.0)
- ヘッダーが追加された(Host, User-Agent, Accept)

といった変更がなされている

## HTTPの先祖(電子メール)


## HTTPの先祖(ニュースグループ)


## リダイレクト


## URL(Uniform Resource Locators)


## ボディ

