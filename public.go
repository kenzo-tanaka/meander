package meander

type Facade interface {
	Public() interface{}
}

// 受け取った任意のオブジェクトに対して
// Facadeが実装されているかどうかをチェック
// 実装されているならPublicを呼び出す
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}

	return o
}
