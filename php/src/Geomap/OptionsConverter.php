<?php

namespace Grafana\Foundation\Geomap;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Geomap\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Geomap\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'view(';
        $arg0 = \Grafana\Foundation\Geomap\MapViewConfigConverter::convert($input->view);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'controls(';
        $arg0 = \Grafana\Foundation\Geomap\ControlsOptionsConverter::convert($input->controls);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'basemap(';
        $arg0 = \Grafana\Foundation\Common\MapLayerOptionsConverter::convert($input->basemap);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->layers) >= 1) {
    
        
    $buffer = 'layers(';
        $tmparg0 = [];
        foreach ($input->layers as $arg1) {
        $tmplayersarg1 = \Grafana\Foundation\Common\MapLayerOptionsConverter::convert($arg1);
        $tmparg0[] = $tmplayersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Geomap\TooltipOptionsConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

