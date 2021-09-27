// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import "github.com/gogf/gf/container/gvar"

// SetParam sets custom parameter with key-value pair.
func (r *Request) SetParam(key string, value interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	r.paramsMap[key] = value
}

// GetParam returns custom parameter with given name <key>.
// It returns <def> if <key> does not exist.
// It returns nil if <def> is not passed.
func (r *Request) GetParam(key string, def ...interface{}) *gvar.Var {
	if r.paramsMap != nil {
		return gvar.New(r.paramsMap[key])
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}
