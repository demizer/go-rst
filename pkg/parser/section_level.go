package parser

import (
	"github.com/demizer/go-rst/pkg/log"

	doc "github.com/demizer/go-rst/pkg/document"
	mes "github.com/demizer/go-rst/pkg/messages"
)

// sectionLevel is a single section level. sections contains a list of pointers to doc.SectionNode that are dertermined to be a
// section of the level indicated by level. rChar is the rune character that denotes the section level.
type sectionLevel struct {
	rChar    rune
	level    int
	overLine bool               // True if the section level has an overline.
	sections []*doc.SectionNode // Sections matching level.
}

// sectionLevels contains the encountered section levels in order by level.  levels[0] is section level 1 and levels[1] is
// section level 2.  lastSectionNode is a pointer to the lastSectionNode added to levels.
type sectionLevels struct {
	lastSectionNode *doc.SectionNode
	levels          []*sectionLevel

	logConf log.Config
	log.Logger
}

func newSectionLevels(logConf log.Config) *sectionLevels {
	conf := logConf
	conf.Name = "section_level"
	return &sectionLevels{
		logConf: conf,
		Logger:  log.NewLogger(conf),
	}
}

// FindByRune loops through the sectionLevels to find a section using a Rune as the key. If the section is found, a pointer
// to the SectionNode is returned.
func (s *sectionLevels) FindByRune(rChar rune) *sectionLevel {
	for _, sec := range s.levels {
		if sec.rChar == rChar {
			return sec
		}
	}
	return nil
}

// Add determines if the underline rune in the sec argument matches any existing sectionLevel in sectionLevels. Add also
// checks the section level ordering is correct and returns a mes.SectionErrorTitleLevelInconsistent ParserMessage if inconsistencies
// are found.
func (s *sectionLevels) Add(sec *doc.SectionNode) (err mes.MessageType) {
	level := 1
	secLvl := s.FindByRune(sec.UnderLine.Rune)

	// Local function for creating a sectionLevel
	var newSectionLevel = func() {
		var oLine bool
		if sec.OverLine != nil {
			oLine = true
		}
		s.Msgr("Creating new sectionLevel", "level", level)
		secLvl = &sectionLevel{
			rChar: sec.UnderLine.Rune,
			level: level, overLine: oLine,
		}
		s.levels = append(s.levels, secLvl)
		// secLvl.sections = append(secLvl.sections, sec)
	}

	if secLvl == nil {
		if s.lastSectionNode != nil {
			// Check if the provisional level of sec is already in sectionLevels; if it is and the adornment
			// characters don't match, then we have an inconsistent level error.
			level = s.lastSectionNode.Level + 1
			nextLevel := s.SectionLevelByLevel(level)
			if nextLevel != nil &&
				nextLevel.rChar != sec.UnderLine.Rune {
				return mes.SectionErrorTitleLevelInconsistent
			}
		} else {
			level = len(s.levels) + 1
		}
		newSectionLevel()
	} else {
		if secLvl.overLine && sec.OverLine == nil ||
			!secLvl.overLine && sec.OverLine != nil {
			// If sec has an OverLine, but the matching sectionLevel with the same Rune as sec does not have an
			// OverLine, then they are not in the same sectionLevel, and visa versa.
			level = len(s.levels) + 1
			newSectionLevel()
		} else {
			s.Msgr("using sectionLevel", "sectionLevel", secLvl.level)
			level = secLvl.level
		}
	}

	secLvl.sections = append(secLvl.sections, sec)
	sec.Level = level
	s.lastSectionNode = sec
	return
}

// SectionLevelByLevel returns a pointer to a sectionLevel of level level. Nil is returned if l is greater than the number of
// section levels encountered.
func (s *sectionLevels) SectionLevelByLevel(level int) *sectionLevel {
	if level > len(s.levels) {
		return nil
	}
	return (s.levels)[level-1]
}

// LastSectionByLevel returns a pointer to the last section encountered by level.
func (s *sectionLevels) LastSectionByLevel(level int) (sec *doc.SectionNode) {
exit:
	for i := len(s.levels) - 1; i >= 0; i-- {
		if (s.levels)[i].level != level {
			continue
		}
		for j := len((s.levels)[i].sections) - 1; j >= 0; j-- {
			sec = (s.levels)[i].sections[j]
			if sec.Level == level {
				s.Msgr("found sectionLevel", "sectionLevel", sec.Level)
				break exit
			}
		}
	}
	return
}
