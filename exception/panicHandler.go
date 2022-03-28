package exception

func IsPanic(err interface{}) {
	if err != nil {
		panic(err)
	}
}
