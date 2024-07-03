package main

var (
	mysqlUrl = "mysql://blabla"
	// 全局数据库实例
	db = NewMySQLClient(mysqlUrl)
)

func NewMySQLClient(url string) *MySQLClient {
	return &MySQLClient{url: url}
}

type MySQLClient struct {
	url string
}

func (c *MySQLClient) Exec(query string, args ...interface{}) string {
	return "data"
}

func NewApp() *App {
	return &App{}
}

type App struct {
}

func (a *App) GetData(query string, args ...interface{}) string {
	data := db.Exec(query, args...)
	return data
}

// 不使用依赖注入
func main() {
	app := NewApp()
	rest := app.GetData("select * from table where id = ?", "1")
	println(rest)
}
