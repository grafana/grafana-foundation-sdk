<?php

namespace Grafana\Foundation\Geomap;

final class MapViewConfigConverter
{
    public static function convert(\Grafana\Foundation\Geomap\MapViewConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Geomap\MapViewConfigBuilder())',
        ];
            if ($input->id !== "" && $input->id !== "zero") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lat !== null && $input->lat !== 0) {
    
        
    $buffer = 'lat(';
        $arg0 =\var_export($input->lat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lon !== null && $input->lon !== 0) {
    
        
    $buffer = 'lon(';
        $arg0 =\var_export($input->lon, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->zoom !== null && $input->zoom !== 1) {
    
        
    $buffer = 'zoom(';
        $arg0 =\var_export($input->zoom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->minZoom !== null) {
    
        
    $buffer = 'minZoom(';
        $arg0 =\var_export($input->minZoom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxZoom !== null) {
    
        
    $buffer = 'maxZoom(';
        $arg0 =\var_export($input->maxZoom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->padding !== null) {
    
        
    $buffer = 'padding(';
        $arg0 =\var_export($input->padding, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->allLayers !== null && $input->allLayers !== true) {
    
        
    $buffer = 'allLayers(';
        $arg0 =\var_export($input->allLayers, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lastOnly !== null) {
    
        
    $buffer = 'lastOnly(';
        $arg0 =\var_export($input->lastOnly, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->layer !== null && $input->layer !== "") {
    
        
    $buffer = 'layer(';
        $arg0 =\var_export($input->layer, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->shared !== null) {
    
        
    $buffer = 'shared(';
        $arg0 =\var_export($input->shared, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

