/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_encode_decode.py'
 *
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

// ConstructorList
// RequestList

package mtproto

import (
    // "encoding/binary"
    // "fmt"
    // "github.com/golang/protobuf/proto"
)

{{ range .BaseTypeList }}
///////////////////////////////////////////////////////////////////////////////
// {{.Name}} <--
{{ range .SubMessageList }}//  + TL_{{.Name}}
{{end}}//

func (m *{{.Name}}) Encode() []byte {
    switch m.GetConstructor() {
{{ range $i, $v := .SubMessageList }}    case TLConstructor_CRC32_{{.Predicate}}:
        t := m.To_{{.Name}}()
        return t.Encode()
{{end}}
    default:
        return []byte{}
    }
}

{{ range .SubMessageList }}// {{.Line}}
func (m *{{.ResType}}) To_{{.Name}}() *TL{{.Name}} {
    return &TL{{.Name}}{
        Data2: m.Data2,
    }
}

{{end}}

{{ range $i, $v := .SubMessageList }}// {{.Line}}
func (m *TL{{.Name}}) To_{{.ResType}}() *{{.ResType}} {
    return &{{.ResType}}{
        Constructor: TLConstructor_CRC32_error,
        Data2: m.Data2,
    }
}

{{range $i2, $v2 := .ParamList}}
func (m *TL{{$v.Name}}) Set{{.Name}}(v {{.Type}}) { m.Data2.{{.Name2}} = v }
func (m *TL{{$v.Name}}) Get{{.Name}}() {{.Type}} { return m.Data2.{{.Name2}} }
{{end}}

func (m* TL{{.Name}}) Encode() []byte {
    x := NewEncodeBuf(512)
    x.Int(int32(TLConstructor_CRC32_{{.Predicate}}))

{{range $i2, $v2 :=    .EncodeCodeList}}    {{$v2}}
{{end}}
    return x.buf
}

func (m* TL{{.Name}}) Decode(dbuf *DecodeBuf) error {
{{range $i2, $v2 :=    .DecodeCodeList}}    {{$v2}}
{{end}}
    return dbuf.err
}
{{end}}
{{end}}
