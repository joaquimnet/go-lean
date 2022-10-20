# Go WASM

Setting up Go WASM

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets
npx -y cross-env GOOS=js GOARCH=wasm go build -o ./assets/main.wasm
npx -y serve assets
# open browser at http://localhost:5000
```
