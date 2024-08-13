package main

type FrequencyTable map[rune]int

func GenerateFrequencyTable(str string) FrequencyTable {
	freqTable := make(FrequencyTable, 0)
	for _, val := range str {
		freqTable[val] += 1
	}
	return freqTable
}
