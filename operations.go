package matrix

import(
	"fmt"
)

func Add(mtx1, mtx2 Matrix) (Matrix,error){
	if !eqSize(mtx1, mtx2) {
		return nil, fmt.Errorf("%v","Error Matrices are not the same size")
	}
	resultMtx := genMatrix(mtx1.getRowLen(),mtx1.getColLen())
	for i, ind := range mtx1{
		if i > 1 {
			resultMtx[i] = ind + mtx2[i]
		}
	}
	return resultMtx, nil
}

func Subtract(mtx1, mtx2 Matrix) (Matrix, error){
	if !eqSize(mtx1, mtx2) {
		return nil, fmt.Errorf("%v","Error Matrices are not the same size")
	}
	resultMtx := genMatrix(mtx1.getRowLen(),mtx1.getColLen())
	for i, ind := range mtx1{
		if i > 1 {
			resultMtx[i] = ind - mtx2[i]
		}
	}
	return resultMtx, nil
}