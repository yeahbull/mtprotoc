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

package parser

import (
	"strings"
	"hash/crc32"
	"regexp"
	"os"
	"bufio"
	"io"
	"github.com/golang/glog"
	"fmt"
)

func calculateCRC32(name, params, restype string) uint32 {
	cleanline := fmt.Sprintf("%s%s= %s", name, params, restype)
	cleanline = regexp.MustCompile(" [a-zA-Z0-9_]+\\:flags\\.[0-9]+\\?true").ReplaceAllString(cleanline, "")
	cleanline = strings.Replace(cleanline, "<", " ", -1)
	cleanline = strings.Replace(cleanline, ">", " ", -1)
	cleanline = strings.Replace(cleanline, "  ", " ", -1)
	cleanline = regexp.MustCompile("^ ").ReplaceAllString(cleanline, "")
	cleanline = regexp.MustCompile(" $").ReplaceAllString(cleanline, "")
	cleanline = strings.Replace(cleanline, ":bytes ", ":string ", -1)
	cleanline = strings.Replace(cleanline, "?bytes ", "?string ", -1)
	cleanline = strings.Replace(cleanline, "{", "", -1)
	cleanline = strings.Replace(cleanline, "}", "", -1)

	// 通过cleanline计算出typeid并进行验证
	return crc32.ChecksumIEEE([]byte(cleanline))
}

func makeVectorType(schemas *Schemas, vecType string, vectemplate string) Type {
	var t Type
	matched1, _ := regexp.MatchString("^[A-Z]", vectemplate)
	matched2, _ := regexp.MatchString("^[a-zA-Z0-9]+_[A-Z]", vectemplate)
	if matched1 || matched2 {
		// CustomType
		if vecType == "vector" {
			t = BuiltInVectorType{CustomType{vectemplate}}
		} else {
			t = TVectorType{CustomType{vectemplate}}
		}
	} else if IsBuiltInTypeByName(vectemplate) {
		// glog.Info(vectemplate)
		// built-in type
		if vecType == "vector" {
			t = BuiltInVectorType{MakeBuiltInType(vectemplate)}
		} else {
			t = TVectorType{MakeBuiltInType(vectemplate)}
		}
	} else {
		// future_salts#ae500895 req_msg_id:long now:int salts:vector<future_salt> = FutureSalts;
		for _, metatype := range schemas.ConstructorList {
			if vectemplate == metatype.Name() {
				if vecType == "vector" {
					t = BuiltInVectorType{metatype}
				} else {
					t = TVectorType{metatype}
				}
				glog.Infof("foundmeta: %s", vectemplate)
				break
			}
		}

		if t == nil {
			glog.Errorf("invalid vectemplate: %s", vectemplate)
		}
	}
	return t
}

