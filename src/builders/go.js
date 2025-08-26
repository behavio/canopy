'use strict'

const { basename, join } = require('path')
const Base = require('./base')

const toCamelCase = (string) => {
  return string.split('_').map((word) => {
    if (word.length === 0) return ''
    return word.charAt(0).toUpperCase() + word.slice(1)
  }).join('')
}

class Builder extends Base {
  _quote (string) {
    return JSON.stringify(string)
  }

  comment (lines) {
    return lines.map((line) => '// ' + line)
  }

  package_ (name, actions, block) {
    this._grammarName = basename(name)

    const actionNames = actions.map(toCamelCase)

    this._currentBuffer = join(this._outputPath, 'parser.go')
    this._buffers.set(this._currentBuffer, '')
    this._template('go', 'parser.go.tpl', { name: this._grammarName })

    this._currentBuffer = join(this._outputPath, 'treenode.go')
    this._buffers.set(this._currentBuffer, '')
    this._template('go', 'treenode.go.tpl', { name: this._grammarName })

    this._currentBuffer = join(this._outputPath, 'actions.go')
    this._buffers.set(this._currentBuffer, '')
    this._template('go', 'actions.go.tpl', {
      name: this._grammarName,
      actions: actionNames
    })

    this._currentBuffer = join(this._outputPath, 'go.mod')
    this._buffers.set(this._currentBuffer, '')
    this._line('module ' + this._grammarName, false)
    this._line('go 1.22.0', false)
  }

  grammarModule_ (block) {}
  parserClass_ (root) {}
  class_ (name, parent, block) {}
  constructor_ (args, block) {}
  function_ (name, args, block) {}
  method_ (name, args, block) {}
  cache_ (name, block) {}
  attribute_ (name, value) {}
  localVars_ (vars) { return {} }
  localVar_ (name, value) {}
  chunk_ (length) {}
  syntaxNode_ (address, start, end, elements, action, nodeClass) {}
  ifNode_ (address, block, else_) {}
  unlessNode_ (address, block, else_) {}
  ifNull_ (elements, block, else_) {}
  extendNode_ (address, nodeType) {}
  failure_ (address, expected) {}
  jump_ (address, rule) {}
  if_ (condition, block, else_) {}
  loop_ (block) {}
  break_ () {}
  sizeInRange_ (address, [min, max]) {}
  stringMatch_ (expression, string) {}
  stringMatchCI_ (expression, string) {}
  regexMatch_ (regex, string) {}
  arrayLookup_ (expression, offset) {}
  append_ (list, value, index) {}
  hasChars_ () {}
  nullNode_ () {}
  offset_ () {}
  emptyList_ (size) {}
  _emptyString () {}
  null_ () {}
}

module.exports = Builder
