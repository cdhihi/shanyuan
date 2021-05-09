package shanyuan

func If(b bool, ret1, ret2 interface{}) interface{} {
	if b {
		return ret1
	}

	return ret2
}

func Verson() string {

	return "cdhihi => v0.0.1"
}
