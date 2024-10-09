<?php

namespace Grafana\Foundation\Geomap;

final class ControlsOptionsConverter
{
    public static function convert(\Grafana\Foundation\Geomap\ControlsOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Geomap\ControlsOptionsBuilder())',
        ];
            if ($input->showZoom !== null) {
    
        
    $buffer = 'showZoom(';
        $arg0 =\var_export($input->showZoom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->mouseWheelZoom !== null) {
    
        
    $buffer = 'mouseWheelZoom(';
        $arg0 =\var_export($input->mouseWheelZoom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showAttribution !== null) {
    
        
    $buffer = 'showAttribution(';
        $arg0 =\var_export($input->showAttribution, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showScale !== null) {
    
        
    $buffer = 'showScale(';
        $arg0 =\var_export($input->showScale, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showDebug !== null) {
    
        
    $buffer = 'showDebug(';
        $arg0 =\var_export($input->showDebug, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showMeasure !== null) {
    
        
    $buffer = 'showMeasure(';
        $arg0 =\var_export($input->showMeasure, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

