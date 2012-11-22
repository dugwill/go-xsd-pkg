//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		kbcafe.com/rss/atom.xsd.xml
package go_AtomXml

//	This version of the Atom schema is based on version 1.0 of the format specifications,
//	found here http://www.atomenabled.org/developers/syndication/atom-format-spec.php.
//	An Atom document may have two root elements, feed and entry, as defined in section 2.


import (
	xsdt "github.com/metaleap/go-xsd/types"
	xml "github.com/metaleap/go-xsd-pkg/www.w3.org/2001/03/xml.xsd_go"
)

type XsdGoPkgHasAtts_CommonAttributes struct {
	xml.XsdGoPkgHasAttr_Base

	xml.XsdGoPkgHasAttr_Lang

}

//	The Atom feed construct is defined in section 4.1.1 of the format spec.
//	The Atom text construct is defined in section 3.1 of the format spec.
type TxsdTextTypeType xsdt.Token

//	Since TxsdTextTypeType is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdTextTypeType) SetFromString (s string)  { (*xsdt.Token)(me).SetFromString(s) }

//	Returns true if the value of this enumerated TxsdTextTypeType is "xhtml".
func (me TxsdTextTypeType) IsXhtml () bool { return me == "xhtml" }

//	Returns true if the value of this enumerated TxsdTextTypeType is "text".
func (me TxsdTextTypeType) IsText () bool { return me == "text" }

//	Since TxsdTextTypeType is just a simple String type, this merely returns the current string value.
func (me TxsdTextTypeType) String () string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TxsdTextTypeType's alias type xsdt.Token.
func (me TxsdTextTypeType) ToXsdtToken () xsdt.Token { return xsdt.Token(me) }

//	Returns true if the value of this enumerated TxsdTextTypeType is "html".
func (me TxsdTextTypeType) IsHtml () bool { return me == "html" }

type XsdGoPkgHasAttr_Type_TxsdTextTypeType_ struct {
	Type TxsdTextTypeType `xml:"http://www.w3.org/2005/Atom type,attr"`

}

type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`

}

//	If the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
func (me *XsdGoPkgHasCdata) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasCdata; fn != nil { fn(me) }
 }

type TtextType struct {
	XsdGoPkgHasAttr_Type_TxsdTextTypeType_

	XsdGoPkgHasCdata

	XsdGoPkgHasAtts_CommonAttributes

}

//	If the WalkHandlers.TtextType function is not nil (ie. was set by outside code), calls it with this TtextType instance as the single argument. Then calls the Walk() method on 1/3 embed(s) and 0/0 field(s) belonging to this TtextType instance.
func (me *TtextType) Walk ()  { 
	if fn := WalkHandlers.TtextType; fn != nil { fn(me) }
	me.XsdGoPkgHasCdata.Walk()
 }

type XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_ struct {
	Titles []*TtextType `xml:"http://www.w3.org/2005/Atom title"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_ instance.
func (me *XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_; fn != nil { fn(me) }
	for _, x := range me.Titles { x.Walk() }
 }

//	The Atom icon construct is defined in section 4.2.5 of the format spec.
type TiconType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	Simply returns the value of its XsdGoPkgValue field.
func (me *TiconType) ToXsdtAnyURI () xsdt.AnyURI { return me.XsdGoPkgValue }

