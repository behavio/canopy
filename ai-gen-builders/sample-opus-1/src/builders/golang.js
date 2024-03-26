'use strict'

const Base = require('./base')

class Builder extends Base {
  constructor (...args) {
    super(...args)
  } 

  _tab () {
    return '\t'  
  }

  comment (lines) {
    return lines.map((line) => '// ' + line)
  }

  package_ (name, actions, block) {
    const pkgName = name.split('.').pop()
    this.assign_('package', pkgName)
    this._newline()
    block()
  }

  syntaxNodeClass_ () {
    return 'TreeNode'
  }
  
  grammarModule_ (block) {
    this._write(`
type TreeNode struct {
\tText    string
\tOffset  int
\tElements []TreeNode
\tNodeType string
}

func (n *TreeNode) String() string {
\treturn n.Text
}
    `)
    block()    
  }

  parserClass_ (root) {
    this._write(`
func Parse(input string) (*TreeNode, error) {
\tp := &parser{Buffer: input}
\tp.Init()
\tif err := p.Parse(); err != nil {
\t\treturn nil, err
\t}
\treturn p.`+ root +`(), nil
}
    `)
  }

  compileRegex_ (charClass, name) {
    let regex  = charClass.regex,
        source = regex.source.replace(/^\^/, '`')

    this.assign_(name, '`'+source+'`')
    charClass.constName = name
  }

  stringMatch_ (expression, string) {
    return expression + ' == "' + string + '"'
  }

  regexMatch_ (regex, string) {
    return string + ' =~ ' + regex  
  }

  localVar_ (name, value) {
    this.assign_(name, value)
    return name
  }

  localVars_ (vars) {
    let names = {}
    for (let name in vars)
      names[name] = this.localVar_(name, vars[name])
    return names
  }

  // More builder methods will be implemented here
}

module.exports = Builder  
