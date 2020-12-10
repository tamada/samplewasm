# HTTP GET sample in WASM created by Go lang

Go言語のWASM (Web Assembly) で HTTP GET を呼び出すと，以下のようなデッドロックが発生する．

```
all goroutines are asleep - deadlock!
```

解決策として，goroutine を使って呼び出せば良いとされている．
それで実際にやってみたのがこのリポジトリの内容である．

* https://github.com/golang/go/issues/26382
    * https://github.com/golang/go/issues/26382#issuecomment-416926829

## 使い方

* `make` を実行すると，`server` と `site/main.wasm` が作成される．
* `server`を実行すると，`localhost:8080` に HTTP サーバが建てられる（`make start`を実行しても良い）．
    * `site` 以下のファイルをそのまま返す簡易HTTPサーバ．`cmd/server/main.go` にソースコードあり．
* `http://localhost:8080` にアクセスし，テキストフィールドに適当なURLを入力し，「Get Content」ボタンを押すと「Content」の下に内容が表示される．

## 構成

* `Makefile`
    * ビルドファイル．
* `httpget.go`
    * 単純に HTTP GET しているファイル．
* `src/wasm/main.go`
    * wasmから HTTP GET を goroutine で呼び出している．
* `src/server/main.go`
    * 簡易 HTTP サーバのプログラム．
* `site/index.html`
    * サンプルの html ファイル．
* `site/sample.js`
    * `site/index.html` のアクションを定義し，wasm を呼び出している．
* `site/wasm_exec.js`
    * `${GO_INSTALLED_DIR}/misc/wasm/wasm_exec.js` からコピーしたもの．