//	If the WalkHandlers.TiconType function is not nil (ie. was set by outside code), calls it with this TiconType instance as the single argument. Then calls the Walk() method on 0/1 embed(s) and 0/1 field(s) belonging to this TiconType instance.
func (me *TiconType) Walk ()  { 
	if fn := WalkHandlers.TiconType; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_ struct {
	Icons []*TiconType `xml:"http://www.w3.org/2005/Atom icon"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_ instance.
func (me *XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_; fn != nil { fn(me) }
	for _, x := range me.Icons { x.Walk() }
 }

//	The Atom person construct is defined in section 3.2 of the format spec.
//	Schema definition for an email address.
type TemailType xsdt.NormalizedString

//	Since TemailType is just a simple String type, this merely returns the current string value.
func (me TemailType) String () string { return xsdt.NormalizedString(me).String() }

//	This convenience method just performs a simple type conversion to TemailType's alias type xsdt.NormalizedString.
func (me TemailType) ToXsdtNormalizedString () xsdt.NormalizedString { return xsdt.NormalizedString(me) }

//	Since TemailType is just a simple String type, this merely sets the current value from the specified string.
func (me *TemailType) SetFromString (s string)  { (*xsdt.NormalizedString)(me).SetFromString(s) }

type XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_ struct {
	Emails []TemailType `xml:"http://www.w3.org/2005/Atom email"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_ instance.
func (me *XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_ struct {
	Names []xsdt.String `xml:"http://www.w3.org/2005/Atom name"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_ instance.
func (me *XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_; fn != nil { fn(me) }
 }

type TuriType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	Simply returns the value of its XsdGoPkgValue field.
func (me *TuriType) ToXsdtAnyURI () xsdt.AnyURI { return me.XsdGoPkgValue }

//	If the WalkHandlers.TuriType function is not nil (ie. was set by outside code), calls it with this TuriType instance as the single argument. Then calls the Walk() method on 0/1 embed(s) and 0/1 field(s) belonging to this TuriType instance.
func (me *TuriType) Walk ()  { 
	if fn := WalkHandlers.TuriType; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_ struct {
	Uris []*TuriType `xml:"http://www.w3.org/2005/Atom uri"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_ instance.
func (me *XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_; fn != nil { fn(me) }
	for _, x := range me.Uris { x.Walk() }
 }

type TpersonType struct {
	XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_

	XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_

	XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_

	XsdGoPkgHasAtts_CommonAttributes

}

//	If the WalkHandlers.TpersonType function is not nil (ie. was set by outside code), calls it with this TpersonType instance as the single argument. Then calls the Walk() method on 3/4 embed(s) and 0/0 field(s) belonging to this TpersonType instance.
func (me *TpersonType) Walk ()  { 
	if fn := WalkHandlers.TpersonType; fn != nil { fn(me) }
	me.XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_.Walk()
	me.XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_.Walk()
	me.XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_.Walk()
 }

type XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_ struct {
	Contributors []*TpersonType `xml:"http://www.w3.org/2005/Atom contributor"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_ instance.
func (me *XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_; fn != nil { fn(me) }
	for _, x := range me.Contributors { x.Walk() }
 }

//	The Atom id construct is defined in section 4.2.6 of the format spec.
type TidType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	Simply returns the value of its XsdGoPkgValue field.
func (me *TidType) ToXsdtAnyURI () xsdt.AnyURI { return me.XsdGoPkgValue }

//	If the WalkHandlers.TidType function is not nil (ie. was set by outside code), calls it with this TidType instance as the single argument. Then calls the Walk() method on 0/1 embed(s) and 0/1 field(s) belonging to this TidType instance.
func (me *TidType) Walk ()  { 
	if fn := WalkHandlers.TidType; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_ struct {
	Ids []*TidType `xml:"http://www.w3.org/2005/Atom id"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_ instance.
func (me *XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_; fn != nil { fn(me) }
	for _, x := range me.Ids { x.Walk() }
 }

//	The Atom cagegory construct is defined in section 4.2.2 of the format spec.
type XsdGoPkgHasAttr_Label_XsdtString_ struct {
	Label xsdt.String `xml:"http://www.w3.org/2005/Atom label,attr"`

}

type XsdGoPkgHasAttr_Scheme_XsdtAnyURI_ struct {
	Scheme xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom scheme,attr"`

}

type XsdGoPkgHasAttr_Term_XsdtString_ struct {
	Term xsdt.String `xml:"http://www.w3.org/2005/Atom term,attr"`

}

type TcategoryType struct {
	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Scheme_XsdtAnyURI_

	XsdGoPkgHasAttr_Term_XsdtString_

	XsdGoPkgHasAttr_Label_XsdtString_

}

//	If the WalkHandlers.TcategoryType function is not nil (ie. was set by outside code), calls it with this TcategoryType instance as the single argument. Then calls the Walk() method on 0/4 embed(s) and 0/0 field(s) belonging to this TcategoryType instance.
func (me *TcategoryType) Walk ()  { 
	if fn := WalkHandlers.TcategoryType; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_ struct {
	Categories []*TcategoryType `xml:"http://www.w3.org/2005/Atom category"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_ instance.
func (me *XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_; fn != nil { fn(me) }
	for _, x := range me.Categories { x.Walk() }
 }

//	The Atom generator element is defined in section 4.2.4 of the format spec.
type XsdGoPkgHasAttr_Uri_XsdtAnyURI_ struct {
	Uri xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom uri,attr"`

}

type XsdGoPkgHasAttr_Version_XsdtString_ struct {
	Version xsdt.String `xml:"http://www.w3.org/2005/Atom version,attr"`

}

type TgeneratorType struct {
	XsdGoPkgValue xsdt.String `xml:",chardata"`

	XsdGoPkgHasAttr_Uri_XsdtAnyURI_

	XsdGoPkgHasAttr_Version_XsdtString_

	XsdGoPkgHasAtts_CommonAttributes

}

//	Simply returns the value of its XsdGoPkgValue field.
func (me *TgeneratorType) ToXsdtString () xsdt.String { return me.XsdGoPkgValue }

//	If the WalkHandlers.TgeneratorType function is not nil (ie. was set by outside code), calls it with this TgeneratorType instance as the single argument. Then calls the Walk() method on 0/3 embed(s) and 0/1 field(s) belonging to this TgeneratorType instance.
func (me *TgeneratorType) Walk ()  { 
	if fn := WalkHandlers.TgeneratorType; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_ struct {
	Generators []*TgeneratorType `xml:"http://www.w3.org/2005/Atom generator"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_ instance.
func (me *XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_; fn != nil { fn(me) }
	for _, x := range me.Generators { x.Walk() }
 }

type XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_ struct {
	Subtitles []*TtextType `xml:"http://www.w3.org/2005/Atom subtitle"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_ instance.
func (me *XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_; fn != nil { fn(me) }
	for _, x := range me.Subtitles { x.Walk() }
 }

type XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_ struct {
	Rightses []*TtextType `xml:"http://www.w3.org/2005/Atom rights"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_ instance.
func (me *XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_; fn != nil { fn(me) }
	for _, x := range me.Rightses { x.Walk() }
 }

//	The Atom logo construct is defined in section 4.2.8 of the format spec.
type TlogoType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	If the WalkHandlers.TlogoType function is not nil (ie. was set by outside code), calls it with this TlogoType instance as the single argument. Then calls the Walk() method on 0/1 embed(s) and 0/1 field(s) belonging to this TlogoType instance.
func (me *TlogoType) Walk ()  { 
	if fn := WalkHandlers.TlogoType; fn != nil { fn(me) }
 }

//	Simply returns the value of its XsdGoPkgValue field.
func (me *TlogoType) ToXsdtAnyURI () xsdt.AnyURI { return me.XsdGoPkgValue }

type XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_ struct {
	Logos []*TlogoType `xml:"http://www.w3.org/2005/Atom logo"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_ instance.
func (me *XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_; fn != nil { fn(me) }
	for _, x := range me.Logos { x.Walk() }
 }

type XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_ struct {
	Authors []*TpersonType `xml:"http://www.w3.org/2005/Atom author"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_ instance.
func (me *XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_; fn != nil { fn(me) }
	for _, x := range me.Authors { x.Walk() }
 }

type TdateTimeType struct {
	XsdGoPkgValue xsdt.DateTime `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	Simply returns the value of its XsdGoPkgValue field.
func (me *TdateTimeType) ToXsdtDateTime () xsdt.DateTime { return me.XsdGoPkgValue }

//	If the WalkHandlers.TdateTimeType function is not nil (ie. was set by outside code), calls it with this TdateTimeType instance as the single argument. Then calls the Walk() method on 0/1 embed(s) and 0/1 field(s) belonging to this TdateTimeType instance.
func (me *TdateTimeType) Walk ()  { 
	if fn := WalkHandlers.TdateTimeType; fn != nil { fn(me) }
 }

type XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_ struct {
	Updateds []*TdateTimeType `xml:"http://www.w3.org/2005/Atom updated"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_ instance.
func (me *XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_; fn != nil { fn(me) }
	for _, x := range me.Updateds { x.Walk() }
 }

//	The Atom link construct is defined in section 3.4 of the format spec.
type XsdGoPkgHasAttr_Rel_XsdtString_ struct {
	Rel xsdt.String `xml:"http://www.w3.org/2005/Atom rel,attr"`

}

type XsdGoPkgHasAttr_Hreflang_XsdtNmtoken_ struct {
	Hreflang xsdt.Nmtoken `xml:"http://www.w3.org/2005/Atom hreflang,attr"`

}

type XsdGoPkgHasAttr_Title_XsdtString_ struct {
	Title xsdt.String `xml:"http://www.w3.org/2005/Atom title,attr"`

}

type XsdGoPkgHasAttr_Length_XsdtPositiveInteger_ struct {
	Length xsdt.PositiveInteger `xml:"http://www.w3.org/2005/Atom length,attr"`

}

type XsdGoPkgHasAttr_Href_XsdtAnyURI_ struct {
	Href xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom href,attr"`

}

type XsdGoPkgHasAttr_Type_XsdtString_ struct {
	Type xsdt.String `xml:"http://www.w3.org/2005/Atom type,attr"`

}

type TlinkType struct {
	XsdGoPkgHasAttr_Type_XsdtString_

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Rel_XsdtString_

	XsdGoPkgHasAttr_Hreflang_XsdtNmtoken_

	XsdGoPkgHasAttr_Title_XsdtString_

	XsdGoPkgHasAttr_Length_XsdtPositiveInteger_

	XsdGoPkgHasAttr_Href_XsdtAnyURI_

	XsdGoPkgHasCdata

}

//	If the WalkHandlers.TlinkType function is not nil (ie. was set by outside code), calls it with this TlinkType instance as the single argument. Then calls the Walk() method on 1/8 embed(s) and 0/0 field(s) belonging to this TlinkType instance.
func (me *TlinkType) Walk ()  { 
	if fn := WalkHandlers.TlinkType; fn != nil { fn(me) }
	me.XsdGoPkgHasCdata.Walk()
 }

type XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_ struct {
	Links []*TlinkType `xml:"http://www.w3.org/2005/Atom link"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_ instance.
func (me *XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_; fn != nil { fn(me) }
	for _, x := range me.Links { x.Walk() }
 }

//	The Atom entry construct is defined in section 4.1.2 of the format spec.
//	The Atom content construct is defined in section 4.1.3 of the format spec.
type XsdGoPkgHasAttr_Src_XsdtAnyURI_ struct {
	Src xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom src,attr"`

}

type TcontentType struct {
	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasCdata

	XsdGoPkgHasAttr_Src_XsdtAnyURI_

	XsdGoPkgHasAttr_Type_XsdtString_

}

//	If the WalkHandlers.TcontentType function is not nil (ie. was set by outside code), calls it with this TcontentType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TcontentType instance.
func (me *TcontentType) Walk ()  { 
	if fn := WalkHandlers.TcontentType; fn != nil { fn(me) }
	me.XsdGoPkgHasCdata.Walk()
 }

type XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_ struct {
	Contents []*TcontentType `xml:"http://www.w3.org/2005/Atom content"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_ instance.
func (me *XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_; fn != nil { fn(me) }
	for _, x := range me.Contents { x.Walk() }
 }

type XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_ struct {
	Summaries []*TtextType `xml:"http://www.w3.org/2005/Atom summary"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_ instance.
func (me *XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_; fn != nil { fn(me) }
	for _, x := range me.Summaries { x.Walk() }
 }

type XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_ struct {
	Sources []*TtextType `xml:"http://www.w3.org/2005/Atom source"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_ instance.
func (me *XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_; fn != nil { fn(me) }
	for _, x := range me.Sources { x.Walk() }
 }

type XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_ struct {
	Publisheds []*TdateTimeType `xml:"http://www.w3.org/2005/Atom published"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_ instance.
func (me *XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_; fn != nil { fn(me) }
	for _, x := range me.Publisheds { x.Walk() }
 }

type TentryType struct {
	XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_

	XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_

	XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_

	XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_

	XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_

	XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_

	XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_

	XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_

	XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_

	XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_

	XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_

}

//	If the WalkHandlers.TentryType function is not nil (ie. was set by outside code), calls it with this TentryType instance as the single argument. Then calls the Walk() method on 12/13 embed(s) and 0/0 field(s) belonging to this TentryType instance.
func (me *TentryType) Walk ()  { 
	if fn := WalkHandlers.TentryType; fn != nil { fn(me) }
	me.XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_.Walk()
	me.XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_.Walk()
	me.XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_.Walk()
	me.XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_.Walk()
	me.XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_.Walk()
	me.XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_.Walk()
	me.XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_.Walk()
	me.XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_.Walk()
	me.XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_.Walk()
	me.XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_.Walk()
	me.XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_.Walk()
	me.XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_.Walk()
 }

type XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_ struct {
	Entries []*TentryType `xml:"http://www.w3.org/2005/Atom entry"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_ function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_ instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_ instance.
func (me *XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_; fn != nil { fn(me) }
	for _, x := range me.Entries { x.Walk() }
 }

type TfeedType struct {
	XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_

	XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_

	XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_

	XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_

	XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_

	XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_

	XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_

	XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_

	XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_

	XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_

	XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_

	XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_

}

//	If the WalkHandlers.TfeedType function is not nil (ie. was set by outside code), calls it with this TfeedType instance as the single argument. Then calls the Walk() method on 13/14 embed(s) and 0/0 field(s) belonging to this TfeedType instance.
func (me *TfeedType) Walk ()  { 
	if fn := WalkHandlers.TfeedType; fn != nil { fn(me) }
	me.XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_.Walk()
	me.XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_.Walk()
	me.XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_.Walk()
	me.XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_.Walk()
	me.XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_.Walk()
	me.XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_.Walk()
	me.XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_.Walk()
	me.XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_.Walk()
	me.XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_.Walk()
	me.XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_.Walk()
	me.XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_.Walk()
	me.XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_.Walk()
	me.XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_.Walk()
 }

type XsdGoPkgHasElems_Feed struct {
	Feeds []*TfeedType `xml:"http://www.w3.org/2005/Atom feed"`

}

//	If the WalkHandlers.XsdGoPkgHasElems_Feed function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_Feed instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_Feed instance.
func (me *XsdGoPkgHasElems_Feed) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElems_Feed; fn != nil { fn(me) }
	for _, x := range me.Feeds { x.Walk() }
 }

type XsdGoPkgHasElem_Feed struct {
	Feed *TfeedType `xml:"http://www.w3.org/2005/Atom feed"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Feed function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Feed instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_Feed instance.
func (me *XsdGoPkgHasElem_Feed) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Feed; fn != nil { fn(me) }
	me.Feed.Walk()
 }

type XsdGoPkgHasElem_Entry struct {
	Entry *TentryType `xml:"http://www.w3.org/2005/Atom entry"`

}

//	If the WalkHandlers.XsdGoPkgHasElem_Entry function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_Entry instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_Entry instance.
func (me *XsdGoPkgHasElem_Entry) Walk ()  { 
	if fn := WalkHandlers.XsdGoPkgHasElem_Entry; fn != nil { fn(me) }
	me.Entry.Walk()
 }

//	Provides 37 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
var WalkHandlers = &XsdGoPkgWalkHandlers {}

//	Provides 37 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
type XsdGoPkgWalkHandlers struct {
	TlogoType func (o *TlogoType)
	XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_ func (o *XsdGoPkgHasElems_entrychoicefeedTypeschema_Entry_TentryType_)
	TdateTimeType func (o *TdateTimeType)
	XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_ func (o *XsdGoPkgHasElems_emailchoicepersonTypeschema_Email_TemailType_)
	TlinkType func (o *TlinkType)
	TiconType func (o *TiconType)
	TfeedType func (o *TfeedType)
	XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_ func (o *XsdGoPkgHasElems_iconchoicefeedTypeschema_Icon_TiconType_)
	XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_ func (o *XsdGoPkgHasElems_logochoicefeedTypeschema_Logo_TlogoType_)
	XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_ func (o *XsdGoPkgHasElems_linkchoicefeedTypeschema_Link_TlinkType_)
	XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_ func (o *XsdGoPkgHasElems_generatorchoicefeedTypeschema_Generator_TgeneratorType_)
	TcategoryType func (o *TcategoryType)
	XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_ func (o *XsdGoPkgHasElems_authorchoicefeedTypeschema_Author_TpersonType_)
	TuriType func (o *TuriType)
	XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_ func (o *XsdGoPkgHasElems_rightschoicefeedTypeschema_Rights_TtextType_)
	XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_ func (o *XsdGoPkgHasElems_idchoicefeedTypeschema_Id_TidType_)
	XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_ func (o *XsdGoPkgHasElems_sourcechoiceentryTypeschema_Source_TtextType_)
	XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_ func (o *XsdGoPkgHasElems_updatedchoicefeedTypeschema_Updated_TdateTimeType_)
	XsdGoPkgHasCdata func (o *XsdGoPkgHasCdata)
	TentryType func (o *TentryType)
	XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_ func (o *XsdGoPkgHasElems_namechoicepersonTypeschema_Name_XsdtString_)
	XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_ func (o *XsdGoPkgHasElems_contributorchoicefeedTypeschema_Contributor_TpersonType_)
	TpersonType func (o *TpersonType)
	TgeneratorType func (o *TgeneratorType)
	XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_ func (o *XsdGoPkgHasElems_contentchoiceentryTypeschema_Content_TcontentType_)
	XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_ func (o *XsdGoPkgHasElems_subtitlechoicefeedTypeschema_Subtitle_TtextType_)
	TtextType func (o *TtextType)
	TidType func (o *TidType)
	TcontentType func (o *TcontentType)
	XsdGoPkgHasElem_Entry func (o *XsdGoPkgHasElem_Entry)
	XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_ func (o *XsdGoPkgHasElems_publishedchoiceentryTypeschema_Published_TdateTimeType_)
	XsdGoPkgHasElems_Feed func (o *XsdGoPkgHasElems_Feed)
	XsdGoPkgHasElem_Feed func (o *XsdGoPkgHasElem_Feed)
	XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_ func (o *XsdGoPkgHasElems_titlechoicefeedTypeschema_Title_TtextType_)
	XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_ func (o *XsdGoPkgHasElems_summarychoiceentryTypeschema_Summary_TtextType_)
	XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_ func (o *XsdGoPkgHasElems_urichoicepersonTypeschema_Uri_TuriType_)
	XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_ func (o *XsdGoPkgHasElems_categorychoicefeedTypeschema_Category_TcategoryType_)
}
