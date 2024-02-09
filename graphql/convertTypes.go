package graphql

func ConvertStringSliceToPointerSlice(strSlice []string) []*string {
	var pointerSlice []*string
	for _, str := range strSlice {
		strPointer := &str
		pointerSlice = append(pointerSlice, strPointer)
	}
	return pointerSlice
}
