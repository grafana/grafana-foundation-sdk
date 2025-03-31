<?php

namespace Grafana\Foundation\Dashboard;

/**
 * +k8s:deepcopy-gen=true
 */
class AnnotationActions implements \JsonSerializable
{
    public ?bool $canAdd;

    public ?bool $canDelete;

    public ?bool $canEdit;

    /**
     * @param bool|null $canAdd
     * @param bool|null $canDelete
     * @param bool|null $canEdit
     */
    public function __construct(?bool $canAdd = null, ?bool $canDelete = null, ?bool $canEdit = null)
    {
        $this->canAdd = $canAdd;
        $this->canDelete = $canDelete;
        $this->canEdit = $canEdit;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{canAdd?: bool, canDelete?: bool, canEdit?: bool} $inputData */
        $data = $inputData;
        return new self(
            canAdd: $data["canAdd"] ?? null,
            canDelete: $data["canDelete"] ?? null,
            canEdit: $data["canEdit"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->canAdd)) {
            $data["canAdd"] = $this->canAdd;
        }
        if (isset($this->canDelete)) {
            $data["canDelete"] = $this->canDelete;
        }
        if (isset($this->canEdit)) {
            $data["canEdit"] = $this->canEdit;
        }
        return $data;
    }
}
