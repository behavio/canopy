'use strict';

const { basename, join } = require('path');
const Base = require('./base');

const toPascalCase = (source) => {
  return source
    .split(/[^A-Za-z0-9]+/)
    .filter(Boolean)
    .map((part) => {
      return part.charAt(0).toUpperCase() + part.slice(1);
    })
    .join('');
};

const toPackageName = (source) => {
  let cleaned = source.replace(/[^A-Za-z0-9]+/g, '').toLowerCase();
  if (!cleaned) cleaned = 'grammar';
  if (!cleaned.endsWith('parser')) cleaned += 'parser';
  return cleaned;
};

const TYPES = {
  address: 'TreeNode',
  index: 'int',
  elements: '[]TreeNode',
  chunk: 'string',
  max: 'int',
  cache: 'map[int]cacheEntry',
};

class Builder extends Base {
  constructor(...args) {
    super(...args);
    this._parserImports = new Set();
    this._currentClass = null;
  }

  _tab() {
    return '    ';
  }

  _line(source) {
    super._line(source, false);
  }

  _quote(string) {
    return JSON.stringify(string);
  }

  comment(lines) {
    return lines.map((line) => '// ' + line);
  }

  package_(name, actions, block) {
    this._grammarName = name;
    this._baseName = basename(this._outputPath);
    this._packageName = toPackageName(this._baseName);
    this._structName = toPascalCase(this._baseName) + 'Parser';
    this._actionMap = new Map();
    this._actionNames = actions.map((actionName) => {
      const methodName = toPascalCase(actionName);
      this._actionMap.set(actionName, methodName);
      return methodName;
    });
    this._parserImports = new Set(['fmt']);

    this._currentBuffer = join(this._outputPath, 'parser.go');
    this._buffers.set(this._currentBuffer, '');
    this._writePrelude();

    this._currentBuffer = join(this._outputPath, 'treenode.go');
    this._buffers.set(this._currentBuffer, '');
    this._template('go', 'treenode.go.tpl', { name: this._packageName });

    this._currentBuffer = join(this._outputPath, 'actions.go');
    this._buffers.set(this._currentBuffer, '');
    this._template('go', 'actions.go.tpl', {
      name: this._packageName,
      actions: this._actionNames,
    });

    this._currentBuffer = join(this._outputPath, 'go.mod');
    this._buffers.set(this._currentBuffer, '');
    this._line('module ' + this._packageName);
    this._line('go 1.22.0');

    this._currentBuffer = join(this._outputPath, 'parser.go');
    block();
  }

  _writePrelude() {
    this._line('package ' + this._packageName);
    this._newline();
    this._line('import (');
    this._indent(() => {
      this._line('// __IMPORTS__');
    });
    this._line(')');
    this._newline();

    this._line('type NodeExtender func(TreeNode) TreeNode');
    this._newline();

    this._line('type expectation struct {');
    this._indent(() => {
      this._line('rule string');
      this._line('expected string');
    });
    this._line('}');
    this._newline();

    this._line('type failureState struct {');
    this._indent(() => {
      this._line('offset int');
      this._line('expected []expectation');
    });
    this._line('}');
    this._newline();

    this._line('type cacheEntry struct {');
    this._indent(() => {
      this._line('node TreeNode');
      this._line('offset int');
    });
    this._line('}');
    this._newline();

    this._line('type ParseError struct {');
    this._indent(() => {
      this._line('Input string');
      this._line('Offset int');
      this._line('Line int');
      this._line('Column int');
      this._line('Expected []expectation');
      this._line('Message string');
    });
    this._line('}');
    this._newline();

    this._line('func (e *ParseError) Error() string {');
    this._indent(() => {
      this._line('return e.Message');
    });
    this._line('}');
    this._newline();

    this._line('type ' + this._structName + ' struct {');
    this._indent(() => {
      this._line('input []rune');
      this._line('inputString string');
      this._line('actions Actions');
      this._line('types map[string]NodeExtender');
      this._line('offset int');
      this._line('cache map[string]map[int]cacheEntry');
      this._line('failure failureState');
      this._line('actionErr error');
    });
    this._line('}');
    this._newline();
  }

  syntaxNodeClass_() {
    return 'Node';
  }

  grammarModule_(block) {
    this._newline();
    block();
  }

  parserClass_(root) {
    this._writeParserHelpers(root);
  }

  class_(name, parent, block) {
    this._currentClass = {
      name,
      fields: new Map(),
      assignments: [],
    };
    block();
    this._appendNodeClass(this._currentClass);
    this._currentClass = null;
  }

