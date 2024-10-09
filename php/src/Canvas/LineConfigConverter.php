<?php

namespace Grafana\Foundation\Canvas;

final class LineConfigConverter
{
    public static function convert(\Grafana\Foundation\Canvas\LineConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\LineConfigBuilder())',
        ];
            if ($input->color !== null) {
    
        
    $buffer = 'color(';
        $arg0 = \Grafana\Foundation\Common\ColorDimensionConfigConverter::convert($input->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->width !== null) {
    
        
    $buffer = 'width(';
        $arg0 =\var_export($input->width, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

