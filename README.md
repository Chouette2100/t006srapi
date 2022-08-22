# t006srapi
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

		以上4行は、Githubからソースをダウンロードしてます。vn.n.nのところは、ソースのバージョンを指定します。
		バージョンは、Githubのリリースページで確認してください。
		ダウンロードはどんな方法でも構わなくて、 とくにWindowsの場合、以上の操作はGUIを使った方が簡単です。

		極端な話Githubのソース
			https://github.com/Chouette2100/t006srapi/t006.go
		をコピペでエディターに貼り付けてもOKです。
		詳細は上に紹介した三つの記事にあります。

	$ go mod init
	$ go mod tidy
	$ go build t006srapi.go
	$ ./t006srapi

	ここでブラウザを起動し

	　　		http://localhost:8080/top

	と入力してください。