func Parse(filePath string) (*MTProtoSchemas, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	mtProtoSchemas := NewMTProtoSchemas()

	var funcsNow = false
	var schemas *Schemas

	bfRd := bufio.NewReader(f)
	for {
		line2, err := bfRd.ReadBytes('\n')
		if err != nil {
			//遇到任何错误立即返回，并忽略 EOF 错误信息
			// log.Println(schemas)
			if err == io.EOF {
				break
				// return schemas, nil
			}
			return nil, err
		}

		line := strings.TrimSpace(string(line2))

		///////////////////// Layer cons
		//LayerCons *Schemas
		///// Authorization key creation
		//Handshake *Schemas
		//////////////// System messages
		//Transport *Schemas
		/////////// Main application API
		//Sync *Schemas
		// layer
		layerline := regexp.MustCompile("// LAYER (\\d+)").FindStringSubmatch(line)
		if len(layerline) > 1 {
			mtProtoSchemas.Layer = layerline[1]
			continue
		}

		nocomment := regexp.MustCompile("^(.*?)//").FindStringSubmatch(line)
		if len(nocomment) > 1 {
			// Split schemas
			if strings.Index(line, "Layer cons") >= 0 {
				schemas = mtProtoSchemas.GetSchemas("LayerCons")
			} else if strings.Index(line, "Authorization key creation") >= 0 {
				schemas = mtProtoSchemas.GetSchemas("Handshake")
			} else if strings.Index(line, "System messages") >= 0 {
				schemas = mtProtoSchemas.GetSchemas("Transport")
			} else if strings.Index(line, "Main application API") >= 0 {
				schemas = mtProtoSchemas.GetSchemas("Sync")
			}

			line = nocomment[1]
		}

		// functions
		if matched, err := regexp.MatchString("\\-\\-\\-functions\\-\\-\\-", line); err == nil && matched {
			funcsNow = true
			continue
		}

		// types
		if matched, err := regexp.MatchString("\\-\\-\\-types\\-\\-\\-", line); err == nil && matched {
			funcsNow = false
			continue
		}

		if matched, err := regexp.MatchString("^\\s*$", line); err == nil && matched {
			// 空行
			continue
		}

		nametype := regexp.MustCompile("([a-zA-Z\\.0-9_]+)#([0-9a-f]+)([^=]*)=\\s*([a-zA-Z\\.<>0-9_]+);").FindStringSubmatch(line)
		if nametype == nil {
			// 特殊处理 vector#1cb5c415 {t:Type} # [ t ] = Vector t;
			if matched, err := regexp.MatchString("vector#1cb5c415 \\{t:Type\\} # \\[ t \\] = Vector t;", line); err == nil && !matched {
				fmt.Printf("Bad line found: %s\n", line)
			}
			continue
		}

		// resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
		// name为: resPQ
		// contest.saveDeveloperInfo#9a5f6e95 vk_id:int name:string phone_number:string age:int city:string = Bool;
		// name为: contest_saveDeveloperInfo
		name := string(nametype[1])

		// typeid为: 05162463
		typeid := string(nametype[2])
		// 去掉前面的0
		for len(typeid) > 0 && typeid[0] == '0' {
			typeid = typeid[1:]
		}
		if len(typeid) == 0 {
			typeid = "0"
		}
		typeid = "0x" + typeid

		// fmt.Printf("%s#%s\n", name, typeid)

		// cleanline := nametype[1] + nametype[3] + "= " + nametype[4]
		// 归一化并计算constractor
		crc32 := calculateCRC32(nametype[1], nametype[3], nametype[4])
		countTypeId := fmt.Sprintf("0x%x", crc32)

		if typeid != countTypeId {
			glog.Warning("Warning: counted ", countTypeId, " mismatch with provided ", typeid, " (", nametype[1], nametype[3], nametype[4], ")")
			continue
		}

		restype := nametype[4]
		var ResType Type

		if funcsNow {
			// ResType只有两种可能性, CustomType and Vector
			// functions resType
			if strings.Index(restype, "<") >= 0 {
				templ := regexp.MustCompile("^([vV]ector<)([A-Za-z0-9\\._]+)>$").FindStringSubmatch(restype)
				if templ == nil {
					glog.Errorf("make function vector restype error: %s, in line: %s\n", err, line)
					continue
				}

				ResType = makeVectorType(schemas, templ[1], templ[2])
			} else {
				ResType = MakeCustomType(restype)
			}
		} else {
			ResType = MakeCustomType(restype)
		}

		if _, ok := schemas.TypeMap[restype]; !ok {
			schemas.TypeMap[restype] = ResType
		}

		// 参数处理
		// params为: nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long>
		// params :=
		paramsList := strings.Split(strings.TrimSpace(nametype[3]), " ")
		var isTemplate, hasFlags, hasTemplate string

		var params []Param
		for _, param := range paramsList {
			if matched, _ := regexp.MatchString("^\\s*$", param); matched {
				// glog.Errorf("bad param found in param: %s, in line: %s", param, line)
				continue
			}

			// invokeAfterMsg#cb9f372d {X:Type} msg_id:long query:!X = X;
			templ := regexp.MustCompile("^{([A-Za-z]+):Type}$").FindStringSubmatch(param)
			if templ != nil {
				hasTemplate = templ[1]
				glog.Infof("hasTemplate in line: %s", line)
				continue
			}

			pnametype := regexp.MustCompile("([a-z_][a-z0-9_]*):([A-Za-z0-9<>\\._]+|![a-zA-Z]+|\\#|[a-z_][a-z0-9_]*\\.[0-9]+\\?[A-Za-z0-9<>\\._]+)$").FindStringSubmatch(param)
			if pnametype == nil {
				glog.Errorf("Bad param found: %s, in line: %s", param, line)
				continue
			}

			// 参数名, 参数类型
			pname := pnametype[1]
			ptypewide := pnametype[2]

			var ptype Type
			// _ = ptype

			if matched, _ := regexp.MatchString("^!([a-zA-Z]+)$", ptypewide); matched {
				if "!"+hasTemplate == ptypewide {
					// 模板类型
					isTemplate = pname
					ptype = TemplateType{}
					// ptype = "TQueryType"
					// print('template param name: ' + pname + ', type: TQueryType');
				} else {
					glog.Error("Bad template param name: ", param, ", in line: ", line)
					continue
				}
			} else if ptypewide == "#" {
				// # flags, 类似protobuf的optional字段
				hasFlags = pname
				// ptype = "int32"
				ptype = FlagsType{}
			} else {
				ptype2 := ptypewide
				if strings.Index(ptype2, "?") >= 0 {
					// # flags.0?string
					pmasktype := regexp.MustCompile("([a-z_][a-z0-9_]*)\\.([0-9]+)\\?([A-Za-z0-9<>\\._]+)").FindStringSubmatch(ptype2)
					if pmasktype == nil || pmasktype[1] != hasFlags {
						glog.Error("Bad param found: \"", param, "\" in line: ", line)
						continue
					}
					ptype2 = pmasktype[3]
					if strings.Index(ptype2, "<") >= 0 {
						// # inputMediaUploadedPhoto#630c9af1 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> = InputMedia;
						// # print('flags\'s template type: ' + ptype);
						templ = regexp.MustCompile("^([vV]ector<)([A-Za-z0-9\\._]+)>$").FindStringSubmatch(ptype2)
						if templ != nil {
							ptype = SubFlagsType{
								Mask: pmasktype[2],
								Type: makeVectorType(schemas, templ[1], templ[2]),
							}
							// ptype = makeVectorType(schemas, templ[1], templ[2])
						} else {
							glog.Error("Bad template type: ", ptype)
							continue
						}
					} else {
						// TODO(@benqi): 检测Invalid type
						if IsBuiltInTypeByName(ptype2) {
							ptype = SubFlagsType{
								Mask: pmasktype[2],
								Type: MakeBuiltInType(ptype2),
							}
						} else {
							ptype = SubFlagsType{
								Mask: pmasktype[2],
								Type: MakeCustomType(ptype2),
							}
						}
					}
				} else if strings.Index(ptype2, "<") >= 0 {
					templ = regexp.MustCompile("^([vV]ector<)([A-Za-z0-9\\._]+)>$").FindStringSubmatch(ptype2)
					if templ != nil {
						ptype = makeVectorType(schemas, templ[1], templ[2])
					} else {
						glog.Error("Bad template type: ", ptype)
						continue
					}
				} else {
					// TODO(@benqi): 检测Invalid type
					if IsBuiltInTypeByName(ptype2) {
						ptype = MakeBuiltInType(ptype2)
					} else {
						ptype =  MakeCustomType(ptype2)
					}
				}
			}

			params = append(params, Param{pname, ptype})
			// glog.Info(params, ", line: ", line)
			// prmsList = append(prmsList, pname)
		}

		if isTemplate == "" && restype == "X" {
			glog.Error("Bad response type \"X\" in \"", name, "\" in line: ", line)
			continue
		}

		if funcsNow {
			f := Function{
				Method: name,
				Id: int32(crc32),
				ParamList: params,
				ResType: ResType,
				Line: line,
			}
			schemas.FunctionList = append(schemas.FunctionList, f)
		} else {
			c := Constructor{
				Predicate: name,
				Id: int32(crc32),
				ParamList: params,
				BaseType: ResType,
				Line: line,
			}
			schemas.ConstructorList = append(schemas.ConstructorList, c)
		}
	}
	return mtProtoSchemas, nil
}
