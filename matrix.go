package matrix

import (
	"fmt"
)

type Matrix []float64
type Row []float64
type Builder []Row


func Convert(build Builder) Matrix{
	mtx := genMatrix(len(build[0]),len(build))
	return mtx
}


func genMatrix(rowlen, colmlen int) Matrix{
	mtx := make(Matrix, rowlen*colmlen+2)
	mtx[0] = float64(rowlen)
	mtx[1] = float64(colmlen)
	return mtx
}

func (mtx Matrix)build(builder Builder) (Matrix,error){
	var resultmtx Matrix
	if len(builder) == 0 || len(builder[0]) == 0 {
		err := fmt.Errorf("Error Matrix empty matrix")
		return resultmtx,err
	}

	resultmtx = genMatrix(len(builder), len(builder[0]))
	for i, row := range builder {
		for j, value := range row {
			resultmtx[resultmtx.GetNewPos(i, j)] = value
		}
	}
	return resultmtx,nil

}

// Position in row * column length + position in column + 2([0] and [1] are column and row length)
func (mtx Matrix) GetNewPos(rowPos,colPos int) int{
	return rowPos * int(mtx[1]) + colPos +2   
}

func (mtx Matrix) getRowLen() int{
	return int(mtx[0])
} 

func (mtx Matrix) getColLen() int{
	return int(mtx[1])
} 

func 