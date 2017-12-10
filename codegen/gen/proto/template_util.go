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

package genproto

import (
	"strings"
	mtproto_parser "github.com/nebulaim/mtprotoc/codegen/parser"
	"strconv"
	"fmt"
)

// types
type TplTypesDataList struct {
	BaseTypeList []TplBaseTypeData
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

// 参数列表
type TplParam struct {
	Type string
	Name string
	Index string
}

// 对应生成proto消息
type TplMessageData struct{
	Name string
	// 碰撞的字段名，特殊处理
	ParamList []TplParam
	ResType string
	Line string
}

// Base类型
type TplBaseTypeData struct {
	Name string
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

	// boolFalse#bc799737 = Bool;
	// boolTrue#997275b5 = Bool;
	//
	// true#3fedd339 = True;
	//
	// vector#1cb5c415 {t:Type} # [ t ] = Vector t;
	//
	// error#c4b9f9bb code:int text:string = Error;
	//
	// null#56730bcc = Null;
	coreTypeList = []string {"boolFalse", "boolTrue", "true", "vector", "error", "null"}
)

func isCoreTypeList(t string) bool {
	for _, n := range coreTypeList {
		if n == t {
			return true
		}
	}
	return false
}

func toProtoType(t mtproto_parser.Type) (n string) {
	switch t.(type) {
	case mtproto_parser.BoolType:
		n = "bool"
	case mtproto_parser.IntType:
		n = "int32"
	case mtproto_parser.LongType:
		n = "int64"
	case mtproto_parser.DoubleType:
		n = "double"
	case mtproto_parser.Int128Type:
		n = "bytes"
	case mtproto_parser.Int256Type:
		n = "bytes"
	case mtproto_parser.StringType:
		n = "string"
	case mtproto_parser.BytesType:
		n = "bytes"
	case mtproto_parser.FlagsType:
		n = ""
	case mtproto_parser.SubFlagsType:
		t2, _ := t.(mtproto_parser.SubFlagsType)
		// glog.Info(t2)
		n = toProtoType(t2.Type)
	case mtproto_parser.BuiltInVectorType:
		t2, _ := t.(mtproto_parser.BuiltInVectorType)
		n = fmt.Sprintf("repeated %s", toProtoType(t2.Type))
	case mtproto_parser.TVectorType:
		t2, _ := t.(mtproto_parser.TVectorType)
		n = fmt.Sprintf("repeated %s", toProtoType(t2.Type))
	case mtproto_parser.CustomType:
		n = strings.Replace(t.Name(), ".", "_", -1)
		//if n == "Bool" {
		//	n = "bool"
		//}
	case mtproto_parser.Constructor:
		t2, _ := t.(mtproto_parser.Constructor)
		n = "TL_" + strings.Replace(t2.Predicate, ".", "_", -1)
	case mtproto_parser.TemplateType:
		n = "bytes"
	default:
		panic(fmt.Errorf("Bad type: %v", t))
	}

	return
}

func checkByParamList(params []TplParam, p2 TplParam) bool {
	for _, p := range params {
		if p.Name == p2.Name && p.Type == p2.Type {
			return true
		}
	}
	return false
}

//func checkNameByParamList(params []TplParam, p2 TplParam) bool {
//	for _, p := range params {
//		if p.Name == p2.Name && p.Type == p2.Type {
//			return true
//		}
//	}
//	return false
//}

func checkByStringList(strList []string, s string) bool {
	for _, p := range strList {
		if p == s {
			return true
		}
	}
	return false
}

func makeBaseTypeListTpl(schemas *mtproto_parser.Schemas, isCoreTypes bool) (types *TplTypesDataList) {
	baseTypeMap := make(map[string]*TplBaseTypeData)

	for _, c := range schemas.ConstructorList {
		if isCoreTypes {
			if !isCoreTypeList(c.Predicate) {
				continue
			}
		} else {
			if isCoreTypeList(c.Predicate) {
				continue
			}
		}

		baseTypeName := toMessageName(c.BaseType.Name())
		// baseMessage := &TplBaseMessage{Name: baseTypeName}
		baseType, ok := baseTypeMap[baseTypeName]
		if !ok {
			baseType = &TplBaseTypeData{
				Name: baseTypeName,
			}
			baseTypeMap[baseTypeName] = baseType
		}

		messageData := TplMessageData{
			Name: toMessageName(c.Predicate),
			Line: c.Line,
			ResType: baseTypeName,
			// toMessageName(c.BaseType.Name()),
		}
		//glog.Info(c.Line)
		//
		for _, p := range c.ParamList {
			param := TplParam{}
			param.Name = p.Name
			// param.Index = conv.I64toa(int64(i+1))
			param.Type = toProtoType(p.Type)
			if param.Type == "" {
				continue
			}
			messageData.ParamList = append(messageData.ParamList, param)

			if !checkByParamList(baseType.ParamList, param) {
				baseType.ParamList = append(baseType.ParamList, param)
			} else {
				// logs.Info(param)
			}
		}
		baseType.SubMessageList = append(baseType.SubMessageList, messageData)
		// glog.Info(baseType)
		// messages.MessageList = append(messages.MessageList, message)
	}

	types = &TplTypesDataList{}
	for _, v := range baseTypeMap {
		// param.Name有重复，要修改Name
		names := make(map[string][]int)

		for i, p := range v.ParamList {
			v.ParamList[i].Index = strconv.Itoa(i+1)
			if _, ok := names[p.Name]; !ok {
				names[p.Name] = []int{i}
			} else {
				names[p.Name] = append(names[p.Name], i)
			}
		}

		for _, v2 := range names {
			if len(v2) > 1 {
				for _, idx := range v2 {
					v.ParamList[idx].Name = v.ParamList[idx].Name + "_" + v.ParamList[idx].Index
				}
			}
		}

		// glog.Info(v)
		types.BaseTypeList = append(types.BaseTypeList, *v)
	}

	// glog.Info(types)

	return
}

func makeFunctionDataListTpl(schemas *mtproto_parser.Schemas) (funcs *TplFunctionDataList) {
	// messages := make(map[string]*TplBaseMessage)
	funcs = &TplFunctionDataList{}

	serviceTypeMap := make(map[string]*TplBaseTypeData)

	// RequestList
	for _, c := range schemas.FunctionList {
		rpcName := strings.Split(c.Method, ".")[0]
		// baseMessage := &TplBaseMessage{Name: baseTypeName}
		service, ok := serviceTypeMap[rpcName]
		if !ok {
			service = &TplBaseTypeData{
				Name: rpcName,
			}
			serviceTypeMap[rpcName] = service
		}

		message := TplMessageData{
			Name: strings.Replace(c.Method, ".", "_", -1),
			Line: c.Line,
		}

		serviceMessage := TplMessageData{
			Name: strings.Replace(c.Method, ".", "_", -1),
			Line: c.Line,
		}

		// glog.Info(c.Line)
		for i, p := range c.ParamList {
			param := TplParam{}
			param.Name = p.Name
			param.Index = strconv.Itoa(i+1)
			param.Type = toProtoType(p.Type)
			if param.Type == "" {
				continue
			}
			message.ParamList = append(message.ParamList, param)
		}
		funcs.RequestList = append(funcs.RequestList, message)

		switch c.ResType.(type) {
		case mtproto_parser.TVectorType:
			vectorType := TplParam{
				Type: toProtoType(c.ResType),
				Name: toMessageName(c.ResType.(mtproto_parser.TVectorType).Type.Name()),
			}
			if !checkByParamList(funcs.VectorResList, vectorType) {
				funcs.VectorResList = append(funcs.VectorResList, vectorType)
			}
			serviceMessage.ResType = "Vector_" + vectorType.Name
		case mtproto_parser.BuiltInVectorType:
			vectorType := TplParam{
				Type: toProtoType(c.ResType),
				Name: toMessageName(c.ResType.(mtproto_parser.TVectorType).Type.Name()),
			}
			if !checkByParamList(funcs.VectorResList, vectorType) {
				funcs.VectorResList = append(funcs.VectorResList, vectorType)
			}
			serviceMessage.ResType = "Vector_" + vectorType.Name
		default:
			serviceMessage.ResType = toMessageName(c.ResType.Name())
		}
		service.SubMessageList = append(service.SubMessageList, serviceMessage)
	}

	for _, v := range serviceTypeMap {
		if checkByStringList(ignoreRpcList, v.Name) {
			continue
		}
		funcs.ServiceList = append(funcs.ServiceList, *v)
		// glog.Info(v)
	}

	return
}

//func makeRpcFunctions(schemas *mtproto_parser.Schemas) (tplMessages []TplBaseMessage) {
//	messages := make(map[string]*TplBaseMessage)
//
//	for _, c := range schemas.FunctionList {
//		if checkByStringList(ignoreRpcList, c.Method) {
//			continue
//		}
//
//		rpcName := strings.Split(c.Method, ".")[0]
//		// baseMessage := &TplBaseMessage{Name: baseTypeName}
//		baseMessages, ok := messages[rpcName]
//		if !ok {
//			baseMessages = &TplBaseMessage{Name: rpcName}
//			messages[rpcName] = baseMessages
//		}
//
//		message := TplMessage{
//			Name: strings.Replace(c.Method, ".", "_", -1),
//			Line: c.Line,
//		}
//
//		switch c.ResType.(type) {
//		case mtproto_parser.TVectorType:
//			message.HasVector = "Vector"
//			message.ResType = "Vector_" + strings.Replace(c.ResType.(mtproto_parser.TVectorType).Type.Name(), ".", "_", -1)
//		default:
//			message.ResType = strings.Replace(c.ResType.Name(), ".", "_", -1)
//		}
//
//		baseMessages.MessageList2 = append(baseMessages.MessageList2, message)
//	}
//
//	for _, v := range messages {
//		tplMessages = append(tplMessages, *v)
//	}
//	return
//}
