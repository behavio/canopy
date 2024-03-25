# obsoleted by workflow.sh using llm cli
#
# use as intreactive notbook, do not run as a script

#
# OpenAI section
#
from openai import OpenAI

# load up sample files
with open('site/langs/java.md', encoding='utf-8') as f:
    java_md = f.read()

with open('site/langs/python.md', encoding='utf-8') as f:
    python_md = f.read()

client = OpenAI()

model="gpt-3.5-turbo"     # context 16k, max 4k output
# model="gpt-4-turbo-preview" # context 128k
# model="gpt-4" # context 8k

completion = client.chat.completions.create(
  # the complexity of requesting the right max is not worth it now (gpt-3.5-turbo makes it hard with the output limit less than context window)
  # max_tokens=2*len(java_md),
  n=3,
  model=model,
  messages=[
    {"role": "system", "content": "You are an expert programmer, skilled in writing code and documentation for Go, Java, and Python."},
    {"role": "user", "content": "I need you to write a documentation file for a Golang parser builder. It is a part of a parser genertor system called Canopy. "
                                "It genrates parsers in multiple languages. As an example, I'll provide documentation for Python and Java parser builders."},
    {"role": "user", "content": f"java.md\n```markdown\n{java_md}```"},
    {"role": "user", "content": f"python.md\n```markdown\n{python_md}```"},
    {"role": "user", "content": "Please generate a file that I can use as golang.md. "
                                "Generate all the sections present in the example files, with content appropriate for the Go language."
                                "Provide a full implementation of the examples, not just placeholder comments."}
  ]
)

print(completion.usage)

# we can have more choices for each model
for i, choice in enumerate(completion.choices):
  with open(f'ai-gen-builders/samples-go/doc-{model}-{i}.md', 'w', encoding='utf-8') as fo:
      fo.write(choice.message.content)
