<?php

namespace Grafana\Foundation\Folderv1beta1;

class Folder implements \JsonSerializable
{
    public string $title;

    public ?string $description;

    /**
     * @param string|null $title
     * @param string|null $description
     */
    public function __construct(?string $title = null, ?string $description = null)
    {
        $this->title = $title ?: "";
        $this->description = $description;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, description?: string} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->title = $this->title;
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        return $data;
    }
}
