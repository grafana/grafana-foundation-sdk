<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TargetConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TargetBuilder())',
        ];
            
    
        {
    $buffer = 'query(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertDataQueryKindToCode($input->spec->query, $input->spec->query->group, );
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->refId !== "" && $input->spec->refId !== "A") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'hidden(';
        $arg0 =\var_export($input->spec->hidden, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

