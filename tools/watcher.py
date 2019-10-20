from shutil import copyfile, rmtree
from distutils.dir_util import copy_tree
from subprocess import check_output
import subprocess
import os



def quit(code=0):
    input("Press any key to continue...")
    exit(code)

def cmd(command) -> str:
    out = subprocess.run(command, shell=True, check=True, capture_output=True, encoding='utf-8').stdout
    if out and out[-1] == '\n':
        return out[:-1]
    return out

def error(msg):
    print("ERROR: '" + e + "'.")

def main():

    print('\n---------------------------------------------')

    print("GOROOT: '", end='')
    GOROOT_PATH         = os.path.abspath(cmd('go env GOROOT'))
    print(GOROOT_PATH + "'")

    SRC_PACKAGE_PATH = os.path.abspath(GOROOT_PATH + '/src/github.com/KrzysztofZawisla/WasmQuery' )
    PROJECT_PATH     = os.path.dirname(os.path.abspath(__file__  + '/..'))
    SRC_PATH         = os.path.abspath(PROJECT_PATH + '/src/go-src')
    TEMP_PATH        = os.path.abspath(PROJECT_PATH + '/.temp')
    
    WASM_TEST_PATH   = os.path.abspath(TEMP_PATH + '/WasmQuery.wasm')
    WASM_DIST_PATH   = os.path.abspath(PROJECT_PATH + '/dist/WasmQuery.wasm')
    WASM_SERVER_PATH = os.path.abspath(PROJECT_PATH + '/tools/server/WasmQuery.wasm')

    print('SRC_PACKAGE_PATH: ' + SRC_PACKAGE_PATH)
    print('PROJECT_PATH: ' + PROJECT_PATH)
    print('SRC_PATH: ' + SRC_PATH)
    print('WASM_TEST_PATH: ' + WASM_TEST_PATH)
    print('WASM_DIST_PATH: ' + WASM_DIST_PATH)
    print('WASM_SERVER_PATH: ' + WASM_SERVER_PATH)

    print('---------------------------------------------')

    print("Checking if package path '" + SRC_PACKAGE_PATH + "' exists: ", end="")
    ret = os.path.isdir(SRC_PACKAGE_PATH)
    print(ret)

    if not ret:
        print("Copying source files from '" + SRC_PATH + "' to '" + SRC_PACKAGE_PATH + "'.", end="")
        copy_tree(SRC_PATH, SRC_PACKAGE_PATH)
        quit()

    print("Building package from '" + SRC_PACKAGE_PATH + "' to '" + WASM_TEST_PATH + "'.")
    ret = cmd("set GOOS=js& set GOARCH=wasm& go build -o " + WASM_TEST_PATH + " " + SRC_PACKAGE_PATH)
    
    print("Testing: ", end="")
    ret = True
    print(ret)
    
    if not ret:
        print("Tests failed")
        quit(1)
    
    print("All tests passed.")
    
    print("Copying dist from '" + WASM_TEST_PATH + "' to '" + WASM_SERVER_PATH + "'.")
    copyfile(WASM_TEST_PATH, WASM_SERVER_PATH)
    
    print("Copying dist from '" + WASM_TEST_PATH + "' to '" + WASM_DIST_PATH + "'.")
    copyfile(WASM_TEST_PATH, WASM_DIST_PATH)
    
    print("Removing files from '" + WASM_TEST_PATH + "'.")
    rmtree(TEMP_PATH, ignore_errors=True)

    print("Removing source files from '" + SRC_PATH + "'.")
    rmtree(TEMP_PATH, ignore_errors=True)
    
    print("Copying source files from '" + SRC_PACKAGE_PATH + "' to '" + SRC_PATH + "'.")
    copy_tree(SRC_PACKAGE_PATH, SRC_PATH)

if __name__ == "__main__":
    main()

