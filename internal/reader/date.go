package reader

import (
	"encoding/xml"
	"time"
)

type Date time.Time

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(d).Format(time.RFC1123Z), start)
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := dec.DecodeElement(&v, &start); err != nil {
		return err
	}

	t, err := time.Parse(time.RFC1123Z, v)
	if err != nil {
		return err
	}

	*d = Date(t)

	return nil
}
