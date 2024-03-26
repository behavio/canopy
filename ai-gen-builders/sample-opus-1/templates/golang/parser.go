package {{{name}}}

import (
  "regexp"
)

{{{classes}}}

type parser struct {
  Buffer string
  Offset int
  Failure int
  Expected []string
}

func (p *parser) Init() {
  p.Offset = 0
  p.Failure = -1
  p.Expected = []string{}
}

func (p *parser) Parse() error {
  node := p.parse_{{{root}}}()
  if node == nil || p.Offset != len(p.Buffer) {
    // TODO: Generate detailed error message
    return fmt.Errorf("Parsing failed at offset %d", p.Failure)
  }
  return nil    
}

func (p *parser) Expect(msg string) {
  p.Expected = append(p.Expected, msg)
}

{{{methods}}}
