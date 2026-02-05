<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class DataLinkConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataLink $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\DataLinkBuilder())',
        ];
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->url !== "") {
    
        
    $buffer = 'url(';
        $arg0 =\var_export($input->url, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->targetBlank !== null) {
    
        
    $buffer = 'targetBlank(';
        $arg0 =\var_export($input->targetBlank, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

