# HTTP/2のシンタックス

HTTP/2では通信の高速化がなされ、かつ相乗りプロトコルも増えた

## HTTP/2

https://techblog.yahoo.co.jp/infrastructure/http2/introduction_to_http2/
https://zenlogic.jp/aossl/basic/response/

- ストリーム
  - HTTP/1.0ではリクエストを1つ送ってはレスポンスを待つ、という方式だったので、ネットワークの遅延時間（レイテンシ）が大きいと、応答がどんどん遅くなる
  - HTTP/1.1のパイプラインでは、複数のリクエストを同時に送信でき、ネットワークの遅延を（見掛け上）隠ぺいできる
    - だがHTTPのリクエストは逐次処理されるため、あるリクエストの応答が滞ると、後続の応答も止まってしまう
  - HTTP/2のストリームでは処理が完了したものから順に（順不同で）レスポンスが返されるため、停滞の影響を受けにくい
- サーバープッシュ
  - サーバーがHTMLを生成しているあいだにCSSやJavaScriptのファイルをプッシュしておけば、HTMLが届いた瞬間にWebページを表示できる
- ヘッダーの圧縮
  - 1度送信したヘッダーは基本的には再度送信することはなく、新たに送信が必要なヘッダーのみを差分として抽出して送信することで転送量を削減

## Server-Send Events

`Content-type:text/event-stream` を利用し、HTTPレスポンスを閉じずにChunked形式でデータをクライアントに返し続ける

## WebSocket

ポーリングやCometやServer-Send Eventsと違い、双方向な通信ができる

UpgradeヘッダーやConnectionヘッダーでWebSocketへプロトコルアップデートを指定する

クライアント側は以下のリクエストヘッダーを投げると

```
Upgrade : websocket
Connection: Upgrade
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: hogehoge
Sec-WebSocket-Protocol: chat
```

サーバー側は以下のレスポンスを返す

```
Upgrade : websocket
Connection: Upgrade
Sec-WebSocket-Version: 13
Sec-WebSocket-Accept: fugafuga
Sec-WebSocket-Protocol: chat
```

クライアントとサーバーは `Sec-WebSocket-Key` `Sec-WebSocket-Accept` でリクエストとレスポンスの接続のハンドシェイクをする

ハンドシェイク後にWebSocketによる双方向通信を行うことができる

`Sec-WebSocket-Protocol` でWebSocketのサブプロトコルを指定する

WebSocketはソケット通信機能だけを提供するため、どのようなデータ形式(文字列やらbinaryやらファイルやらjsonやら)を使うかはアプリケーションで決める必要がある。
