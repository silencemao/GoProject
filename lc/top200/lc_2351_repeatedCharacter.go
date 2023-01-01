package main

/*
第一个出现两次的字符
*/
func repeatedCharacter(s string) byte {
	set := map[byte]int{}
	for i := 0; ; i++ {
		set[s[i]] += 1
		if set[s[i]] == 2 {
			return s[i]
		}
	}
}

func main() {

}
