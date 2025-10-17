'use strict';

const Compiler = require('./compiler');

module.exports = {
  builders: {
    go: require('./builders/go'),
    java: require('./builders/java'),
    javascript: require('./builders/javascript'),
    python: require('./builders/python'),
    ruby: require('./builders/ruby'),
  },

  compile(grammar, builder) {
    let compiler = new Compiler(grammar, builder);
    return compiler.toSource();
  },
};
