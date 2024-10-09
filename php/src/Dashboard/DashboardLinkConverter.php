<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardLinkConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardLink $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardLinkBuilder('.\var_export($input->title, true).'))',
        ];
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Dashboard\DashboardLinkType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->icon !== "") {
    
        
    $buffer = 'icon(';
        $arg0 =\var_export($input->icon, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->tooltip !== "") {
    
        
    $buffer = 'tooltip(';
        $arg0 =\var_export($input->tooltip, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->url !== null && $input->url !== "") {
    
        
    $buffer = 'url(';
        $arg0 =\var_export($input->url, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->tags) >= 1) {
    
        
    $buffer = 'tags(';
        $tmparg0 = [];
        foreach ($input->tags as $arg1) {
        $tmptagsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptagsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->asDropdown !== false) {
    
        
    $buffer = 'asDropdown(';
        $arg0 =\var_export($input->asDropdown, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->targetBlank !== false) {
    
        
    $buffer = 'targetBlank(';
        $arg0 =\var_export($input->targetBlank, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->includeVars !== false) {
    
        
    $buffer = 'includeVars(';
        $arg0 =\var_export($input->includeVars, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->keepTime !== false) {
    
        
    $buffer = 'keepTime(';
        $arg0 =\var_export($input->keepTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

