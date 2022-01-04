// btn.addEventListener("click", () )
const go = new Go();

WebAssembly.instantiateStreaming(
	fetch('hello.wasm'),
	go.importObject
).then((result) => {
	// console.log(result.instance.exports.add(4, 5));
	go.run(result.instance);
});


// DOM elements
const btn = document.querySelector("#btn")

btn.addEventListener("click", () => {
	const str = document.querySelector("#str")
	const p = document.querySelector("#textCapitalized")
	p.textContent = mettreEnMajuscule(str.value)

	// Effacer le contenu de l'input
	str.value = ''
})


