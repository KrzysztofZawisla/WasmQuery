import brotli
import sys
import argparse

try:
    parser = argparse.ArgumentParser()
    parser.add_argument("--file", help="Here you put file to compress")
    try:
        parser.add_argument("--outfile", help="Filename after compression")
    except:
        pass
    args = parser.parse_args()
except Exception as e:
    print(e)
    sys.exit(1)

with open(args.file, "rb") as data:
    bytesFromFile = data.read()

compressedBytes = brotli.compress(bytesFromFile)

if bool(args.outfile):
    with open(args.outfile + ".br", "wb") as data:
        data.write(compressedBytes)
else: 
    with open(args.file + ".br", "wb") as data:
        data.write(compressedBytes)
        
sys.exit(0)