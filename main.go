package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := NewProduct(
		WithParam("Name", "Android"),
		WithParam("ID", 2),
		WithParam("IsLatest", true))
	fmt.Println(p)
}

type Option func(*Product)

type Product struct {
	Name     string
	ID       int
	IsLatest bool
}

func NewProduct(opts ...Option) *Product {

	p := &Product{}
	for _, opt := range opts {
		opt(p)
	}

	return p
}

type Param interface {
	int | string | bool
}

func WithParam[T Param](param string, value T) Option {
	return func(p *Product) {
		// p のフィールドを取得
		v := reflect.ValueOf(p).Elem()
		v.FieldByName(param).Set(reflect.ValueOf(value))
		fmt.Println(v)
	}

}
