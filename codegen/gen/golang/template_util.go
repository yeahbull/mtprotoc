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

import (
	"fmt"
	mtproto_parser "github.com/nebulaim/mtprotoc/codegen/parser"
	"strconv"
	"strings"
	// "github.com/golang/glog"
)

type TplCRC32Info struct {
	Name string
	Type string
}

type TplCRC32 struct {
	CRC32List []TplCRC32Info
}

func makeTplCRC32List(schemas *mtproto_parser.Schemas) (tplCRC32List []TplCRC32Info) {
	for _, c := range schemas.ConstructorList {
		// glog.Info(c)
		tplCRC32List = append(tplCRC32List, TplCRC32Info{
			Name: strings.Replace(c.Predicate, ".", "_", -1),
			Type: toProtoGoName(strings.Replace(c.Predicate, ".", "_", -1)),
		})
	}
	for _, f := range schemas.FunctionList {
		// glog.Info(f)
		tplCRC32List = append(tplCRC32List, TplCRC32Info{
			Name: strings.Replace(f.Method, ".", "_", -1),
			Type: toProtoGoName(strings.Replace(f.Method, ".", "_", -1)),
		})
	}
	return
}

// types
type TplCodecDataList struct {
	BaseTypeList []TplBaseTypeData
	RequestList []TplMessageData
	CRC32List []TplCRC32Info
}

// functions
type TplFunctionDataList struct {
	RequestList []TplMessageData
	// Vector是非标准proto消息，故要自动生成一个Vector包装proto消息
	// 注意去重
	VectorResList []TplParam
	// RpcList
	// service RPCAuth {
	//	rpc auth_checkPhone(TL_auth_checkPhone) returns (auth_CheckedPhone) {}
	// }
	ServiceList []TplBaseTypeData
}

type TplServiceData struct {
	RPCName string
	MethodName string
	RequestName string
	ResponseName string
	Line string
}
// functions
type TplServiceDataList struct {
	ServiceList map[string][]TplServiceData
}

// 参数列表
type TplParam struct {
	Type  string
	Name  string
	Name2 string
	Index int
}

// 对应生成proto消息
type TplMessageData struct {
	Predicate string
	Name      string
	// 碰撞的字段名，特殊处理
	ParamList []TplParam
	ResType   string
	Line      string

	EncodeCodeList []string
	DecodeCodeList []string
}

// Base类型
type TplBaseTypeData struct {
	Name      string
	ParamList []TplParam
	// Line string
	ResType string
	// 所有的子类
	SubMessageList []TplMessageData
}

func toMessageName(n string) string {
	return strings.Replace(n, ".", "_", -1)
}

var (
	ignoreRpcList = []string{"invokeAfterMsg", "invokeAfterMsgs", "initConnection", "invokeWithLayer", "invokeWithoutUpdates"}
)

func toGolangType2(t mtproto_parser.Type) (n string) {
	n = toGolangType(t)
	if n != "" {
		if n[0] == '*' {
			return n[1:]
		}
	}
	return
}

func toGolangType(t mtproto_parser.Type) (n string) {
	switch t.(type) {
	case mtproto_parser.BoolType:
		n = "bool"
	case mtproto_parser.IntType:
		n = "int32"
	case mtproto_parser.LongType:
		n = "int64"
	case mtproto_parser.DoubleType:
		n = "float64"
	case mtproto_parser.Int128Type:
		n = "[]byte"
	case mtproto_parser.Int256Type:
		n = "[]byte"
	case mtproto_parser.StringType:
		n = "string"
	case mtproto_parser.BytesType:
		n = "[]byte"
	case mtproto_parser.FlagsType:
		n = ""
	case mtproto_parser.SubFlagsType:
		t2, _ := t.(mtproto_parser.SubFlagsType)
		// glog.Info(t2)
		n = toGolangType(t2.Type)
	case mtproto_parser.BuiltInVectorType:
		t2, _ := t.(mtproto_parser.BuiltInVectorType)
		n = fmt.Sprintf("[]%s", toGolangType(t2.Type))
	case mtproto_parser.TVectorType:
		t2, _ := t.(mtproto_parser.TVectorType)
		n = fmt.Sprintf("[]%s", toGolangType(t2.Type))
	case mtproto_parser.CustomType:
		n = "*" + toProtoGoName(strings.Replace(t.Name(), ".", "_", -1))
		//if n == "Bool" {
		//	n = "bool"
		//}
	case mtproto_parser.Constructor:
		t2, _ := t.(mtproto_parser.Constructor)
		n = "*TL" + toProtoGoName(strings.Replace(t2.Predicate, ".", "_", -1))
	case mtproto_parser.TemplateType:
		n = "[]byte"
	default:
		panic(fmt.Errorf("Bad type: %v", t))
	}

	return
}

