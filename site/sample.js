'use strict'

const availableUrl = (urlString) => {
    try {
        new URL(urlString)
        return true
    } catch (e) {
        return false
    }
}

const initGo = () => {
    const go = new Go()
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
        .then((result) => {
            go.run(result.instance);
        });
}

const initListeners = () => {
    const button = document.getElementById('contentGetButton')
    const clear = document.getElementById('clearButton')
    const text = document.getElementById('url')
    button.addEventListener('click', (e) => {
        wasmCallback(text.value)
    })
    clear.addEventListener('click', (e) => {
        text.value = ""
        button.disabled = true
        const area = document.getElementById('app')
        area.innerText = ""
    })
    button.disabled = true
    text.addEventListener('input', (e) => {
        button.disabled = !availableUrl(text.value)
    })
}

const init = () => {
    initGo()
    initListeners()
}

window.onload = () => {
    init()
    console.log("initialize done")
}