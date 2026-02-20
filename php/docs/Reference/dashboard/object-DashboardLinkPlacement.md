---
title: <span class="badge object-type-enum"></span> DashboardLinkPlacement
---
# <span class="badge object-type-enum"></span> DashboardLinkPlacement

Dashboard Link placement. Defines where the link should be displayed.

- "inControlsMenu" renders the link in bottom part of the dashboard controls dropdown menu

## Definition

```php
final class DashboardLinkPlacement implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DashboardLinkPlacement>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function inControlsMenu(): self
    {
        if (!isset(self::$instances["inControlsMenu"])) {
            self::$instances["inControlsMenu"] = new self("inControlsMenu");
        }

        return self::$instances["inControlsMenu"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "inControlsMenu") {
            return self::inControlsMenu();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DashboardLinkPlacement");
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
