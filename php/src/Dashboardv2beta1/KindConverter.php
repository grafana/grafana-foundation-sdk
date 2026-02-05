<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class KindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Kind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\KindBuilder())',
        ];
            if ($input->kind !== "") {
    
        
    $buffer = 'kind(';
        $arg0 =\var_export($input->kind, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null) {
    
        
    $buffer = 'spec(';
        $arg0 =\var_export($input->spec, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metadata !== null) {
    
        
    $buffer = 'metadata(';
        $arg0 =\var_export($input->metadata, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

