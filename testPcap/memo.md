# gopacket test

goでpcap使ってうはうはしたーい為のメモ。

## 初期導入手順
goのインストール(chocolateyを使用)

```power
admin> choco install golang
```

### gopacketのインストール

参考: [golang pcap on Windows - Qiita](http://qiita.com/kwi/items/ed1261fb53f78c244e6d)

powershellに環境変数GOPATHを設定する。
この時に絶対パスを指定する必要あり。

```power
> $ENV:GOPATH="C:\Users\Tomoaki\repo\github\StudyGo"
$ENV:PATH+="$PATH;C:\tools\mingw64\bin\"
```

gopcaketのインストール
```power
go get code.google.com/p/gopacket
```

```power
PS C:\Users\Tomoaki\repo\github\StudyGo> go run .\src\pcapTest.go
# code.google.com/p/gopacket/pcap
src\code.google.com\p\gopacket\pcap\pcap.go:18:18: fatal error: pcap.h: No such file or directory
 #include <pcap.h>
                  ^
```

### WinPcap developer

gopacketがWinPcapのCライブラリを利用しビルドし動作するため、WinPcap developer版の導入が必要。
合わせてCのビルド環境としてgccが必要。

Reference: [WindowsPCAP - gopacket - How to build gopacket PCAP library on Windows - Packet decoding for the Go programming language. - Google Project Hosting](https://code.google.com/p/gopacket/wiki/WindowsPCAP)

#### gccのインストール
```
choco install mingw
```

#### WinPcap developer resourcesのダウンロード

https://www.winpcap.org/devel.htm

展開したzip内のWpdPackフォルダをC:\に移動させる。
(gopacket内にヘッダパスが直書きされているため…)

#### (間違い)WinPcapのインストール

※tsumura 環境での問題
chocolateyでWiresharkインストール時にWinPcapのインストールには失敗していたため入れなおしが必要だった。

%%普通はWiresharkが入っていれば導入されているはず。%%

C:\Users\Tomoaki\repo\github\StudyGo\src\code.google.com\p\gopacket

#### サンプルコードの実行

NICの一覧を取得するコード

```go
package main

import (
        "code.google.com/p/gopacket/pcap"
        "fmt"
)

func main() {
        ifs, err := pcap.FindAllDevs()
        if err == nil {
                for _, ife := range ifs {
                        fmt.Println(ife.Name, ife.Description)
                }
        }
}
```

```power
PS C:\Users\Tomoaki\repo\github\StudyGo> go run .\src\pcapTest.go
\Device\NPF_{0EFED103-BBE3-491C-958F-C1E21FB26E43} Microsoft
\Device\NPF_{9A12C697-E382-4EC2-A7F0-ED2CE8EB155A} Realtek PCIe GBE Family Controller
\Device\NPF_{083CF7C6-D1C8-4AA5-97DF-9E34A72E18C2} Microsoft
\Device\NPF_{94F8E4C5-FB80-455C-B93F-EA2ADC62D261} Microsoft
\Device\NPF_{DA933A4E-DAB2-404F-8B22-474D955F54A3} Sun
```


## Reference

* [Big Sky :: THE GO TOOL](http://mattn.kaoriya.net/software/lang/go/20120216093718.htm)
