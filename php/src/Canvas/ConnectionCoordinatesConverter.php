<?php

namespace Grafana\Foundation\Canvas;

final class ConnectionCoordinatesConverter
{
    public static function convert(\Grafana\Foundation\Canvas\ConnectionCoordinates $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\ConnectionCoordinatesBuilder())',
        ];
            
    
        {
    $buffer = 'x(';
        $arg0 =\var_export($input->x, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'y(';
        $arg0 =\var_export($input->y, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

