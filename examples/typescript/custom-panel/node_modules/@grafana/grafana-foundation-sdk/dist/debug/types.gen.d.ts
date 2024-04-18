export interface UpdateConfig {
    render: boolean;
    dataChanged: boolean;
    schemaChanged: boolean;
}
export declare const defaultUpdateConfig: () => UpdateConfig;
export declare enum DebugMode {
    Render = "render",
    Events = "events",
    Cursor = "cursor",
    State = "State",
    ThrowError = "ThrowError"
}
export declare const defaultDebugMode: () => DebugMode;
export interface Options {
    mode: DebugMode;
    counters?: UpdateConfig;
}
export declare const defaultOptions: () => Options;
