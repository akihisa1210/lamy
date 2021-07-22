# lamy

『独学大全』（読書猿著、ダイヤモンド社、2020）の技法17「ラミのトポス」の質問を行うCLIツールです。

## 使い方

### 通常の質問

```sh
$ lamy 猫
(類) 猫 は何の一種か？

(種差) 猫 は、同じグループの中で他と何が違うのか？

(部分) 猫 を構成する部分を列挙すると？

(定義) 猫 とは何か？

(語源) 猫 の語源は？

(相反) 猫 の反対は？

(原因・由来) 猫 を生じさせる（た）ものは？

(結果・派生) 猫 から生じる（た）ものは？
```

### 技術的なトピック用の質問

```sh
$ lamy -t DNS
(類) DNS は何の一種か？

(種差) DNS は、同じグループの中で他と何が違うのか？

(部分) DNS を構成する部分を列挙すると？

(連携) DNS と連携するものは？

(定義) DNS とは何か？

(ユースケース) DNS の具体的なユースケースは？

(利点) DNS を使うと何が嬉しいのか？

(欠点) DNS の欠点は？
```

### インタラクティブモード

```sh
$ lamy -i 犬
? (類) 犬 は何の一種か？
...
```

## インストール

### homebrew

```sh
brew install akihisa1210/lamy/lamy
```

### Download releases

[Release](https://github.com/akihisa1210/lamy/releases)ページから各OS用のバイナリを含むZIPファイルをダウンロードできます。解凍して `lamy` もしくは `lamy.exe` を任意のディレクトリに配置してください。

## 開発

### Git Hook の設定

`.githooks` ディレクトリ以下に本プロジェクト用の Git Hook が配置されているので設定する。

```sh
git config --local core.hooksPath .githooks
```

### 動作確認

```sh
go run ./main
```

### ビルド（ローカル）

```sh
goreleaser build --rm-dist --snapshot
```

### リリース

```sh
git tag vX.X.X
git push origin vX.X.X
```
