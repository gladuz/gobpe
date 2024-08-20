package main

import (
	"fmt"
	"unicode/utf8"
)

type Pair struct{
	A, B int	
}

func getStats(tokens []int) map[Pair]int {
	counts := make(map[Pair]int)
	for i := 0; i < len(tokens)-1; i++ {
		counts[Pair{tokens[i], tokens[i+1]}] += 1
	}
	return counts
}

func findMostCommonPair(tokens []int) Pair{
	counts := getStats(tokens)
	mostCommon := Pair{}
	maxCount := -1
	for k, v := range counts{
		if(v > maxCount){
			maxCount = v
			mostCommon = k
		}
	}
	return mostCommon
}

func merge(ids []int, pair Pair, idx int) []int{
	newIds := make([]int, 0)
	i := 0;
	for i < len(ids){
		if i < len(ids)-1 && ids[i] == pair.A && ids[i+1] == pair.B{
			newIds = append(newIds, idx)
			i += 2
		}else{
			newIds = append(newIds, ids[i])
			i += 1	
		}
	}
	return newIds
}

func main(){
	text := "Ｕｎｉｃｏｄｅ! 🅤🅝🅘🅒🅞🅓🅔‽ 🇺‌🇳‌🇮‌🇨‌🇴‌🇩‌🇪! 😄 The very name strikes fear and awe into the hearts of programmers worldwide. We all know we ought to “support Unicode” in our software (whatever that means—like using wchar_t for all the strings, right?). But Unicode can be abstruse, and diving into the thousand-page Unicode Standard plus its dozens of supplementary annexes, reports, and notes can be more than a little intimidating. I don’t blame programmers for still finding the whole thing mysterious, even 30 years after Unicode’s inception."
	//text = "안녕하세요 ✋ (Hello in Korean)"
	tokens := make([]int, len([]byte(text)))
	for i, b := range []byte(text) {
		tokens[i] = int(b)
	}
	fmt.Printf("Text length: %d, Byte length: %d \n", utf8.RuneCountInString(text), len(tokens))
	mostCommon := findMostCommonPair(tokens)
	fmt.Println(mostCommon)
	fmt.Println("Replaced sequence: {", string(mostCommon.A), string(mostCommon.B), "}")
	fmt.Println(merge(tokens, mostCommon, 256))
}
