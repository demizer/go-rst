package document

import (
	"bytes"
	"encoding/json"
)

// NodeList is a list of parser nodes that implement Node.
type NodeList []Node

func (l *NodeList) Append(n ...Node) {
	for _, node := range n {
		*l = append(*l, node)
	}
}

// last returns the last item added to the slice
func (l *NodeList) LastNode(n ...Node) Node { return (*l)[len(*l)-1] }

// MarshalJSON satisfies the Marshaler interface.
func (l *NodeList) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("[")
	length := len(*l)
	count := 0
	for _, value := range *l {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(string(jsonValue))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}
