# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import publicdashboard


class PublicDashboard(cogbuilder.Builder[publicdashboard.PublicDashboard]):    
    _internal: publicdashboard.PublicDashboard

    def __init__(self):
        self._internal = publicdashboard.PublicDashboard()

    def build(self) -> publicdashboard.PublicDashboard:
        return self._internal    
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Unique public dashboard identifier
        """
            
        self._internal.uid = uid
    
        return self
    
    def dashboard_uid(self, dashboard_uid: str) -> typing.Self:    
        """
        Dashboard unique identifier referenced by this public dashboard
        """
            
        self._internal.dashboard_uid = dashboard_uid
    
        return self
    
    def access_token(self, access_token: str) -> typing.Self:    
        """
        Unique public access token
        """
            
        self._internal.access_token = access_token
    
        return self
    
    def is_enabled(self, is_enabled: bool) -> typing.Self:    
        """
        Flag that indicates if the public dashboard is enabled
        """
            
        self._internal.is_enabled = is_enabled
    
        return self
    
    def annotations_enabled(self, annotations_enabled: bool) -> typing.Self:    
        """
        Flag that indicates if annotations are enabled
        """
            
        self._internal.annotations_enabled = annotations_enabled
    
        return self
    
    def time_selection_enabled(self, time_selection_enabled: bool) -> typing.Self:    
        """
        Flag that indicates if the time range picker is enabled
        """
            
        self._internal.time_selection_enabled = time_selection_enabled
    
        return self
    