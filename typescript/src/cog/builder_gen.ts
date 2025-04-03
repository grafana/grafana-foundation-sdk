// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface Builder<T> {
  build: () => T;
}

export function isBuilder<T>(input: Builder<T> | any): input is Builder<T> {
  if (input === null) {
    return false;
  }
  if (!input?.build) {
    return false;
  }

  return typeof input.build === "function";
}
