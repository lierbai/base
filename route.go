package base

// Route 表示包含方法和路径及其处理程序的请求路由的规范.
type Route struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

// Routes defines a RouteInfo array.
type Routes []Route
