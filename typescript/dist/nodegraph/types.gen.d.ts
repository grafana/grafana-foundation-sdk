export interface ArcOption {
    field?: string;
    color?: string;
}
export declare const defaultArcOption: () => ArcOption;
export interface NodeOptions {
    mainStatUnit?: string;
    secondaryStatUnit?: string;
    arcs?: ArcOption[];
}
export declare const defaultNodeOptions: () => NodeOptions;
export interface EdgeOptions {
    mainStatUnit?: string;
    secondaryStatUnit?: string;
}
export declare const defaultEdgeOptions: () => EdgeOptions;
export declare enum ZoomMode {
    Cooperative = "cooperative",
    Greedy = "greedy"
}
export declare const defaultZoomMode: () => ZoomMode;
export interface Options {
    nodes?: NodeOptions;
    edges?: EdgeOptions;
    zoomMode?: ZoomMode;
}
export declare const defaultOptions: () => Options;
