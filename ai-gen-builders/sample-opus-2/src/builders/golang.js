"use strict";
const { sep } = require("path");
const Base = require("./base");

const TYPES = {
  address: "*node",
  chunk: "string",
  elements: "[]node",
  index: "int",
  max: "int",
};

class Builder extends Base {
  constructor(...args) {
    super(...args);
    this._labels = new Set();
  }

  _tab() {
    return "\t";
  }

  _initBuffer(pathname) {
    let namespace = pathname.split(sep);
    namespace.pop();
    return "package " + namespace[namespace.length - 1] + "\n\n";
  }

  _quote(string) {
    string = string
      .replace(/\\/g, "\\\\")
      .replace(/"/g, '\\"')
      .replace(/\x08/g, "\\b")
      .replace(/\t/g, "\\t")
      .replace(/\n/g, "\\n")
      .replace(/\f/g, "\\f")
      .replace(/\r/g, "\\r");
    return '"' + string + '"';
  }

  comment(lines) {
    lines = lines.map((line) => "// " + line);
    return lines.join("\n");
  }

  package_(name, actions, block) {
    this._grammarName = name;
    this._newBuffer("go");
    this._line('import "strings"');
    this._newline();
    block();
  }

  syntaxNodeClass_() {
    return "node";
  }

  grammarModule_(block) {
    this._line("const FAILURE = &node{}");
    this._newline();
    this._line("type node struct {");
    this._indent(() => {
      this._line("text     string");
      this._line("offset   int");
      this._line("elements []node");
    });
    this._line("}");
    this._newline();
    this._newline();
    this.class_("grammar", "struct", () => {
      this._line("input    string");
      this._line("size     int");
      this._line("offset   int");
      this._line("cache    map[int]map[int]*node");
      this._line("failure  int");
      this._line("expected []string");
    });
    this._newline();
    this.method_("grammar", "newGrammar", ["input string"], () => {
      this._line("g := &grammar{");
      this._indent(() => {
        this._line("input: input,");
        this._line("size:  len(input),");
        this._line("cache: make(map[int]map[int]*node),");
      });
      this._line("}");
      this._line("return g");
    });
    this._newline();
    block();
  }

  compileRegex_(charClass, name) {
    let regex = charClass.regex;
    this._line(
      "var " +
        name +
        " = regexp.MustCompile(" +
        this._quote("^" + regex.source) +
        ")"
    );
    charClass.constName = name;
    this._newline();
  }

  parserClass_(root) {
    this._newline();
    this.method_("grammar", "parse", [], () => {
      this._line("tree := g._read_" + root + "()");
      this.if_("tree != FAILURE && g.offset == g.size", () => {
        this._line("return tree, nil");
      });
      this.if_("len(g.expected) == 0", () => {
        this._line("g.failure = g.offset");
        this._line('g.expected = append(g.expected, "<EOF>")');
      });
      this._line(
        'return nil, fmt.Errorf("' +
          this._grammarName +
          '.ParseError: Line %d: expected %s",'
      );
      this._line(
        '                        g.lineNumber(), strings.Join(g.expected, ", "))'
      );
    });
    this._newline();
    this.method_("grammar", "lineNumber", [], () => {
      this._line('lines := strings.Split(g.input[0:g.failure], "\\n")');
      this._line("return len(lines)");
    });
  }

  class_(name, parent, block) {
    this._newline();
    this._line("type " + name + " " + parent);
    block();
  }

  constructor_(args, block) {
    // No constructors in Go
  }

  method_(receiver, name, args, block) {
    args = ["g *" + receiver].concat(args).join(", ");
    this._newline();
    this._line("func (" + args + ") " + name + "() {");
    this._indent(() => block());
    this._line("}");
  }

  cache_(name, block) {
    this._labels.add(name);
    let temp = this.localVars_({ address: "tree", index: "g.offset" }),
      address = temp.address,
      offset = temp.index;
    this._line("if _, ok := g.cache[" + offset + "]; !ok {");
    this._line("   g.cache[" + offset + "] = make(map[int]*node)");
    this._line("}");
    this.if_(
      "tree, ok := g.cache[" + offset + "][" + this._quote(name) + "]; ok",
      () => {
        this._line("g.offset = tree.offset");
        this._line("return tree");
      }
    );
    block(address);
    this._line(
      "g.cache[" + offset + "][" + this._quote(name) + "] = " + address
    );
    this._return(address);
  }

  attributes_(names) {
    for (let name of names) {
      this._line(name + " node");
      this._labels.add(name);
    }
  }

  attribute_(name, value) {
    this._line("tree." + name + " = " + value);
  }

  localVars_(vars) {
    let names = {};
    for (let name in vars) names[name] = this.localVar_(name, vars[name]);
    return names;
  }

  localVar_(name, value) {
    let varName = this._varName(name);
    if (value === undefined) value = "FAILURE";
    this._line("var " + varName + " " + TYPES[name] + " = " + value);
    return varName;
  }

  chunk_(length) {
    let input = "g.input",
      ofs = "g.offset",
      temp = this.localVars_({ chunk: '""', max: "end" }),
      end = this.localVar_("int", ofs + " + " + length);
    this.if_(end + " <= g.size", () => {
      this._line(temp.chunk + " = " + input + "[" + ofs + ":" + end + "]");
    });
    return temp.chunk;
  }

  syntaxNode_(address, start, end, elements, action, nodeClass) {
    nodeClass = nodeClass || "node";
    let args = [this._quote(""), start, end];
    if (elements) args.push(elements);
    else args.push("[]node{}");
    this._line(
      address +
        " = &" +
        nodeClass +
        "{text: g.input[" +
        start +
        ":" +
        end +
        "], offset: " +
        start +
        ", elements: " +
        args[3] +
        "}"
    );
    if (action) {
      this._line("g.actions." + action + "(" + address + ")");
    }
    this._line("g.offset = " + end);
  }

  ifNode_(address, block, else_) {
    this.if_(address + " != FAILURE", block, else_);
  }

  unlessNode_(address, block, else_) {
    this.if_(address + " == FAILURE", block, else_);
  }

  ifNull_(elements, block, else_) {
    this.if_(elements + " == nil", block, else_);
  }

  extendNode_(address, nodeType) {
    // TODO: Implement node extension if needed
  }

  assign_(name, value) {
    this._line(name + " = " + value);
  }

  jump_(address, name) {
    this._line(address + " = g._read_" + name + "()");
  }

  conditional_(keyword, condition, block, else_) {
    this._line(keyword + " " + condition + " {");
    this._indent(() => block());
    if (else_) {
      this._line("} else {");
      this._indent(() => else_());
    }
    this._line("}");
  }

  if_(condition, block, else_) {
    this.conditional_("if", condition, block, else_);
  }

  unless_(condition, block, else_) {
    this.conditional_("if", condition, block, else_);
  }

  loop_(block) {
    this._line("for {");
    this._indent(() => block());
    this._line("}");
  }

  break_() {
    this._line("break");
  }

  return_(expression) {
    this._line("return " + expression);
  }

  string_() {
    return '""';
  }

  stringMatch_(expression, string) {
    return expression + " == " + this._quote(string);
  }

  stringMatchCI_(expression, string) {
    // TODO: Implement case-insensitive matching if needed
    return expression + " == " + this._quote(string);
  }

  regexMatch_(regex, string) {
    return string + ' != "" && ' + regex + ".MatchString(" + string + ")";
  }

  anyChar_() {
    return "g.offset < g.size";
  }

  appendToList_(list, element) {
    this._line(list + " = append(" + list + ", " + element + ")");
  }

  decodeChar(text) {
    if (text === "\\n") return "\n";
    if (text === "\\r") return "\r";
    if (text === "\\t") return "\t";
    if (text === "\\'") return "'";
    if (text === '\\"') return '"';
    if (text === "\\\\") return "\\";
    return text;
  }

  listLength_(list, length) {
    if (length) return "len(" + list + ") == " + length;
    else return list + " != nil";
  }

  action_(action) {
    this._line("g.actions." + action);
  }

  inRange_(value, range) {
    return range[0] + " <= " + value + " && " + value + " < " + range[1];
  }

  unpack_(...elements) {
    return elements.join(", ");
  }

  offset_() {
    return "g.offset";
  }
}

module.exports = Builder;
