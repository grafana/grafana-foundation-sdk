import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class DashboardBuilder implements cog.Builder<dashboardv2beta1.Dashboard> {
    protected readonly internal: dashboardv2beta1.Dashboard;
    constructor(title: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.Dashboard;
    annotations(annotations: cog.Builder<dashboardv2beta1.AnnotationQueryKind>[]): this;
    cursorSync(cursorSync: dashboardv2beta1.DashboardCursorSync): this;
    description(description: string): this;
    editable(editable: boolean): this;
    elements(elements: Record<string, cog.Builder<dashboardv2beta1.Element>>): this;
    element(key: string, element: cog.Builder<dashboardv2beta1.Element>): this;
    gridLayout(gridLayoutKind: cog.Builder<dashboardv2beta1.GridLayoutKind>): this;
    rowsLayout(rowsLayoutKind: cog.Builder<dashboardv2beta1.RowsLayoutKind>): this;
    autoGridLayout(autoGridLayoutKind: cog.Builder<dashboardv2beta1.AutoGridLayoutKind>): this;
    tabsLayout(tabsLayoutKind: cog.Builder<dashboardv2beta1.TabsLayoutKind>): this;
    links(links: cog.Builder<dashboardv2beta1.DashboardLink>[]): this;
    liveNow(liveNow: boolean): this;
    preload(preload: boolean): this;
    revision(revision: number): this;
    tags(tags: string[]): this;
    timeSettings(timeSettings: cog.Builder<dashboardv2beta1.TimeSettingsSpec>): this;
    title(title: string): this;
    variables(variables: cog.Builder<dashboardv2beta1.VariableKind>[]): this;
    variable(variable: cog.Builder<dashboardv2beta1.VariableKind>): this;
}
