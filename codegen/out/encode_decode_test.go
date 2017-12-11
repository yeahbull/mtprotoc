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

package mtproto

import "github.com/nebulaim/telegramd/mtproto"

// error#c4b9f9bb code:int text:string = Error;
//type TLError struct {
//	Data *mtproto `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
//}
//

/*
///////////////////////////////////////////////////////////////////////////////
// Error <--
//  + TL_error
//
message Error_Data {
	int32 code = 1;
	string text = 2;
}

message Error {
	TLConstructor constructor = 1;
	Error_Data data = 2;
}

// error#c4b9f9bb code:int text:string = Error;
message TL_error {
	Error_Data data = 2;
}
*/

func (m *Error) To_TLError() *TLError {
	return &TLError{
		Data2: m.Data2,
	}
}

func (m *TLError) To_Error() *Error {
	return &Error{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLError) SetCode(v int32) { m.Data2.Code = v }
func (m *TLError) GetCode() int32  { return m.Data2.Code }

func (m *TLError) SetText(v string) { m.Data2.Text = v }
func (m *TLError) GetText() string  { return m.Data2.Text }

func (m *TLError) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_error))
	x.Int(m.Data2.Code)
	x.String(m.Data2.Text)
	return x.buf
}

func (m *TLError) Decode(dbuf *DecodeBuf) error {
	m.Data2.Code = dbuf.Int()
	m.Data2.Text = dbuf.String()
	return dbuf.err
}
