<?php

namespace Grafana\Foundation\Dashboardv2;

class LibraryPanelKindSpec implements \JsonSerializable
{
    /**
     * Panel ID for the library panel in the dashboard
     */
    public int $id;

    /**
     * Title for the library panel in the dashboard
     */
    public string $title;

    public \Grafana\Foundation\Dashboardv2\LibraryPanelRef $libraryPanel;

    /**
     * @param int|null $id
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboardv2\LibraryPanelRef|null $libraryPanel
     */
    public function __construct(?int $id = null, ?string $title = null, ?\Grafana\Foundation\Dashboardv2\LibraryPanelRef $libraryPanel = null)
    {
        $this->id = $id ?: 0;
        $this->title = $title ?: "";
        $this->libraryPanel = $libraryPanel ?: new \Grafana\Foundation\Dashboardv2\LibraryPanelRef();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: int, title?: string, libraryPanel?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            title: $data["title"] ?? null,
            libraryPanel: isset($data["libraryPanel"]) ? (function($input) {
    	/** @var array{name?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\LibraryPanelRef::fromArray($val);
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
