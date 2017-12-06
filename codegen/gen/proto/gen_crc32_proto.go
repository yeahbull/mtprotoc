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
)

type TplCRC32Info struct {
	Name string
	Id int32
}

type TplCRC32 struct {
	CRC32List []TplCRC32Info
}

func makeTplCRC32List(schemas *mtproto_parser.Schemas) (tplCRC32List []TplCRC32Info) {
	for _, c := range schemas.ConstructorList {
		tplCRC32List = append(tplCRC32List, TplCRC32Info{
			Name: strings.Replace(c.Predicate, ".", "_", -1),
			Id: c.Id,
		})
	}
	for _, f := range schemas.FunctionList {
		tplCRC32List = append(tplCRC32List, TplCRC32Info{
			Name: strings.Replace(f.Method, ".", "_", -1),
			Id: f.Id,
		})
	}
	return
}

func genCRC32Proto(schemas *mtproto_parser.MTProtoSchemas, outFilePath string) {
	tplCRC32 := &TplCRC32{}
	tplCRC32.CRC32List = append(tplCRC32.CRC32List, makeTplCRC32List(schemas.LayerCons)...)
	tplCRC32.CRC32List = append(tplCRC32.CRC32List, makeTplCRC32List(schemas.Handshake)...)
	tplCRC32.CRC32List = append(tplCRC32.CRC32List, makeTplCRC32List(schemas.Transport)...)
	tplCRC32.CRC32List = append(tplCRC32.CRC32List, makeTplCRC32List(schemas.Sync)...)

	// glog.Info(tplCRC32)
	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("./gen/tpl/schema.tl.crc32.proto.tpl"))
	t.Execute(&buf, tplCRC32)
	// glog.Info(fmt.Sprintf("%s/out/schema.tl.crc32.proto", outFilePath))
	err := ioutil.WriteFile(fmt.Sprintf("%s/out/schema.tl.crc32.proto", outFilePath), buf.Bytes(), 0666)
	if err != nil {
		glog.Fatal("genProtoCRC32 error: ", err)
	}
}

