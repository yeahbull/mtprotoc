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
	mtproto_parser "github.com/nebulaim/mtprotoc/codegen/parser"
	"bytes"
	"io/ioutil"
	"fmt"
	"github.com/golang/glog"
	"text/template"
	"strings"
	"github.com/cosiner/gohper/conv"
	"github.com/astaxie/beego/logs"
)

type TplParam struct {
	Type string
	Name string
	Index string
}

type TplMessage struct {
	Name string
	ParamList []TplParam
	Line string
	ResType string
}

// TODO(@benqi): 独立出types, request, rpc
type TplBaseMessage struct {
	Name string
	ParamList []TplParam
	MessageList2 []TplMessage
	ResType string
	Line string
}

type TplMessages struct {
	ConstructoList []TplBaseMessage
	FunctionList []TplBaseMessage
	RPCList []TplBaseMessage
}

func findByParamList(params []TplParam, p2 TplParam) bool {
	for _, p := range params {
		if p.Name == p2.Name && p.Type == p2.Type {
			return true
		}
	}
	return false
}

func findByStringList(strList []string, s string) bool {
	for _, p := range strList {
		if p == s {
			return true
		}
	}
	return false
}

func makeTplConstructors(schemas *mtproto_parser.Schemas, isCoreTypes bool) (tplMessages []TplBaseMessage) {
	messages := make(map[string]*TplBaseMessage)

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

		baseTypeName := strings.Replace(c.BaseType.Name(), ".", "_", -1)
		// baseMessage := &TplBaseMessage{Name: baseTypeName}
		baseMessages, ok := messages[baseTypeName]
		if !ok {
			baseMessages = &TplBaseMessage{Name: baseTypeName}
			messages[baseTypeName] = baseMessages
		}

		message := TplMessage{
			Name: strings.Replace(c.Predicate, ".", "_", -1),
			Line: c.Line,
			ResType: strings.Replace(c.BaseType.Name(), ".", "_", -1),
		}
		glog.Info(c.Line)

		for _, p := range c.ParamList {
			param := TplParam{}
			param.Name = p.Name
			// param.Index = conv.I64toa(int64(i+1))
			param.Type = toProtoType(p.Type)
			if param.Type == "" {
				continue
			}
			message.ParamList = append(message.ParamList, param)

			if !findByParamList(baseMessages.ParamList, param) {
				baseMessages.ParamList = append(baseMessages.ParamList, param)

				// param.Index
			} else {
				logs.Info(param)
			}
		}
		baseMessages.MessageList2 = append(baseMessages.MessageList2, message)

		// messages.MessageList = append(messages.MessageList, message)
	}

	// tplMessages = &TplMessageList{}
	for _, v := range messages {
		for i, _ := range v.ParamList {
			v.ParamList[i].Index = conv.I64toa(int64(i+1))
		}
		// glog.Info(v)
		tplMessages = append(tplMessages, *v)
	}
	return
}

func makeTplFunctions(schemas *mtproto_parser.Schemas) (tplMessages []TplBaseMessage) {
	// messages := make(map[string]*TplBaseMessage)

	for _, c := range schemas.FunctionList {
		message := TplBaseMessage{
			Name: strings.Replace(c.Method, ".", "_", -1),
			Line: c.Line,
		}
		glog.Info(c.Line)
		for i, p := range c.ParamList {
			param := TplParam{}
			param.Name = p.Name
			param.Index = conv.I64toa(int64(i+1))
			param.Type = toProtoType(p.Type)
			if param.Type == "" {
				continue
			}
			message.ParamList = append(message.ParamList, param)
		}
		tplMessages = append(tplMessages, message)
	}
	return
}

func makeRpcFunctions(schemas *mtproto_parser.Schemas) (tplMessages []TplBaseMessage) {
	messages := make(map[string]*TplBaseMessage)

	for _, c := range schemas.FunctionList {
		rpcName := strings.Split(c.Method, ".")[0]
		// baseMessage := &TplBaseMessage{Name: baseTypeName}
		baseMessages, ok := messages[rpcName]
		if !ok {
			baseMessages = &TplBaseMessage{Name: rpcName}
			messages[rpcName] = baseMessages
		}

		message := TplMessage{
			Name: strings.Replace(c.Method, ".", "_", -1),
			Line: c.Line,
			ResType: strings.Replace(c.ResType.Name(), ".", "_", -1),
		}
		baseMessages.MessageList2 = append(baseMessages.MessageList2, message)
	}

	for _, v := range messages {
		tplMessages = append(tplMessages, *v)
	}
	return
}

func genCoreTypesProto(schemas *mtproto_parser.MTProtoSchemas, outFilePath string) {
	messages := &TplMessages{}
	messages.ConstructoList = makeTplConstructors(schemas.Sync, true)
	// glog.Info(messages)

	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("./gen/tpl/schema.tl.core_types.proto.tpl"))
	t.Execute(&buf, messages)
	err := ioutil.WriteFile(fmt.Sprintf("%s/out/schema.tl.core_types.proto", outFilePath), buf.Bytes(), 0666)
	if err != nil {
		glog.Fatal("genCoreTypesProto error: ", err)
	}
}
