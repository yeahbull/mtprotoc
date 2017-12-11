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
    "fmt"
    // "github.com/golang/protobuf/proto"
    "github.com/golang/glog"
)

type newTLObjectFunc func() TLObject

var registers2 = map[int32]newTLObjectFunc{
    int32(TLConstructor_CRC32_message2) : func() (TLObject) { return &TLMessage2{} },
    int32(TLConstructor_CRC32_msg_container) : func() (TLObject) { return &TLMsgContainer{} },
    int32(TLConstructor_CRC32_msg_copy) : func() (TLObject) { return &TLMsgCopy{} },
    int32(TLConstructor_CRC32_gzip_packed) : func() (TLObject) { return &TLGzipPacked{} },
    int32(TLConstructor_CRC32_rpc_result) : func() (TLObject) { return &TLRpcResult{} },
{{ range .CRC32List }}    int32(TLConstructor_CRC32_{{.Name}}):  func() TLObject { return NewTL{{.Type}}() },
{{end}}
}

func NewTLObjectByClassID(classId int32) TLObject {
m, ok := registers2[classId]
if !ok {
return nil
}
return m()
}

//////////////////////////////////////////////////////////////////////////////////////////

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
        glog.Error("Constructor error: ",  m.GetConstructor())
        return nil
    }
}

func (m *{{.Name}}) Decode(dbuf *DecodeBuf) error {
    m.Constructor = TLConstructor(dbuf.Int())
    switch m.Constructor {
{{ range $i, $v := .SubMessageList }}    case TLConstructor_CRC32_{{.Predicate}}:
        m2 := &TL{{.Name}}{&{{.ResType}}_Data{}}
        m2.Decode(dbuf)
        m.Data2 = m2.Data2
{{end}}
    default:
        return fmt.Errorf("Invalid constructorId: %d", int32(m.Constructor))
    }
    return dbuf.err
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
        Constructor: TLConstructor_CRC32_{{.Predicate}},
        Data2: m.Data2,
    }
}

{{range $i2, $v2 := .ParamList}}
func (m *TL{{$v.Name}}) Set{{.Name}}(v {{.Type}}) { m.Data2.{{.Name2}} = v }
func (m *TL{{$v.Name}}) Get{{.Name}}() {{.Type}} { return m.Data2.{{.Name2}} }
{{end}}

func NewTL{{.Name}}() * TL{{.Name}} {
    return &TL{{.Name}}{ Data2: &{{.ResType}}_Data{} }
}

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

{{range .RequestList}}
func NewTL{{.Name}}() * TL{{.Name}} {
    return &TL{{.Name}}{}
}

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