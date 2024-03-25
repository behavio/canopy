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
  llm -m 4t \
    -s "$( <$PROMPT head -2 )" \
> samples-go/doc-4t.md

PROMPT=prompts/go-new-doc.md.tpl
jinja2 --strict $PROMPT \
    -D java_md="$( cat ../site/langs/java.md )" \
    -D python_md="$( cat ../site/langs/python.md )" |
  llm -m gpt-4 \
    -s "$( <$PROMPT head -2 )" \
> samples-go/doc-gpt-4.md

PROMPT=prompts/go-new-doc.markup.tpl
jinja2 --strict $PROMPT \
    -D java_md="$( cat ../site/langs/java.md )" \
    -D python_md="$( cat ../site/langs/python.md )" |
  llm -m claude-3-opus \
    -s "$( <$PROMPT head -2 )" \
> samples-go/doc-opus.md

# after some checkging on how to do the actions API:
# - ideal solution would be to keep the parser 'in package'
# - Actions implemented as methods on a struct, checked via an interface
#   - go does not like `_` in names, names should be fixed (non-exported camel case)
# - best results come from opus and 4t

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