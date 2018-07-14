package matrix

import (
	"reflect"
	"testing"
)

func TestMatrix_GetRow(t *testing.T) {
	mtx, err := Build(
		Builder{
			Row{1, 1, 1},
			Row{2, 2, 2},
			Row{2, 2, 2},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	wantRow := Row{1, 2, 2}
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		mtx  Matrix
		args args
		want Row
	}{
		{
			name: "test",
			mtx:  mtx,
			args: args{
				rowIndex: 1,
			},
			want: wantRow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mtx.GetRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_GetColumn(t *testing.T) {
	mtx, err := Build(
		Builder{
			Row{1, 1, 1},
			Row{2, 2, 2},
			Row{2, 2, 2},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	wantRow := Row{1, 1, 1}
	type args struct {
		colIndex int
	}
	tests := []struct {
		name string
		mtx  Matrix
		args args
		want Row
	}{
		{
			name: "test",
			mtx:  mtx,
			args: args{
				colIndex: 0,
			},
			want: wantRow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mtx.GetColumn(tt.args.colIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.GetColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
