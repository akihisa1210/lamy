# lamy

『独学大全』（読書猿著、ダイヤモンド社、2020）の技法17「ラミのトポス」の質問を行うCLIツールです。

## 使い方

```sh
lamy [質問の対象]
```

## インストール

### homebrew

```sh
brew install akihisa1210/lamy/lamy
```

### Download releases

[Release](https://github.com/akihisa1210/lamy/releases)ページから各OS用のバイナリを含むZIPファイルをダウンロードできます。解凍して `lamy` もしくは `lamy.exe` を任意のディレクトリに配置してください。

## 開発

### ビルド（ローカル）

```sh
goreleaser build --rm-dist --snapshot
```

### リリース

```sh
git tag vX.X.X
git push origin vX.X.X
```
