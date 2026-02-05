export interface Dashboard {
    annotations: AnnotationQueryKind[];
    cursorSync: DashboardCursorSync;
    description?: string;
    editable?: boolean;
    elements: Record<string, Element>;
    layout: GridLayoutKind | RowsLayoutKind | AutoGridLayoutKind | TabsLayoutKind;
    links: DashboardLink[];
    liveNow?: boolean;
    preload: boolean;
    revision?: number;
    tags: string[];
    timeSettings: TimeSettingsSpec;
    title: string;
    variables: VariableKind[];
}
export declare const defaultDashboard: () => Dashboard;
export interface AnnotationQueryKind {
    kind: "AnnotationQuery";
    spec: AnnotationQuerySpec;
}
export declare const defaultAnnotationQueryKind: () => AnnotationQueryKind;
export interface AnnotationQuerySpec {
    query: DataQueryKind;
    enable: boolean;
    hide: boolean;
    iconColor: string;
    name: string;
    builtIn?: boolean;
    filter?: AnnotationPanelFilter;
    placement?: "inControlsMenu";
    legacyOptions?: Record<string, any>;
}
export declare const defaultAnnotationQuerySpec: () => AnnotationQuerySpec;
export interface DataQueryKind {
    kind: "DataQuery";
    group: string;
    version: string;
    datasource?: {
        name?: string;
    };
    spec: any;
}
export declare const defaultDataQueryKind: () => DataQueryKind;
export interface AnnotationPanelFilter {
    exclude?: boolean;
    ids: number[];
}
export declare const defaultAnnotationPanelFilter: () => AnnotationPanelFilter;
export declare const AnnotationQueryPlacement = "inControlsMenu";
export declare enum DashboardCursorSync {
    Crosshair = "Crosshair",
    Tooltip = "Tooltip",
    Off = "Off"
}
export declare const defaultDashboardCursorSync: () => DashboardCursorSync;
export type Element = PanelKind | LibraryPanelKind;
export declare const defaultElement: () => Element;
export interface PanelKind {
    kind: "Panel";
    spec: PanelSpec;
}
export declare const defaultPanelKind: () => PanelKind;
export interface PanelSpec {
    id: number;
    title: string;
    description: string;
    links: DataLink[];
    data: QueryGroupKind;
    vizConfig: VizConfigKind;
    transparent?: boolean;
}
export declare const defaultPanelSpec: () => PanelSpec;
export interface DataLink {
    title: string;
    url: string;
    targetBlank?: boolean;
}
export declare const defaultDataLink: () => DataLink;
export interface QueryGroupKind {
    kind: "QueryGroup";
    spec: QueryGroupSpec;
}
export declare const defaultQueryGroupKind: () => QueryGroupKind;
export interface QueryGroupSpec {
    queries: PanelQueryKind[];
    transformations: TransformationKind[];
    queryOptions: QueryOptionsSpec;
}
export declare const defaultQueryGroupSpec: () => QueryGroupSpec;
export interface PanelQueryKind {
    kind: "PanelQuery";
    spec: PanelQuerySpec;
}
export declare const defaultPanelQueryKind: () => PanelQueryKind;
export interface PanelQuerySpec {
    query: DataQueryKind;
    refId: string;
    hidden: boolean;
}
export declare const defaultPanelQuerySpec: () => PanelQuerySpec;
export interface TransformationKind {
    kind: string;
    spec: DataTransformerConfig;
}
export declare const defaultTransformationKind: () => TransformationKind;
export interface DataTransformerConfig {
    id: string;
    disabled?: boolean;
    filter?: MatcherConfig;
    topic?: DataTopic;
    options: any;
}
export declare const defaultDataTransformerConfig: () => DataTransformerConfig;
export interface MatcherConfig {
    id: string;
    options?: any;
}
export declare const defaultMatcherConfig: () => MatcherConfig;
export declare enum DataTopic {
    Series = "series",
    Annotations = "annotations",
    AlertStates = "alertStates"
}
export declare const defaultDataTopic: () => DataTopic;
export interface QueryOptionsSpec {
    timeFrom?: string;
    maxDataPoints?: number;
    timeShift?: string;
    queryCachingTTL?: number;
    interval?: string;
    cacheTimeout?: string;
    hideTimeOverride?: boolean;
    timeCompare?: string;
}
export declare const defaultQueryOptionsSpec: () => QueryOptionsSpec;
export interface VizConfigKind {
    kind: "VizConfig";
    group: string;
    version: string;
    spec: VizConfigSpec;
}
export declare const defaultVizConfigKind: () => VizConfigKind;
export interface VizConfigSpec {
    options: any;
    fieldConfig: FieldConfigSource;
}
export declare const defaultVizConfigSpec: () => VizConfigSpec;
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
    links?: any[];
    actions?: Action[];
    noValue?: string;
    custom?: any;
}
export declare const defaultFieldConfig: () => FieldConfig;
export type ValueMapping = ValueMap | RangeMap | RegexMap | SpecialValueMap;
export declare const defaultValueMapping: () => ValueMapping;
export interface ValueMap {
    type: MappingType.Value;
    options: Record<string, ValueMappingResult>;
}
export declare const defaultValueMap: () => ValueMap;
export declare enum MappingType {
    Value = "value",
    Range = "range",
    Regex = "regex",
    Special = "special"
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
    type: MappingType.Range;
    options: {
        from: number | null;
        to: number | null;
        result: ValueMappingResult;
    };
}
export declare const defaultRangeMap: () => RangeMap;
export interface RegexMap {
    type: MappingType.Regex;
    options: {
        pattern: string;
        result: ValueMappingResult;
    };
}
export declare const defaultRegexMap: () => RegexMap;
export interface SpecialValueMap {
    type: MappingType.Special;
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
    NullAndNaN = "null+nan",
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
    value: number;
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
export interface Action {
    type: ActionType;
    title: string;
    fetch?: FetchOptions;
    infinity?: InfinityOptions;
    confirmation?: string;
    oneClick?: boolean;
    variables?: ActionVariable[];
    style?: {
        backgroundColor?: string;
    };
}
export declare const defaultAction: () => Action;
export declare enum ActionType {
    Fetch = "fetch",
    Infinity = "infinity"
}
export declare const defaultActionType: () => ActionType;
export interface FetchOptions {
    method: HttpRequestMethod;
    url: string;
    body?: string;
    queryParams?: string[][];
    headers?: string[][];
}
export declare const defaultFetchOptions: () => FetchOptions;
export declare enum HttpRequestMethod {
    GET = "GET",
    PUT = "PUT",
    POST = "POST",
    DELETE = "DELETE",
    PATCH = "PATCH"
}
export declare const defaultHttpRequestMethod: () => HttpRequestMethod;
export interface InfinityOptions {
    method: HttpRequestMethod;
    url: string;
    body?: string;
    queryParams?: string[][];
    datasourceUid: string;
    headers?: string[][];
}
export declare const defaultInfinityOptions: () => InfinityOptions;
export interface ActionVariable {
    key: string;
    name: string;
    type: "string";
}
export declare const defaultActionVariable: () => ActionVariable;
export declare const ActionVariableType = "string";
export interface DynamicConfigValue {
    id: string;
    value?: any;
}
export declare const defaultDynamicConfigValue: () => DynamicConfigValue;
export interface LibraryPanelKind {
    kind: "LibraryPanel";
    spec: LibraryPanelKindSpec;
}
export declare const defaultLibraryPanelKind: () => LibraryPanelKind;
export interface LibraryPanelKindSpec {
    id: number;
    title: string;
    libraryPanel: LibraryPanelRef;
}
export declare const defaultLibraryPanelKindSpec: () => LibraryPanelKindSpec;
export interface LibraryPanelRef {
    name: string;
    uid: string;
}
export declare const defaultLibraryPanelRef: () => LibraryPanelRef;
export interface GridLayoutKind {
    kind: "GridLayout";
    spec: GridLayoutSpec;
}
export declare const defaultGridLayoutKind: () => GridLayoutKind;
export interface GridLayoutSpec {
    items: GridLayoutItemKind[];
}
export declare const defaultGridLayoutSpec: () => GridLayoutSpec;
export interface GridLayoutItemKind {
    kind: "GridLayoutItem";
    spec: GridLayoutItemSpec;
}
export declare const defaultGridLayoutItemKind: () => GridLayoutItemKind;
export interface GridLayoutItemSpec {
    x: number;
    y: number;
    width: number;
    height: number;
    element: ElementReference;
    repeat?: RepeatOptions;
}
export declare const defaultGridLayoutItemSpec: () => GridLayoutItemSpec;
export interface ElementReference {
    kind: "ElementReference";
    name: string;
}
export declare const defaultElementReference: () => ElementReference;
export interface RepeatOptions {
    mode: RepeatMode.Variable;
    value: string;
    direction?: "h" | "v";
    maxPerRow?: number;
}
export declare const defaultRepeatOptions: () => RepeatOptions;
export declare enum RepeatMode {
    Variable = "variable"
}
export declare const defaultRepeatMode: () => RepeatMode;
export interface RowsLayoutKind {
    kind: "RowsLayout";
    spec: RowsLayoutSpec;
}
export declare const defaultRowsLayoutKind: () => RowsLayoutKind;
export interface RowsLayoutSpec {
    rows: RowsLayoutRowKind[];
}
export declare const defaultRowsLayoutSpec: () => RowsLayoutSpec;
export interface RowsLayoutRowKind {
    kind: "RowsLayoutRow";
    spec: RowsLayoutRowSpec;
}
export declare const defaultRowsLayoutRowKind: () => RowsLayoutRowKind;
export interface RowsLayoutRowSpec {
    title?: string;
    collapse?: boolean;
    hideHeader?: boolean;
    fillScreen?: boolean;
    conditionalRendering?: ConditionalRenderingGroupKind;
    repeat?: RowRepeatOptions;
    layout: GridLayoutKind | AutoGridLayoutKind | TabsLayoutKind | RowsLayoutKind;
}
export declare const defaultRowsLayoutRowSpec: () => RowsLayoutRowSpec;
export interface ConditionalRenderingGroupKind {
    kind: "ConditionalRenderingGroup";
    spec: ConditionalRenderingGroupSpec;
}
export declare const defaultConditionalRenderingGroupKind: () => ConditionalRenderingGroupKind;
export interface ConditionalRenderingGroupSpec {
    visibility: "show" | "hide";
    condition: "and" | "or";
    items: (ConditionalRenderingVariableKind | ConditionalRenderingDataKind | ConditionalRenderingTimeRangeSizeKind)[];
}
export declare const defaultConditionalRenderingGroupSpec: () => ConditionalRenderingGroupSpec;
export interface ConditionalRenderingVariableKind {
    kind: "ConditionalRenderingVariable";
    spec: ConditionalRenderingVariableSpec;
}
export declare const defaultConditionalRenderingVariableKind: () => ConditionalRenderingVariableKind;
export interface ConditionalRenderingVariableSpec {
    variable: string;
    operator: "equals" | "notEquals" | "matches" | "notMatches";
    value: string;
}
export declare const defaultConditionalRenderingVariableSpec: () => ConditionalRenderingVariableSpec;
export interface ConditionalRenderingDataKind {
    kind: "ConditionalRenderingData";
    spec: ConditionalRenderingDataSpec;
}
export declare const defaultConditionalRenderingDataKind: () => ConditionalRenderingDataKind;
export interface ConditionalRenderingDataSpec {
    value: boolean;
}
export declare const defaultConditionalRenderingDataSpec: () => ConditionalRenderingDataSpec;
export interface ConditionalRenderingTimeRangeSizeKind {
    kind: "ConditionalRenderingTimeRangeSize";
    spec: ConditionalRenderingTimeRangeSizeSpec;
}
export declare const defaultConditionalRenderingTimeRangeSizeKind: () => ConditionalRenderingTimeRangeSizeKind;
export interface ConditionalRenderingTimeRangeSizeSpec {
    value: string;
}
export declare const defaultConditionalRenderingTimeRangeSizeSpec: () => ConditionalRenderingTimeRangeSizeSpec;
export interface RowRepeatOptions {
    mode: RepeatMode.Variable;
    value: string;
}
export declare const defaultRowRepeatOptions: () => RowRepeatOptions;
export interface AutoGridLayoutKind {
    kind: "AutoGridLayout";
    spec: AutoGridLayoutSpec;
}
export declare const defaultAutoGridLayoutKind: () => AutoGridLayoutKind;
export interface AutoGridLayoutSpec {
    maxColumnCount?: number;
    columnWidthMode: "narrow" | "standard" | "wide" | "custom";
    columnWidth?: number;
    rowHeightMode: "short" | "standard" | "tall" | "custom";
    rowHeight?: number;
    fillScreen?: boolean;
    items: AutoGridLayoutItemKind[];
}
export declare const defaultAutoGridLayoutSpec: () => AutoGridLayoutSpec;
export interface AutoGridLayoutItemKind {
    kind: "AutoGridLayoutItem";
    spec: AutoGridLayoutItemSpec;
}
export declare const defaultAutoGridLayoutItemKind: () => AutoGridLayoutItemKind;
export interface AutoGridLayoutItemSpec {
    element: ElementReference;
    repeat?: AutoGridRepeatOptions;
    conditionalRendering?: ConditionalRenderingGroupKind;
}
export declare const defaultAutoGridLayoutItemSpec: () => AutoGridLayoutItemSpec;
export interface AutoGridRepeatOptions {
    mode: RepeatMode.Variable;
    value: string;
}
export declare const defaultAutoGridRepeatOptions: () => AutoGridRepeatOptions;
export interface TabsLayoutKind {
    kind: "TabsLayout";
    spec: TabsLayoutSpec;
}
export declare const defaultTabsLayoutKind: () => TabsLayoutKind;
export interface TabsLayoutSpec {
    tabs: TabsLayoutTabKind[];
}
export declare const defaultTabsLayoutSpec: () => TabsLayoutSpec;
export interface TabsLayoutTabKind {
    kind: "TabsLayoutTab";
    spec: TabsLayoutTabSpec;
}
export declare const defaultTabsLayoutTabKind: () => TabsLayoutTabKind;
export interface TabsLayoutTabSpec {
    title?: string;
    layout: GridLayoutKind | RowsLayoutKind | AutoGridLayoutKind | TabsLayoutKind;
    conditionalRendering?: ConditionalRenderingGroupKind;
    repeat?: TabRepeatOptions;
}
export declare const defaultTabsLayoutTabSpec: () => TabsLayoutTabSpec;
export interface TabRepeatOptions {
    mode: RepeatMode.Variable;
    value: string;
}
export declare const defaultTabRepeatOptions: () => TabRepeatOptions;
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
    placement?: "inControlsMenu";
}
export declare const defaultDashboardLink: () => DashboardLink;
export declare enum DashboardLinkType {
    Link = "link",
    Dashboards = "dashboards"
}
export declare const defaultDashboardLinkType: () => DashboardLinkType;
export declare const DashboardLinkPlacement = "inControlsMenu";
export interface TimeSettingsSpec {
    timezone?: string;
    from: string;
    to: string;
    autoRefresh: string;
    autoRefreshIntervals: string[];
    quickRanges?: TimeRangeOption[];
    hideTimepicker: boolean;
    weekStart?: "saturday" | "monday" | "sunday";
    fiscalYearStartMonth: number;
    nowDelay?: string;
}
export declare const defaultTimeSettingsSpec: () => TimeSettingsSpec;
export interface TimeRangeOption {
    display: string;
    from: string;
    to: string;
}
export declare const defaultTimeRangeOption: () => TimeRangeOption;
export type VariableKind = QueryVariableKind | TextVariableKind | ConstantVariableKind | DatasourceVariableKind | IntervalVariableKind | CustomVariableKind | GroupByVariableKind | AdhocVariableKind | SwitchVariableKind;
export declare const defaultVariableKind: () => VariableKind;
export interface QueryVariableKind {
    kind: "QueryVariable";
    spec: QueryVariableSpec;
}
export declare const defaultQueryVariableKind: () => QueryVariableKind;
export interface QueryVariableSpec {
    name: string;
    current: VariableOption;
    label?: string;
    hide: VariableHide;
    refresh: VariableRefresh;
    skipUrlSync: boolean;
    description?: string;
    query: DataQueryKind;
    regex: string;
    sort: VariableSort;
    definition?: string;
    options: VariableOption[];
    multi: boolean;
    includeAll: boolean;
    allValue?: string;
    placeholder?: string;
    allowCustomValue: boolean;
    staticOptions?: VariableOption[];
    staticOptionsOrder?: "before" | "after" | "sorted";
}
export declare const defaultQueryVariableSpec: () => QueryVariableSpec;
export interface VariableOption {
    selected?: boolean;
    text: string | string[];
    value: string | string[];
}
export declare const defaultVariableOption: () => VariableOption;
export declare enum VariableHide {
    DontHide = "dontHide",
    HideLabel = "hideLabel",
    HideVariable = "hideVariable",
    InControlsMenu = "inControlsMenu"
}
export declare const defaultVariableHide: () => VariableHide;
export declare enum VariableRefresh {
    Never = "never",
    OnDashboardLoad = "onDashboardLoad",
    OnTimeRangeChanged = "onTimeRangeChanged"
}
export declare const defaultVariableRefresh: () => VariableRefresh;
export declare enum VariableSort {
    Disabled = "disabled",
    AlphabeticalAsc = "alphabeticalAsc",
    AlphabeticalDesc = "alphabeticalDesc",
    NumericalAsc = "numericalAsc",
    NumericalDesc = "numericalDesc",
    AlphabeticalCaseInsensitiveAsc = "alphabeticalCaseInsensitiveAsc",
    AlphabeticalCaseInsensitiveDesc = "alphabeticalCaseInsensitiveDesc",
    NaturalAsc = "naturalAsc",
    NaturalDesc = "naturalDesc"
}
export declare const defaultVariableSort: () => VariableSort;
export interface TextVariableKind {
    kind: "TextVariable";
    spec: TextVariableSpec;
}
export declare const defaultTextVariableKind: () => TextVariableKind;
export interface TextVariableSpec {
    name: string;
    current: VariableOption;
    query: string;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
}
export declare const defaultTextVariableSpec: () => TextVariableSpec;
export interface ConstantVariableKind {
    kind: "ConstantVariable";
    spec: ConstantVariableSpec;
}
export declare const defaultConstantVariableKind: () => ConstantVariableKind;
export interface ConstantVariableSpec {
    name: string;
    query: string;
    current: VariableOption;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
}
export declare const defaultConstantVariableSpec: () => ConstantVariableSpec;
export interface DatasourceVariableKind {
    kind: "DatasourceVariable";
    spec: DatasourceVariableSpec;
}
export declare const defaultDatasourceVariableKind: () => DatasourceVariableKind;
export interface DatasourceVariableSpec {
    name: string;
    pluginId: string;
    refresh: VariableRefresh;
    regex: string;
    current: VariableOption;
    options: VariableOption[];
    multi: boolean;
    includeAll: boolean;
    allValue?: string;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
    allowCustomValue: boolean;
}
export declare const defaultDatasourceVariableSpec: () => DatasourceVariableSpec;
export interface IntervalVariableKind {
    kind: "IntervalVariable";
    spec: IntervalVariableSpec;
}
export declare const defaultIntervalVariableKind: () => IntervalVariableKind;
export interface IntervalVariableSpec {
    name: string;
    query: string;
    current: VariableOption;
    options: VariableOption[];
    auto: boolean;
    auto_min: string;
    auto_count: number;
    refresh: VariableRefresh;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
}
export declare const defaultIntervalVariableSpec: () => IntervalVariableSpec;
export interface CustomVariableKind {
    kind: "CustomVariable";
    spec: CustomVariableSpec;
}
export declare const defaultCustomVariableKind: () => CustomVariableKind;
export interface CustomVariableSpec {
    name: string;
    query: string;
    current: VariableOption;
    options: VariableOption[];
    multi: boolean;
    includeAll: boolean;
    allValue?: string;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
    allowCustomValue: boolean;
}
export declare const defaultCustomVariableSpec: () => CustomVariableSpec;
export interface GroupByVariableKind {
    kind: "GroupByVariable";
    group: string;
    datasource?: {
        name?: string;
    };
    spec: GroupByVariableSpec;
}
export declare const defaultGroupByVariableKind: () => GroupByVariableKind;
export interface GroupByVariableSpec {
    name: string;
    defaultValue?: VariableOption;
    current: VariableOption;
    options: VariableOption[];
    multi: boolean;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
}
export declare const defaultGroupByVariableSpec: () => GroupByVariableSpec;
export interface AdhocVariableKind {
    kind: "AdhocVariable";
    group: string;
    datasource?: {
        name?: string;
    };
    spec: AdhocVariableSpec;
}
export declare const defaultAdhocVariableKind: () => AdhocVariableKind;
export interface AdhocVariableSpec {
    name: string;
    baseFilters: AdHocFilterWithLabels[];
    filters: AdHocFilterWithLabels[];
    defaultKeys: MetricFindValue[];
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
    allowCustomValue: boolean;
}
export declare const defaultAdhocVariableSpec: () => AdhocVariableSpec;
export interface AdHocFilterWithLabels {
    key: string;
    operator: string;
    value: string;
    values?: string[];
    keyLabel?: string;
    valueLabels?: string[];
    forceEdit?: boolean;
    origin?: "dashboard";
    condition?: string;
}
export declare const defaultAdHocFilterWithLabels: () => AdHocFilterWithLabels;
export declare const FilterOrigin = "dashboard";
export interface MetricFindValue {
    text: string;
    value?: string | number;
    group?: string;
    expandable?: boolean;
}
export declare const defaultMetricFindValue: () => MetricFindValue;
export interface SwitchVariableKind {
    kind: "SwitchVariable";
    spec: SwitchVariableSpec;
}
export declare const defaultSwitchVariableKind: () => SwitchVariableKind;
export interface SwitchVariableSpec {
    name: string;
    current: string;
    enabledValue: string;
    disabledValue: string;
    label?: string;
    hide: VariableHide;
    skipUrlSync: boolean;
    description?: string;
}
export declare const defaultSwitchVariableSpec: () => SwitchVariableSpec;
export interface Kind {
    kind: string;
    spec: any;
    metadata?: any;
}
export declare const defaultKind: () => Kind;
export type VariableValue = VariableValueSingle | VariableValueSingle[];
export declare const defaultVariableValue: () => VariableValue;
export type VariableValueSingle = string | boolean | number | CustomVariableValue;
export declare const defaultVariableValueSingle: () => VariableValueSingle;
export interface CustomVariableValue {
    formatter: string;
}
export declare const defaultCustomVariableValue: () => CustomVariableValue;
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
    Snapshot = "snapshot",
    Switch = "switch"
}
export declare const defaultVariableType: () => VariableType;
export interface CustomFormatterVariable {
    name: string;
    type: VariableType;
    multi: boolean;
    includeAll: boolean;
}
export declare const defaultCustomFormatterVariable: () => CustomFormatterVariable;
export interface VariableValueOption {
    label: string;
    value: VariableValueSingle;
    group?: string;
}
export declare const defaultVariableValueOption: () => VariableValueOption;
