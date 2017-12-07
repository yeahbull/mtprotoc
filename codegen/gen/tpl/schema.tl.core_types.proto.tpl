/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_proto.py'
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

syntax = "proto3";

package mtproto;

option java_package = "com.nebulaim.engine.mtproto";
option java_outer_classname = "MTProto";
option optimize_for = CODE_SIZE;


{{ range .BaseTypeList }}
///////////////////////////////////////////////////////////////////////////////
// {{.Name}} <--
{{ range .SubMessageList }}//  + TL_{{.Name}}
{{end}}//
message {{.Name}}_Data {
{{range .ParamList}}    {{.Type}} {{.Name}} = {{.Index}};
{{end}}}

message {{.Name}} {
    int32 constructor = 1;
    {{.Name}}_Data data = 2;
}

{{ range .SubMessageList }}// {{.Line}}
message TL_{{.Name}} {
    {{.ResType}} data = 2;
}

{{end}}{{end}}
