package main

func NewMySQLClient(url string) *MySQLClient {
	return &MySQLClient{url: url}
}

type MySQLClient struct {
	url string
}

func (c *MySQLClient) Exec(query string, args ...interface{}) string {
	return "data"
}

func NewApp(client *MySQLClient) *App {
	return &App{client: client}
}

type App struct {
	// App 持有唯一的 MySQLClient 实例
	client *MySQLClient
}

func (a *App) GetData(query string, args ...interface{}) string {
	data := a.client.Exec(query, args...)
	return data
}

// 手动依赖注入
func main() {
	client := NewMySQLClient("mysql://blabla")
	app := NewApp(client)
	rest := app.GetData("select * from table where id = ?", "1")
	println(rest)
}
