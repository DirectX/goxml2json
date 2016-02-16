package xml2json

import (
	"encoding/xml"
	"io"
)

// A Decoder reads and decodes XML objects from an input stream.
type Decoder struct {
	r          io.Reader
	err        error
	attrPrefix string
}

type element struct {
	parent *element
	n      *Node
	label  string
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r, attrPrefix: "-"}
}

// NewDecoder returns a new decoder that reads from r with custom attribute prefiх.
func NewDecoderWithAttrPrefix(r io.Reader, attrPrefix string) *Decoder {
	return &Decoder{r: r, attrPrefix: attrPrefix}
}

// Decode reads the next JSON-encoded value from its
// input and stores it in the value pointed to by v.
func (dec *Decoder) Decode(root *Node) error {
	xmlDec := xml.NewDecoder(dec.r)

	// Create first element from the root node
	elem := &element{
		parent: nil,
		n:      root,
	}

	for {
		t, _ := xmlDec.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			// Build new a new current element and link it to its parent
			elem = &element{
				parent: elem,
				n:      &Node{},
				label:  se.Name.Local,
			}

			// Extract attributes as children
			for _, a := range se.Attr {
				elem.n.AddChild(dec.attrPrefix+a.Name.Local, &Node{Data: a.Value})
			}
		case xml.CharData:
			// Extract XML data (if any)
			elem.n.Data = string(xml.CharData(se))
		case xml.EndElement:
			// And add it to its parent list
			if elem.parent != nil {
				elem.parent.n.AddChild(elem.label, elem.n)
			}

			// Then change the current element to its parent
			elem = elem.parent
		}
	}

	return nil
}
