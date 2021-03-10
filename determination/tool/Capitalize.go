package tool

func Capitalize(str string) string {
	first := []rune(str[:1])
	if first[0] >= 97 && first[0] <= 122 {  
		first[0] -= 32 
		return string(first[0])+str[1:]
	}else{
		return str
	}
}