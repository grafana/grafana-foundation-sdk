<?php

namespace Grafana\Foundation\Dashboard;

final class ValueMappingResultConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\ValueMappingResult $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ValueMappingResultBuilder())',
        ];
            if ($input->text !== null && $input->text !== "") {
    
        
    $buffer = 'text(';
        $arg0 =\var_export($input->text, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->color !== null && $input->color !== "") {
    
        
    $buffer = 'color(';
        $arg0 =\var_export($input->color, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->icon !== null && $input->icon !== "") {
    
        
    $buffer = 'icon(';
        $arg0 =\var_export($input->icon, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->index !== null) {
    
        
    $buffer = 'index(';
        $arg0 =\var_export($input->index, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

