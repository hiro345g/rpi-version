# rpi-version

これは Raspberry Pi のリビジョンコードに含まれる情報を表示するコマンドツールです。

Raspberry Pi のリビジョンコードに含まれる情報についての公式ドキュメントについては次の URL を参照してください。

- [Raspberry Pi hardware new style revision codes](https://www.raspberrypi.com/documentation/computers/raspberry-pi.html#new-style-revision-codes)

また、このプログラムは、下記の Python 版を移植したものとなります。

- <https://qiita.com/nak435/items/3b16e7a7a300ee891421> (author: <https://qiita.com/nak435>)

## ビルドとインストール

`build.sh` を実行します。これで `bin/rpi-version` が作成されます。

```sh
sh build.sh
```

ビルドした `rpi-version` を Raspberry Pi OS へコピーします。

```sh
ssh pi@raspberrypi.local sh -c "if [ ! -e /home/pi/.local/bin/ ]; then mkdir -p /home/pi/.local/bin/; fi"
scp bin/rpi-version pi@raspberrypi.local:.local/bin/
```

## 実行

Raspberry Pi にログインして実行します。

```sh
rpi-version
```

SSH 経由で実行する場合は次のようにします。

```sh
ssh pi@raspberrypi.local /home/pi/.local/bin/rpi-version
```

実行例

```console
$ rpi-version
Revision: c03112
        Model: 4B
        Revision: 1.2
        Memory Size: 4GB
        Processor: BCM2711
        Manufacturer: Sony UK
```
