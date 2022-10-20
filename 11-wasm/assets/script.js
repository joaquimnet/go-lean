const goWasm = new Go();

WebAssembly.instantiateStreaming(fetch('main.wasm'), goWasm.importObject).then((result) => {
  goWasm.run(result.instance);

  document.getElementById('get-html').addEventListener('click', () => {
    const h4 = document.createElement('h4');
    h4.innerHTML = getHtml();
    document.body.appendChild(h4);
  });
});
