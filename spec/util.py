def read_file(file: str) -> str:
    """
    读取文件
    Raise OSError upon failure.
    """
    stream = open(file, encoding="utf8")
    text = stream.read()
    stream.close()
    return text


# __name__ 模块名
# https://docs.python.org/3/library/__main__.html
if __name__ == "__main__":
    print(read_file("python.md"))
