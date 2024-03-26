
'use strict'

const { sep } = require('path')
const Base = require('./base')

const TYPES = {
  address:  'interface{}',
  chunk:    'string',
  elements: '[]interface{}', 
  index:    'int',
  max:      'int'
}

class Builder extends Base {
  _tab() {
    return '\t'
  }

  _initBuffer(pathname) {
    let namespace = pathname.split(sep)
    namespace.pop()
    return 'package ' + namespace.join('/') + '\n'
  }

  _quote(string) {
    let quoted = string.replace(/\\/g, '\\\\')
                       .replace(/"/g, '\\"')
                       .replace(/\n/g, '\\n')
                       .replace(/\r/g, '\\r')
    return '"' + quoted + '"'
  }

  comment(lines) {
    return ['/*'].concat(lines.map((line) => ' * ' + line)).concat([' */'])
  }

  package_(name, actions, block) {
    this._grammarName = name

    this._newBuffer('go', 'actions')
    this._template('go', 'actions.go', { actions })

    this._newBuffer('go', 'position')
    this._template('go', 'position.go')

    block()
  }

  syntaxNodeClass_() {
    let name = 'Node'

    this._newBuffer('go', name)
    this._template('go', 'node.go', { name })

    return name
  }

  grammarModule_(block) {
    this._newBuffer('go', 'grammar')

    this._line('import "regexp"')
    this._newline()

    this._line('var FAILURE = new(Node)')
    this._newline()

    this._line('type Grammar struct {', false)
    this._indent(() => {
      this._line('Buffer []rune')
      this._line('Position')
      this._line('Actions actions')
      this._line('cache map[string]*Entry')
    })
    this._line('}', false)
    this._newline()

    block()
  }

  compileRegex_(charClass, name) {
    let regex = charClass.regex,
        source = regex.source

    this.assign_('var ' + name, 'regexp.MustCompile(' + this._quote(source) + ')')
    charClass.constName = name
  }

  parserClass_(root) {
    let grammar = this._quote(this._grammarName)
    this._newBuffer('go', 'parser')
    this._template('go', 'parser.go', { grammar, root })
  }

  class_(name, parent, block) {
    this._newline()
    this._line('type ' + name + ' ' + parent + ' {', false)
    this._scope(block, name, parent)
    this._line('}', false)
    this._newline()
  }

  constructor_(args, block) {
    this._line('func new' + this._currentScope.name + '() *' + this._currentScope.name + ' {', false)
    this._indent(() => {
      this._line('return &' + this._currentScope.name + '{}')
    })
    this._line('}', false)
  }

  method_(name, args, block) {
    let argsDecl = ['g *Grammar'].concat(args.map((arg) => arg + ' ' + TYPES[arg])).join(', ')
    this._newline()
    this._line('func (n *' + this._currentScope.name + ') ' + name + '(' + argsDecl + ') ' + TYPES.address + ' {', false)
    this._scope(block)
    this._line('}', false)
  }

  cache_(name, block) {
    let address = this.localVar_('start'), 
        cacheKey = this._quote(name)

    this.assign_('node', 'FAILURE')

    this.if_('g.cache == nil', () => {
      this.assign_('g.cache', 'make(map[string]*Entry)')
    })

    this.if_('entry, ok := g.cache[' + cacheKey + ']; ok', () => {
      this.assign_('node, g.Position', 'entry.Node, entry.Position')
    }, () => {
      block(address)
      this._line('g.cache[' + cacheKey + '] = &Entry{Node: ' + address + ', Position: g.Position}')
    })

    this._return(address)
  }

  attribute_(name, value) {
    this._line('n.' + name + ' = ' + value)
  }

  localVars_(vars) {
    let names = {}
    for (let name in vars)
      names[name] = this.localVar_(name, vars[name])
    return names
  }

  localVar_(name, value) {
    let varName = this._varName(name)

    if (value === undefined) value = this.nullNode_()
    this.assign_(TYPES[name] + ' ' + varName, value)

    return varName
  }

  chunk_(length) {
    let input = 'string(g.Buffer)', 
        ofs = 'g.Position.Offset',
        temp = this.localVars_({ chunk: this.null_(), max: ofs + ' + ' + length })

    this.if_(temp.max + ' <= len(g.Buffer)', () => {
      this.assign_(temp.chunk, input + '[' + ofs + ':' + temp.max + ']') 
    })

    return temp.chunk
  }

  syntaxNode_(address, start, end, elements, action, nodeClass) {
    let args = [this.chunk_(end + ' - ' + start), start]

    if (elements) {
      args.push(elements)
    } else {
      args.push('nil')
    }

    if (action) {
      args.push('g.Actions.' + action)
    } else {
      args.push('nil')
    }

    if (nodeClass) {
      args.push('new' + nodeClass + '()')
    } else {
      args.push('new(Node)')
    }

    this.assign_(address, 'makeNode(' + args.join(', ') + ')')
    this.assign_('g.Position.Offset', end)
  }

  ifNode_(address, block, else_) {
    this._conditional('if ' + address + ' != nil', block, else_)
  }

  unlessNode_(address, block, else_) {
    this._conditional('if ' + address + ' == nil', block, else_)
  }
  
  ifNull_(elements, block, else_) {
    this.if_(elements + ' == nil', block, else_)
  }

  extendNode_(address, nodeType) {
    // Not currently supported. Could potentially implement via anonymous composition.
  }

  failure_(address, expected) {
    let rule = this._quote(this._grammarName + '::' + this._ruleName)
    expected = this._quote(expected)

    this.assign_(address, 'FAILURE')

    this.if_('g.Position.furthestFailure() <= g.Position.Offset', () => {
      this.assign_('g.Position.furthestFailure()', 'g.Position.Offset')
      this._line('g.Position.expected = []Expected{Expected{' + rule + ', ' + expected + '}}')
    }, () => {
      this.if_('g.Position.Offset == g.Position.furthestFailure()', () => {
        this.append_('g.Position.expected', 'Expected{' + rule + ', ' + expected + '}')
      })
    })    
  }

  jump_(address, rule) {
    this.assign_(address, 'g.' + rule + '()')
  }

  _conditional(kwd, condition, block, else_) {
    this._line(kwd + ' {', false)
    this._indent(block)
    if (else_) {
      this._line('} else {', false)
      this._indent(else_)
    }
    this._line('}', false)
  }

  if_(condition, block, else_) {
    this._conditional('if ' + condition, block, else_)
  }

  loop_(block) {
    this._conditional('for {', block)
  }

  break_() {
    this._line('break')
  }

  sizeInRange_(address, [min, max]) {
    if (max === -1) {
      return 'len(' + address + ') >= ' + min
    } else if (max === 0) {
      return 'len(' + address + ') == ' + min
    } else {
      return 'len(' + address + ') >= ' + min + ' && len(' + address + ') <= ' + max
    }
  }

  stringMatch_(expression, string) {
    return expression + ' == ' + this._quote(string)
  }

  stringMatchCI_(expression, string) {
    return 'strings.EqualFold(' + expression + ', ' + this._quote(string) + ')'
  }

  regexMatch_(regex, string) {
    return regex + '.MatchString(' + string + ')'
  }

  arrayLookup_(expression, offset) {
    return expression + '[' + offset + ']'
  }

  append_(list, value) {
    this._line(list + ' = append(' + list + ', ' + value + ')')
  }

  hasChars_() {
    return 'g.Position.Offset < len(g.Buffer)'
  }

  nullNode_() {
    return 'FAILURE'
  }

  offset_() {
    return 'g.Position.Offset'
  }

  emptyList_() {
    return 'nil'
  }

  _emptyString() {
    return '""'
  }

  null_() {
    return 'nil'
  }
}

module.exports = Builder
