"""Pack or unpack text files into one stream for use in Anthropic models.
The modes are switched by the first argument, tar-like: a to pack, x to unpack.

Pack mode (a):
Input:
  list of file names on stdin, one per line

Output:
  Contents of the files, wrapped in <document index="1">.
  Inside <source> contains the path and <document_content> contains the actual bytes.
  The whole stream is wrapped in <documents>.

Unpack mode (x):
Input:
  The stream as described above.

Output:
  Writes the files to the current directory, using the <source> as the file name.
  The script ensures not to traverse above the current directory.
"""
import sys
import argparse
import os


def pack(input_files):
    print("<documents>")

    for index, line in enumerate(input_files):
        filename = line.strip()
        with open(filename, encoding="utf-8") as f:
            content = f.read()
            print(f'<document index="{index+1}">')
            print(f'<source>{filename}</source>')
            print(f'<document_content>{content}</document_content>')
            print("</document>")

    print("</documents>")


class InputStream:
    """A simple implementation of input stream on top of memory buffer.
    Reads all the data at once, keeps track of the current position.
    """
    def __init__(self, input):
        self.data = input.read()
        self.pos = 0

    def read_until(self, target):
        new_pos = self.data.find(target, self.pos)
        if new_pos == -1:
            raise ValueError(f"Target '{target}' not found")
        old_pos = self.pos
        self.pos = new_pos + len(target)

        return self.data[old_pos:new_pos]


def safe_path(path):
    """Normalize the path and compare it to cwd. Return True if it is safe.
    """
    path = os.path.abspath(path)
    cwd = os.getcwd()

    if os.path.commonpath([path, cwd]) != cwd:
        raise ValueError(f"Path '{path}' is not safe")

    return path


def process_document(input_stream):
    input_stream.read_until('<source>')
    path = input_stream.read_until('</source>')

    # check if path is safe
    out_path = safe_path(path)

    # get the file content
    input_stream.read_until('<document_content>')
    content = input_stream.read_until('</document_content>')

    # make sure the directories exists
    os.makedirs(os.path.dirname(out_path), exist_ok=True)

    # write out the data
    with open(out_path, "w", encoding="utf-8") as f:
        f.write(content)


def unpack(input_data):

    input_stream = InputStream(input_data)

    # skip the contents until the first <documents>
    input_stream.read_until('<documents>')

    while True:
        try:
            input_stream.read_until('<document')
            process_document(input_stream)
        except ValueError:
            break

    # won't bother checking if the closing tag is present

def main():
    parser = argparse.ArgumentParser(description="Pack or unpack text files into one stream for use in Anthropic models.")
    parser.add_argument("mode", choices=["a", "x"], help="Mode: 'a' for pack, 'x' for unpack")
    args = parser.parse_args()

    if args.mode == "a":
        pack(sys.stdin)
    elif args.mode == "x":
        unpack(sys.stdin)


if __name__ == "__main__":
    main()
