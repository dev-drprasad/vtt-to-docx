// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_Shape struct {
	// Use Background Fill
	UseBgFillAttr *bool
	// Non-Visual Properties for a Shape
	NvSpPr *CT_ShapeNonVisual
	SpPr   *dml.CT_ShapeProperties
	// Shape Style
	Style *dml.CT_ShapeStyle
	// Shape Text Body
	TxBody *dml.CT_TextBody
	ExtLst *CT_ExtensionListModify
}

func NewCT_Shape() *CT_Shape {
	ret := &CT_Shape{}
	ret.NvSpPr = NewCT_ShapeNonVisual()
	ret.SpPr = dml.NewCT_ShapeProperties()
	return ret
}

func (m *CT_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UseBgFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useBgFill"},
			Value: fmt.Sprintf("%d", b2i(*m.UseBgFillAttr))})
	}
	e.EncodeToken(start)
	senvSpPr := xml.StartElement{Name: xml.Name{Local: "p:nvSpPr"}}
	e.EncodeElement(m.NvSpPr, senvSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "p:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "p:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.TxBody != nil {
		setxBody := xml.StartElement{Name: xml.Name{Local: "p:txBody"}}
		e.EncodeElement(m.TxBody, setxBody)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvSpPr = NewCT_ShapeNonVisual()
	m.SpPr = dml.NewCT_ShapeProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "useBgFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseBgFillAttr = &parsed
			continue
		}
	}
lCT_Shape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "nvSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "nvSpPr"}:
				if err := d.DecodeElement(m.NvSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "style"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "style"}:
				m.Style = dml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "txBody"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "txBody"}:
				m.TxBody = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxBody, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Shape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Shape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Shape and its children
func (m *CT_Shape) Validate() error {
	return m.ValidateWithPath("CT_Shape")
}

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (m *CT_Shape) ValidateWithPath(path string) error {
	if err := m.NvSpPr.ValidateWithPath(path + "/NvSpPr"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.TxBody != nil {
		if err := m.TxBody.ValidateWithPath(path + "/TxBody"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
