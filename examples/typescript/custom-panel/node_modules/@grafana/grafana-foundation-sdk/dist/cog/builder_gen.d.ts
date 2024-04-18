export interface Builder<T> {
    build: () => T;
}
export declare function isBuilder<T>(input: Builder<T> | any): input is Builder<T>;
