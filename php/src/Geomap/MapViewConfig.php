<?php

namespace Grafana\Foundation\Geomap;

class MapViewConfig implements \JsonSerializable
{
    public string $id;

    public ?int $lat;

    public ?int $lon;

    public ?int $zoom;

    public ?int $minZoom;

    public ?int $maxZoom;

    public ?int $padding;

    public ?bool $allLayers;

    public ?bool $lastOnly;

    public ?string $layer;

    public ?bool $shared;

    /**
     * @param string|null $id
     * @param int|null $lat
     * @param int|null $lon
     * @param int|null $zoom
     * @param int|null $minZoom
     * @param int|null $maxZoom
     * @param int|null $padding
     * @param bool|null $allLayers
     * @param bool|null $lastOnly
     * @param string|null $layer
     * @param bool|null $shared
     */
    public function __construct(?string $id = null, ?int $lat = null, ?int $lon = null, ?int $zoom = null, ?int $minZoom = null, ?int $maxZoom = null, ?int $padding = null, ?bool $allLayers = null, ?bool $lastOnly = null, ?string $layer = null, ?bool $shared = null)
    {
        $this->id = $id ?: "zero";
        $this->lat = $lat;
        $this->lon = $lon;
        $this->zoom = $zoom;
        $this->minZoom = $minZoom;
        $this->maxZoom = $maxZoom;
        $this->padding = $padding;
        $this->allLayers = $allLayers;
        $this->lastOnly = $lastOnly;
        $this->layer = $layer;
        $this->shared = $shared;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, lat?: int, lon?: int, zoom?: int, minZoom?: int, maxZoom?: int, padding?: int, allLayers?: bool, lastOnly?: bool, layer?: string, shared?: bool} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            lat: $data["lat"] ?? null,
            lon: $data["lon"] ?? null,
            zoom: $data["zoom"] ?? null,
            minZoom: $data["minZoom"] ?? null,
            maxZoom: $data["maxZoom"] ?? null,
            padding: $data["padding"] ?? null,
            allLayers: $data["allLayers"] ?? null,
            lastOnly: $data["lastOnly"] ?? null,
            layer: $data["layer"] ?? null,
            shared: $data["shared"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
        ];
        if (isset($this->lat)) {
            $data["lat"] = $this->lat;
        }
        if (isset($this->lon)) {
            $data["lon"] = $this->lon;
        }
        if (isset($this->zoom)) {
            $data["zoom"] = $this->zoom;
        }
        if (isset($this->minZoom)) {
            $data["minZoom"] = $this->minZoom;
        }
        if (isset($this->maxZoom)) {
            $data["maxZoom"] = $this->maxZoom;
        }
        if (isset($this->padding)) {
            $data["padding"] = $this->padding;
        }
        if (isset($this->allLayers)) {
            $data["allLayers"] = $this->allLayers;
        }
        if (isset($this->lastOnly)) {
            $data["lastOnly"] = $this->lastOnly;
        }
        if (isset($this->layer)) {
            $data["layer"] = $this->layer;
        }
        if (isset($this->shared)) {
            $data["shared"] = $this->shared;
        }
        return $data;
    }
}
