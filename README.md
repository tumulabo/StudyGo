# StudyGo

## Try tutorial

* [A Tour of Go](http://go-tour-jp.appspot.com/)

### memo

```
Goでは、最初の文字が大文字で始まる場合は、その名前はエクスポートされています。
たとえば、 Pi は外部へ公開される名前ですが、 pi では公開されません。
```

```
型が変数名の 後 にくることに注意してください。
（型をなぜこのように宣言するのか、についての詳細な情報は、 記事「Go's declaration syntax」 を参照してください。）
Go's Declaration Syntax - The Go Blog http://blog.golang.org/gos-declaration-syntax
```

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```

Goは複数の戻り値パラメータを返すことが可能
```
package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
```

var ステートメントは変数を宣言します。 関数の引数リストと同じように、２つ以上連続した 関数パラメータの最後に型を書くことで、変数のリストを宣言できます。
```
package main

import "fmt"

var i int
var c, python, java bool

func main() {
    fmt.Println(i, c, python, java)
}
```

var 宣言では、変数ひとつひとつに初期化子( initializer )を与えることができます。

もし初期化子が指定されている場合、型を省略できます。 その変数は初期化子の型になります。
```go
package main

import "fmt"

var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
    fmt.Println(i, j, c, python, java)
}
```

#### 関数内での暗黙的な型宣言

関数内では、 var 宣言の代わりに、暗黙的な型宣言ができる := の代入文を使うことができます。

なお、関数外でのすべての宣言にはキーワードでの宣言(`var`, func, など)が必要で、 := での暗黙的な宣言は利用できません。

```go
package main

import "fmt"

func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
```

#### 組み込み型
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 の別名

rune // int32 の別名
     // Unicode のコードポイントを表す

float32 float64

complex64 complex128

#### 定数 Constants

```go
定数( Constant )は、 const キーワードを使って変数のように宣言します。

定数は、character、string、boolean、数値(numeric)のみで使えます。

なお、定数は := を使って宣言できません。
```

#### forループ

Goは、 for ループだけを繰り返し文として使います。 Goには while 文はありません！

基本的には、C言語 や Java と同じですが、括弧 ( ) は不要で（付けてはいけません）、中括弧 { } は必要です。

セミコロンを抜くこともできます。つまり、C言語での while は、Goでは for だけを使います

#### if ステートメント

if ステートメントは、 for のように、実行のための短いステートメントを条件の前に書くことができます。

ここで宣言された変数は、 if のスコープだけで有効です。

(ためしに最後の return 文で、 v を使ってみてください。 使えなかったでしょ？)

ちなみに if ステートメントで宣言された変数は、 else ブロック内でも使うことができます。

便利ですね！

```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
```

#### Loops and Function

関数とループを使った簡単な練習として、 ニュートン法 を使った平方根の計算を実装してみましょう。

この問題では、ニュートン法は、 開始点 z を選び、以下の式を繰り返すことによって、 Sqrt(x) を近似します。


最初は、その計算式を10回だけ繰り返し、 x を(1, 2, 3, ...)と様々な値に対する結果がどれだけ正解値に近いかを確認してみてください。

次に、ループを回すときの直前に求めたzの値がこれ以上変化しなくなったとき （または、差がとても小さくなったとき） に停止するようにループを変更してみてください。 この変更により、ループ回数が多くなったか、少なくなったのか見てみてください。 math.Sqrt と比べてどれくらい近似できましたか？

ヒント：浮動小数点を宣言し、値を初期化するには、型のキャストか、浮動小数点を使ってください：

z := float64(1)
z := 1.0

```go
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for i:=0 ; i <= 10 ; i++ {
    	z = z - ((z*z-x)/(2*z))
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
    
}
```

#### struct

構造体のフィールドは、ドット(.)を用いてアクセスします。

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

#### Pointers
Go言語にはポインタがありますが、ポインタ演算はありません。

構造体のフィールドは、構造体のポインタを通してアクセスできます。このポインタを通じた間接的なアクセスで、とてもわかりやすくなります。

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    p := Vertex{1, 2}
    q := &p
    q.X = 1e9
    fmt.Println(p)
}
```


#### struct literals

structリテラルは、フィールドの値を列挙することによって、構造体の初期値の割り当てを示しています。

Name: 構文を使って、フィールドの一部だけを記述することができます。 （この方法での名前のフィールドの指定順序は無関係です。）

訳注：例では X: 1 として X だけを初期化しています。

特別な接頭辞 & は、新しく割り当てられたstructへのポインタを示します。

```go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    p = Vertex{1, 2}  // has type Vertex
    q = &Vertex{1, 2} // has type *Vertex
    r = Vertex{X: 1}  // Y:0 is implicit
    s = Vertex{}      // X:0 and Y:0
)

func main() {
    fmt.Println(p, q, r, s)
}
```

#### The new function

new(T) という表現は、ゼロ初期化した( zeroed ) T の値をメモリに割り当て、そのポインタを返します。

'''
var t *T = new(T)
'''
または、

t := new(T)
と記述できます。

```go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main() {
    v := new(Vertex)
    fmt.Println(v)
    v.X, v.Y = 11, 9
    fmt.Println(v)
}
```

#### Slices

'``go
package main

import "fmt"

func main() {
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n", i, p[i])
    }
}
'``

```go
package main

import "fmt"

func main() {
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)
    fmt.Println("p[1:4] ==", p[1:4])

    // missing low index implies 0
    fmt.Println("p[:3] ==", p[:3])

    // missing high index implies len(s)
    fmt.Println("p[4:] ==", p[4:])
}
```

sliceは、 make 関数で生成します。 これは、ゼロに初期化した配列をメモリに割り当て、その配列を参照したsliceを返す働きをします：

a := make([]int, 5)  // len(a)=5
sliceは、長さと容量( capacity )を持っています。 sliceの容量は、sliceが基礎となる配列で拡大できる最大の長さです。

容量を指定するためには、 make の３番目の引数に渡します：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5
sliceは　再スライス　( re-slicing )によって拡大縮小させることができます（最大容量まで）：

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

```go
package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
```

#### range
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```

#### Range continued
```go
package main

import "fmt"

func main() {
    pow := make([]int, 256)
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
```
