# goping

Go で実装した ICMP ping コマンドラインツールです。ホストへの到達性確認と往復時間（RTT）の計測ができます。

## 特徴

- ホスト名または IP アドレスへの ICMP ping
- 送信回数の指定（`-c` / `--count`）
- `Ctrl+C` で実行中の ping を中断
- 重複応答（DUP）の検出と表示
- 終了時に送信・受信パケット数、パケットロス、RTT 統計を表示

## 必要条件

- [Go](https://go.dev/) 1.26 以降
- ICMP を送受信するため、多くの環境では **管理者権限**（Linux/macOS では `sudo`）が必要です

## インストール

### ソースからビルド

```bash
git clone https://github.com/<your-username>/goping.git
cd goping
go build -o goping .
```

### go install

```bash
go install github.com/<your-username>/goping@latest
```

> `go install` を使う場合は、リポジトリのモジュールパス（`go.mod` の `module` 行）を GitHub の実際のパスに合わせてください。

## 使い方

```bash
# デフォルト（4 回 ping）
sudo ./goping example.com

# 送信回数を指定
sudo ./goping -c 10 192.168.1.1
sudo ./goping --count 3 google.com
```

### オプション

| オプション | 短縮形 | デフォルト | 説明 |
|-----------|--------|-----------|------|
| `--count` | `-c` | `4` | 送信する ping の回数 |

### 出力例

```
PING example.com (93.184.216.34):
64 bytes from 93.184.216.34: icmp_seq=1 time=12.3ms
64 bytes from 93.184.216.34: icmp_seq=2 time=11.8ms
...

--- example.com ping statistics ---
4 packets transmitted, 4 packets received, 0% packet loss
round-trip min/avg/max/stddev = 11.8ms/12.1ms/12.5ms/0.3ms
```

## 権限について

ICMP は raw ソケットを使うため、OS によっては root 権限が必要です。

- **macOS / Linux**: `sudo ./goping <host>` で実行
- 権限なしで動かすには、OS ごとの `cap_net_raw` 付与などの設定が必要な場合があります

## 依存ライブラリ

- [github.com/go-ping/ping](https://github.com/go-ping/ping) — ICMP ping の実装
- [github.com/spf13/cobra](https://github.com/spf13/cobra) — CLI フレームワーク

## ライセンス

[MIT License](LICENSE)
