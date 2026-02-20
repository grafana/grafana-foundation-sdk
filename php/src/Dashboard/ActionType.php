<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Dashboard action type
 */
final class ActionType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ActionType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function fetch(): self
    {
        if (!isset(self::$instances["fetch"])) {
            self::$instances["fetch"] = new self("fetch");
        }

        return self::$instances["fetch"];
    }

    public static function infinity(): self
    {
        if (!isset(self::$instances["infinity"])) {
            self::$instances["infinity"] = new self("infinity");
        }

        return self::$instances["infinity"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "fetch") {
            return self::fetch();
        }

        if ($value === "infinity") {
            return self::infinity();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ActionType");
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

