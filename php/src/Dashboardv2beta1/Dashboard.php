<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboard implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind>
     */
    public array $annotations;

    /**
     * Configuration of dashboard cursor sync behavior.
     * "Off" for no shared crosshair or tooltip (default).
     * "Crosshair" for shared crosshair.
     * "Tooltip" for shared crosshair AND shared tooltip.
     */
    public \Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync $cursorSync;

    /**
     * Description of dashboard.
     */
    public ?string $description;

    /**
     * Whether a dashboard is editable or not.
     */
    public ?bool $editable;

    /**
     * @var array<string, \Grafana\Foundation\Dashboardv2beta1\PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>
     */
    public array $elements;

    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind
     */
    public $layout;

    /**
     * Links with references to other dashboards or external websites.
     * @var array<\Grafana\Foundation\Dashboardv2beta1\DashboardLink>
     */
    public array $links;

    /**
     * When set to true, the dashboard will redraw panels at an interval matching the pixel width.
     * This will keep data "moving left" regardless of the query refresh rate. This setting helps
     * avoid dashboards presenting stale live data.
     */
    public ?bool $liveNow;

    /**
     * When set to true, the dashboard will load all panels in the dashboard when it's loaded.
     */
    public bool $preload;

    /**
     * Plugins only. The version of the dashboard installed together with the plugin.
     * This is used to determine if the dashboard should be updated when the plugin is updated.
     */
    public ?int $revision;

    /**
     * Tags associated with dashboard.
     * @var array<string>
     */
    public array $tags;

    public \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $timeSettings;

    /**
     * Title of dashboard.
     */
    public string $title;

    /**
     * Configured template variables.
     * @var array<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind|\Grafana\Foundation\Dashboardv2beta1\TextVariableKind|\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind|\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind|\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind|\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind|\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind|\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>
     */
    public array $variables;

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind>|null $annotations
     * @param \Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync|null $cursorSync
     * @param string|null $description
     * @param bool|null $editable
     * @param array<string, \Grafana\Foundation\Dashboardv2beta1\PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>|null $elements
     * @param \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind|null $layout
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DashboardLink>|null $links
     * @param bool|null $liveNow
     * @param bool|null $preload
     * @param int|null $revision
     * @param array<string>|null $tags
     * @param \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec|null $timeSettings
     * @param string|null $title
     * @param array<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind|\Grafana\Foundation\Dashboardv2beta1\TextVariableKind|\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind|\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind|\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind|\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind|\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind|\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>|null $variables
     */
    public function __construct(?array $annotations = null, ?\Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync $cursorSync = null, ?string $description = null, ?bool $editable = null, ?array $elements = null,  $layout = null, ?array $links = null, ?bool $liveNow = null, ?bool $preload = null, ?int $revision = null, ?array $tags = null, ?\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $timeSettings = null, ?string $title = null, ?array $variables = null)
    {
        $this->annotations = $annotations ?: [];
        $this->cursorSync = $cursorSync ?: \Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync::off();
        $this->description = $description;
        $this->editable = $editable;
        $this->elements = $elements ?: [];
        $this->layout = $layout ?: new \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind();
        $this->links = $links ?: [];
        $this->liveNow = $liveNow;
        $this->preload = $preload ?: false;
        $this->revision = $revision;
        $this->tags = $tags ?: [];
        $this->timeSettings = $timeSettings ?: new \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec();
        $this->title = $title ?: "";
        $this->variables = $variables ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{annotations?: array<mixed>, cursorSync?: string, description?: string, editable?: bool, elements?: array<string, mixed|mixed>, layout?: mixed|mixed|mixed|mixed, links?: array<mixed>, liveNow?: bool, preload?: bool, revision?: int, tags?: array<string>, timeSettings?: mixed, title?: string, variables?: array<mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed>} $inputData */
        $data = $inputData;
        return new self(
            annotations: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind::fromArray($val);
    }), $data["annotations"] ?? [])),
            cursorSync: isset($data["cursorSync"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync::fromValue($input); })($data["cursorSync"]) : null,
            description: $data["description"] ?? null,
            editable: $data["editable"] ?? null,
            elements: isset($data["elements"]) ? (function($input) {
        /** @var array<string, \Grafana\Foundation\Dashboardv2beta1\PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = isset($val) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "LibraryPanel":
            return LibraryPanelKind::fromArray($input);
        case "Panel":
            return PanelKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($val) : null;
        }
        return array_filter($results);
    })($data["elements"]) : null,
            layout: isset($data["layout"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "AutoGridLayout":
            return AutoGridLayoutKind::fromArray($input);
        case "GridLayout":
            return GridLayoutKind::fromArray($input);
        case "RowsLayout":
            return RowsLayoutKind::fromArray($input);
        case "TabsLayout":
            return TabsLayoutKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["layout"]) : null,
            links: array_filter(array_map((function($input) {
    	/** @var array{title?: string, type?: string, icon?: string, tooltip?: string, url?: string, tags?: array<string>, asDropdown?: bool, targetBlank?: bool, includeVars?: bool, keepTime?: bool, placement?: "inControlsMenu"} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DashboardLink::fromArray($val);
    }), $data["links"] ?? [])),
            liveNow: $data["liveNow"] ?? null,
            preload: $data["preload"] ?? null,
            revision: $data["revision"] ?? null,
            tags: $data["tags"] ?? null,
            timeSettings: isset($data["timeSettings"]) ? (function($input) {
    	/** @var array{timezone?: string, from?: string, to?: string, autoRefresh?: string, autoRefreshIntervals?: array<string>, quickRanges?: array<mixed>, hideTimepicker?: bool, weekStart?: string, fiscalYearStartMonth?: int, nowDelay?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec::fromArray($val);
    })($data["timeSettings"]) : null,
            title: $data["title"] ?? null,
            variables: !empty($data["variables"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "AdhocVariable":
            return AdhocVariableKind::fromArray($input);
        case "ConstantVariable":
            return ConstantVariableKind::fromArray($input);
        case "CustomVariable":
            return CustomVariableKind::fromArray($input);
        case "DatasourceVariable":
            return DatasourceVariableKind::fromArray($input);
        case "GroupByVariable":
            return GroupByVariableKind::fromArray($input);
        case "IntervalVariable":
            return IntervalVariableKind::fromArray($input);
        case "QueryVariable":
            return QueryVariableKind::fromArray($input);
        case "SwitchVariable":
            return SwitchVariableKind::fromArray($input);
        case "TextVariable":
            return TextVariableKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["variables"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->annotations = $this->annotations;
        $data->cursorSync = $this->cursorSync;
        $data->elements = $this->elements;
        $data->layout = $this->layout;
        $data->links = $this->links;
        $data->preload = $this->preload;
        $data->tags = $this->tags;
        $data->timeSettings = $this->timeSettings;
        $data->title = $this->title;
        $data->variables = $this->variables;
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->editable)) {
            $data->editable = $this->editable;
        }
        if (isset($this->liveNow)) {
            $data->liveNow = $this->liveNow;
        }
        if (isset($this->revision)) {
            $data->revision = $this->revision;
        }
        return $data;
    }
}
