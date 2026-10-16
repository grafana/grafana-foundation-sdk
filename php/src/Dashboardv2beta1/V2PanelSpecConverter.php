<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class V2PanelSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\V2PanelSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\V2PanelSpecBuilder())',
        ];
            
    
        {
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->subtitle !== null && $input->subtitle !== "") {
    
        
    $buffer = 'subtitle(';
        $arg0 =\var_export($input->subtitle, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->links) >= 1) {
    
        
    $buffer = 'links(';
        $tmparg0 = [];
        foreach ($input->links as $arg1) {
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
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKindConverter::convert($input->data);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'vizConfig(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertPanelToCode($input->vizConfig, $input->vizConfig->group, );
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->transparent !== null) {
    
        
    $buffer = 'transparent(';
        $arg0 =\var_export($input->transparent, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

