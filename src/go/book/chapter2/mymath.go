package chapter2

func Plus(args ...int) (ret int) {
	ret = 0
	for _, val := range args {
		ret += val
	}
	return
}
