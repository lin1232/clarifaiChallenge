package util

import (
	"sort"
	"fmt"
	"strings"
)

type Pair struct{
	Key string
	Value float64
}

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

type PairList []Pair


func Sort_by_value(p PairList) PairList {
    sort.Sort(p)
    return p
}

func Merge_map(m1, m2 map[string]map[string]float64){
	
	for key, value := range m2{
		m1[key] = value
	}
}

func Output_map(m map[string]map[string]float64){
	i:=0
	for key,value := range m{
		fmt.Printf("%d: %s:", i, key)
		i++
		j := 0
		for key1, value1:= range value{
			fmt.Printf("%d %s %f ",j, key1, value1)
			j++
		}
		fmt.Print("\n")
	}
}

func Split(str string, token string )[]string{
	strs := strings.Split(str, token)
	//strs[len(strs)-1] = strs[len(strs)-1][0:len(strs[len(strs)-1])-2]
	return strs
	
}