  attributes_(iterable) {
    if (!this._currentClass) return;
    for (let name of iterable) {
      const fieldName = toPascalCase(name);
      if (!fieldName) continue;
      this._currentClass.fields.set(name, fieldName);
    }
  }

  constructor_(args, block) {
    if (!this._currentClass) return;
    block();
  }

  attribute_(name, value) {
    if (!this._currentClass) return;
    this._currentClass.assignments.push({ name, value });
  }

  _appendNodeClass(cls) {
    this._newline();
    this._line('type ' + cls.name + ' struct {');
    this._indent(() => {
      this._line('BaseNode');
      for (let fieldName of cls.fields.values()) {
        if (!fieldName) continue;
        this._line(fieldName + ' TreeNode');
      }
    });
    this._line('}');
    this._newline();

    this._line('var _ TreeNode = (*' + cls.name + ')(nil)');
    this._newline();

    this._line(
      'func new' +
        cls.name +
        '(text string, start int, elements []TreeNode) TreeNode {'
    );
    this._indent(() => {
      this._line('node := &' + cls.name + '{');
      this._indent(() => {
        this._line(
          'BaseNode: BaseNode{text: text, offset: start, children: elements},'
        );
      });
      this._line('}');
      for (let assignment of cls.assignments) {
        let fieldName = cls.fields.get(assignment.name);
        if (fieldName)
          this._line('node.' + fieldName + ' = ' + assignment.value);
      }
      this._line('return node');
    });
    this._line('}');
    this._newline();
  }

  method_(name, args, block) {
    this._startFunction(name, args, block);
  }

  function_(name, args, block) {
    this._startFunction(name, args, block);
  }

  _startFunction(name, args, block) {
    this._newline();
    this._line(
      'func (p *' +
        this._structName +
        ') ' +
        name +
        '(' +
        args.join(', ') +
        ') TreeNode {'
    );
    this._indent(() => {
      block();
    });
    this._line('}');
  }

  cache_(name, block) {
    let vars = this.localVars_({
      address: this.nullNode_(),
      index: this.offset_(),
    });
    let address = vars.address;
    let start = vars.index;
    let cacheVar = this.localVar_(
      'cache',
      'p.cache[' + this._quote(name) + ']'
    );

    this.if_(cacheVar + ' == nil', () => {
      this.assign_(cacheVar, 'make(map[int]cacheEntry)');
      this.assign_('p.cache[' + this._quote(name) + ']', cacheVar);
    });

    this._line('if entry, ok := ' + cacheVar + '[' + start + ']; ok {');
    this._indent(() => {
      this.assign_('p.offset', 'entry.offset');
      this._return('entry.node');
    });
    this._line('}');

    block(address);

    this.assign_(
      cacheVar + '[' + start + ']',
      'cacheEntry{node: ' + address + ', offset: p.offset}'
    );
    this._return(address);
  }

  localVars_(vars) {
    let names = {};
    for (let key in vars) names[key] = this.localVar_(key, vars[key]);
    return names;
  }

  localVar_(name, value) {
    let varName = this._varName(name);
    let typeName = TYPES[name] || 'TreeNode';
    this._line('var ' + varName + ' ' + typeName);
    if (value !== undefined) this.assign_(varName, value);
    return varName;
  }

  chunk_(length) {
    let vars = this.localVars_({
      chunk: this._emptyString(),
      max: this.offset_() + ' + ' + length,
    });
    let chunk = vars.chunk;
    let max = vars.max;
    this.if_(max + ' <= len(p.input)', () => {
      this.assign_(chunk, 'string(p.input[p.offset:' + max + '])');
    });
    return chunk;
  }

  syntaxNode_(address, start, end, elements, action, nodeClass) {
    let textExpr = 'p.slice(' + start + ', ' + end + ')';
    let elementsExpr = elements || 'nil';

    if (action) {
      const methodName =
        (this._actionMap && this._actionMap.get(action)) ||
        toPascalCase(action);
      this._line('if p.actions == nil {');
      this._indent(() => {
        this.assign_(
          'p.actionErr',
          'fmt.Errorf(' + this._quote('missing actions for ' + action) + ')'
        );
        this._return('nil');
      });
      this._line('}');
      this._line(
        'node, err := p.actions.' +
          methodName +
          '(p.inputString, ' +
          start +
          ', ' +
          end +
          ', ' +
          elementsExpr +
          ')'
      );
      this._line('if err != nil {');
      this._indent(() => {
        this.assign_('p.actionErr', 'err');
        this._return('nil');
      });
      this._line('}');
      this.assign_(address, 'node');
    } else if (nodeClass) {
      this.assign_(
        address,
        'new' +
          nodeClass +
          '(' +
          textExpr +
          ', ' +
          start +
          ', ' +
          elementsExpr +
          ')'
      );
    } else {
      this.assign_(
        address,
        '&BaseNode{text: ' +
          textExpr +
          ', offset: ' +
          start +
          ', children: ' +
          elementsExpr +
          '}'
      );
    }
    this.assign_('p.offset', end);
  }

