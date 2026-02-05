<?php

namespace Grafana\Foundation\Common;

final class MapLayerOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\MapLayerOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\MapLayerOptionsBuilder())',
        ];
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->config !== null) {
    
        
    $buffer = 'config(';
        $arg0 =\var_export($input->config, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->location !== null) {
    
        
    $buffer = 'location(';
        $arg0 = \Grafana\Foundation\Common\FrameGeometrySourceConverter::convert($input->location);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filterData !== null) {
    
        
    $buffer = 'filterData(';
        $arg0 =\var_export($input->filterData, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->opacity !== null) {
    
        
    $buffer = 'opacity(';
        $arg0 =\var_export($input->opacity, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->tooltip !== null) {
    
        
    $buffer = 'tooltip(';
        $arg0 =\var_export($input->tooltip, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

