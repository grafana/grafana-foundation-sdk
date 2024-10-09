<?php

namespace Grafana\Foundation\Common;

final class FrameGeometrySourceConverter
{
    public static function convert(\Grafana\Foundation\Common\FrameGeometrySource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\FrameGeometrySourceBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\FrameGeometrySourceMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->geohash !== null && $input->geohash !== "") {
    
        
    $buffer = 'geohash(';
        $arg0 =\var_export($input->geohash, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->latitude !== null && $input->latitude !== "") {
    
        
    $buffer = 'latitude(';
        $arg0 =\var_export($input->latitude, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->longitude !== null && $input->longitude !== "") {
    
        
    $buffer = 'longitude(';
        $arg0 =\var_export($input->longitude, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->wkt !== null && $input->wkt !== "") {
    
        
    $buffer = 'wkt(';
        $arg0 =\var_export($input->wkt, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lookup !== null && $input->lookup !== "") {
    
        
    $buffer = 'lookup(';
        $arg0 =\var_export($input->lookup, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gazetteer !== null && $input->gazetteer !== "") {
    
        
    $buffer = 'gazetteer(';
        $arg0 =\var_export($input->gazetteer, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

