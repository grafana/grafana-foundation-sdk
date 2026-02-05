<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class LibraryPanelKindSpec implements \JsonSerializable
{
    /**
     * Panel ID for the library panel in the dashboard
     */
    public float $id;

    /**
     * Title for the library panel in the dashboard
     */
    public string $title;

    public \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef $libraryPanel;

    /**
     * @param float|null $id
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef|null $libraryPanel
     */
    public function __construct(?float $id = null, ?string $title = null, ?\Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef $libraryPanel = null)
    {
        $this->id = $id ?: 0;
        $this->title = $title ?: "";
        $this->libraryPanel = $libraryPanel ?: new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: float, title?: string, libraryPanel?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            title: $data["title"] ?? null,
            libraryPanel: isset($data["libraryPanel"]) ? (function($input) {
    	/** @var array{name?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef::fromArray($val);
    })($data["libraryPanel"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->id = $this->id;
        $data->title = $this->title;
        $data->libraryPanel = $this->libraryPanel;
        return $data;
    }
}
