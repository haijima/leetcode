package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	step := 2*numRows - 2

	output := ""
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += step {
			output += string(s[j])
			if i != 0 && i != numRows-1 && j+step-2*i < len(s) {
				output += string(s[j+step-2*i])
			}
		}
	}
	return output
}
