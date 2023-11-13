package conv

type NilCastRule = int

const (
	NilToZero  NilCastRule = iota // 直接返回错误
	NilToError                    // 转换为对应类型的零值
)

type ValueCastRule int

const (
	ValueToZero  ValueCastRule = iota // 直接返回错误
	ValueToError                      // 转换为对应类型的零值
)

type Options struct {
	Nil   NilCastRule   // 如果被转换的类型是nil时如何处理
	Empty ValueCastRule // 如果被转换的类型是空字符串时如何处理
}
