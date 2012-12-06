//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/TR/2002/WD-SVG11-20020108/xml.xsd
package go_Xml

import (
	xsdt "github.com/metaleap/go-xsd/types"
)

//	The content of the lang attribute is a hyphen-separated list of case-insensitive tokens where the first token, if one character, is "i" or "x"; if two characters, is an ISO 639-1 language code; if three characters, is an ISO 639-2 language code. The second token, if two characters, is an ISO 3166-1 country code; if from three to eight characters, is an IANA-registered language code. The meaning of other tokens is undefined. This documentation briefly summarizes  RFC3066, which obseletes RFC1766.
//	Examples: en-GB ja-JP no-bok sgn-US i-navajo cel-gaulish
//	XML specification, Language Identification (in the second edition as modified by errattum E11)
//	RFC 3066
//	Registration authority for ISO 639-1 and -2
//	Registration authority for ISO 3166
//	IANA-registered language codes
type XsdGoPkgHasAttr_Lang struct {
	//	The content of the lang attribute is a hyphen-separated list of case-insensitive tokens where the first token, if one character, is "i" or "x"; if two characters, is an ISO 639-1 language code; if three characters, is an ISO 639-2 language code. The second token, if two characters, is an ISO 3166-1 country code; if from three to eight characters, is an IANA-registered language code. The meaning of other tokens is undefined. This documentation briefly summarizes  RFC3066, which obseletes RFC1766.
	//	Examples: en-GB ja-JP no-bok sgn-US i-navajo cel-gaulish
	//	XML specification, Language Identification (in the second edition as modified by errattum E11)
	//	RFC 3066
	//	Registration authority for ISO 639-1 and -2
	//	Registration authority for ISO 3166
	//	IANA-registered language codes
	Lang xsdt.Nmtoken `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
}

type TxsdSpace xsdt.String

//	This convenience method just performs a simple type conversion to TxsdSpace's alias type xsdt.String.
func (me TxsdSpace) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdSpace is "default".
func (me TxsdSpace) IsDefault() bool { return me == "default" }

//	Returns true if the value of this enumerated TxsdSpace is "preserve".
func (me TxsdSpace) IsPreserve() bool { return me == "preserve" }

//	Since TxsdSpace is just a simple String type, this merely returns the current string value.
func (me TxsdSpace) String() string { return xsdt.String(me).String() }

//	Since TxsdSpace is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdSpace) SetFromString(s string) { (*xsdt.String)(me).SetFromString(s) }

type XsdGoPkgHasAttr_Space struct {
	Space TxsdSpace `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
}

type XsdGoPkgHasAttr_Base struct {
	Base xsdt.AnyURI `xml:"http://www.w3.org/XML/1998/namespace base,attr"`
}
