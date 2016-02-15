package logfmt

import (
	"bytes"
	"errors"
	"testing"

	"github.com/Sirupsen/logrus"
)

// Not a real test, this skews the test coverage percentage. Precise testing of logging output is not important in this
// instance.
func TestPrintForceColored(t *testing.T) {
	var buf bytes.Buffer
	tf := &TextFormatter{ForceColors: true}
	var log2 = &logrus.Logger{
		Out:       &buf,
		Formatter: tf,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	log2.Debug("Yep")
	log2.Info("Yep")
	tf.FullTimestamp = true
	log2.Warn("Yep")
	log2.Error("Yep")
	tf.DisableColors = true
	log2.WithFields(logrus.Fields{
		"aField":   1,
		"anError":  errors.New("foobar"),
		"anError2": errors.New(`/foobar x, y "`),
	}).Info("Yep")
	tf.DisableColors = false
	log2.WithFields(logrus.Fields{
		"time":    1,
		"msg":     "test",
		"level":   "1",
		"anError": errors.New(`/foobar x, y "`),
	}).Error("Yep")
	// log2.Fatal("Yep") // Not sure how to test this yet
	// fmt.Println("buf:", buf.String())
}
