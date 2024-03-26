# dependencies
pip install openai anthropic llm ttok jinja2-cli
conda env config vars set ANTHROPIC_API_KEY="$( cat ~/.anthropic.key )"
conda env config vars set OPENAI_API_KEY="$( cat ~/.openai.key )"

# "prompts" encoding
# - first 2 lines is system (empty lines if not needed)

#
# Plan 1
#
# The original idea, no AI input;) Willing to try GPT and Claude.
#
# - start with golang, as it's more popular than R
# - use language docs as samples to generate Go docs
# - fix up the doc manually
# - use pairs of doc - implementation as examples and generate the golang generator
# - try to run the generator on the samples and iterate with model over the errors
# - generate the test runner for golang
# - run the included tests, iterate with errors

# see gen-builder.py for initial tests

# shell version follows:
PROMPT=prompts/go-new-doc.md.tpl
jinja2 --strict $PROMPT \
    -D java_md="$( cat ../site/langs/java.md )" \
    -D python_md="$( cat ../site/langs/python.md )" |
  tail -n+3 |
  llm -m 4t \
    -s "$( <$PROMPT head -2 )" \
> samples-go/doc-4t.md

PROMPT=prompts/go-new-doc.md.tpl
jinja2 --strict $PROMPT \
    -D java_md="$( cat ../site/langs/java.md )" \
    -D python_md="$( cat ../site/langs/python.md )" |
  tail -n+3 |
  llm -m gpt-4 \
    -s "$( <$PROMPT head -2 )" \
> samples-go/doc-gpt-4.md

PROMPT=prompts/go-new-doc.markup.tpl
jinja2 --strict $PROMPT \
    -D java_md="$( cat ../site/langs/java.md )" \
    -D python_md="$( cat ../site/langs/python.md )" |
  tail -n+3 |
  llm -m claude-3-opus \
    -s "$( <$PROMPT head -2 )" \
> samples-go/doc-opus.md

# after some checkging on how to do the actions API:
# - ideal solution would be to keep the parser 'in package'
# - Actions implemented as methods on a struct, checked via an interface
#   - go does not like `_` in names, names should be fixed (non-exported camel case)
# - best results come from opus and 4t (latest gpt 4 turbo)
# - won't implement name shortcuts yet, it just makes things more complicated

# use the best doc and update it a little
cp samples-go/doc-opus.md ../site/langs/golang.md

#
# Plan 2
# mostly Claude with it's 200k context.. maybe 128k GPT would do
#
# - ask for guidance first
# - feed all the relevant source code to the model, ask it to generate the needed files

# trying claude to suggest the best course of action
PROMPT=prompts/go-course-of-action.txt

tail -n+3 $PROMPT |
  llm -m claude-3-opus \
    -s "$( <$PROMPT head -2 )" \
> samples/course-of-action-opus.md

tail -n+3 $PROMPT |
  llm -m 4t \
    -s "$( <$PROMPT head -2 )" \
> samples/course-of-action-gpt-4-turbo.md

# the suggested course of action is rather for a human
find .. -type f |
  grep -v -e .git -e node_modules -e ai-gen -e examples \
    -e 'test/grammars.*\.[^p]' -e test/java/target -e __pycache \
    -e 'test/grammars.*\.py' \
    -e meta_grammar.js -e site/langs/golang \
    -e ruby -e javascript \
    -e site/[^l][^a] \
> tmp-files.txt

<tmp-files.txt xargs -I{} ttok -i {} > tmp-tokens.txt

paste tmp-tokens.txt tmp-files.txt |
  sort -rn |
  head -20

# all tokens at once
<tmp-files.txt xargs cat | ttok

# whole repo over 200k ;( got to ~50k

# pack the context, nice paths..
cd ..
<ai-gen-builders/tmp-files.txt sed 's,../,,' |
  python ai-gen-builders/ai-tar.py a \
> ai-gen-builders/tmp-repo.markup
cd ai-gen-builders/

# java has ~3k tokens, some changes in CI, readme, docs .. 10k?
# data is too big to fit on the command line, so we'll use a file
# cygwin cant do file handles on command line, so we'll use a file
<tmp-repo.markup jq --raw-input --slurp '{"doc_dump": .}' > tmp-jinja.json

PROMPT=prompts/go-big-bang-2.markup.tpl
jinja2 --strict --format json $PROMPT tmp-jinja.json |
  tail -n+3 |
  llm -m claude-3-opus \
    -s "$( <$PROMPT head -2 )" \
> samples-go/bang-opus-2.markup

jinja2 --strict --format json $PROMPT tmp-jinja.json |
  tail -n+3 |
  (cat; echo -e "\nAssistant:"; cat samples-go/bang-opus-2.markup ) |
  llm -m claude-3-opus \
    -s "$( <$PROMPT head -2 )" \
> samples-go/bang-opus-2-1.markup

# this required manual editing, as the 'assistant' prompt is hacky in llm
cat samples-go/bang-opus-2{,-1}.markup > samples-go/bang-opus-2-full.markup
mkdir sample-opus-2
cd sample-opus-2
<../samples-go/bang-opus-2-full.markup python ../ai-tar.py x

#s
# cheaper model, more turns
PROMPT=prompts/go-big-bang-2.markup.tpl
jinja2 --strict --format json $PROMPT tmp-jinja.json |
  tail -n+3 |
  llm -m claude-3-sonnet \
    -s "$( <$PROMPT head -2 )" \
> samples-go/bang-sonnet-1.markup

# maybe something like llm -c would do, but i don't like the stateful nature of it
jinja2 --strict --format json $PROMPT tmp-jinja.json |
  tail -n+3 |
  (cat; echo -e "\nAssistant:"; cat samples-go/bang-sonnet-1.markup ) |
  llm -m claude-3-sonnet \
    -s "$( <$PROMPT head -2 )" \
> samples-go/bang-sonnet-1-1.markup

mkdir sample-sonnet
cd sample-sonnet
cat ../samples-go/bang-sonnet-1{,-1}.markup | python ../ai-tar.py x
