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
	"fmt"
)

/////////////////// Layer cons
type Schemas struct{
	TypeMap map[string]Type
	ConstructorList []Constructor
	FunctionList []Function
}

func NewSchemas() *Schemas {
	return &Schemas{
		TypeMap: make(map[string]Type),
		ConstructorList : make([]Constructor, 0),
		FunctionList : make([]Function, 0),
	}
}

type MTProtoSchemas struct{
	Layer string		// layer

	/////////////////// Layer cons
	LayerCons *Schemas
	/// Authorization key creation
	Handshake *Schemas
	////////////// System messages
	Transport *Schemas
	///////// Main application API
	Sync *Schemas
}

func NewMTProtoSchemas() *MTProtoSchemas {
	return &MTProtoSchemas{
		LayerCons:NewSchemas(),
		Handshake:NewSchemas(),
		Transport:NewSchemas(),
		Sync:NewSchemas(),
	}
}

func (this *MTProtoSchemas) GetSchemas(schemaType string) *Schemas {
	switch schemaType {
	case "LayerCons":
		return this.LayerCons
	case "Handshake":
		return this.Handshake
	case "Transport":
		return this.Transport
	case "Sync":
		return this.Sync
	default:
		panic("Bad schemaType: " + schemaType)
	}
	return nil
}

type Type interface {
	Name() string
}

type Param struct {
	Name string
	Type
}

func parserParam(n, t string) Param {
	return Param{
		Name:n,
	}
}

// funcsDict[restype].append([name, typeid, prmsList, prms, hasFlags, conditionsList, conditions, trivialConditions, line]);
type Function struct {
	Method string
	Id int32
	//PrmsList []string
	//Prms map[string]Type
	//HasFlags bool
	//ConditionsList []string
	//Conditions map[string]string
	//TrivialConditions []string
	ParamList []Param
	ResType Type
	Line string
}

func NewFunction(id int32, method, params, typeName string) Function {
	return Function{
		Id: id,
		Method: method,
		// Type: BaseCustomType{typeName},
	}
}

// typesDict[restype].append([name, typeid, prmsList, prms, hasFlags, conditionsList, conditions, trivialConditions, line]);
type Constructor struct {
	Predicate string
	Id int32
	//PrmsList []string
	//Prms map[string]Type
	//HasFlags bool
	//ConditionsList []string
	//Conditions map[string]string
	//TrivialConditions []string
	ParamList []Param
	BaseType Type
	Line string
}

func NewConstructor(id int32, predicate, params, typeName string) Constructor {
	return Constructor{
		Id: id,
		Predicate: predicate,
		BaseType: CustomType{typeName},
	}
}

func (this Constructor) Name() string {
	return this.Predicate
}

type CustomType struct {
	name string
}

func (this CustomType) Name() string {
	return this.name
}

func MakeCustomType(restype string) CustomType {
	return CustomType{restype}
}

///////////////////////////////////////////////////////////////////////////////
type FlagsType struct {
}

func (this FlagsType) Name() string {
	return "#"
}

type SubFlagsType struct {
	// name string	// flags.0:true
	Mask string
	Type
}

func (this SubFlagsType) Name() string {
	return fmt.Sprintf("flags.%d:%s", this.Mask, this.Type.Name())
}

type BuiltInVectorType struct {
	Type
}

func (this BuiltInVectorType) Name() string {
	return fmt.Sprintf("vector<%s>", this.Type.Name())
}

type TVectorType struct {
	Type
}

func (this TVectorType) Name() string {
	return fmt.Sprintf("Vector<%s>", this.Type.Name())
}

//////////////////////////////////////////////////////////////////////////
type TemplateType struct {
}

func (this TemplateType) Name() string {
	return "X"
}

type BoolType struct {
}

func (this BoolType) Name() string {
	return "bool"
}

type IntType struct {
}

func (this IntType) Name() string {
	return "int"
}

type LongType struct {
}

func (this LongType) Name() string {
	return "long"
}

type Int128Type struct {
}

func (this Int128Type) Name() string {
	return "int128"
}

type Int256Type struct {
}

func (this Int256Type) Name() string {
	return "int256"
}

type StringType struct {
}

func (this StringType) Name() string {
	return "string"
}

type BytesType struct {
}

func (this BytesType) Name() string {
	return "bytes"
}

func IsBuiltInType(t Type) bool {
	switch t.(type) {
	case BoolType, IntType, LongType, Int128Type, Int256Type, StringType, BytesType:
		return true
	default:
		return false
	}
}

func IsBuiltInTypeByName(n string) bool {
	switch n {
	case "true", "bool", "int", "long", "int128", "int256", "string", "bytes":
		return true
	default:
		return false
	}
}

func MakeBuiltInType(n string) Type {
	switch n {
	case "bool", "true":
		return BoolType{}
	case "int":
		return IntType{}
	case "long":
		return LongType{}
	case "int128":
		return Int128Type{}
	case "int256":
		return Int256Type{}
	case "string":
		return StringType{}
	case "bytes":
		return BytesType{}
	default:
		// 不可能执行
		return nil
	}
}

