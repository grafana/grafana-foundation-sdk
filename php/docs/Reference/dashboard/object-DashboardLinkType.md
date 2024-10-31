---
title: <span class="badge object-type-enum"></span> DashboardLinkType
---
# <span class="badge object-type-enum"></span> DashboardLinkType

Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

## Definition

```php
final class DashboardLinkType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DashboardLinkType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function link(): self
    {
        if (!isset(self::$instances["link"])) {
            self::$instances["link"] = new self("link");
        }

        return self::$instances["link"];
    }

    public static function dashboards(): self
    {
        if (!isset(self::$instances["dashboards"])) {
            self::$instances["dashboards"] = new self("dashboards");
        }

        return self::$instances["dashboards"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "link") {
            return self::link();
        }

        if ($value === "dashboards") {
            return self::dashboards();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DashboardLinkType");
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

```
