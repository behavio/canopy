"""Pack text files into one stream for use in Anthropic models.

Input:
  list of file names on stdin, one per line

Output:
  Contents of the files, wrapped in <document index="1">.
  Inside <source> contains the path and <document_content> contains the actual bytes.
  The whole stream is wrapped in <documents>.
"""
import sys

def main():

    print("<documents>")

    for index, line in enumerate(sys.stdin):
        filename = line.strip()
        with open(filename, encoding="utf-8") as f:
            content = f.read()
            print(f'<document index="{index+1}">')
            print(f'<source>{filename}</source>')
            print(f'<document_content>{content}</document_content>')
            print("</document>")

    print("</documents>")

if __name__ == "__main__":
    main()
