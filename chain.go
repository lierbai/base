package base

// HandlerFunc 处理器
type HandlerFunc func(*Context)

// HandlersChain 定义HandlerFunc数组.
type HandlersChain []HandlerFunc

// Next()前的所有操作处理好以后处理Last  然后处理Next()后的操作

// Last 返回链的最后一个Handler程序. 即. 最后一个Handler是主要的那个.
func (chain HandlersChain) Last() HandlerFunc {
	if length := len(chain); length > 0 {
		return chain[length-1]
	}
	return nil
}
