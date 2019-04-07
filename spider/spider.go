package spider

func Trans(word string) string {
	body := Post(word)
	return getMeans(body)
}
