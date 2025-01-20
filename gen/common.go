// Package gen
// @Author: yinwei
// @File: common
// @Version: 1.0.0
// @Date: 2025/1/19 10:44

package gen

import "golang.org/x/exp/constraints"

type Address interface {
	constraints.Integer | constraints.Float | ~string | ~bool
}

// BaseToPtr 基本类型转指针
//
// Parameters:
//   - a: T
//
// Returns:
//   - *T
func BaseToPtr[T Address](a T) *T {
	return &a
}
func BaseErrToPtr[T Address](a T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	return BaseToPtr(a), nil
}
