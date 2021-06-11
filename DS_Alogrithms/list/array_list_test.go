package list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSqList_Get(t *testing.T) {
	type fields struct {
		data   []elementType
		length int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"empty sqlist",
			fields{[]elementType{}, 0},
			args{4},
			nil,
			true,
		}, {"case 2",
			fields{[]elementType{1, 2, 3}, 3},
			args{4},
			nil,
			true,
		}, {"case 3",
			fields{[]elementType{1, 2, 3, 4, 5}, 5},
			args{4},
			elementType(5),
			false,
		}, {"case 4",
			fields{[]elementType{}, 0},
			args{-1},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			got, err := s.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v[type: %T], want %v[type: %T]", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestSqList_Add(t *testing.T) {
	type fields struct {
		data   []elementType
		length int
	}
	type args struct {
		values []elementType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"case 1:emtpy list add", fields{[]elementType{}, 0}, args{[]elementType{1, 2, 3}}, false},
		{"case 2:emtpy list add", fields{[]elementType{1}, 1}, args{[]elementType{1, 2, 3}}, false},
		{"case 3:add empty list", fields{[]elementType{1}, 1}, args{[]elementType{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			fmt.Println(s.length)
			if err := s.Add(tt.args.values...); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(s.length)
		})
	}
}

func TestSqList_FindByValue(t *testing.T) {
	type fields struct {
		data   []elementType
		length int
	}
	type args struct {
		elem elementType
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex int
		wantErr   bool
	}{
		{"case 1", fields{[]elementType{}, 0}, args{1}, -1, false},
		{"case 2", fields{[]elementType{1, 2, 3}, 3}, args{1}, 0, false},
		{"case 3", fields{[]elementType{1, 2, 3}, 3}, args{4}, -1, false},
		{"case 4", fields{[]elementType{1, 2, 2, 3}, 4}, args{2}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			gotIndex, err := s.FindByValue(tt.args.elem)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("FindByValue() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestSqList_Remove(t *testing.T) {
	type fields struct {
		data   []elementType
		length int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"case 1", fields{[]elementType{}, 0}, args{1}, true},
		{"case 2", fields{[]elementType{1, 2, 3}, 3}, args{4}, true},
		{"case 3", fields{[]elementType{1, 2, 3}, 3}, args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			fmt.Println(s)
			if err := s.Remove(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(s)
		})
	}
}

func TestSqList_Insert(t *testing.T) {
	type fields struct {
		data   []elementType
		length int
	}
	type args struct {
		index int
		value elementType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"case 1", fields{[]elementType{1, 2, 3}, 3}, args{4, 1}, true},
		{"case 2", fields{[]elementType{1, 2, 3}, 3}, args{1, 1}, false},
		{"case 3", fields{[]elementType{1, 2, 3}, 3}, args{2, 4}, false},
		{"case 4", fields{[]elementType{1, 2, 3}, 3}, args{0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			fmt.Println(s)
			if err := s.Insert(tt.args.index, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(s)
		})
	}
}
