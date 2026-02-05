<?php

namespace Grafana\Foundation\Common;

class FrameGeometrySource implements \JsonSerializable
{
    public \Grafana\Foundation\Common\FrameGeometrySourceMode $mode;

    /**
     * Field mappings
     */
    public ?string $geohash;

    public ?string $latitude;

    public ?string $longitude;

    public ?string $wkt;

    public ?string $lookup;

    /**
     * Path to Gazetteer
     */
    public ?string $gazetteer;

    /**
     * @param \Grafana\Foundation\Common\FrameGeometrySourceMode|null $mode
     * @param string|null $geohash
     * @param string|null $latitude
     * @param string|null $longitude
     * @param string|null $wkt
     * @param string|null $lookup
     * @param string|null $gazetteer
     */
    public function __construct(?\Grafana\Foundation\Common\FrameGeometrySourceMode $mode = null, ?string $geohash = null, ?string $latitude = null, ?string $longitude = null, ?string $wkt = null, ?string $lookup = null, ?string $gazetteer = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\FrameGeometrySourceMode::Auto();
        $this->geohash = $geohash;
        $this->latitude = $latitude;
        $this->longitude = $longitude;
        $this->wkt = $wkt;
        $this->lookup = $lookup;
        $this->gazetteer = $gazetteer;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, geohash?: string, latitude?: string, longitude?: string, wkt?: string, lookup?: string, gazetteer?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\FrameGeometrySourceMode::fromValue($input); })($data["mode"]) : null,
            geohash: $data["geohash"] ?? null,
            latitude: $data["latitude"] ?? null,
            longitude: $data["longitude"] ?? null,
            wkt: $data["wkt"] ?? null,
            lookup: $data["lookup"] ?? null,
            gazetteer: $data["gazetteer"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        if (isset($this->geohash)) {
            $data->geohash = $this->geohash;
        }
        if (isset($this->latitude)) {
            $data->latitude = $this->latitude;
        }
        if (isset($this->longitude)) {
            $data->longitude = $this->longitude;
        }
        if (isset($this->wkt)) {
            $data->wkt = $this->wkt;
        }
        if (isset($this->lookup)) {
            $data->lookup = $this->lookup;
        }
        if (isset($this->gazetteer)) {
            $data->gazetteer = $this->gazetteer;
        }
        return $data;
    }
}
