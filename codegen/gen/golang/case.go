/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gengolang

// ToSnake string, XxYy to xx_yy, X_Y to x_y
func ToSnake(s string) string {
	num := len(s)
	need := false // need determin if it's necessery to add a '_'

	snake := make([]byte, 0, len(s)*2)
	for i := 0; i < num; i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
			if need {
				snake = append(snake, '_')
			}
		} else {
			// if previous is '_' or ' ',
			// there is no need to add extra '_' before
			need = (c != '_' && c != ' ')
		}

		snake = append(snake, c)
	}

	return string(snake)
}

// ToCamel string, xx_yy to XxYy, xx__yy to Xx_Yy
// xx _yy to Xx Yy, the rule is that a lower case letter
// after '_' will combine to a upper case letter
func ToCamel(s string) string {
	num := len(s)
	need := true

	var prev byte = ' '
	camel := make([]byte, 0, len(s))
	for i := 0; i < num; i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			if need {
				c = c - 'a' + 'A'
				need = false
			}
		} else {
			if prev == '_' {
				camel = append(camel, '_')
			}
			need = (c == '_' || c == ' ')
			if c == '_' {
				prev = '_'
				continue
			}
		}

		prev = c
		camel = append(camel, c)
	}

	return string(camel)
}

// ToAbridge extract first letter and all upper case letter
// from string as it's abridge case
func ToAbridge(str string) string {
	l := len(str)
	if l == 0 {
		return ""
	}

	arbi := []byte{str[0]}
	for i := 1; i < l; i++ {
		b := str[i]
		if IsUpper(b) {
			arbi = append(arbi, b)
		}
	}

	return string(arbi)
}

// ToLowerAbridge extract first letter and all upper case letter
// from string as it's abridge case, and convert it to lower case
func ToLowerAbridge(str string) (s string) {
	l := len(str)
	if l == 0 {
		return ""
	}

	arbi := []byte{ToLower(str[0])}
	for i := 1; i < l; i++ {
		b := str[i]
		if IsUpper(b) {
			arbi = append(arbi, ToLower(b))
		}
	}

	return string(arbi)
}
