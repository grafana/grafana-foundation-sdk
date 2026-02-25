<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AnnotationQueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AnnotationQueryBuilder())',
        ];
            
    
        {
    $buffer = 'query(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertDataQueryKindToCode($input->spec->query, $input->spec->query->group, );
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'enable(';
        $arg0 =\var_export($input->spec->enable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->iconColor !== "") {
    
        
    $buffer = 'iconColor(';
        $arg0 =\var_export($input->spec->iconColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->spec->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->builtIn !== null && $input->spec->builtIn !== false) {
    
        
    $buffer = 'builtIn(';
        $arg0 =\var_export($input->spec->builtIn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->filter !== null) {
    
        
    $buffer = 'filter(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilterConverter::convert($input->spec->filter);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->placement !== null) {
    
        
    $buffer = 'placement(';
        $arg0 =\var_export($input->spec->placement, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->legacyOptions !== null) {
    
        
    $buffer = 'legacyOptions(';
        $arg0 = "[";
        foreach ($input->spec->legacyOptions as $key => $arg1) {
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

