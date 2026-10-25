package main 

func TrimUnicodeSpace(s string) string{
	if s == ""{
		return ""
	}
	sNew := []rune(s)
	sOld := sNew
	for {
		if sNew[0] == ' '{
			sNew = append(sNew[:0], sNew[1:]...)
		}
		if sNew[len(sNew)-1] == ' '{
			sNew = sNew[:len(sNew)-1]
		}
		if string(sNew) == string(sOld){
			return string(sNew)
		}
		sOld = sNew
	}
	
}