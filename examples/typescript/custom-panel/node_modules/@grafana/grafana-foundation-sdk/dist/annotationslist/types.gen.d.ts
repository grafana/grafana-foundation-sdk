export interface Options {
    onlyFromThisDashboard: boolean;
    onlyInTimeRange: boolean;
    tags: string[];
    limit: number;
    showUser: boolean;
    showTime: boolean;
    showTags: boolean;
    navigateToPanel: boolean;
    navigateBefore: string;
    navigateAfter: string;
}
export declare const defaultOptions: () => Options;
