package representation

type IResp interface {
	GetError() error
	GetData() interface{}
}

type Resp struct {
	Data interface{}
	Err  error
}

func (r Resp) GetData() interface{} {
	return r.Data
}

func (r Resp) GetError() error {
	return r.Err
}
