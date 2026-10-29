<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class NotificationSettingsType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, NotificationSettingsType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function simplifiedRouting(): self
    {
        if (!isset(self::$instances["SimplifiedRouting"])) {
            self::$instances["SimplifiedRouting"] = new self("SimplifiedRouting");
        }

        return self::$instances["SimplifiedRouting"];
    }

    public static function namedRoutingTree(): self
    {
        if (!isset(self::$instances["NamedRoutingTree"])) {
            self::$instances["NamedRoutingTree"] = new self("NamedRoutingTree");
        }

        return self::$instances["NamedRoutingTree"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "SimplifiedRouting") {
            return self::simplifiedRouting();
        }

        if ($value === "NamedRoutingTree") {
            return self::namedRoutingTree();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum NotificationSettingsType");
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

