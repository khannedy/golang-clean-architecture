package exception

func PanicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
}
