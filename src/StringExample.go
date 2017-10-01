package main

import (
	"fmt"
	"strings"
	"unicode"
)

type StringExample struct {

}


func (runner *StringExample) example1() {
	var str1 = "foo"
	var str2 = "bar"
	var strCombined = strings.Join([] string{str1 , str2}, " ")
	var contains_foo = strings.Contains(strCombined, str1)
	var contains_bar = strings.Contains(strCombined , str2)
	var contains_bim = strings.Contains(strCombined, "bim")
	var str1_compare_str2 = strings.Compare(str1, str2)
	var has_prefix_foo = strings.HasPrefix(strCombined, str1)
	var has_suffix_bar = strings.HasSuffix(strCombined, str2)
	var repeat = strings.Repeat("*", 10)
	var replace = strings.Replace(strCombined, "foo", "FOO", -1)
	//strings.Map()
	var title = strings.ToTitle(strCombined)
	var titleSpecial = strings.ToTitleSpecial(unicode.AzeriCase, strCombined)
	var lower = strings.ToLower(strCombined)
	var lowerSpecial = strings.ToLowerSpecial(unicode.AzeriCase, strCombined)
	var upper = strings.ToUpper(strCombined)
	var upperSpecial = strings.ToUpperSpecial(unicode.AzeriCase, strCombined)
	var splitted = strings.Split(strCombined , " ")
	fmt.Println("<<String methods>>")
	fmt.Println("str1 = ", str1)
	fmt.Println("str2 = ", str2)
	fmt.Println("strCombined = ", strCombined)
	fmt.Println("contains foo = ", contains_foo)
	fmt.Println("contains bar = ", contains_bar)
	fmt.Println("contains bim = ", contains_bim)
	fmt.Println("str1 compare str2 = ", str1_compare_str2)
	fmt.Println("has prefix foo = ", has_prefix_foo)
	fmt.Println("has suffix bar = ", has_suffix_bar)
	fmt.Println("repeat = ", repeat)
	fmt.Println("replace = ", replace)
	fmt.Println("title = ", title)
	fmt.Println("titleSpecial = ", titleSpecial)
	fmt.Println("lower = ", lower)
	fmt.Println("lowerSpecial = ", lowerSpecial)
	fmt.Println("upper = ", upper)
	fmt.Println("upperSpecial = ", upperSpecial)
	fmt.Println("<<splitted>>")
	for _, el := range(splitted) {
		fmt.Printf("%s ; ", el)
	}
	fmt.Println()


	myUpper := func(r rune) rune {
		return r  - 32  //to upper case A(65) , a(97)
	}

	var mapped = strings.Map(myUpper , strCombined)
	var size = len(strCombined)
	fmt.Println("mapped = " , mapped)
	fmt.Println("size = ", size)
	var ch uint8 = strCombined[0]
	fmt.Println("first char of strCombined = ", string(ch))
	fmt.Println("<<iterating over strCombined>>")
	for i, c := range(strCombined) {
		fmt.Printf("%d => %s\n", i , string(c))
	}
	fmt.Println()
}
