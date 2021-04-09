package tool

func Implode(field interface{}) (Rts string){
	switch field.(type) {
		case []string :
			for index,val := range field.([]string) {
				if index != 0 {
					Rts = Rts+","+val
				}else{
					Rts = Rts+val
				}
				
			}
	}
	return Rts
}