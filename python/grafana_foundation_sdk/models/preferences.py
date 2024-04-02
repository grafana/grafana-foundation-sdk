# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Preferences:
    # UID for the home dashboard
    home_dashboard_uid: typing.Optional[str]
    # The timezone selection
    # TODO: this should use the timezone defined in common
    timezone: typing.Optional[str]
    # day of the week (sunday, monday, etc)
    week_start: typing.Optional[str]
    # light, dark, empty is default
    theme: typing.Optional[str]
    # Selected language (beta)
    language: typing.Optional[str]
    # Explore query history preferences
    query_history: typing.Optional['QueryHistoryPreference']
    # Cookie preferences
    cookie_preferences: typing.Optional['CookiePreferences']

    def __init__(self, home_dashboard_uid: typing.Optional[str] = None, timezone: typing.Optional[str] = None, week_start: typing.Optional[str] = None, theme: typing.Optional[str] = None, language: typing.Optional[str] = None, query_history: typing.Optional['QueryHistoryPreference'] = None, cookie_preferences: typing.Optional['CookiePreferences'] = None):
        self.home_dashboard_uid = home_dashboard_uid
        self.timezone = timezone
        self.week_start = week_start
        self.theme = theme
        self.language = language
        self.query_history = query_history
        self.cookie_preferences = cookie_preferences

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.home_dashboard_uid is not None:
            payload["homeDashboardUID"] = self.home_dashboard_uid
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        if self.week_start is not None:
            payload["weekStart"] = self.week_start
        if self.theme is not None:
            payload["theme"] = self.theme
        if self.language is not None:
            payload["language"] = self.language
        if self.query_history is not None:
            payload["queryHistory"] = self.query_history
        if self.cookie_preferences is not None:
            payload["cookiePreferences"] = self.cookie_preferences
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "homeDashboardUID" in data:
            args["home_dashboard_uid"] = data["homeDashboardUID"]
        if "timezone" in data:
            args["timezone"] = data["timezone"]
        if "weekStart" in data:
            args["week_start"] = data["weekStart"]
        if "theme" in data:
            args["theme"] = data["theme"]
        if "language" in data:
            args["language"] = data["language"]
        if "queryHistory" in data:
            args["query_history"] = QueryHistoryPreference.from_json(data["queryHistory"])
        if "cookiePreferences" in data:
            args["cookie_preferences"] = CookiePreferences.from_json(data["cookiePreferences"])        

        return cls(**args)


class QueryHistoryPreference:
    # one of: '' | 'query' | 'starred';
    home_tab: typing.Optional[str]

    def __init__(self, home_tab: typing.Optional[str] = None):
        self.home_tab = home_tab

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.home_tab is not None:
            payload["homeTab"] = self.home_tab
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "homeTab" in data:
            args["home_tab"] = data["homeTab"]        

        return cls(**args)


class CookiePreferences:
    analytics: typing.Optional[object]
    performance: typing.Optional[object]
    functional: typing.Optional[object]

    def __init__(self, analytics: typing.Optional[object] = None, performance: typing.Optional[object] = None, functional: typing.Optional[object] = None):
        self.analytics = analytics
        self.performance = performance
        self.functional = functional

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.analytics is not None:
            payload["analytics"] = self.analytics
        if self.performance is not None:
            payload["performance"] = self.performance
        if self.functional is not None:
            payload["functional"] = self.functional
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "analytics" in data:
            args["analytics"] = data["analytics"]
        if "performance" in data:
            args["performance"] = data["performance"]
        if "functional" in data:
            args["functional"] = data["functional"]        

        return cls(**args)



