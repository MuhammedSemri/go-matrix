package matrix

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	builder1 := Builder{
		Row{1, 1, 1},
		Row{1, 1, 1},
	}
	builder2 := Builder{
		Row{1, 1, 1},
		Row{1, 1, 1},
	}
	mtx1, err := Build(builder1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	mtx2, err := Build(builder2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	wantmtx := Matrix{2, 3, 2, 2, 2, 2, 2, 2}
	type args struct {
		mtx1 Matrix
		mtx2 Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				mtx1: mtx1,
				mtx2: mtx2,
			},
			want:    wantmtx,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.mtx1, tt.args.mtx2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {

	builder1 := Builder{
		Row{1, 1, 1},
		Row{1, 1, 1},
	}
	builder2 := Builder{
		Row{1, 1, 1},
		Row{1, 1, 1},
	}
	mtx1, err := Build(builder1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	mtx2, err := Build(builder2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	wantmtx := Matrix{2, 3, 0, 0, 0, 0, 0, 0}

	type args struct {
		mtx1 Matrix
		mtx2 Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				mtx1: mtx1,
				mtx2: mtx2,
			},
			want:    wantmtx,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Subtract(tt.args.mtx1, tt.args.mtx2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
