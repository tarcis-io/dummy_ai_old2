'use strict';

const wasmStart = function(wasmFile) {

	const go = new Go();
	WebAssembly.instantiateStreaming(fetch(wasmFile), go.importObject);
};
