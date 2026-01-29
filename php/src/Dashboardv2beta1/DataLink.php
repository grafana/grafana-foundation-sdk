<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class DataLink implements \JsonSerializable
{
    public string $title;

    public string $url;

    public ?bool $targetBlank;

    /**
     * @param string|null $title
     * @param string|null $url
     * @param bool|null $targetBlank
     */
    public function __construct(?string $title = null, ?string $url = null, ?bool $targetBlank = null)
    {
        $this->title = $title ?: "";
        $this->url = $url ?: "";
        $this->targetBlank = $targetBlank;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, url?: string, targetBlank?: bool} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            url: $data["url"] ?? null,
            targetBlank: $data["targetBlank"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->title = $this->title;
        $data->url = $this->url;
        if (isset($this->targetBlank)) {
            $data->targetBlank = $this->targetBlank;
        }
        return $data;
    }
}
