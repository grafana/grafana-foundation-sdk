<?php

namespace Grafana\Foundation\News;

class Options implements \JsonSerializable
{
    /**
     * empty/missing will default to grafana blog
     */
    public ?string $feedUrl;

    public ?bool $showImage;

    /**
     * @param string|null $feedUrl
     * @param bool|null $showImage
     */
    public function __construct(?string $feedUrl = null, ?bool $showImage = null)
    {
        $this->feedUrl = $feedUrl;
        $this->showImage = $showImage;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{feedUrl?: string, showImage?: bool} $inputData */
        $data = $inputData;
        return new self(
            feedUrl: $data["feedUrl"] ?? null,
            showImage: $data["showImage"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->feedUrl)) {
            $data["feedUrl"] = $this->feedUrl;
        }
        if (isset($this->showImage)) {
            $data["showImage"] = $this->showImage;
        }
        return $data;
    }
}
