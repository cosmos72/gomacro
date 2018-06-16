// this file was generated by gomacro command: import _b "crypto/x509/pkix"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	"crypto/x509/pkix"
	. "reflect"
)

// reflection: allow interpreted code to import "crypto/x509/pkix"
func init() {
	Packages["crypto/x509/pkix"] = Package{
		Types: map[string]Type{
			"AlgorithmIdentifier":          TypeOf((*pkix.AlgorithmIdentifier)(nil)).Elem(),
			"AttributeTypeAndValue":        TypeOf((*pkix.AttributeTypeAndValue)(nil)).Elem(),
			"AttributeTypeAndValueSET":     TypeOf((*pkix.AttributeTypeAndValueSET)(nil)).Elem(),
			"CertificateList":              TypeOf((*pkix.CertificateList)(nil)).Elem(),
			"Extension":                    TypeOf((*pkix.Extension)(nil)).Elem(),
			"Name":                         TypeOf((*pkix.Name)(nil)).Elem(),
			"RDNSequence":                  TypeOf((*pkix.RDNSequence)(nil)).Elem(),
			"RelativeDistinguishedNameSET": TypeOf((*pkix.RelativeDistinguishedNameSET)(nil)).Elem(),
			"RevokedCertificate":           TypeOf((*pkix.RevokedCertificate)(nil)).Elem(),
			"TBSCertificateList":           TypeOf((*pkix.TBSCertificateList)(nil)).Elem(),
		},
	}
}
