<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureResourceGraphQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureResourceGraphQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureResourceGraphQueryBuilder())',
        ];
            if ($input->query !== null && $input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resultFormat !== null && $input->resultFormat !== "") {
    
        
    $buffer = 'resultFormat(';
        $arg0 =\var_export($input->resultFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

