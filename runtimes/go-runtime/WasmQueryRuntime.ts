interface WasmQueryWindow extends Window {
  checkWasmQueryError: Function;
}

(window as WasmQueryWindow & typeof globalThis).checkWasmQueryError = <T>(value: T) => {
  if(value instanceof Error) {
    throw value;
  } else {
    return value;
  }
}