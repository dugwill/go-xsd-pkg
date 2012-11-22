//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		schemas.opengis.net/kml/2.2.0/atom-author-link.xsd
package go_AtomAuthorLink

//	There is no official atom XSD. This XSD is created based on:
//	http://atompub.org/2005/08/17/atom.rnc. A subset of Atom as used in the
//	ogckml22.xsd is defined here.


import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasElems_Name struct {
	Names []xsdt.String `xml:"http://www.w3.org/2005/Atom name"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_Name function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Name instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Name instance.
func (me *XsdGoPkgHasElems_Name) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_Name; fn != nil { fn(me) }
 }

type XsdGoPkgHasElem_Name struct {
	Name xsdt.String `xml:"http://www.w3.org/2005/Atom name"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Name function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Name instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_Name instance.
func (me *XsdGoPkgHasElem_Name) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Name; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_Uri struct {
	Uris []xsdt.String `xml:"http://www.w3.org/2005/Atom uri"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_Uri function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Uri instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Uri instance.
func (me *XsdGoPkgHasElems_Uri) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_Uri; fn != nil { fn(me) }
 }

type XsdGoPkgHasElem_Uri struct {
	Uri xsdt.String `xml:"http://www.w3.org/2005/Atom uri"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Uri function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Uri instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_Uri instance.
func (me *XsdGoPkgHasElem_Uri) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Uri; fn != nil { fn(me) }
 }

type TatomEmailAddress xsdt.String

//	This convenience method just performs a simple type conversion to TatomEmailAddress's alias type xsdt.String.
func (me TatomEmailAddress) ToXsdtString () xsdt.String { return xsdt.String(me) }

//	Since TatomEmailAddress is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomEmailAddress) SetFromString (s string)  { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomEmailAddress is just a simple String type, this merely returns the current string value.
func (me TatomEmailAddress) String () string { return xsdt.String(me).String() }

type XsdGoPkgHasElems_Email struct {
	Emails []TatomEmailAddress `xml:"http://www.w3.org/2005/Atom email"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_Email function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Email instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Email instance.
func (me *XsdGoPkgHasElems_Email) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_Email; fn != nil { fn(me) }
 }

type XsdGoPkgHasElem_Email struct {
	Email TatomEmailAddress `xml:"http://www.w3.org/2005/Atom email"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Email function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Email instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_Email instance.
func (me *XsdGoPkgHasElem_Email) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Email; fn != nil { fn(me) }
 }

type TatomPersonConstruct struct {
	XsdGoPkgHasElems_Name

	XsdGoPkgHasElems_Email

	XsdGoPkgHasElems_Uri

}

//	If the WalkHandlers.TatomPersonConstruct function is not nil (ie. was set by outside code), calls it with this TatomPersonConstruct instance as the single argument. Then calls the Walk() method on 3/3 embed(s) and 0/0 field(s) belonging to this TatomPersonConstruct instance.
func (me *TatomPersonConstruct) Walk ()  { 
	if fn := WalkHandlers.TatomPersonConstruct; fn != nil { fn(me) }
	me.XsdGoPkgHasElems_Email.Walk()
	me.XsdGoPkgHasElems_Uri.Walk()
	me.XsdGoPkgHasElems_Name.Walk()
 }

type XsdGoPkgHasElem_Author struct {
	Author *TatomPersonConstruct `xml:"http://www.w3.org/2005/Atom author"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Author function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Author instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_Author instance.
func (me *XsdGoPkgHasElem_Author) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Author; fn != nil { fn(me) }
	me.Author.Walk()
 }

type XsdGoPkgHasElems_Author struct {
	Authors []*TatomPersonConstruct `xml:"http://www.w3.org/2005/Atom author"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_Author function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Author instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Author instance.
func (me *XsdGoPkgHasElems_Author) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_Author; fn != nil { fn(me) }
	for _, x := range me.Authors { x.Walk() }
 }

type XsdGoPkgHasAttr_Rel_XsdtString_ struct {
	Rel xsdt.String `xml:"http://www.w3.org/2005/Atom rel,attr"`

}

type XsdGoPkgHasAttr_Length_XsdtString_ struct {
	Length xsdt.String `xml:"http://www.w3.org/2005/Atom length,attr"`

}

type TatomMediaType xsdt.String

//	Since TatomMediaType is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomMediaType) SetFromString (s string)  { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomMediaType is just a simple String type, this merely returns the current string value.
func (me TatomMediaType) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TatomMediaType's alias type xsdt.String.
func (me TatomMediaType) ToXsdtString () xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Type_TatomMediaType_ struct {
	Type TatomMediaType `xml:"http://www.w3.org/2005/Atom type,attr"`

}

type TatomLanguageTag xsdt.String

//	Since TatomLanguageTag is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomLanguageTag) SetFromString (s string)  { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomLanguageTag is just a simple String type, this merely returns the current string value.
func (me TatomLanguageTag) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TatomLanguageTag's alias type xsdt.String.
func (me TatomLanguageTag) ToXsdtString () xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Hreflang_TatomLanguageTag_ struct {
	Hreflang TatomLanguageTag `xml:"http://www.w3.org/2005/Atom hreflang,attr"`

}

type XsdGoPkgHasAttr_Title_XsdtString_ struct {
	Title xsdt.String `xml:"http://www.w3.org/2005/Atom title,attr"`

}

type XsdGoPkgHasAttr_Href_XsdtString_ struct {
	Href xsdt.String `xml:"http://www.w3.org/2005/Atom href,attr"`

}

type TxsdLink struct {
	XsdGoPkgHasAttr_Href_XsdtString_

	XsdGoPkgHasAttr_Rel_XsdtString_

	XsdGoPkgHasAttr_Length_XsdtString_

	XsdGoPkgHasAttr_Type_TatomMediaType_

	XsdGoPkgHasAttr_Hreflang_TatomLanguageTag_

	XsdGoPkgHasAttr_Title_XsdtString_

}

//	If the WalkHandlers.TxsdLink function is not nil (ie. was set by outside code), calls it with this TxsdLink instance as the single argument. Then calls the Walk() method on 0/6 embed(s) and 0/0 field(s) belonging to this TxsdLink instance.
func (me *TxsdLink) Walk ()  { 
	if fn := WalkHandlers.TxsdLink; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_Link struct {
	Links []*TxsdLink `xml:"http://www.w3.org/2005/Atom link"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_Link function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Link instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Link instance.
func (me *XsdGoPkgHasElems_Link) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_Link; fn != nil { fn(me) }
	for _, x := range me.Links { x.Walk() }
 }

type XsdGoPkgHasElem_Link struct {
	Link *TxsdLink `xml:"http://www.w3.org/2005/Atom link"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Link function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Link instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_Link instance.
func (me *XsdGoPkgHasElem_Link) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Link; fn != nil { fn(me) }
	me.Link.Walk()
 }

//	Provides 12 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
var WalkHandlers = &XsdGoPkgWalkHandlers {}

//	Provides 12 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasElems_Name func (o *XsdGoPkgHasElems_Name)
	XsdGoPkgHasElem_Name func (o *XsdGoPkgHasElem_Name)
	XsdGoPkgHasElem_Uri func (o *XsdGoPkgHasElem_Uri)
	XsdGoPkgHasElem_Email func (o *XsdGoPkgHasElem_Email)
	XsdGoPkgHasElems_Author func (o *XsdGoPkgHasElems_Author)
	XsdGoPkgHasElems_Uri func (o *XsdGoPkgHasElems_Uri)
	TxsdLink func (o *TxsdLink)
	XsdGoPkgHasElems_Link func (o *XsdGoPkgHasElems_Link)
	TatomPersonConstruct func (o *TatomPersonConstruct)
	XsdGoPkgHasElem_Link func (o *XsdGoPkgHasElem_Link)
	XsdGoPkgHasElems_Email func (o *XsdGoPkgHasElems_Email)
	XsdGoPkgHasElem_Author func (o *XsdGoPkgHasElem_Author)
}
