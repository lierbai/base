package base

// Param 单个URL参数，由键和值组成.
type Param struct {
	Key   string
	Value string
}

// Params 数字定义参数数组
type Params []Param

// Get 返回键与给定名称匹配的第一个参数的值
// 没有找到匹配的参数，则返回空字符串.
func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

// ByName 返回键与给定名称匹配的第一个参数的值.
// 没有找到匹配的参数，则返回空字符串.
func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}
