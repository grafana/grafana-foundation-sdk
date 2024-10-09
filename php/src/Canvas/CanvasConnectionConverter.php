<?php

namespace Grafana\Foundation\Canvas;

final class CanvasConnectionConverter
{
    public static function convert(\Grafana\Foundation\Canvas\CanvasConnection $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\CanvasConnectionBuilder())',
        ];
            
    
        {
    $buffer = 'source(';
        $arg0 = \Grafana\Foundation\Canvas\ConnectionCoordinatesConverter::convert($input->source);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'target(';
        $arg0 = \Grafana\Foundation\Canvas\ConnectionCoordinatesConverter::convert($input->target);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->targetName !== null && $input->targetName !== "") {
    
        
    $buffer = 'targetName(';
        $arg0 =\var_export($input->targetName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'path(';
        $arg0 ='\Grafana\Foundation\Canvas\ConnectionPath::fromValue("'.$input->path.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->color !== null) {
    
        
    $buffer = 'color(';
        $arg0 = \Grafana\Foundation\Common\ColorDimensionConfigConverter::convert($input->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->size !== null) {
    
        
    $buffer = 'size(';
        $arg0 = \Grafana\Foundation\Common\ScaleDimensionConfigConverter::convert($input->size);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->vertices !== null && count($input->vertices) >= 1) {
    
        
    $buffer = 'vertices(';
        $tmparg0 = [];
        foreach ($input->vertices as $arg1) {
        $tmpverticesarg1 = \Grafana\Foundation\Canvas\ConnectionCoordinatesConverter::convert($arg1);
        $tmparg0[] = $tmpverticesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sourceOriginal !== null) {
    
        
    $buffer = 'sourceOriginal(';
        $arg0 = \Grafana\Foundation\Canvas\ConnectionCoordinatesConverter::convert($input->sourceOriginal);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->targetOriginal !== null) {
    
        
    $buffer = 'targetOriginal(';
        $arg0 = \Grafana\Foundation\Canvas\ConnectionCoordinatesConverter::convert($input->targetOriginal);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

