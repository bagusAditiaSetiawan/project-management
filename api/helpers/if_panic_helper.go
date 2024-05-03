package helpers

func IfPanicHelper(err error) {
	if err != nil {
		panic(err)
	}
}
