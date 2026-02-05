<?php

namespace Grafana\Foundation\Common;

final class FillConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\FillConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\FillConfigBuilder())',
        ];
            if ($input->fillColor !== null && $input->fillColor !== "") {
    
        
    $buffer = 'fillColor(';
        $arg0 =\var_export($input->fillColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillOpacity !== null) {
    
        
    $buffer = 'fillOpacity(';
        $arg0 =\var_export($input->fillOpacity, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillBelowTo !== null && $input->fillBelowTo !== "") {
    
        
    $buffer = 'fillBelowTo(';
        $arg0 =\var_export($input->fillBelowTo, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

