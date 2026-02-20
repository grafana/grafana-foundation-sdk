<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Action variable type
 */
final class ActionVariableType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ActionVariableType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function string(): self
    {
        if (!isset(self::$instances["string"])) {
            self::$instances["string"] = new self("string");
        }

        return self::$instances["string"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "string") {
            return self::string();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ActionVariableType");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

