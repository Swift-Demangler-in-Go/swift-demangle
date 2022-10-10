// @File: Err.go
// @Author: Jason
// @Date: 2022/8/10

package demangle

import "errors"

// Err 含错误信息的AST结点
type Err struct {
	err string
}

// Error 实现error接口
func (e *Err) Error() string {
	return e.err
}

// Print Err结点的打印规则
func (e *Err) Print(ps *PrintState) {
	ps.WriteString(e.err)
}

var (
	// ErrInvalidInput 输入不合法
	ErrInvalidInput = errors.New("get invalid swift mangle parser")
)
