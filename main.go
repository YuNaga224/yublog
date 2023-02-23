package main

import (
	"time"
	"net/http"
	"strconv"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
// テンプレートファイルを配置するディレクトリへの相対パス
const tmplPath = "src/template/"

// アプリケーションをグローバル変数に格納
var e = createMux()

func main() {
	e.GET("/", articleIndex)
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow)
	e.GET("/:id/edit",articleEdit)
	e.Logger.Fatal(e.Start(":8080"))
}


// アプリケーションのインスタンスを生成
func createMux() *echo.Echo {
	e := echo.New()
	//各種ミドルウェアの設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	//src/cssディレクトリ以下のファイルを/cssのパスで静的ファイルとして配信
	e.Static("/css", "src/css")
	// jsを静的ファイルとして配信
	e.Static("/js", "src/js")

	return e
}



// HTTPハンドラ
func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "articleIndexページです",
		"Now": time.Now(),
	}
	return render(c, "article/index.html", data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "articleNewのページです",
		"Now": time.Now(),

	}
	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	// pathパラメータ:idから値を抽出して、strconv.Atoiで数値型にキャスト
	id,_ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "articleShowのページです",
		"Now": time.Now(),
		"ID": id,
	}
	return render(c, "article/show.html",data)
}

func articleEdit(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "articleEditのページです",
		"Now": time.Now(),
		"ID": id,
	}
	return render(c, "article/edit.html",data)
}

//テンプレートファイルとデータをもとにHTMLファイルを生成
func htmlBlob(file string, data map[string]interface{}) ([]byte, error){
	//生成したHTMLをバイトデータとして返却
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

// HTMLのレンダリング
func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file,data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
