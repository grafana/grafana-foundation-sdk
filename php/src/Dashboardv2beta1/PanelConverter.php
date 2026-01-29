<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class PanelConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\PanelKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\PanelBuilder())',
        ];
            
    
        {
    $buffer = 'id(';
        $arg0 =\var_export($input->spec->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->spec->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->spec->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->links) >= 1) {
    
        
    $buffer = 'links(';
        $tmparg0 = [];
        foreach ($input->spec->links as $arg1) {
        $tmplinksarg1 = \Grafana\Foundation\Dashboardv2beta1\DataLinkConverter::convert($arg1);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'data(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\QueryGroupConverter::convert($input->spec->data);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'visualization(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertPanelToCode($input->spec->vizConfig, $input->spec->vizConfig->group, );
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->transparent !== null) {
    
        
    $buffer = 'transparent(';
        $arg0 =\var_export($input->spec->transparent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

