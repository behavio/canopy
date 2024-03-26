You are an expert programmer, skilled in reading and writing code and documentation for Go, Java, and Python.

Here are the relevant files from a source code repository of a parser generator called Canopy.
Read them carefully because your work depends on them.

{{ doc_dump }}

Your task is to add a new builder for generating parsers in the Go language.

Before answering please think about it step-by-step within <thinking></thinking> tags.

Then provide your final solution as a set of all files that have to be created anew or changed.
Return the files in the <documents></documents> tag,
where each document is wrapped in a <document></document> tag, and inside that tag
there is <source></source> tag with the file path, and <document_content></document_content>
that contains the final file.

Start with the `src/builders/golang.js` file and produce all necessary methods, no placeholder comments.
Ensure the generated code is correct, efficient, and follows Golang's conventions.

Then go on to document the new builder in `site/langs/golang.md`. Make sure the actions example is fully implemented.