  ifNode_(address, block, else_) {
    this.if_(address + ' != nil', block, else_);
  }

  unlessNode_(address, block, else_) {
    this.if_(address + ' == nil', block, else_);
  }

  ifNull_(elements, block, else_) {
    this.if_(elements + ' == nil', block, else_);
  }

  extendNode_(address, nodeType) {
    this.assign_(
      address,
      'p.extendNode(' + address + ', ' + this._quote(nodeType) + ')'
    );
  }

  failure_(address, expected) {
    let rule = this._grammarName + '::' + this._ruleName;
    this.assign_(address, this.nullNode_());
    this.if_('p.offset > p.failure.offset', () => {
      this.assign_('p.failure.offset', 'p.offset');
      this.assign_('p.failure.expected', 'nil');
    });
    this.if_('p.offset == p.failure.offset', () => {
      this.assign_(
        'p.failure.expected',
        'append(p.failure.expected, expectation{rule: ' +
          this._quote(rule) +
          ', expected: ' +
          this._quote(expected) +
          '})'
      );
    });
  }

  jump_(address, rule) {
    this.assign_(address, 'p._read_' + rule + '()');
  }

  if_(condition, block, else_) {
    this._line('if ' + condition + ' {');
    this._indent(block);
    if (else_) {
      this._line('} else {');
      this._indent(else_);
    }
    this._line('}');
  }

  loop_(block) {
    this._line('for {');
    this._indent(block);
    this._line('}');
  }

  break_() {
    this._line('break');
  }

  sizeInRange_(address, [min, max]) {
    if (max === -1) return 'len(' + address + ') >= ' + min;
    if (max === 0) return 'len(' + address + ') == ' + min;
    return (
      'len(' + address + ') >= ' + min + ' && len(' + address + ') <= ' + max
    );
  }

  stringMatch_(expression, string) {
    return expression + ' == ' + this._quote(string);
  }

  stringMatchCI_(expression, string) {
    this._parserImports.add('strings');
    return 'strings.EqualFold(' + expression + ', ' + this._quote(string) + ')';
  }

  regexMatch_(regex, string) {
    this._parserImports.add('regexp');
    return regex + '.MatchString(' + string + ')';
  }

  compileRegex_(charClass, name) {
    let pattern = charClass.regex.source;
    this._line(
      'var ' + name + ' = regexp.MustCompile(' + this._quote(pattern) + ')'
    );
    charClass.constName = name;
    this._parserImports.add('regexp');
  }

  arrayLookup_(expression, offset) {
    return expression + '[' + offset + ']';
  }

  append_(list, value, index) {
    if (index === undefined) {
      this.assign_(list, 'append(' + list + ', ' + value + ')');
    } else {
      this._line(list + '[' + index + '] = ' + value);
    }
  }

  hasChars_() {
    return 'p.offset < len(p.input)';
  }

  nullNode_() {
    return 'nil';
  }

  offset_() {
    return 'p.offset';
  }

  emptyList_(size) {
    if (size) return 'make([]TreeNode, ' + size + ')';
    return 'nil';
  }

  _emptyString() {
    return '""';
  }

  null_() {
    return 'nil';
  }

  pass_() {
    this._line('// pass');
  }