func toProtoGoName(n string) string {
	if len(n) == 0 {
		return n
	}
	var n2 = n

	// TODO(@benqi): add ruler table
	if n == "udp_p2p" {
		n2 = "UdpP2P"
	}
	ss := strings.Split(n2, "_")
	for i, v := range ss {
		if i != 0 && IsUpper(v[0]) {
			ss[i] = "_" + v
		} else {
			// glog.Info(v)
			ss[i] = string(ToUpper(v[0])) + v[1:]
		}
	}
	return strings.Join(ss, "")
}

func findByParamList(params []TplParam, p2 TplParam) int {
	for i, p := range params {
		if p.Name == p2.Name && p.Type == p2.Type {
			return i
		}
	}
	return -1
}

func makeEncodeFlags(params []mtproto_parser.Param) (s string) {
	s = fmt.Sprintf("// flags\n")
	s += fmt.Sprintf("    var flags uint32 = 0\n")
	for _, p := range params {
		switch p.Type.(type) {
		case mtproto_parser.SubFlagsType:
			t2, _ := p.Type.(mtproto_parser.SubFlagsType)
			// TODO(@benqi): other type
			switch t2.Type.(type) {
			case mtproto_parser.BoolType:
				s += fmt.Sprintf("    if m.Get%s() == true { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			case mtproto_parser.IntType, mtproto_parser.LongType:
				s += fmt.Sprintf("    if m.Get%s() != 0 { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			case mtproto_parser.StringType:
				s += fmt.Sprintf("    if m.Get%s() != \"\" { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			default:
				s += fmt.Sprintf("    if m.Get%s() != nil { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			}
		}
	}
	s += fmt.Sprintf("	x.UInt(flags)\n")
	return
}

func makeCodecCodeByNameType(n string, t mtproto_parser.Type, idx int) (e string, d string) {
	switch t.(type) {
	case mtproto_parser.BoolType:
		// e = fmt.Sprintf("// x.Bool()")
		d = fmt.Sprintf("(true)", toProtoGoName(n))
	case mtproto_parser.IntType:
		e = fmt.Sprintf("x.Int(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Int())", toProtoGoName(n))
	case mtproto_parser.LongType:
		e = fmt.Sprintf("x.Long(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Long())", toProtoGoName(n))
	case mtproto_parser.DoubleType:
		e = fmt.Sprintf("x.Double(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Double())", toProtoGoName(n))
	case mtproto_parser.Int128Type:
		e = fmt.Sprintf("x.Bytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Bytes(16))", toProtoGoName(n))
	case mtproto_parser.Int256Type:
		e = fmt.Sprintf("x.Bytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Bytes(32))", toProtoGoName(n))
	case mtproto_parser.BytesType:
		e = fmt.Sprintf("x.StringBytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.StringBytes())", toProtoGoName(n))
	case mtproto_parser.StringType:
		e = fmt.Sprintf("x.String(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.String())", toProtoGoName(n))
	case mtproto_parser.Constructor:
		e = fmt.Sprintf("x.Bytes(m.Get%s().Encode())", toProtoGoName(n))
		d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.Set%s(m%d)", idx, toProtoGoName(n), idx, toProtoGoName(n), idx)
	case mtproto_parser.CustomType:
		e = fmt.Sprintf("x.Bytes(m.Get%s().Encode())", toProtoGoName(n))
		d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.Set%s(m%d)", idx, toProtoGoName(n), idx, toProtoGoName(n), idx)
	}

	return
}

// Encode
// Decode
func makeCodecCode(params []mtproto_parser.Param, n string, t mtproto_parser.Type, idx int) (e string, d string) {
	switch t.(type) {
	case mtproto_parser.BoolType:
		// e = fmt.Sprintf("// x.Bool()")
		d = fmt.Sprintf("m.Set%s(true)", toProtoGoName(n))
	case mtproto_parser.IntType:
		e = fmt.Sprintf("x.Int(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Int())", toProtoGoName(n))
	case mtproto_parser.LongType:
		e = fmt.Sprintf("x.Long(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Long())", toProtoGoName(n))
	case mtproto_parser.DoubleType:
		e = fmt.Sprintf("x.Double(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Double())", toProtoGoName(n))
	case mtproto_parser.Int128Type:
		e = fmt.Sprintf("x.Bytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Bytes(16))", toProtoGoName(n))
	case mtproto_parser.Int256Type:
		e = fmt.Sprintf("x.Bytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.Bytes(32))", toProtoGoName(n))
	case mtproto_parser.StringType:
		e = fmt.Sprintf("x.String(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.String())", toProtoGoName(n))
	case mtproto_parser.BytesType:
		e = fmt.Sprintf("x.StringBytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(dbuf.StringBytes())", toProtoGoName(n))
	case mtproto_parser.FlagsType:
		e = makeEncodeFlags(params)
		d  = fmt.Sprintf("flags := dbuf.UInt()\n")
		d += fmt.Sprintf("_ = flags ")
		// glog.Info(e, " ==> ", d)
	case mtproto_parser.SubFlagsType:
		t2, _ := t.(mtproto_parser.SubFlagsType)
		// TODO(@benqi): other type
		switch t2.Type.(type) {
		case mtproto_parser.BoolType:
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 { m.Set%s(true) }", t2.Mask, toProtoGoName(n))
		case mtproto_parser.IntType,
			mtproto_parser.LongType:
			e2, d2 := makeCodecCode(params, n, t2.Type, idx)
			e = fmt.Sprintf("if m.Get%s() != 0 { %s }", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 { %s }", t2.Mask, d2)
		case mtproto_parser.StringType:
			e2, d2 := makeCodecCode(params, n, t2.Type, idx)
			e = fmt.Sprintf("if m.Get%s() != \"\" { %s }", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 { %s }", t2.Mask, d2)
		case mtproto_parser.Constructor,
			mtproto_parser.CustomType,
			mtproto_parser.BytesType:
			e2, d2 := makeCodecCode(params, n, t2.Type, idx)
			e = fmt.Sprintf("if m.Get%s() != nil {\n %s \n}", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 {\n %s \n}", t2.Mask, d2)
		case mtproto_parser.BuiltInVectorType:
			t3, _ := t.(mtproto_parser.BuiltInVectorType)
			e2, d2 := makeCodecCode(params, n, t3.Type, idx)
			e = fmt.Sprintf("if m.Get%s() != nil {\n %s \n}", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 {\n %s \n}", t2.Mask, d2)
		case mtproto_parser.TVectorType:
			t3, _ := t.(mtproto_parser.TVectorType)
			e2, d2 := makeCodecCode(params, n, t3.Type, idx)
			e = fmt.Sprintf("if m.Get%s() != nil {\n %s \n}", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 {\n %s \n}", t2.Mask, d2)
		default:
		}

	case mtproto_parser.BuiltInVectorType:
		t2, _ := t.(mtproto_parser.BuiltInVectorType)
		n2 := toProtoGoName(n)
		switch t2.Type.(type) {
		case mtproto_parser.IntType:
			e = fmt.Sprintf("x.VectorInt(m.Get%s())\n", n2)
			d = fmt.Sprintf("m.Set%s(dbuf.VectorInt())", n2)
		case mtproto_parser.LongType:
			e = fmt.Sprintf("x.VectorLong(m.Get%s())\n", n2)
			d = fmt.Sprintf("m.Set%s(dbuf.VectorLong())", n2)
		case mtproto_parser.StringType:
			e = fmt.Sprintf("x.VectorString(m.Get%s())\n", n2)
			d = fmt.Sprintf("m.Set%s(dbuf.VectorString())", n2)
		case mtproto_parser.CustomType, mtproto_parser.Constructor:
			// e  = fmt.Sprintf("x.Int(int32(TLConstructor_CRC32_vector))\n")
			e += fmt.Sprintf("x.Int(int32(m.Get%s()))\n", n2)
			e += fmt.Sprintf("for _, v := range m.Get%s() {\n", n2)
			e += fmt.Sprintf("  x.buf = append(x.buf, (*v).Encode()...)\n")
			e += fmt.Sprintf("}")

			// d += fmt.Sprintf("c%d := dbuf.Int()\n", idx)
			// d += fmt.Sprintf("if c%d != TLConstructor_CRC32_vector {\n", idx)
			// d += fmt.Sprintf("	dbuf.err = fmt.Sprintf(\"Invalid CRC32_vector, c%d := %d\", %d, c%d)\n", idx, idx, idx, idx)
			// d += fmt.Sprintf("	return dbuf.err\n")
			// d += fmt.Sprintf("}\n")
			d += fmt.Sprintf("l%d := dbuf.Int()\n", idx)
			d += fmt.Sprintf("v%d := make([]*%s, l%d)\n", idx, toGolangType2(t2.Type), idx)
			d += fmt.Sprintf("for i := int32(0); i < l%d; i++ {\n", idx)
			d += fmt.Sprintf("	v%d[i] = &%s{}\n", idx, toGolangType2(t2.Type))
			d += fmt.Sprintf("	v%d[i].Decode(dbuf)\n", idx)
			d += fmt.Sprintf("}\n")
			d += fmt.Sprintf("m.Set%s(v%d)\n", n2, idx)

			// d += fmt.Sprintf("m.Set%s(make([]*%s, ln))", n2, toGolangType(t2.Type))
			// d += fmt.Sprintf("for i < ln {\n m.Get%s()[i] = &%s\n (*m.Data2.%s[i]).Decode(dbuf)\n}", n2, toGolangType(t2.Type)) // ln := len(m.Data2.%s)\n", n2)
		}
	case mtproto_parser.TVectorType:
		t2, _ := t.(mtproto_parser.TVectorType)
		n2 := toProtoGoName(n)
		switch t2.Type.(type) {
		case mtproto_parser.IntType:
			e = fmt.Sprintf("x.VectorInt(m.Get%s())\n", n2)
			d = fmt.Sprintf("m.Set%s(dbuf.VectorInt())", n2)
		case mtproto_parser.LongType:
			e = fmt.Sprintf("x.VectorLong(m.Get%s())\n", n2)
			d = fmt.Sprintf("m.Set%s(dbuf.VectorLong())", n2)
		case mtproto_parser.StringType:
			e = fmt.Sprintf("x.VectorString(m.Get%s())\n", n2)
			d = fmt.Sprintf("m.Set%s(dbuf.VectorString())", n2)
		case mtproto_parser.CustomType, mtproto_parser.Constructor:
			e = "x.Int(int32(TLConstructor_CRC32_vector))\n"
			e += fmt.Sprintf("x.Int(int32(len(m.Get%s())))\n", n2)
			e += fmt.Sprintf("for _, v := range m.Get%s() {\n", n2)
			e += fmt.Sprintf("  x.buf = append(x.buf, (*v).Encode()...)\n")
			e += fmt.Sprintf("}")

			// d = "dbuf.Int()\n"
			// d += fmt.Sprintf("ln := dbuf.Int()\n", n2)
			// d += fmt.Sprintf("m.Data2.%s = make([]%s, ln)", n2, toGolangType(t2.Type))
			// d += fmt.Sprintf("for i < ln {\n m.Data2.%s[i] = &%s\n (*m.Data2.%s[i]).Decode(dbuf)\n}", n2, toGolangType2(t2.Type)) // ln := len(m.Data2.%s)\n", n2)
			d += fmt.Sprintf("c%d := dbuf.Int()\n", idx)
			d += fmt.Sprintf("if c%d != int32(TLConstructor_CRC32_vector) {\n", idx)
			d += fmt.Sprintf("	dbuf.err = fmt.Errorf(\"Invalid CRC32_vector, c%%d: %%d\", %d, c%d)\n", idx, idx)
			d += fmt.Sprintf("	return dbuf.err\n")
			d += fmt.Sprintf("}\n")
			d += fmt.Sprintf("l%d := dbuf.Int()\n", idx)
			d += fmt.Sprintf("v%d := make([]*%s, l%d)\n", idx, toGolangType2(t2.Type), idx)
			d += fmt.Sprintf("for i := int32(0); i < l%d; i++ {\n", idx)
			d += fmt.Sprintf("	v%d[i] = &%s{}\n", idx, toGolangType2(t2.Type))
			d += fmt.Sprintf("	v%d[i].Decode(dbuf)\n", idx)
			d += fmt.Sprintf("}\n")
			d += fmt.Sprintf("m.Set%s(v%d)\n", n2, idx)
		}
	case mtproto_parser.CustomType:
		// t2, _ := t.(mtproto_parser.CustomType)
		e = fmt.Sprintf("x.Int(int32(TLConstructor_CRC32_%s))\n", toMessageName(t.Name()))
		e = fmt.Sprintf("x.Bytes(m.Get%s().Encode())", toProtoGoName(n))

		d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.Set%s(m%d)", idx, toGolangType2(t), idx, toProtoGoName(n), idx)

		//d += fmt.Sprintf("c%d := dbuf.Int()\n", idx)
		//d += fmt.Sprintf("if c%d != int32(TLConstructor_CRC32_%s) {\n", idx, toMessageName(t2.Name()))
		//d += fmt.Sprintf("	dbuf.err = fmt.Sprintf(\"Invalid CRC32_%s, c%%d: %%d\", %d, c%d)\n", toMessageName(t2.Name()), idx, idx)
		//d += fmt.Sprintf("}\n")
	case mtproto_parser.Constructor:
		t2, _ := t.(mtproto_parser.Constructor)
		e = fmt.Sprintf("x.Int(int32(TLConstructor_CRC32_%s))\n", toMessageName(t2.Name()))
		e = fmt.Sprintf("x.Bytes(m.Get%s().Encode())", toProtoGoName(n))

		d += fmt.Sprintf("c%d := dbuf.Int()\n", idx)
		d += fmt.Sprintf("if c%d != int32(TLConstructor_CRC32_%s) {\n", idx, toMessageName(t2.Name()))
		d += fmt.Sprintf("	dbuf.err = fmt.Sprintf(\"Invalid CRC32_%s, c%%d: %%d\", %d, c%d)\n", toMessageName(t2.Name()), idx, idx)
		d += fmt.Sprintf("}\n")

		// d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.Set%s(m%d)", idx, toGolangType2(t), idx, toProtoGoName(n), idx)
	case mtproto_parser.TemplateType:
		// n = "[]byte"
		e = fmt.Sprintf("x.Bytes(m.Get%s())", toProtoGoName(n))
		d = fmt.Sprintf("m.Set%s(x.Bytes())", toProtoGoName(n))
	default:
	}
	return
}

func checkByStringList(strList []string, s string) bool {
	for _, p := range strList {
		if p == s {
			return true
		}
	}
	return false
}

func makeBaseTypeListTpl(schemas *mtproto_parser.Schemas) (types []TplBaseTypeData) {
	baseTypeMap := make(map[string]*TplBaseTypeData)

	for _, c := range schemas.ConstructorList {
		baseTypeName := toProtoGoName(toMessageName(c.BaseType.Name()))
		// baseMessage := &TplBaseMessage{Name: baseTypeName}
		baseType, ok := baseTypeMap[baseTypeName]
		if !ok {
			baseType = &TplBaseTypeData{
				Name: baseTypeName,
			}
			baseTypeMap[baseTypeName] = baseType
		}

		messageData := TplMessageData{
			Predicate: toMessageName(c.Predicate),
			Name:      toProtoGoName(toMessageName(c.Predicate)),
			Line:      c.Line,
			ResType:   baseTypeName,
			// toMessageName(c.BaseType.Name()),
		}
		//glog.Info(c.Line)
		//
		for idx, p := range c.ParamList {
			param := TplParam{}
			param.Name = toProtoGoName(p.Name)
			param.Type = toGolangType(p.Type)

			// flags
			e, d := makeCodecCode(c.ParamList, p.Name, p.Type, idx+1)
			messageData.EncodeCodeList = append(messageData.EncodeCodeList, e)
			messageData.DecodeCodeList = append(messageData.DecodeCodeList, d)

			if param.Type == "" {
				continue
			}

			if idx := findByParamList(baseType.ParamList, param); idx == -1 {
				param.Index = len(baseType.ParamList)
				baseType.ParamList = append(baseType.ParamList, param)
			} else {
				param.Index = idx
			}
			messageData.ParamList = append(messageData.ParamList, param)
		}
		baseType.SubMessageList = append(baseType.SubMessageList, messageData)
		// glog.Info(baseType)
		// messages.MessageList = append(messages.MessageList, message)
	}

	// types = &TplTypesDataList{}
	for _, v := range baseTypeMap {
		// param.Name有重复，要修改Name
		names := make(map[string][]int)

		for i, p := range v.ParamList {
			v.ParamList[i].Index = i + 1
			if _, ok := names[p.Name]; !ok {
				names[p.Name] = []int{i}
			} else {
				names[p.Name] = append(names[p.Name], i)
			}
		}

		for _, v2 := range names {
			if len(v2) > 1 {
				for _, idx := range v2 {
					v.ParamList[idx].Name = v.ParamList[idx].Name + "_" + strconv.Itoa(v.ParamList[idx].Index)
				}
			}
		}

		for i3, v3 := range v.SubMessageList {
			for i4, v4 := range v3.ParamList {
				// glog.Info(i4, " ==> ", v4, ", ", v3.Line)
				// glog.Info(v.ParamList)
				v4.Name2 = v.ParamList[v4.Index].Name
				v3.ParamList[i4] = v4
			}
			v.SubMessageList[i3] = v3
		}
		// glog.Info(v)
		types = append(types, *v)
	}

	// glog.Info(types)

	return
}


func makeEncodeFlags2(params []mtproto_parser.Param) (s string) {
	s = fmt.Sprintf("// flags\n")
	s += fmt.Sprintf("    var flags uint32 = 0\n")
	for _, p := range params {
		switch p.Type.(type) {
		case mtproto_parser.SubFlagsType:
			t2, _ := p.Type.(mtproto_parser.SubFlagsType)
			// TODO(@benqi): other type
			switch t2.Type.(type) {
			case mtproto_parser.BoolType:
				s += fmt.Sprintf("    if m.%s == true { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			case mtproto_parser.IntType, mtproto_parser.LongType:
				s += fmt.Sprintf("    if m.%s != 0 { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			case mtproto_parser.StringType:
				s += fmt.Sprintf("    if m.%s != \"\" { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			default:
				s += fmt.Sprintf("    if m.%s != nil { flags |= 1 << %s }\n", toProtoGoName(p.Name), t2.Mask)
			}
		}
	}
	s += fmt.Sprintf("	x.UInt(flags)\n")
	return
}

func makeCodecCodeByNameType2(n string, t mtproto_parser.Type, idx int) (e string, d string) {
	switch t.(type) {
	case mtproto_parser.BoolType:
		// e = fmt.Sprintf("// x.Bool()")
		d = fmt.Sprintf("(true)", toProtoGoName(n))
	case mtproto_parser.IntType:
		e = fmt.Sprintf("x.Int(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Int()", toProtoGoName(n))
	case mtproto_parser.LongType:
		e = fmt.Sprintf("x.Long(m.%s", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Long()", toProtoGoName(n))
	case mtproto_parser.DoubleType:
		e = fmt.Sprintf("x.Double(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Double()", toProtoGoName(n))
	case mtproto_parser.Int128Type, mtproto_parser.Int256Type, mtproto_parser.BytesType:
		e = fmt.Sprintf("x.Bytes(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Bytes()", toProtoGoName(n))
	case mtproto_parser.StringType:
		e = fmt.Sprintf("x.String(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.String()", toProtoGoName(n))
	case mtproto_parser.Constructor:
		e = fmt.Sprintf("x.Bytes(m.%s.Encode())", toProtoGoName(n))
		d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.%s = m%d", idx, toProtoGoName(n), idx, toProtoGoName(n), idx)
	case mtproto_parser.CustomType:
		e = fmt.Sprintf("x.Bytes(m.%s.Encode())", toProtoGoName(n))
		d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.%s = m%d", idx, toProtoGoName(n), idx, toProtoGoName(n), idx)
	}

	return
}

// Encode
// Decode
func makeCodecCode2(params []mtproto_parser.Param, n string, t mtproto_parser.Type, idx int) (e string, d string) {
	switch t.(type) {
	case mtproto_parser.BoolType:
		// e = fmt.Sprintf("// x.Bool()")
		d = fmt.Sprintf("m.%s = true", toProtoGoName(n))
	case mtproto_parser.IntType:
		e = fmt.Sprintf("x.Int(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Int()", toProtoGoName(n))
	case mtproto_parser.LongType:
		e = fmt.Sprintf("x.Long(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Long()", toProtoGoName(n))
	case mtproto_parser.DoubleType:
		e = fmt.Sprintf("x.Double(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Double()", toProtoGoName(n))
	case mtproto_parser.Int128Type:
		e = fmt.Sprintf("x.Bytes(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Bytes(16)", toProtoGoName(n))
	case mtproto_parser.Int256Type:
		e = fmt.Sprintf("x.Bytes(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.Bytes(32)", toProtoGoName(n))
	case mtproto_parser.StringType:
		e = fmt.Sprintf("x.String(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.String()", toProtoGoName(n))
	case mtproto_parser.BytesType:
		e = fmt.Sprintf("x.StringBytes(m.%s)", toProtoGoName(n))
		d = fmt.Sprintf("m.%s = dbuf.StringBytes()", toProtoGoName(n))
	case mtproto_parser.FlagsType:
		e = makeEncodeFlags2(params)
		d  = fmt.Sprintf("flags := dbuf.UInt()\n")
		d += fmt.Sprintf("_ = flags")
		// glog.Info(e, " ==> ", d)
	case mtproto_parser.SubFlagsType:
		t2, _ := t.(mtproto_parser.SubFlagsType)
		// TODO(@benqi): other type
		switch t2.Type.(type) {
		case mtproto_parser.BoolType:
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 { m.%s = true }", t2.Mask, toProtoGoName(n))
		case mtproto_parser.IntType,
		mtproto_parser.LongType:
			e2, d2 := makeCodecCode2(params, n, t2.Type, idx)
			e = fmt.Sprintf("if m.%s != 0 { %s }", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 { %s }", t2.Mask, d2)
		case mtproto_parser.StringType:
			e2, d2 := makeCodecCode2(params, n, t2.Type, idx)
			e = fmt.Sprintf("if m.%s != \"\" { %s }", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 { %s }", t2.Mask, d2)
		case mtproto_parser.Constructor,
		mtproto_parser.CustomType,
		mtproto_parser.BytesType:
			e2, d2 := makeCodecCode2(params, n, t2.Type, idx)
			e = fmt.Sprintf("if m.%s != nil {\n %s \n}", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 {\n %s \n}", t2.Mask, d2)
		case mtproto_parser.BuiltInVectorType, mtproto_parser.TVectorType:
			t3, _ := t.(mtproto_parser.BuiltInVectorType)
			e2, d2 := makeCodecCode2(params, n, t3.Type, idx)
			e = fmt.Sprintf("if m.%s != nil {\n %s \n}", toProtoGoName(n), e2)
			d = fmt.Sprintf("if (flags & (1 << %s)) != 0 {\n %s \n}", t2.Mask, d2)
		default:
		}

	case mtproto_parser.BuiltInVectorType:
		t2, _ := t.(mtproto_parser.BuiltInVectorType)
		n2 := toProtoGoName(n)
		switch t2.Type.(type) {
		case mtproto_parser.IntType:
			e = fmt.Sprintf("x.VectorInt(m.%s)\n", n2)
			d = fmt.Sprintf("m.%s = dbuf.VectorInt()", n2)
		case mtproto_parser.LongType:
			e = fmt.Sprintf("x.VectorLong(m.%s)\n", n2)
			d = fmt.Sprintf("m.%s = dbuf.VectorLong()", n2)
		case mtproto_parser.StringType:
			e = fmt.Sprintf("x.VectorString(m.%s)\n", n2)
			d = fmt.Sprintf("m.%s = dbuf.VectorString()", n2)
		case mtproto_parser.CustomType, mtproto_parser.Constructor:
			e += fmt.Sprintf("x.Int(int32(m.%s))\n", n2)
			e += fmt.Sprintf("for _, v := range m.%s {\n", n2)
			e += fmt.Sprintf("  x.buf = append(x.buf, (*v).Encode()...)\n")
			e += fmt.Sprintf("}\n")
			d += fmt.Sprintf("l%d := dbuf.Int()\n", idx)
			d += fmt.Sprintf("m.%s = make([]*%s, l%d)\n", n2, toGolangType2(t2.Type), idx)
			d += fmt.Sprintf("for i := int32(0); i < l%d; i++ {\n m.%s[i] = &%s\n (*m.%s[i]).Decode(dbuf)\n}", idx, n2, toGolangType2(t2.Type)) // ln := len(m.Data2.%s)\n", n2)
		}
	case mtproto_parser.TVectorType:
		t2, _ := t.(mtproto_parser.TVectorType)
		n2 := toProtoGoName(n)
		switch t2.Type.(type) {
		case mtproto_parser.IntType:
			e = fmt.Sprintf("x.VectorInt(m.%s)\n", n2)
			d = fmt.Sprintf("m.%s = dbuf.VectorInt()", n2)
		case mtproto_parser.LongType:
			e = fmt.Sprintf("x.VectorLong(m.%s)\n", n2)
			d = fmt.Sprintf("m.%s = dbuf.VectorLong()", n2)
		case mtproto_parser.StringType:
			e = fmt.Sprintf("x.VectorString(m.%s)\n", n2)
			d = fmt.Sprintf("m.%s = dbuf.VectorString()", n2)
		case mtproto_parser.CustomType, mtproto_parser.Constructor:
			e = "x.Int(int32(TLConstructor_CRC32_vector))\n"
			e += fmt.Sprintf("x.Int(int32(len(m.%s)))\n", n2)
			e += fmt.Sprintf("for _, v := range m.%s {\n", n2)
			e += fmt.Sprintf("  x.buf = append(x.buf, (*v).Encode()...)\n")
			e += fmt.Sprintf("}")

			d = "dbuf.Int()  // TODO(@benqi): Check crc32 invalid\n"
			d += fmt.Sprintf("l%d := dbuf.Int()\n", idx)
			d += fmt.Sprintf("m.%s = make([]%s, l%d)\n", n2, toGolangType(t2.Type), idx)
			d += fmt.Sprintf("for i := int32(0); i < l%d; i++ {\n m.%s[i] = &%s{}\n (*m.%s[i]).Decode(dbuf)\n}", idx, n2, toGolangType2(t2.Type), n2) // ln := len(m.Data2.%s)\n", n2)
		}
	case mtproto_parser.CustomType, mtproto_parser.Constructor:
		// e = fmt.Sprintf("x.Int(int32(TLConstructor_CRC32_%s))\n", toMessageName(t.Name()))
		e += fmt.Sprintf("x.Bytes(m.%s.Encode())", toProtoGoName(n))
		d = fmt.Sprintf("m%d := &%s{}\n    m%d.Decode(dbuf)\n    m.%s = m%d", idx, toGolangType2(t), idx, toProtoGoName(n), idx)
	case mtproto_parser.TemplateType:
		// n = "[]byte"
		e = fmt.Sprintf("x.Bytes(m.%s)", toProtoGoName(n))
		// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好
		//d := o2 := dbuf.Object()
		//m.Query = o2.Encode()
		d  = "// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好\n"
		d += fmt.Sprintf("o%d := dbuf.Object()\n", idx)
		d += fmt.Sprintf("m.%s = o%d.Encode()", toProtoGoName(n), idx)
	default:
	}
	return
}

func makeFunctionDataListTpl(schemas *mtproto_parser.Schemas) (funcs []TplMessageData) {
	// messages := make(map[string]*TplBaseMessage)
	// funcs = &TplFunctionDataList{}

	// serviceTypeMap := make(map[string]*TplBaseTypeData)

	// RequestList
	for _, c := range schemas.FunctionList {
		//rpcName := strings.Split(c.Method, ".")[0]
		//// baseMessage := &TplBaseMessage{Name: baseTypeName}
		//service, ok := serviceTypeMap[rpcName]
		//if !ok {
		//	service = &TplBaseTypeData{
		//		Name: rpcName,
		//	}
		//	serviceTypeMap[rpcName] = service
		//}

		//message := TplMessageData{
		//	Name: strings.Replace(c.Method, ".", "_", -1),
		//	Line: c.Line,
		//}

		messageData := TplMessageData{
			Predicate: toMessageName(c.Method),
			Name:      toProtoGoName(toMessageName(c.Method)),
			Line:      c.Line,
			// ResType:   baseTypeName,
			// toMessageName(c.BaseType.Name()),
		}

		//serviceMessage := TplMessageData{
		//	Name: strings.Replace(c.Method, ".", "_", -1),
		//	Line: c.Line,
		//}

		// glog.Info(c.Line)
		for i, p := range c.ParamList {
			//param := TplParam{}
			//param.Name = p.Name
			//param.Index = i + 1
			//param.Type = toGolangType(p.Type)
			//if param.Type == "" {
			//	continue
			//}
			//message.ParamList = append(message.ParamList, param)


			param := TplParam{}
			param.Name = toProtoGoName(p.Name)
			param.Type = toGolangType(p.Type)
			param.Index = i + 1

			// flags
			e, d := makeCodecCode2(c.ParamList, p.Name, p.Type, i+1)
			messageData.EncodeCodeList = append(messageData.EncodeCodeList, e)
			messageData.DecodeCodeList = append(messageData.DecodeCodeList, d)

			if param.Type == "" {
				continue
			}

			//if idx := findByParamList(baseType.ParamList, param); idx == -1 {
			//	param.Index = len(baseType.ParamList)
			//	baseType.ParamList = append(baseType.ParamList, param)
			//} else {
			//	param.Index = idx
			//}
			messageData.ParamList = append(messageData.ParamList, param)

		}
		funcs = append(funcs, messageData)

		//switch c.ResType.(type) {
		//case mtproto_parser.TVectorType:
		//	vectorType := TplParam{
		//		Type: toGolangType(c.ResType),
		//		Name: toMessageName(c.ResType.(mtproto_parser.TVectorType).Type.Name()),
		//	}
		//	if -1 == findByParamList(funcs.VectorResList, vectorType) {
		//		funcs.VectorResList = append(funcs.VectorResList, vectorType)
		//	}
		//	serviceMessage.ResType = "Vector_" + vectorType.Name
		//case mtproto_parser.BuiltInVectorType:
		//	vectorType := TplParam{
		//		Type: toGolangType(c.ResType),
		//		Name: toMessageName(c.ResType.(mtproto_parser.TVectorType).Type.Name()),
		//	}
		//	if -1 == findByParamList(funcs.VectorResList, vectorType) {
		//		funcs.VectorResList = append(funcs.VectorResList, vectorType)
		//	}
		//	serviceMessage.ResType = "Vector_" + vectorType.Name
		//default:
		//	serviceMessage.ResType = toMessageName(c.ResType.Name())
		//}
		//service.SubMessageList = append(service.SubMessageList, serviceMessage)
	}

	//for _, v := range serviceTypeMap {
	//	if checkByStringList(ignoreRpcList, v.Name) {
	//		continue
	//	}
	//	funcs.ServiceList = append(funcs.ServiceList, *v)
	//	// glog.Info(v)
	//}

	return
}

/*
type TplServiceData struct {
	RPCName string
	RequestName string
	ResponseName string
	Line string
}
 */

func makeServiceDataListTpl(schemas *mtproto_parser.Schemas) (serviceList map[string][]TplServiceData) {
	// RequestList
	serviceList = make(map[string][]TplServiceData)

	for _, c := range schemas.FunctionList {
		rpcName := strings.Split(c.Method, ".")[0]
		if checkByStringList(ignoreRpcList, rpcName) {
			continue
		}

		if _, ok := serviceList[rpcName]; !ok {
			serviceList[rpcName] = []TplServiceData{}
		}
		service := &TplServiceData{
			RPCName: toProtoGoName(rpcName),
			MethodName: c.Method,
			RequestName: toProtoGoName(toMessageName(c.Method)),
			Line: c.Line,
			ResponseName: toProtoGoName(toMessageName(c.ResType.Name())),
		}

		switch c.ResType.(type) {
		case mtproto_parser.TVectorType:
		//case mtproto_parser.IntType:
		//	e = fmt.Sprintf("x.VectorInt(m.%s)\n", n2)
		//	d = fmt.Sprintf("m.%s = dbuf.VectorInt()", n2)
		//case mtproto_parser.LongType:
		//	e = fmt.Sprintf("x.VectorLong(m.%s)\n", n2)
		//	d = fmt.Sprintf("m.%s = dbuf.VectorLong()", n2)

			service.ResponseName = "Vector_" + toProtoGoName(toMessageName(c.ResType.(mtproto_parser.TVectorType).Type.Name()))
			// TODO(@benqi): 很土的办法
			if service.ResponseName == "Vector_Int" {
				service.ResponseName = "VectorInt"
			} else if service.ResponseName == "Vector_Long" {
				service.ResponseName = "VectorLong"
			}
		case mtproto_parser.BuiltInVectorType:
			service.ResponseName = "Vector_" + toProtoGoName(toMessageName(c.ResType.(mtproto_parser.BuiltInVectorType).Type.Name()))
			// TODO(@benqi): 很土的办法
			if service.ResponseName == "Vector_Int" {
				service.ResponseName = "VectorInt"
			} else if service.ResponseName == "Vector_Long" {
				service.ResponseName = "VectorLong"
			}
		default:
			service.ResponseName = toProtoGoName(toMessageName(c.ResType.Name()))
		}
		serviceList[rpcName] = append(serviceList[rpcName], *service)
	}

	return
}
