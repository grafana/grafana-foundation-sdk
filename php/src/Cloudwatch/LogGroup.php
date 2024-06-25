<?php

namespace Grafana\Foundation\Cloudwatch;

class LogGroup implements \JsonSerializable
{
    /**
     * ARN of the log group
     */
    public string $arn;

    /**
     * Name of the log group
     */
    public string $name;

    /**
     * AccountId of the log group
     */
    public ?string $accountId;

    /**
     * Label of the log group
     */
    public ?string $accountLabel;

    /**
     * @param string|null $arn
     * @param string|null $name
     * @param string|null $accountId
     * @param string|null $accountLabel
     */
    public function __construct(?string $arn = null, ?string $name = null, ?string $accountId = null, ?string $accountLabel = null)
    {
        $this->arn = $arn ?: "";
        $this->name = $name ?: "";
        $this->accountId = $accountId;
        $this->accountLabel = $accountLabel;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{arn?: string, name?: string, accountId?: string, accountLabel?: string} $inputData */
        $data = $inputData;
        return new self(
            arn: $data["arn"] ?? null,
            name: $data["name"] ?? null,
            accountId: $data["accountId"] ?? null,
            accountLabel: $data["accountLabel"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "arn" => $this->arn,
            "name" => $this->name,
        ];
        if (isset($this->accountId)) {
            $data["accountId"] = $this->accountId;
        }
        if (isset($this->accountLabel)) {
            $data["accountLabel"] = $this->accountLabel;
        }
        return $data;
    }
}
