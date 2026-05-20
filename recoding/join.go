package main

func JoinSlice(slice1, slice2 []string) []string {
	for i := range slice1 {
		slice1[i] += slice2[i]
	}
	return slice1
}
