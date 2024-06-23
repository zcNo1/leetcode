package main

import (
	"log"
)

func main() {
	//log.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	//log.Println(areSentencesSimilar("of", "A lot of words"))
	//log.Println(areSentencesSimilar("Eating right now", "Eating"))
	//log.Println(areSentencesSimilar("Luky", "Lucccky"))
	//log.Println(areSentencesSimilar("B", "ByI BM"))
	//log.Println(areSentencesSimilar("By", "BI BMy"))
	//log.Println(areSentencesSimilar("By", "BI aBy"))
	//log.Println(areSentencesSimilar("By aa", "BI By aa"))
	log.Println(areSentencesSimilar("By aa", "By aa asd"))
}

func areSentencesSimilar(a string, b string) bool {
	lenA := len(a)
	lenB := len(b)
	if lenA == lenB {
		return a == b
	}
	if lenA > lenB {
		a, b = b, a
		lenA, lenB = lenB, lenA
	}

	log.Printf("len a: %d, len b %d", lenA, lenB)
	for left, right := 0, lenB-lenA; right <= lenB; left, right = left+1, right+1 {
		c := b[:left] + b[right:]
		log.Printf("left: %d, right: %d", left, right)
		log.Println("|" + c + "|")
		log.Println("|" + b[left:right] + "|")
		if c == a {
			if left == 0 && b[right-1] == ' ' {
				return true
			} else if right == lenB && b[left] == ' ' {
				return true
			} else if b[left] == ' ' && b[right] == ' ' {
				return true
			}
		}
	}

	return false
}
