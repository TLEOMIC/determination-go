package tool

func Collection(map1 map[string]string,map2 map[string]string) map[string]string{
	for i,v := range map2 {
		map1[i] = v
	}
	return map1
}