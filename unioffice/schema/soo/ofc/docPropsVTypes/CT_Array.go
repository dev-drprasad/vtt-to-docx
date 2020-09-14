// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/unidoc/unioffice"
)

type CT_Array struct {
	LBoundsAttr  int32
	UBoundsAttr  int32
	BaseTypeAttr ST_ArrayBaseType
	Variant      []*Variant
	I1           []int8
	I2           []int16
	I4           []int32
	Int          []int32
	Ui1          []uint8
	Ui2          []uint16
	Ui4          []uint32
	Uint         []uint32
	R4           []float32
	R8           []float64
	Decimal      []float64
	Bstr         []string
	Date         []time.Time
	Bool         []bool
	Error        []string
	Cy           []string
}

func NewCT_Array() *CT_Array {
	ret := &CT_Array{}
	ret.BaseTypeAttr = ST_ArrayBaseType(1)
	return ret
}

func (m *CT_Array) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lBounds"},
		Value: fmt.Sprintf("%v", m.LBoundsAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uBounds"},
		Value: fmt.Sprintf("%v", m.UBoundsAttr)})
	attr, err := m.BaseTypeAttr.MarshalXMLAttr(xml.Name{Local: "baseType"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.Variant != nil {
		sevariant := xml.StartElement{Name: xml.Name{Local: "vt:variant"}}
		for _, c := range m.Variant {
			e.EncodeElement(c, sevariant)
		}
	}
	if m.I1 != nil {
		sei1 := xml.StartElement{Name: xml.Name{Local: "vt:i1"}}
		for _, c := range m.I1 {
			e.EncodeElement(c, sei1)
		}
	}
	if m.I2 != nil {
		sei2 := xml.StartElement{Name: xml.Name{Local: "vt:i2"}}
		for _, c := range m.I2 {
			e.EncodeElement(c, sei2)
		}
	}
	if m.I4 != nil {
		sei4 := xml.StartElement{Name: xml.Name{Local: "vt:i4"}}
		for _, c := range m.I4 {
			e.EncodeElement(c, sei4)
		}
	}
	if m.Int != nil {
		seint := xml.StartElement{Name: xml.Name{Local: "vt:int"}}
		for _, c := range m.Int {
			e.EncodeElement(c, seint)
		}
	}
	if m.Ui1 != nil {
		seui1 := xml.StartElement{Name: xml.Name{Local: "vt:ui1"}}
		for _, c := range m.Ui1 {
			e.EncodeElement(c, seui1)
		}
	}
	if m.Ui2 != nil {
		seui2 := xml.StartElement{Name: xml.Name{Local: "vt:ui2"}}
		for _, c := range m.Ui2 {
			e.EncodeElement(c, seui2)
		}
	}
	if m.Ui4 != nil {
		seui4 := xml.StartElement{Name: xml.Name{Local: "vt:ui4"}}
		for _, c := range m.Ui4 {
			e.EncodeElement(c, seui4)
		}
	}
	if m.Uint != nil {
		seuint := xml.StartElement{Name: xml.Name{Local: "vt:uint"}}
		for _, c := range m.Uint {
			e.EncodeElement(c, seuint)
		}
	}
	if m.R4 != nil {
		ser4 := xml.StartElement{Name: xml.Name{Local: "vt:r4"}}
		for _, c := range m.R4 {
			e.EncodeElement(c, ser4)
		}
	}
	if m.R8 != nil {
		ser8 := xml.StartElement{Name: xml.Name{Local: "vt:r8"}}
		for _, c := range m.R8 {
			e.EncodeElement(c, ser8)
		}
	}
	if m.Decimal != nil {
		sedecimal := xml.StartElement{Name: xml.Name{Local: "vt:decimal"}}
		for _, c := range m.Decimal {
			e.EncodeElement(c, sedecimal)
		}
	}
	if m.Bstr != nil {
		sebstr := xml.StartElement{Name: xml.Name{Local: "vt:bstr"}}
		for _, c := range m.Bstr {
			e.EncodeElement(c, sebstr)
		}
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "vt:date"}}
		for _, c := range m.Date {
			e.EncodeElement(c, sedate)
		}
	}
	if m.Bool != nil {
		sebool := xml.StartElement{Name: xml.Name{Local: "vt:bool"}}
		for _, c := range m.Bool {
			e.EncodeElement(c, sebool)
		}
	}
	if m.Error != nil {
		seerror := xml.StartElement{Name: xml.Name{Local: "vt:error"}}
		for _, c := range m.Error {
			e.EncodeElement(c, seerror)
		}
	}
	if m.Cy != nil {
		secy := xml.StartElement{Name: xml.Name{Local: "vt:cy"}}
		for _, c := range m.Cy {
			e.EncodeElement(c, secy)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Array) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BaseTypeAttr = ST_ArrayBaseType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "uBounds" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.UBoundsAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "baseType" {
			m.BaseTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "lBounds" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.LBoundsAttr = int32(parsed)
			continue
		}
	}
lCT_Array:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "variant"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "variant"}:
				tmp := NewVariant()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Variant = append(m.Variant, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i1"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i1"}:
				var tmp int8
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I1 = append(m.I1, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i2"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i2"}:
				var tmp int16
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I2 = append(m.I2, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i4"}:
				var tmp int32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I4 = append(m.I4, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "int"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "int"}:
				var tmp int32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Int = append(m.Int, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui1"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui1"}:
				var tmp uint8
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui1 = append(m.Ui1, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui2"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui2"}:
				var tmp uint16
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui2 = append(m.Ui2, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui4"}:
				var tmp uint32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui4 = append(m.Ui4, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "uint"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "uint"}:
				var tmp uint32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Uint = append(m.Uint, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r4"}:
				var tmp float32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.R4 = append(m.R4, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r8"}:
				var tmp float64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.R8 = append(m.R8, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "decimal"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "decimal"}:
				var tmp float64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Decimal = append(m.Decimal, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bstr"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bstr = append(m.Bstr, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "date"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "date"}:
				var tmp time.Time
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Date = append(m.Date, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bool"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bool"}:
				var tmp bool
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bool = append(m.Bool, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "error"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "error"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Error = append(m.Error, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "cy"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "cy"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Cy = append(m.Cy, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Array %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Array
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Array and its children
func (m *CT_Array) Validate() error {
	return m.ValidateWithPath("CT_Array")
}

// ValidateWithPath validates the CT_Array and its children, prefixing error messages with path
func (m *CT_Array) ValidateWithPath(path string) error {
	if m.BaseTypeAttr == ST_ArrayBaseTypeUnset {
		return fmt.Errorf("%s/BaseTypeAttr is a mandatory field", path)
	}
	if err := m.BaseTypeAttr.ValidateWithPath(path + "/BaseTypeAttr"); err != nil {
		return err
	}
	for i, v := range m.Variant {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Variant[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, s := range m.Error {
		if !ST_ErrorPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Error[%d] must match '%s' (have %v)`, path, i, ST_ErrorPatternRe, s)
		}
	}
	for i, s := range m.Cy {
		if !ST_CyPatternRe.MatchString(s) {
			return fmt.Errorf(`%s/m.Cy[%d] must match '%s' (have %v)`, path, i, ST_CyPatternRe, s)
		}
	}
	return nil
}
