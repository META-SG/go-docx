/*
   Copyright (c) 2024 mabiao0525 (马飚)

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package docx

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"
)

// SectPr show the properties of the document, like paper size
type SectPr struct {
	XMLName         xml.Name         `xml:"w:sectPr,omitempty"` // properties of the document, including paper size
	PgSz            *PgSz            `xml:"w:pgSz,omitempty"`
	HeaderReference *HeaderReference `xml:"w:headerReference,omitempty"`
	FooterReference *FooterReference `xml:"w:footerReference,omitempty"`
	Type            *Type            `xml:"w:type,omitempty"`
	PgMar           *PgMar           `xml:"w:pgMar,omitempty"`
	Cols            *Cols            `xml:"w:cols,omitempty"`
	DocGrid         *DocGrid         `xml:"w:docGrid,omitempty"`
}

// PgSz show the paper size
type PgSz struct {
	W int `xml:"w:w,attr"` // width of paper
	H int `xml:"w:h,attr"` // high of paper
}

type HeaderReference struct {
	Id   string `xml:"r:id,attr"`
	Type string `xml:"w:type,attr"`
}

func (pgsz *HeaderReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			pgsz.Id = attr.Value
		case "type":
			pgsz.Type = attr.Value
		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}

type FooterReference struct {
	Id   string `xml:"r:id,attr"`
	Type string `xml:"w:type,attr"`
}

func (pgsz *FooterReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			pgsz.Id = attr.Value
		case "type":
			pgsz.Type = attr.Value
		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}

type Type struct {
	Val string `xml:"w:val,attr"`
}

func (pgsz *Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "val":
			pgsz.Val = attr.Value
		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}

type PgMar struct {
	Top    string `xml:"w:top,attr"`
	Right  string `xml:"w:right,attr"`
	Bottom string `xml:"w:bottom,attr"`
	Left   string `xml:"w:left,attr"`
	Header string `xml:"w:header,attr"`
	Footer string `xml:"w:footer,attr"`
	Gutter string `xml:"w:gutter,attr"`
}

func (pgsz *PgMar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "top":
			pgsz.Top = attr.Value
		case "right":
			pgsz.Right = attr.Value
		case "bottom":
			pgsz.Bottom = attr.Value
		case "left":
			pgsz.Left = attr.Value
		case "header":
			pgsz.Header = attr.Value
		case "footer":
			pgsz.Footer = attr.Value
		case "gutter":
			pgsz.Gutter = attr.Value

		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}

type Cols struct {
	Space string `xml:"w:space,attr"`
	Num   string `xml:"w:num,attr"`
}

func (pgsz *Cols) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "space":
			pgsz.Space = attr.Value
		case "num":
			pgsz.Num = attr.Value
		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}

type DocGrid struct {
	LinePitch string `xml:"w:linePitch,attr"`
	CharSpace string `xml:"w:charSpace,attr"`
}

func (pgsz *DocGrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "linePitch":
			pgsz.LinePitch = attr.Value
		case "charSpace":
			pgsz.CharSpace = attr.Value
		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}

// UnmarshalXML ...
func (sect *SectPr) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if tt, ok := t.(xml.StartElement); ok {
			switch tt.Name.Local {
			case "pgSz":
				var value PgSz
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.PgSz = &value
			case "headerReference":
				var value HeaderReference
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.HeaderReference = &value
			case "footerReference":
				var value FooterReference
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.FooterReference = &value
			case "type":
				var value Type
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.Type = &value
			case "pgMar":
				var value PgMar
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.PgMar = &value
			case "cols":
				var value Cols
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.Cols = &value
			case "docGrid":
				var value DocGrid
				err = d.DecodeElement(&value, &tt)
				if err != nil && !strings.HasPrefix(err.Error(), "expected") {
					return err
				}
				sect.DocGrid = &value
			default:
				err = d.Skip() // skip unsupported tags
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// UnmarshalXML ...
func (pgsz *PgSz) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "w":
			pgsz.W, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
		case "h":
			pgsz.H, err = strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
		default:
			// ignore other attributes now
		}
	}
	// Consume the end element
	_, err = d.Token()
	return err
}
