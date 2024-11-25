package beans

type DefaultBeanFactory struct {
	beans map[string]interface{}
}

func NewDefaultBeanFactory() *DefaultBeanFactory {
	return &DefaultBeanFactory{beans: make(map[string]interface{})}
}

func (d *DefaultBeanFactory) GetObject(name string) interface{} {
	return d.beans[name]
}

func (d *DefaultBeanFactory) RegisterObject(name string, bean interface{}) {
	d.beans[name] = bean
}
