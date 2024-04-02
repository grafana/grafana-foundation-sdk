# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import preferences


class Preferences(cogbuilder.Builder[preferences.Preferences]):    
    """
    Spec defines user, team or org Grafana preferences
    swagger:model Preferences
    """
    
    __internal: preferences.Preferences

    def __init__(self):
        self.__internal = preferences.Preferences()

    def build(self) -> preferences.Preferences:
        return self.__internal    
    
    def home_dashboard_uid(self, home_dashboard_uid: str) -> typing.Self:    
        """
        UID for the home dashboard
        """
            
        self.__internal.home_dashboard_uid = home_dashboard_uid
    
        return self
    
    def timezone(self, timezone: str) -> typing.Self:    
        """
        The timezone selection
        TODO: this should use the timezone defined in common
        """
            
        self.__internal.timezone = timezone
    
        return self
    
    def week_start(self, week_start: str) -> typing.Self:    
        """
        day of the week (sunday, monday, etc)
        """
            
        self.__internal.week_start = week_start
    
        return self
    
    def theme(self, theme: str) -> typing.Self:    
        """
        light, dark, empty is default
        """
            
        self.__internal.theme = theme
    
        return self
    
    def language(self, language: str) -> typing.Self:    
        """
        Selected language (beta)
        """
            
        self.__internal.language = language
    
        return self
    
    def query_history(self, query_history: cogbuilder.Builder[preferences.QueryHistoryPreference]) -> typing.Self:    
        """
        Explore query history preferences
        """
            
        query_history_resource = query_history.build()
        self.__internal.query_history = query_history_resource
    
        return self
    
    def cookie_preferences(self, cookie_preferences: cogbuilder.Builder[preferences.CookiePreferences]) -> typing.Self:    
        """
        Cookie preferences
        """
            
        cookie_preferences_resource = cookie_preferences.build()
        self.__internal.cookie_preferences = cookie_preferences_resource
    
        return self
    

class QueryHistoryPreference(cogbuilder.Builder[preferences.QueryHistoryPreference]):    
    __internal: preferences.QueryHistoryPreference

    def __init__(self):
        self.__internal = preferences.QueryHistoryPreference()

    def build(self) -> preferences.QueryHistoryPreference:
        return self.__internal    
    
    def home_tab(self, home_tab: str) -> typing.Self:    
        """
        one of: '' | 'query' | 'starred';
        """
            
        self.__internal.home_tab = home_tab
    
        return self
    

class CookiePreferences(cogbuilder.Builder[preferences.CookiePreferences]):    
    __internal: preferences.CookiePreferences

    def __init__(self):
        self.__internal = preferences.CookiePreferences()

    def build(self) -> preferences.CookiePreferences:
        return self.__internal    
    
    def analytics(self, analytics: object) -> typing.Self:        
        self.__internal.analytics = analytics
    
        return self
    
    def performance(self, performance: object) -> typing.Self:        
        self.__internal.performance = performance
    
        return self
    
    def functional(self, functional: object) -> typing.Self:        
        self.__internal.functional = functional
    
        return self
    