package util

//CheckErr checks and logs errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
