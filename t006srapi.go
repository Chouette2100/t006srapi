/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php

*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//	"net/http/cgi"

	"github.com/dustin/go-humanize"
)

/*
	Webサーバーの作り方の簡単な例です。

	ソースのダウンロード、ビルドについて以下簡単に説明します。詳細は以下の記事を参照してください。

		【Windows】Githubにあるサンプルプログラムの実行方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/e27fc9

		【Unix/Linux】Githubにあるサンプルプログラムの実行方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/220e38

		【Windows】SHOWROOMのAPI関連パッケージ部分を含めたビルドの方法
			https://zenn.dev/chouette2100/books/d8c28f8ff426b7/viewer/fe982a

			（ロードモジュールさえできればいいということでしたらコマンド一つでできます）

	$ cd ~/go/src

	$ curl -OL https://github.com/Chouette2100/t006srapi/archive/refs/tags/v0.0.0.tar.gz
	$ tar xvf v0.0.0.tar.gz
	$ mv t006srapi-0.0.0 t006srapi
	$ cd t006srapi

		以上4行は、Githubからソースをダウンロードしてます。v0.0.0のところは、ソースのバージョンを指定します。
		バージョンは、Githubのリリースページで確認してください（当面 v0.0.0 のままにしておくつもりです）
		ダウンロードはどんな方法でも構わなくて、 とくにWindowsの場合、ZIPをブラウザでダウンロードして
		エクスプローラで解凍した方が簡単でしょう。

		極端な話Githubのソース
			https://github.com/Chouette2100/t006srapi/t006.go
		をコピペでエディターに貼り付けてもOKです。templates/top.gtplとpublic/index.htmlもお忘れなく！
		詳細は上に紹介した三つの記事にあります。

	$ go mod init
	$ go mod tidy
	$ go build t006srapi.go
	$ ./t006srapi

	ここでブラウザを起動し

	　　		http://localhost:8080/top

	と入力してください。


	Ver. 0.0.0

*/

type PointAndRank struct {
	Eventname string
	Roomid    int
	Point     int
	Rank      int
	Gap       int
}

type Top struct {
	Title        string
	ErrMsg       string
	PointAndRank []PointAndRank
}

//	テンプレートに埋め込むデータを作成する。
//	実際にはAPIで取得したり、DBから取得したりする。
func MakeListOfPoints() (pd []PointAndRank, err error) {

	//	pd = make([]PointAndRank, 0)

	pd = []PointAndRank{
		{"イベントA", 100001, 1234567, 1, 34567},
		{"イベントA", 100002, 1200000, 2, 34567},
		{"イベントA", 100003, 999999, 3, 200001},
		{"イベントB", 100004, 2222222, -1, -1},
		{"イベントB", 100005, 2000000, -1, -1},
		{"", 100006, -1, -1, -1},
		{"", 100007, -1, -1, -1},
	}

	return
}

//	"/top"に対するハンドラー
//	http://localhost:8080/top で呼び出される
func HandlerTopForm(
	w http.ResponseWriter,
	r *http.Request,
) {

	var err error

	//	テンプレートで使用する関数を定義する
	funcMap := template.FuncMap{
		"Comma": func(i int) string { return humanize.Comma(int64(i)) }, //	3桁ごとに","を入れる関数。
	}
	// テンプレートをパースする
	tpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/top.gtpl"))

	// テンプレートに埋め込むデータ（ポイントやランク）を作成する
	top := new(Top)
	top.Title = "SHOWROOM イベント獲得ポイント"
	top.PointAndRank, err = MakeListOfPoints()
	if err != nil {
		err = fmt.Errorf("MakeListOfPoints(): %w", err)
		log.Printf("MakeListOfPoints() returned error %s\n", err.Error())
		top.ErrMsg = err.Error()
	}

	// テンプレートへのデータの埋め込みを行う
	if err = tpl.ExecuteTemplate(w, "top.gtpl", top); err != nil {
		log.Printf("tpl.ExecuteTemplate() returned error: %s\n", err.Error())
	}

}

//Webサーバーを起動する。
func main() {

	rootPath := ""

	// URLがホスト名だけのときは public/index.htmlが表示されるようにしておきます。
	http.Handle("/", http.FileServer(http.Dir("public"))) // http://localhost:8080/ で呼び出される。

	//	URLに対するハンドラーを定義する。この例では1行しかないが、実際はURLがあるだけ羅列する。
	http.HandleFunc(rootPath+"/top", HandlerTopForm) //	http://localhost:8080/top で呼び出される。

	//	ポートは8080などを使います。
	//	Webサーバーはroot権限のない（特権昇格ができない）ユーザーで起動した方が安全だと思います。
	//	その場合80や443のポートはlistenできないので、ルータやOSの設定でポートフォワーディングするか
	//	ケーパビリティを設定してください。
	//	# setcap cap_net_bind_service=+ep ShowroomCGI
	//　（設置したあとこの操作を行うこと）
	httpport := "8080"

	err := http.ListenAndServe(":"+httpport, nil) //	Webサーバーとして起動する。
	//	err := cgi.Serve(nil)					  //	CGIとして起動する。
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
