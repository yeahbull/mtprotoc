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
	"fmt"
	"strings"
)

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
var (
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
	case mtproto_parser.Constructor:
		t2, _ := t.(mtproto_parser.Constructor)
		n = strings.Replace(t2.Predicate, ".", "_", -1)
	case mtproto_parser.TemplateType:
		n = "bytes"
	default:
		panic(fmt.Errorf("Bad type: %v", t))
	}

	return
}

func GenProto(schemas *mtproto_parser.MTProtoSchemas, outFilePath string) {
	genCRC32Proto(schemas, outFilePath)
	genCoreTypesProto(schemas, outFilePath)
	genHandshakeProto(schemas, outFilePath)
	genTransportProto(schemas, outFilePath)
	genSyncProto(schemas, outFilePath)
}

