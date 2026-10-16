<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

class Preferences implements \JsonSerializable
{
    /**
     * Explicit home URL (NOTE: this can only be modified in the system settings)
     */
    public ?string $homeURL;

    /**
     * UID for the home dashboard
     */
    public ?string $homeDashboardUID;

    /**
     * The timezone selection
     */
    public ?string $timezone;

    /**
     * day of the week (sunday, monday, etc)
     */
    public ?string $weekStart;

    /**
     * user interface theme
     */
    public ?string $theme;

    /**
     * Selected language
     */
    public ?string $language;

    /**
     * Explore query history preferences
     */
    public ?\Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference $queryHistory;

    /**
     * Navigation preferences
     */
    public ?\Grafana\Foundation\Preferencesv1alpha1\NavbarPreference $navbar;

    /**
     * @param string|null $homeURL
     * @param string|null $homeDashboardUID
     * @param string|null $timezone
     * @param string|null $weekStart
     * @param string|null $theme
     * @param string|null $language
     * @param \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference|null $queryHistory
     * @param \Grafana\Foundation\Preferencesv1alpha1\NavbarPreference|null $navbar
     */
    public function __construct(?string $homeURL = null, ?string $homeDashboardUID = null, ?string $timezone = null, ?string $weekStart = null, ?string $theme = null, ?string $language = null, ?\Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference $queryHistory = null, ?\Grafana\Foundation\Preferencesv1alpha1\NavbarPreference $navbar = null)
    {
        $this->homeURL = $homeURL;
        $this->homeDashboardUID = $homeDashboardUID;
        $this->timezone = $timezone;
        $this->weekStart = $weekStart;
        $this->theme = $theme;
        $this->language = $language;
        $this->queryHistory = $queryHistory;
        $this->navbar = $navbar;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{homeURL?: string, homeDashboardUID?: string, timezone?: string, weekStart?: string, theme?: string, language?: string, queryHistory?: mixed, navbar?: mixed} $inputData */
        $data = $inputData;
        return new self(
            homeURL: $data["homeURL"] ?? null,
            homeDashboardUID: $data["homeDashboardUID"] ?? null,
            timezone: $data["timezone"] ?? null,
            weekStart: $data["weekStart"] ?? null,
            theme: $data["theme"] ?? null,
            language: $data["language"] ?? null,
            queryHistory: isset($data["queryHistory"]) ? (function($input) {
    	/** @var array{homeTab?: string} */
    $val = $input;
    	return \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference::fromArray($val);
    })($data["queryHistory"]) : null,
            navbar: isset($data["navbar"]) ? (function($input) {
    	/** @var array{bookmarkUrls?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Preferencesv1alpha1\NavbarPreference::fromArray($val);
    })($data["navbar"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->homeURL)) {
            $data->homeURL = $this->homeURL;
        }
        if (isset($this->homeDashboardUID)) {
            $data->homeDashboardUID = $this->homeDashboardUID;
        }
        if (isset($this->timezone)) {
            $data->timezone = $this->timezone;
        }
        if (isset($this->weekStart)) {
            $data->weekStart = $this->weekStart;
        }
        if (isset($this->theme)) {
            $data->theme = $this->theme;
        }
        if (isset($this->language)) {
            $data->language = $this->language;
        }
        if (isset($this->queryHistory)) {
            $data->queryHistory = $this->queryHistory;
        }
        if (isset($this->navbar)) {
            $data->navbar = $this->navbar;
        }
        return $data;
    }
}
