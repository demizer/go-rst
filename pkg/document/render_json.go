package document

import (
	"encoding/json"
	"fmt"

	"github.com/demizer/go-rst/pkg/log"
)

// JSON type for rendering the document to JSON.
// Do not initialize this directly. Call JsonRenderer instead.
type JSON struct {
	Messages *NodeList
	Nodes    *NodeList

	logConf log.Config
	log.Logger
}

func (j JSON) Bytes() ([]byte, error) {
	var tmp NodeList

	tmp.Append(NewSystemMessagesNode())
	tmp.LastNode().(*SystemMessagesNode).NodeList.Append(*j.Messages...)
	tmp.Append(*j.Nodes...)

	pJson, err := json.MarshalIndent(tmp, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Error Marshalling JSON: %s", err.Error())
	}
	return pJson, nil
}

// JsonRenderer returns the Renderer interface
func JsonRenderer(logConf log.Config, messages, nodes *NodeList) Renderer {
	conf := logConf
	conf.Name = "document_json"
	ml := make(NodeList, 0)
	return JSON{
		Messages: &ml,
		Nodes:    nodes,
		logConf:  conf,
		Logger:   log.NewLogger(conf),
	}
}
