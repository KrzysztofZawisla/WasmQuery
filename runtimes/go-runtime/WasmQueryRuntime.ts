declare function runWasmQuery(): void;
declare const global;

type checkWasmQueryError = <T>(value: T) => T;
type bindWasmQueryGlobal = () => void;

let wasmQueryIntance;
let wasmQueryModule;

class WasmQueryError extends Error {
  constructor(message: string) {
    super(message);
    Object.setPrototypeOf(this, WasmQueryError.prototype);
  }
}

const checkWasmQueryError: checkWasmQueryError = <T>(value: T) => {
  try {
    if (value instanceof Error) {
      throw value;
    } else {
      return value;
    }
  } catch (err) {
    if (err instanceof WasmQueryError) {
      throw err;
    } else {
      console.warn("Instance reset");
      runWasmQuery();
    }
  }
}

const bindWasmQueryGlobal: bindWasmQueryGlobal = (): void => {
  global.checkWasmQueryError = checkWasmQueryError;
  global.WasmQueryError = WasmQueryError;
}