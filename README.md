# compress-image

JPEG画像のファイルサイズを削減するCLIツール

## 使い方
```bash
go run main.go  
```

### 引数

| 引数 | 説明 |
|------|------|
| 元の画像のパス | 圧縮したいJPEG画像のパス |
| 出力先の画像のパス | 圧縮後の画像の保存先パス |

### 例
```bash
go run main.go input.jpg output.jpg
```
