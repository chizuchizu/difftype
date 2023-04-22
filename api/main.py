from fastapi import FastAPI

app = FastAPI()


def read_file(filepath):
    try:
        with open(filepath, "r") as f:
            file_contents = f.read()
        return file_contents
    except FileNotFoundError:
        print(f"Error: File '{filepath}' not found.")
        return None
    except PermissionError:
        print(f"Error: No permission to read file '{filepath}'.")
        return None


@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.get("/typing_challenge")
async def get_typing_challenge(idx: int):
    src_path = f"challenges/src/{idx}.txt"
    src_str = read_file(src_path)
    dst_path = f"challenges/dst/{idx}.txt"
    dst_str = read_file(dst_path)
    return {"src": src_str, "dst": dst_str}
