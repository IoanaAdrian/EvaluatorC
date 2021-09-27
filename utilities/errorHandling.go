package utilities

func HandleErr(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
	}
}

// HandleErrR TODO: _,err :=
func HandleErrR(r interface{}, err error) interface{} {
	if err!=nil{
		panic(err)
	}
	return r
}

