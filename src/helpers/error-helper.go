package helpers

func GetAllErrors(err []error) []string {
	var values []string

	if err != nil {
		for _, er := range err {
			values = append(values, er.Error())
		}
	}

	return values
}
