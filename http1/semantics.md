# HTTP/1.0のセマンティクス

HTTPを構成する要素である

- メソッドとパス
- ヘッダー
- ボディ
- ステータスコード

これらをブラウザがどう解釈していくかを解説する

```
docker-compose run app go run http1/server.go
```

## シンプルなフォームの送信

```
curl --http1.0 -d title="The Art of Community" -d author="Jono Bacon" http://localhost:18888
```

curlコマンドの-dオプションを使ってフォームで送信するデータPOSTを設定できる

上の送信はencodingを無視しているのでスペースや`&`といった特殊文字を正しく解釈できない, ので下記のようにすると正しく解釈できる

```
curl --http1.0 --data-urlencode title="Head First PHP & MySQL" --data-urlencode author="Lynn Beighley, Michael Morrison" http://localhost:18888
```

## フォームを使ったファイルの送信

curlコマンドを使用してファイルを送信する

```
curl --http1.0 -F title="The Art of Community" -F author="Jono Bacon" -F attachment-file=@test.txt http://localhost:18888
```

ファイルの内容をtest.txtから取得し, 送信ファイル名はローカルファイルと同じ, 形式も自動設定

```
curl --http1.0 -F attachment-file=@test.txt http://localhost:18888
```

ファイルの内容をtest.txtから取得。形式は手動で指定

```
curl --http1.0 -F "attachment-file=@test.txt;type=text/html" http://localhost:18888
```

ファイルの内容をtest.txtから取得. ファイル名は指定したファイル名を利用


```
curl --http1.0 -F "attachment-file=@test.txt;filename=sample.txt" http://localhost:18888
```

## フォームを利用したリダイレクト

300番コードによるリダイレクトとは別に, HTMLのフォームを利用したリダイレクトが存在する

一瞬白紙のページが表示されてしまう, JavaScriptが無効になってると使えない代わりに, データ通信量を節約できる

## コンテントネゴシエーション

| リクエストヘッダー | レスポンス | ネゴシエーション対象 |
| ---- | ---- | ---- |
| Accept | Content-Type | MIMEタイプ |
| Accept-Language | Content-Language | 表示言語 |
| Accept-Charset | Content-Type | 文字のキャラクターセット |
| Accept-Encoding | Content-Encoding | ボディの圧縮 |

## クッキー

ウェブサイトの情報をブラウザ側に保存させる仕組み

```
Set-Cookie: Last_ACCESS_DATE=Jul/31/2016
Set-Cookie: Last_ACCESS_TIME=12:04
```

といったヘッダーの形式でサーバーから送られる. サーバーから送られブラウザに保存された後は

```
Cookie: Last_ACCESS_DATE=Jul/31/2016
Cookie: Last_ACCESS_TIME=12:04
```

といったヘッダーがサーバーに送られる

## 認証とセッション

BASIC認証とDigest認証がある

しかし最近はBASIC認証もDigest認証も使われておらず、フォームを使ったログインとクッキーを使ったセッション管理の組み合わせがよく使われる

ただし、BASIC認証やDigest認証と違いユーザーidやパスワードといった情報を直接サーバーに投げるため、SSL/TLSの利用が必須

## プロキシ

HTTPなどの通信を中継する仕組み. curlの場合はこのようにユーザー名とパスワードを設定して中継する

```
curl --http1.0 -x http://localhost:18888 -U user:pass http://example.com/helloworld
```

## キャッシュ

ページのコンテンツをブラウザにキャッシュしてサーバーとの通信不可を抑制する仕組み

## キャッシュされたファイルとWebサーバーのファイルとが同じかどうかを判断する

- Last-Modified
  - If-Modified-Since
  - ETag

## キャッシュの期限を決めて期限中はキャッシュ参照

- Expires
- Cache-Control

## リファラー

ユーザーがどの経路からページに到達したかをサーバーが把握するために, クライアントがサーバーに送る`Referrer-Policy`ヘッダー

## 検索エンジン向けのコンテンツのアクセス制御

- robots.txt
- sitemap.xml
