package chain

type Chain func(next interface{}) interface{}

func NewChain(v ...Chain) Chain {
	return func(next interface{}) interface{} {
		if len(v) > 0 {
			for i := len(v) - 1; i >= 0; i-- {
				next = v[i](next)
			}
		}
		return next
	}
}
