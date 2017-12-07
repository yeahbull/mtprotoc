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

func genCoreTypesProto(schemas *mtproto_parser.MTProtoSchemas, outFilePath string) {
	baseTypes := makeBaseTypeListTpl(schemas.Sync, true)
	// glog.Info(baseTypes)

	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("./gen/tpl/schema.tl.core_types.proto.tpl"))
	t.Execute(&buf, baseTypes)
	err := ioutil.WriteFile(fmt.Sprintf("%s/out/schema.tl.core_types.proto", outFilePath), buf.Bytes(), 0666)
	if err != nil {
		glog.Fatal("genCoreTypesProto error: ", err)
	}
}
