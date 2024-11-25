package beans

type BeanFactory interface {

	// GetObject 获取bean
	GetObject(name string) interface{}
}
