window.checkWasmQueryError = (value) => {
  if(value instanceof Error) {
    throw value;
  } else {
    return value;
  }
}