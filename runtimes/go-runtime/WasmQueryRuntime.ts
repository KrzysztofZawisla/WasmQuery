type checkWasmQueryError = <T>(value: T) => T;

interface WasmQueryWindow extends Window {
  checkWasmQueryError: checkWasmQueryError;
}

(window as WasmQueryWindow & typeof globalThis).checkWasmQueryError = <T>(value: T) => {
  if(value instanceof Error) {
    throw value;
  } else {
    return value;
  }
}