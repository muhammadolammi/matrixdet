package main

func getDet(m Matrix) int {
	if m.Rows == 2 && m.Columns == 2 {
		return (m.Data[0][0] * m.Data[1][1]) - (m.Data[0][1] * m.Data[1][0])

	}
	det := 0
	for i := range m.Columns {
		elem := m.Data[0][i]
		sign := -1 ^ i

		minorMatrix := getMinorMatrix(0, i, m)
		det += sign * elem * getDet(minorMatrix)

	}
	return det
}

func getMinorMatrix(exludeRow, excludeColumn int, mainMatrix Matrix) Matrix {
	return Matrix{}
}
