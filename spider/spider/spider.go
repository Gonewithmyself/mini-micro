package spider

import "log"

func Trans(word string) string {
	body := Post(word)
	log.Println("trans", word)
	return getMeans(body)
}
