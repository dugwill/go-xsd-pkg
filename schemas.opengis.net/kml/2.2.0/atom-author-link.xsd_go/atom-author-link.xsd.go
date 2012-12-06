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
func (me *XsdGoPkgHasElems_Name) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_Name; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_Name struct {
	Name xsdt.String `xml:"http://www.w3.org/2005/Atom name"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_Name function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Name instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_Name instance.
func (me *XsdGoPkgHasElem_Name) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_Name; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_Uri struct {
	Uris []xsdt.String `xml:"http://www.w3.org/2005/Atom uri"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_Uri function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Uri instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Uri instance.
func (me *XsdGoPkgHasElems_Uri) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_Uri; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_Uri struct {
	Uri xsdt.String `xml:"http://www.w3.org/2005/Atom uri"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_Uri function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Uri instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_Uri instance.
func (me *XsdGoPkgHasElem_Uri) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_Uri; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TatomEmailAddress xsdt.String

//	Since TatomEmailAddress is just a simple String type, this merely returns the current string value.
func (me TatomEmailAddress) String() string { return xsdt.String(me).String() }

//	Since TatomEmailAddress is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomEmailAddress) SetFromString(s string) { (*xsdt.String)(me).SetFromString(s) }

//	This convenience method just performs a simple type conversion to TatomEmailAddress's alias type xsdt.String.
func (me TatomEmailAddress) ToXsdtString() xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasElem_Email struct {
	Email TatomEmailAddress `xml:"http://www.w3.org/2005/Atom email"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_Email function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Email instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElem_Email instance.
func (me *XsdGoPkgHasElem_Email) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_Email; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_Email struct {
	Emails []TatomEmailAddress `xml:"http://www.w3.org/2005/Atom email"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_Email function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Email instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Email instance.
func (me *XsdGoPkgHasElems_Email) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_Email; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TatomPersonConstruct struct {
	XsdGoPkgHasElems_Name

	XsdGoPkgHasElems_Uri

	XsdGoPkgHasElems_Email
}

//	If the WalkHandlers.TatomPersonConstruct function is not nil (ie. was set by outside code), calls it with this TatomPersonConstruct instance as the single argument. Then calls the Walk() method on 3/3 embed(s) and 0/0 field(s) belonging to this TatomPersonConstruct instance.
func (me *TatomPersonConstruct) Walk() (err error) {
	if fn := WalkHandlers.TatomPersonConstruct; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElems_Email.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElems_Name.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElems_Uri.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_Author struct {
	Author *TatomPersonConstruct `xml:"http://www.w3.org/2005/Atom author"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_Author function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Author instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_Author instance.
func (me *XsdGoPkgHasElem_Author) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_Author; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Author.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_Author struct {
	Authors []*TatomPersonConstruct `xml:"http://www.w3.org/2005/Atom author"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_Author function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Author instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Author instance.
func (me *XsdGoPkgHasElems_Author) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_Author; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Authors {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasAttr_Href_XsdtString_ struct {
	Href xsdt.String `xml:"http://www.w3.org/2005/Atom href,attr"`
}

type XsdGoPkgHasAttr_Rel_XsdtString_ struct {
	Rel xsdt.String `xml:"http://www.w3.org/2005/Atom rel,attr"`
}

type TatomLanguageTag xsdt.String

//	Since TatomLanguageTag is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomLanguageTag) SetFromString(s string) { (*xsdt.String)(me).SetFromString(s) }

//	This convenience method just performs a simple type conversion to TatomLanguageTag's alias type xsdt.String.
func (me TatomLanguageTag) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Since TatomLanguageTag is just a simple String type, this merely returns the current string value.
func (me TatomLanguageTag) String() string { return xsdt.String(me).String() }

type XsdGoPkgHasAttr_Hreflang_TatomLanguageTag_ struct {
	Hreflang TatomLanguageTag `xml:"http://www.w3.org/2005/Atom hreflang,attr"`
}

type XsdGoPkgHasAttr_Length_XsdtString_ struct {
	Length xsdt.String `xml:"http://www.w3.org/2005/Atom length,attr"`
}

type TatomMediaType xsdt.String

//	This convenience method just performs a simple type conversion to TatomMediaType's alias type xsdt.String.
func (me TatomMediaType) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Since TatomMediaType is just a simple String type, this merely sets the current value from the specified string.
func (me *TatomMediaType) SetFromString(s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TatomMediaType is just a simple String type, this merely returns the current string value.
func (me TatomMediaType) String() string { return xsdt.String(me).String() }

type XsdGoPkgHasAttr_Type_TatomMediaType_ struct {
	Type TatomMediaType `xml:"http://www.w3.org/2005/Atom type,attr"`
}

type XsdGoPkgHasAttr_Title_XsdtString_ struct {
	Title xsdt.String `xml:"http://www.w3.org/2005/Atom title,attr"`
}

type TxsdLink struct {
	XsdGoPkgHasAttr_Rel_XsdtString_

	XsdGoPkgHasAttr_Hreflang_TatomLanguageTag_

	XsdGoPkgHasAttr_Length_XsdtString_

	XsdGoPkgHasAttr_Type_TatomMediaType_

	XsdGoPkgHasAttr_Title_XsdtString_

	XsdGoPkgHasAttr_Href_XsdtString_
}

//	If the WalkHandlers.TxsdLink function is not nil (ie. was set by outside code), calls it with this TxsdLink instance as the single argument. Then calls the Walk() method on 0/6 embed(s) and 0/0 field(s) belonging to this TxsdLink instance.
func (me *TxsdLink) Walk() (err error) {
	if fn := WalkHandlers.TxsdLink; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_Link struct {
	Link *TxsdLink `xml:"http://www.w3.org/2005/Atom link"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_Link function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Link instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_Link instance.
func (me *XsdGoPkgHasElem_Link) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_Link; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.Link.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_Link struct {
	Links []*TxsdLink `xml:"http://www.w3.org/2005/Atom link"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_Link function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Link instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Link instance.
func (me *XsdGoPkgHasElems_Link) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_Link; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.Links {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 12 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 12 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasElems_Email  func(*XsdGoPkgHasElems_Email, bool) error
	TxsdLink                func(*TxsdLink, bool) error
	XsdGoPkgHasElem_Name    func(*XsdGoPkgHasElem_Name, bool) error
	XsdGoPkgHasElem_Email   func(*XsdGoPkgHasElem_Email, bool) error
	XsdGoPkgHasElem_Link    func(*XsdGoPkgHasElem_Link, bool) error
	TatomPersonConstruct    func(*TatomPersonConstruct, bool) error
	XsdGoPkgHasElems_Link   func(*XsdGoPkgHasElems_Link, bool) error
	XsdGoPkgHasElems_Uri    func(*XsdGoPkgHasElems_Uri, bool) error
	XsdGoPkgHasElems_Name   func(*XsdGoPkgHasElems_Name, bool) error
	XsdGoPkgHasElem_Author  func(*XsdGoPkgHasElem_Author, bool) error
	XsdGoPkgHasElems_Author func(*XsdGoPkgHasElems_Author, bool) error
	XsdGoPkgHasElem_Uri     func(*XsdGoPkgHasElem_Uri, bool) error
}
