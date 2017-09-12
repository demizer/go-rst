package document

import (
	"github.com/demizer/go-rst/pkg/log"
)

// JSON type for rendering the document to JSON.
// Do not initialize this directly. Call JsonRenderer instead.
type JSON struct {
	messages NodeList
	nodes    NodeList

	logConf log.Config
	log.Logger
}

func (j JSON) Bytes() (error, []byte) {
	var tmp NodeList

	tmp.Append(NewSystemMessagesNode())
	tmp.LastNode().(*SystemMessagesNode).NodeList.Append(j.nodes...)

	// pJson, err := json.MarshalIndent(pNodes, "", "    ")
	// if err != nil {
	// return fmt.Errorf("Error Marshalling JSON: %s", err.Error()), nil
	// }
	return nil, nil
}

// JsonRenderer returns the Renderer interface
func JsonRenderer(logConf log.Config, messages, nodes NodeList) Renderer {
	conf := logConf
	conf.Name = "document_json"
	return JSON{
		messages: messages,
		nodes:    nodes,
		logConf:  conf,
		Logger:   log.NewLogger(conf),
	}
}
