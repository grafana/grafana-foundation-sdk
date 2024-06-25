<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Links with references to other dashboards or external resources
 */
class DashboardLink implements \JsonSerializable
{
    /**
     * Title to display with the link
     */
    public string $title;

    /**
     * Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
     */
    public \Grafana\Foundation\Dashboard\DashboardLinkType $type;

    /**
     * Icon name to be displayed with the link
     */
    public string $icon;

    /**
     * Tooltip to display when the user hovers their mouse over it
     */
    public string $tooltip;

    /**
     * Link URL. Only required/valid if the type is link
     */
    public string $url;

    /**
     * List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
     * @var array<string>
     */
    public array $tags;

    /**
     * If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
     */
    public bool $asDropdown;

    /**
     * If true, the link will be opened in a new tab
     */
    public bool $targetBlank;

    /**
     * If true, includes current template variables values in the link as query params
     */
    public bool $includeVars;

    /**
     * If true, includes current time range in the link as query params
     */
    public bool $keepTime;

    /**
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboard\DashboardLinkType|null $type
     * @param string|null $icon
     * @param string|null $tooltip
     * @param string|null $url
     * @param array<string>|null $tags
     * @param bool|null $asDropdown
     * @param bool|null $targetBlank
     * @param bool|null $includeVars
     * @param bool|null $keepTime
     */
    public function __construct(?string $title = null, ?\Grafana\Foundation\Dashboard\DashboardLinkType $type = null, ?string $icon = null, ?string $tooltip = null, ?string $url = null, ?array $tags = null, ?bool $asDropdown = null, ?bool $targetBlank = null, ?bool $includeVars = null, ?bool $keepTime = null)
    {
        $this->title = $title ?: "";
        $this->type = $type ?: \Grafana\Foundation\Dashboard\DashboardLinkType::Link();
        $this->icon = $icon ?: "";
        $this->tooltip = $tooltip ?: "";
        $this->url = $url ?: "";
        $this->tags = $tags ?: [];
        $this->asDropdown = $asDropdown ?: false;
        $this->targetBlank = $targetBlank ?: false;
        $this->includeVars = $includeVars ?: false;
        $this->keepTime = $keepTime ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, type?: string, icon?: string, tooltip?: string, url?: string, tags?: array<string>, asDropdown?: bool, targetBlank?: bool, includeVars?: bool, keepTime?: bool} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Dashboard\DashboardLinkType::fromValue($input); })($data["type"]) : null,
            icon: $data["icon"] ?? null,
            tooltip: $data["tooltip"] ?? null,
            url: $data["url"] ?? null,
            tags: $data["tags"] ?? null,
            asDropdown: $data["asDropdown"] ?? null,
            targetBlank: $data["targetBlank"] ?? null,
            includeVars: $data["includeVars"] ?? null,
            keepTime: $data["keepTime"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "title" => $this->title,
            "type" => $this->type,
            "icon" => $this->icon,
            "tooltip" => $this->tooltip,
            "url" => $this->url,
            "tags" => $this->tags,
            "asDropdown" => $this->asDropdown,
            "targetBlank" => $this->targetBlank,
            "includeVars" => $this->includeVars,
            "keepTime" => $this->keepTime,
        ];
        return $data;
    }
}
