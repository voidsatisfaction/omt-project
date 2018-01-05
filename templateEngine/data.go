package templateEngine

type Data map[string]interface{}

func NewData() Data {
	return make(Data)
}

func (d Data) Add(key string, value interface{}) {
	d[key] = value
}