  _writeParserHelpers(root) {
    this._newline();
    this._line(
      'func New(input string, actions Actions) *' + this._structName + ' {'
    );
    this._indent(() => {
      this._line('return &' + this._structName + '{');
      this._indent(() => {
        this._line('input: []rune(input),');
        this._line('inputString: input,');
        this._line('actions: actions,');
        this._line('cache: make(map[string]map[int]cacheEntry),');
      });
      this._line('}');
    });
    this._line('}');
    this._newline();

    this._line(
      'func (p *' +
        this._structName +
        ') WithTypes(types map[string]NodeExtender) *' +
        this._structName +
        ' {'
    );
    this._indent(() => {
      this._line('p.types = types');
      this._line('return p');
    });
    this._line('}');
    this._newline();

    this._line(
      'func Parse(input string, actions Actions, types map[string]NodeExtender) (TreeNode, error) {'
    );
    this._indent(() => {
      this._line('parser := New(input, actions)');
      this._line('if types != nil {');
      this._indent(() => {
        this._line('parser.types = types');
      });
      this._line('}');
      this._line('return parser.Parse()');
    });
    this._line('}');
    this._newline();

    this._line(
      'func (p *' + this._structName + ') Parse() (TreeNode, error) {'
    );
    this._indent(() => {
      this._line('node := p._read_' + root + '()');
      this._line('if p.actionErr != nil {');
      this._indent(() => {
        this._line('return nil, p.actionErr');
      });
      this._line('}');
      this._line('if node != nil && p.offset == len(p.input) {');
      this._indent(() => {
        this._line('return node, nil');
      });
      this._line('}');
      this._line('if len(p.failure.expected) == 0 {');
      this._indent(() => {
        this._line('p.failure.offset = p.offset');
        this._line(
          'p.failure.expected = append(p.failure.expected, expectation{rule: ' +
            this._quote(this._grammarName) +
            ', expected: "<EOF>"})'
        );
      });
      this._line('}');
      this._line('return nil, p.newParseError()');
    });
    this._line('}');
    this._newline();

    this._line('func (p *' + this._structName + ') newParseError() error {');
    this._indent(() => {
      this._line('line, column := 1, 1');
      this._line('for i, r := range p.input {');
      this._indent(() => {
        this._line('if i >= p.failure.offset {');
        this._indent(() => {
          this._line('break');
        });
        this._line('}');
        this._line("if r == '\\n' {");
        this._indent(() => {
          this._line('line++');
          this._line('column = 1');
        });
        this._line('} else {');
        this._indent(() => {
          this._line('column++');
        });
        this._line('}');
      });
      this._line('}');
      this._line(
        'message := fmt.Sprintf("parse error at line %d, column %d", line, column)'
      );
      this._line('if len(p.failure.expected) > 0 {');
      this._indent(() => {
        this._line('message += ": expected "');
        this._line('for i, exp := range p.failure.expected {');
        this._indent(() => {
          this._line('if i > 0 {');
          this._indent(() => {
            this._line('if i == len(p.failure.expected)-1 {');
            this._indent(() => {
              this._line('message += " or "');
            });
            this._line('} else {');
            this._indent(() => {
              this._line('message += ", "');
            });
            this._line('}');
          });
          this._line('}');
          this._line(
            'message += fmt.Sprintf("%s from %s", exp.expected, exp.rule)'
          );
        });
        this._line('}');
      });
      this._line('}');
      this._line('expected := make([]expectation, len(p.failure.expected))');
      this._line('copy(expected, p.failure.expected)');
      this._line('return &ParseError{');
      this._indent(() => {
        this._line('Input: p.inputString,');
        this._line('Offset: p.failure.offset,');
        this._line('Line: line,');
        this._line('Column: column,');
        this._line('Expected: expected,');
        this._line('Message: message,');
      });
      this._line('}');
    });
    this._line('}');
    this._newline();

    this._line(
      'func (p *' + this._structName + ') slice(start, end int) string {'
    );
    this._indent(() => {
      this._line('if start < 0 { start = 0 }');
      this._line('if end > len(p.input) { end = len(p.input) }');
      this._line('if start > end { start = end }');
      this._line('return string(p.input[start:end])');
    });
    this._line('}');
    this._newline();

    this._line(
      'func (p *' +
        this._structName +
        ') extendNode(node TreeNode, name string) TreeNode {'
    );
    this._indent(() => {
      this._line('if node == nil {');
      this._indent(() => {
        this._line('return nil');
      });
      this._line('}');
      this._line('if p.types == nil {');
      this._indent(() => {
        this._line('return node');
      });
      this._line('}');
      this._line('if extender, ok := p.types[name]; ok && extender != nil {');
      this._indent(() => {
        this._line('return extender(node)');
      });
      this._line('}');
      this._line('return node');
    });
    this._line('}');
    this._newline();
  }

  serialize() {
    let buffers = super.serialize();
    let parserPath = join(this._outputPath, 'parser.go');
    if (buffers.has(parserPath)) {
      let imports = Array.from(this._parserImports).sort();
      let source = buffers.get(parserPath);
      let block = imports.map((name) => '\t"' + name + '"').join('\n');
      source = source.replace('// __IMPORTS__', block);
      buffers.set(parserPath, source);
    }
    return buffers;
  }
}

module.exports = Builder;
