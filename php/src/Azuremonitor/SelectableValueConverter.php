<?php

namespace Grafana\Foundation\Azuremonitor;

final class SelectableValueConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\SelectableValue $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\SelectableValueBuilder())',
        ];
            if ($input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

