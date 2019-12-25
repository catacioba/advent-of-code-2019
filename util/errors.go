package util

func Must(err error) {
	if err != nil {
		panic("error is not nil")
	}
}
