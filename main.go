package main

import (
	"fmt"
	_ "image/png"
	"io"
	"log"
	"reflect"
)

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fmt.Println(key)
			// display(fmt.Sprintf("%s[%s]", path,
			// 	formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s =%T", path, v)
		// fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abcabcabc"))
}

// 至多包含两个不同字符的最长子串
// https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/
// 滑动窗口
func lengthOfLongestSubstringTwoDistinct(s string) int {
	// 记录窗口内数据
	mapp := make(map[byte]int)
	max := 0
	for i, j := 0, 0; j < len(s); j++ {
		mapp[s[j]]++
		for len(mapp) > 2 {
			// 当字符种类超过 2,从左往右删
			mapp[s[i]]--
			if mapp[s[i]] == 0 {
				delete(mapp, s[i])
			}
			i++
		}
		if max < j-i+1 {
			max = j - i + 1
		}
	}
	return max
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
