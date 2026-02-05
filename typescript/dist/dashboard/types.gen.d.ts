import * as cog from '../cog';
import * as common from '../common';
export interface Dashboard {
    id?: number | null;
    uid?: string;
    title?: string;
    description?: string;
    revision?: number;
    gnetId?: string;
    tags?: string[];
    timezone?: string;
    editable?: boolean;
    graphTooltip?: DashboardCursorSync;
    time?: {
        from: string;
        to: string;
    };
    timepicker?: TimePickerConfig;
    fiscalYearStartMonth?: number;
    liveNow?: boolean;
    weekStart?: string;
    refresh?: string;
    schemaVersion: number;
    version?: number;
    panels?: (Panel | RowPanel)[];
    templating: {
        list?: VariableModel[];
    };
    annotations: AnnotationContainer;
    links?: DashboardLink[];
    snapshot?: Snapshot;
    preload?: boolean;
}
export declare const defaultDashboard: () => Dashboard;
export declare enum DashboardCursorSync {
    Off = 0,
    Crosshair = 1,
    Tooltip = 2
}
export declare const defaultDashboardCursorSync: () => DashboardCursorSync;
export interface TimePickerConfig {
    hidden?: boolean;
    refresh_intervals?: string[];
    quick_ranges?: TimeOption[];
    nowDelay?: string;
}
export declare const defaultTimePickerConfig: () => TimePickerConfig;
export interface TimeOption {
    display: string;
    from: string;
    to: string;
}
export declare const defaultTimeOption: () => TimeOption;
export interface Panel {
    type: string;
    id?: number;
    pluginVersion?: string;
    targets?: cog.Dataquery[];
    title?: string;
    description?: string;
    transparent?: boolean;
    datasource?: common.DataSourceRef;
    gridPos?: GridPos;
    links?: DashboardLink[];
    repeat?: string;
    repeatDirection?: "h" | "v";
    maxPerRow?: number;
    maxDataPoints?: number;
    transformations?: DataTransformerConfig[];
    interval?: string;
    timeFrom?: string;
    timeShift?: string;
    hideTimeOverride?: boolean;
    libraryPanel?: LibraryPanelRef;
    cacheTimeout?: string;
    queryCachingTTL?: number;
    options?: any;
    fieldConfig?: FieldConfigSource;
}
export declare const defaultPanel: () => Panel;
export type DataSourceRef = common.DataSourceRef;
export declare const defaultDataSourceRef: () => DataSourceRef;
export interface GridPos {
    h: number;
    w: number;
    x: number;
    y: number;
    static?: boolean;
}
export declare const defaultGridPos: () => GridPos;
export interface DashboardLink {
    title: string;
    type: DashboardLinkType;
    icon: string;
    tooltip: string;
    url?: string;
    tags: string[];
    asDropdown: boolean;
    targetBlank: boolean;
    includeVars: boolean;
    keepTime: boolean;
}
export declare const defaultDashboardLink: () => DashboardLink;
export declare enum DashboardLinkType {
    Link = "link",
    Dashboards = "dashboards"
}
export declare const defaultDashboardLinkType: () => DashboardLinkType;
export interface DataTransformerConfig {
    id: string;
    disabled?: boolean;
    filter?: MatcherConfig;
    topic?: "series" | "annotations" | "alertStates";
    options: any;
}
export declare const defaultDataTransformerConfig: () => DataTransformerConfig;
export interface MatcherConfig {
    id: string;
    options?: any;
}
export declare const defaultMatcherConfig: () => MatcherConfig;
export interface LibraryPanelRef {
    name: string;
    uid: string;
}
export declare const defaultLibraryPanelRef: () => LibraryPanelRef;
export interface FieldConfigSource {
    defaults: FieldConfig;
    overrides: {
        matcher: MatcherConfig;
        properties: DynamicConfigValue[];
    }[];
}
export declare const defaultFieldConfigSource: () => FieldConfigSource;
export interface FieldConfig {
    displayName?: string;
    displayNameFromDS?: string;
    description?: string;
    path?: string;
    writeable?: boolean;
    filterable?: boolean;
    unit?: string;
    decimals?: number;
    min?: number;
    max?: number;
    mappings?: ValueMapping[];
    thresholds?: ThresholdsConfig;
    color?: FieldColor;
    links?: DashboardLink[];
    noValue?: string;
    custom?: any;
}
export declare const defaultFieldConfig: () => FieldConfig;
export type ValueMapping = ValueMap | RangeMap | RegexMap | SpecialValueMap;
export declare const defaultValueMapping: () => ValueMapping;
export interface ValueMap {
    type: MappingType.ValueToText;
    options: Record<string, ValueMappingResult>;
}
export declare const defaultValueMap: () => ValueMap;
export declare enum MappingType {
    ValueToText = "value",
    RangeToText = "range",
    RegexToText = "regex",
    SpecialValue = "special"
}
export declare const defaultMappingType: () => MappingType;
export interface ValueMappingResult {
    text?: string;
    color?: string;
    icon?: string;
    index?: number;
}
export declare const defaultValueMappingResult: () => ValueMappingResult;
export interface RangeMap {
    type: MappingType.RangeToText;
    options: {
        from: number | null;
        to: number | null;
        result: ValueMappingResult;
    };
}
export declare const defaultRangeMap: () => RangeMap;
export interface RegexMap {
    type: MappingType.RegexToText;
    options: {
        pattern: string;
        result: ValueMappingResult;
    };
}
export declare const defaultRegexMap: () => RegexMap;
export interface SpecialValueMap {
    type: MappingType.SpecialValue;
    options: {
        match: SpecialValueMatch;
        result: ValueMappingResult;
    };
}
export declare const defaultSpecialValueMap: () => SpecialValueMap;
export declare enum SpecialValueMatch {
    True = "true",
    False = "false",
    Null = "null",
    NotANumber = "nan",
    NullAndNan = "null+nan",
    Empty = "empty"
}
export declare const defaultSpecialValueMatch: () => SpecialValueMatch;
export interface ThresholdsConfig {
    mode: ThresholdsMode;
    steps: Threshold[];
}
export declare const defaultThresholdsConfig: () => ThresholdsConfig;
export declare enum ThresholdsMode {
    Absolute = "absolute",
    Percentage = "percentage"
}
export declare const defaultThresholdsMode: () => ThresholdsMode;
export interface Threshold {
    value: number | null;
    color: string;
}
export declare const defaultThreshold: () => Threshold;
export interface FieldColor {
    mode: FieldColorModeId;
    fixedColor?: string;
    seriesBy?: FieldColorSeriesByMode;
}
export declare const defaultFieldColor: () => FieldColor;
export declare enum FieldColorModeId {
    Thresholds = "thresholds",
    PaletteClassic = "palette-classic",
    PaletteClassicByName = "palette-classic-by-name",
    ContinuousGrYlRd = "continuous-GrYlRd",
    ContinuousRdYlGr = "continuous-RdYlGr",
    ContinuousBlYlRd = "continuous-BlYlRd",
    ContinuousYlRd = "continuous-YlRd",
    ContinuousBlPu = "continuous-BlPu",
    ContinuousYlBl = "continuous-YlBl",
    ContinuousBlues = "continuous-blues",
    ContinuousReds = "continuous-reds",
    ContinuousGreens = "continuous-greens",
    ContinuousPurples = "continuous-purples",
    Fixed = "fixed",
    Shades = "shades"
}
export declare const defaultFieldColorModeId: () => FieldColorModeId;
export declare enum FieldColorSeriesByMode {
    Min = "min",
    Max = "max",
    Last = "last"
}
export declare const defaultFieldColorSeriesByMode: () => FieldColorSeriesByMode;
export interface DynamicConfigValue {
    id: string;
    value?: any;
}
export declare const defaultDynamicConfigValue: () => DynamicConfigValue;
export interface RowPanel {
    type: "row";
    collapsed: boolean;
    title?: string;
    datasource?: common.DataSourceRef;
    gridPos?: GridPos;
    id: number;
    panels: Panel[];
    repeat?: string;
}
export declare const defaultRowPanel: () => RowPanel;
export interface VariableModel {
    type: VariableType;
    name: string;
    label?: string;
    hide?: VariableHide;
    skipUrlSync?: boolean;
    description?: string;
    query?: string | Record<string, any>;
    datasource?: common.DataSourceRef;
    current?: VariableOption;
    multi?: boolean;
    allowCustomValue?: boolean;
    options?: VariableOption[];
    refresh?: VariableRefresh;
    sort?: VariableSort;
    includeAll?: boolean;
    allValue?: string;
    regex?: string;
    auto?: boolean;
    auto_min?: string;
    auto_count?: number;
}
export declare const defaultVariableModel: () => VariableModel;
export declare enum VariableType {
    Query = "query",
    Adhoc = "adhoc",
    Groupby = "groupby",
    Constant = "constant",
    Datasource = "datasource",
    Interval = "interval",
    Textbox = "textbox",
    Custom = "custom",
    System = "system",
    Snapshot = "snapshot"
}
export declare const defaultVariableType: () => VariableType;
export declare enum VariableHide {
    DontHide = 0,
    HideLabel = 1,
    HideVariable = 2
}
export declare const defaultVariableHide: () => VariableHide;
export interface VariableOption {
    selected?: boolean;
    text: string | string[];
    value: string | string[];
}
export declare const defaultVariableOption: () => VariableOption;
export declare enum VariableRefresh {
    Never = 0,
    OnDashboardLoad = 1,
    OnTimeRangeChanged = 2
}
export declare const defaultVariableRefresh: () => VariableRefresh;
export declare enum VariableSort {
    Disabled = 0,
    AlphabeticalAsc = 1,
    AlphabeticalDesc = 2,
    NumericalAsc = 3,
    NumericalDesc = 4,
    AlphabeticalCaseInsensitiveAsc = 5,
    AlphabeticalCaseInsensitiveDesc = 6,
    NaturalAsc = 7,
    NaturalDesc = 8
}
export declare const defaultVariableSort: () => VariableSort;
export interface AnnotationContainer {
    list?: AnnotationQuery[];
}
export declare const defaultAnnotationContainer: () => AnnotationContainer;
export interface AnnotationQuery {
    name: string;
    datasource: common.DataSourceRef;
    enable: boolean;
    hide?: boolean;
    iconColor: string;
    filter?: AnnotationPanelFilter;
    target?: AnnotationTarget;
    type?: string;
    builtIn?: number;
    expr?: string;
}
export declare const defaultAnnotationQuery: () => AnnotationQuery;
export interface AnnotationPanelFilter {
    exclude?: boolean;
    ids: number[];
}
export declare const defaultAnnotationPanelFilter: () => AnnotationPanelFilter;
export interface AnnotationTarget {
    limit: number;
    matchAny: boolean;
    tags: string[];
    type: string;
}
export declare const defaultAnnotationTarget: () => AnnotationTarget;
export interface Snapshot {
    created: string;
    expires: string;
    external: boolean;
    externalUrl: string;
    originalUrl: string;
    id: number;
    key: string;
    name: string;
    orgId: number;
    updated: string;
    url?: string;
    userId: number;
    dashboard?: Dashboard;
}
export declare const defaultSnapshot: () => Snapshot;
export interface AnnotationActions {
    canAdd?: boolean;
    canDelete?: boolean;
    canEdit?: boolean;
}
export declare const defaultAnnotationActions: () => AnnotationActions;
export interface AnnotationPermission {
    dashboard?: AnnotationActions;
    organization?: AnnotationActions;
}
export declare const defaultAnnotationPermission: () => AnnotationPermission;
export interface DashboardMeta {
    annotationsPermissions?: AnnotationPermission;
    apiVersion?: string;
    canAdmin?: boolean;
    canDelete?: boolean;
    canEdit?: boolean;
    canSave?: boolean;
    canStar?: boolean;
    created?: string;
    createdBy?: string;
    expires?: string;
    folderId?: number;
    folderTitle?: string;
    folderUid?: string;
    folderUrl?: string;
    hasAcl?: boolean;
    isFolder?: boolean;
    isSnapshot?: boolean;
    isStarred?: boolean;
    provisioned?: boolean;
    provisionedExternalId?: string;
    publicDashboardEnabled?: boolean;
    slug?: string;
    type?: string;
    updated?: string;
    updatedBy?: string;
    url?: string;
    version?: number;
}
export declare const defaultDashboardMeta: () => DashboardMeta;
