package domain

type success struct {
	Data interface{} `json:"data"`
}

type failed struct {
	Msg interface{} `json:"msg"`
}

func ResponseSuccess(data interface{}) success {
	return success{
		Data: data,
	}
}

func ResponseFailed(msg interface{}) failed {
	return failed{
		Msg: msg,
	}
}
