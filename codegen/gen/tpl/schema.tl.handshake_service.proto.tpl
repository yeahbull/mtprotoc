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

// import "schema.tl.core_types.proto";

{{ range .RequestList }}
///////////////////////////////////////////////////////////////////////////////
// {{.Line}}
message TL_{{.Name}} {
{{range .ParamList}}    {{.Type}} {{.Name}} = {{.Index}};
{{end}}}
{{end}}

///////////////////////////////////////////////////////////////////////////////
// rpc
service RPCAuthKey {
    // req_pq#60469778 nonce:int128 = ResPQ;
    rpc req_pq(TL_req_pq) returns (ResPQ) {}

    // req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
    rpc req_DH_params(TL_req_DH_params) returns (Server_DH_Params) {}

    // set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
    rpc set_client_DH_params(TL_set_client_DH_params) returns (Set_client_DH_params_answer) {}

    // destroy_auth_key#d1435160 = DestroyAuthKeyRes;
    rpc destroy_auth_key(TL_destroy_auth_key) returns (DestroyAuthKeyRes) {}
}
