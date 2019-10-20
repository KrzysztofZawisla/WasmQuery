# WasmQuery
Experimental WASM library which simplifies writing JavaScript code and share some standard libraries from Golang.  
Note: Library have problem with speed on chromium based browsers.
## General Roadmap
| Goals | Status |
| --- | --- |
| Create user friendly API with error coverage | ☐ | 
| Create build with Tinygo | ☐ | 
| Create three versions (in GO, Rust and AssemblyScript) | ☐ | 
| Stable build for shortQuery module | ☐ | 
| Write complete documentation | ☐ | 
| Optimize the code | ☐ | 
## Wiki
https://github.com/KrzysztofZawisla/WasmQuery/wiki
## Working Environment
Windows 10 Professional 64-bit version 1903 build 18362.356  
Version of Golang: 1.12.9 windows/amd64  
Version of Python: 3.6.8 windows/amd64  
Version of Node.js: 12.10.0 windows/amd64  
Version of TypeScript: 3.6.3 windows/amd64  
Version of Flow: - windows/amd64
## Used packages
encoding/hex  
crypto/sha512  
crypto/sha256  
crypto/sha1  
crypto/md5  
golang.org/x/crypto/md4  
golang.org/x/crypto/sha3  
hash/adler32  
time  
math  
syscall/js  
## Programmers
Krzysztof Zawisła - https://github.com/KrzysztofZawisla  
Szymon Bendorz - https://github.com/dsonyy
### Build commands:
Dev build: go build -o WasmQuery.wasm  
Production build: go build -o WasmQuery.wasm -ldflags="-s -w"
