package exception

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
