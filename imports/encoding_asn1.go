// this file was generated by gomacro command: import _b "encoding/asn1"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	asn1 "encoding/asn1"
)

// reflection: allow interpreted code to import "encoding/asn1"
func init() {
	Packages["encoding/asn1"] = Package{
		Name: "asn1",
		Binds: map[string]Value{
			"ClassApplication":	ValueOf(asn1.ClassApplication),
			"ClassContextSpecific":	ValueOf(asn1.ClassContextSpecific),
			"ClassPrivate":	ValueOf(asn1.ClassPrivate),
			"ClassUniversal":	ValueOf(asn1.ClassUniversal),
			"Marshal":	ValueOf(asn1.Marshal),
			"MarshalWithParams":	ValueOf(asn1.MarshalWithParams),
			"NullBytes":	ValueOf(&asn1.NullBytes).Elem(),
			"NullRawValue":	ValueOf(&asn1.NullRawValue).Elem(),
			"TagBMPString":	ValueOf(asn1.TagBMPString),
			"TagBitString":	ValueOf(asn1.TagBitString),
			"TagBoolean":	ValueOf(asn1.TagBoolean),
			"TagEnum":	ValueOf(asn1.TagEnum),
			"TagGeneralString":	ValueOf(asn1.TagGeneralString),
			"TagGeneralizedTime":	ValueOf(asn1.TagGeneralizedTime),
			"TagIA5String":	ValueOf(asn1.TagIA5String),
			"TagInteger":	ValueOf(asn1.TagInteger),
			"TagNull":	ValueOf(asn1.TagNull),
			"TagNumericString":	ValueOf(asn1.TagNumericString),
			"TagOID":	ValueOf(asn1.TagOID),
			"TagOctetString":	ValueOf(asn1.TagOctetString),
			"TagPrintableString":	ValueOf(asn1.TagPrintableString),
			"TagSequence":	ValueOf(asn1.TagSequence),
			"TagSet":	ValueOf(asn1.TagSet),
			"TagT61String":	ValueOf(asn1.TagT61String),
			"TagUTCTime":	ValueOf(asn1.TagUTCTime),
			"TagUTF8String":	ValueOf(asn1.TagUTF8String),
			"Unmarshal":	ValueOf(asn1.Unmarshal),
			"UnmarshalWithParams":	ValueOf(asn1.UnmarshalWithParams),
		}, Types: map[string]Type{
			"BitString":	TypeOf((*asn1.BitString)(nil)).Elem(),
			"Enumerated":	TypeOf((*asn1.Enumerated)(nil)).Elem(),
			"Flag":	TypeOf((*asn1.Flag)(nil)).Elem(),
			"ObjectIdentifier":	TypeOf((*asn1.ObjectIdentifier)(nil)).Elem(),
			"RawContent":	TypeOf((*asn1.RawContent)(nil)).Elem(),
			"RawValue":	TypeOf((*asn1.RawValue)(nil)).Elem(),
			"StructuralError":	TypeOf((*asn1.StructuralError)(nil)).Elem(),
			"SyntaxError":	TypeOf((*asn1.SyntaxError)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"ClassApplication":	"int:1",
			"ClassContextSpecific":	"int:2",
			"ClassPrivate":	"int:3",
			"ClassUniversal":	"int:0",
			"TagBMPString":	"int:30",
			"TagBitString":	"int:3",
			"TagBoolean":	"int:1",
			"TagEnum":	"int:10",
			"TagGeneralString":	"int:27",
			"TagGeneralizedTime":	"int:24",
			"TagIA5String":	"int:22",
			"TagInteger":	"int:2",
			"TagNull":	"int:5",
			"TagNumericString":	"int:18",
			"TagOID":	"int:6",
			"TagOctetString":	"int:4",
			"TagPrintableString":	"int:19",
			"TagSequence":	"int:16",
			"TagSet":	"int:17",
			"TagT61String":	"int:20",
			"TagUTCTime":	"int:23",
			"TagUTF8String":	"int:12",
		}, 
	}
}
