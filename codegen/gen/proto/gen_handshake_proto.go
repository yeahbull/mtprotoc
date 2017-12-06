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
)

func genHandshakeProto(schemas *mtproto_parser.MTProtoSchemas, outFilePath string) {
	messages := &TplMessages{}
	messages.ConstructoList = makeTplConstructors(schemas.Handshake, false)
	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("./gen/tpl/schema.tl.handshake.proto.tpl"))
	t.Execute(&buf, messages)
	err := ioutil.WriteFile(fmt.Sprintf("%s/out/schema.tl.handshake.proto", outFilePath), buf.Bytes(), 0666)
	if err != nil {
		glog.Fatal("genCoreTypesProto error: ", err)
	}

	messages.FunctionList = makeTplFunctions(schemas.Handshake)
	var buf2 bytes.Buffer
	t = template.Must(template.ParseFiles("./gen/tpl/schema.tl.handshake_service.proto.tpl"))
	t.Execute(&buf2, messages)
	err = ioutil.WriteFile(fmt.Sprintf("%s/out/schema.tl.handshake_service.proto", outFilePath), buf2.Bytes(), 0666)
	if err != nil {
		glog.Fatal("genCoreTypesProto error: ", err)
	}
}

