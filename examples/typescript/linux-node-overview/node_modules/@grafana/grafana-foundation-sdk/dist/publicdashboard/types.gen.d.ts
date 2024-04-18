export interface PublicDashboard {
    uid: string;
    dashboardUid: string;
    accessToken?: string;
    isEnabled: boolean;
    annotationsEnabled: boolean;
    timeSelectionEnabled: boolean;
}
export declare const defaultPublicDashboard: () => PublicDashboard;
