<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AnnotationQuerySpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpecBuilder())',
        ];
            
    
        {
    $buffer = 'query(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertDataQueryKindToCode($input->query, $input->query->group, );
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'enable(';
        $arg0 =\var_export($input->enable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->iconColor !== "") {
    
        
    $buffer = 'iconColor(';
        $arg0 =\var_export($input->iconColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->builtIn !== null && $input->builtIn !== false) {
    
        
    $buffer = 'builtIn(';
        $arg0 =\var_export($input->builtIn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filter !== null) {
    
        
    $buffer = 'filter(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilterConverter::convert($input->filter);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->placement !== null) {
    
        
    $buffer = 'placement(';
        $arg0 =\var_export($input->placement, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->legacyOptions !== null) {
    
        
    $buffer = 'legacyOptions(';
        $arg0 = "[";
        foreach ($input->legacyOptions as $key => $arg1) {
            $tmplegacyOptionsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmplegacyOptionsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

