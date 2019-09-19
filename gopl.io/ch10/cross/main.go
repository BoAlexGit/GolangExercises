// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// cross.go
// See page 295.
//Программа cross выводит операционную систему и архитектуру, для которой она
//была построена:
// The cross command prints the values of GOOS and GOARCH for this target.
package main

import (
	"fmt"
	"runtime"
)

//!+
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

//!-

/*********************************************************
Очень просто выполняется в Go кросс-компиляция, т.е. построение выполнимого
файла, предназначенного для работы с другой операционной системой или процес­
сором. Просто установите переменные среды GOOS и GOARCH на время построения.
Программа cross выводит операционную систему и архитектуру, для которой она
была построена:


$ cat quoteargs.go

package main
import (
"fmt"
"os"
)
func main() {
	fmt.Printf("%q\n", os.Args[1:])
}
$ go build quoteargs.go

$ ./quoteargs one "two three" four\ five
["one" "two three" "four five"]
Часто для разовых программ, таких как эта, мы хотим выполнить и сразу по по­
строении запустить выполнимый файл. Команда go run сочетает в себе эти два шага:
$ go run quoteargs.go one "two three" four\ five
["one" "two three" "four five"]


$ cd $GOPATH/src/gopl.io/ch1/helloworld
$ go build
И
$ cd anywhere
$ go build gopl.io/ch1/helloworld
И
$ cd $GOPATH
$ go build ./src/gopl.io/ch1/helloworld
Но не
$ cd $GOPATH
$ go build src/gopl.io/ch1/helloworld
Ошибка: не найден пакет "src/gopl.io/ch1/helloworld".


 */
// +build linux darwin
перед объявлением пакета (и его документирующим комментарием), go b u i l d будет
компилировать его только при построении для Linux или Mac OS X, а следующий
комментарий указывает, что данный файл никогда не должен компилироваться:
// +build ignore
Более подробную информацию можно найти в разделе Build Constraints документа­
ции пакета g o / b u i l d :
$ go doc go/build


$ go list github.com/go-sql-driver/mysql
github.com/go-sql-driver/mysql
Аргумент go l i s t может содержать символы “ . .
которые соответствуют лю­
бой подстроке пути импорта пакета. Мы можем использовать их для перечисления
всех пакетов в рабочей области Go:
$ go list ...
archive/tar
archive/zip
bufio
bytes
cmd/addr21ine
cmd/api
. . .
и т .д . . . .
Или внутри конкретного поддерева:
$ go list gopl.io/ch3/...
gopl.io/ch3/basenamel
gopl.io/ch3/basename2
gopl.io/ch3/comma
gopl.io/ch3/mandelbrot
gopl.io/ch3/netflag
gopl.io/ch3/printints
gopl.io/ch3/surface
Или связанных с конкретной темой:
$ go list ...xml...
encoding/xml
gopl.io/ch7/xmlselect


$ go list json
hash
{
	"Dir": "/home/gopher/go/src/hash",
	"ImportPath": "hash",
	"Name": "hash",
	"Doc": "Package hash предоставляет интерфейсы для хеш-функций.",
	"Target": "/home/gopher/go/pkg/darwin_amd64/hash.a",
	"Goroot": true,
	"Standard": true,
	"Root": "/home/gopher/go",
	"GoFiles": [
		"hash.go"
	],
	"Imports": [
		"io”
	],
	"Deps": [
		"errors",
		"io",
		"runtime",
		"sync",
		"sync/atomic",
		"unsafe"
	]
}

$ go list -f '{{join .Deps " "}}' strconv
errors math runtime unicode/utf8 unsafe


$ go list -f ’{{.ImportPath}} -> {{join .Imports " "}}' compress/...
compress/bzip2 -> bufio io sort
compress/flate -> bufio fmt io math sort strconv
compress/gzip -> bufio compress/flate errors fmt hash hash/crc32 io time
compress/lzw -> bufio errors fmt io
compress/zlib -> bufio compress/flate errors fmt hash hash/adler32 io
