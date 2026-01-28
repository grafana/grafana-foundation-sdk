<?php

namespace Grafana\Foundation\Annotationslist;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Annotationslist\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Annotationslist\OptionsBuilder())',
        ];
            if ($input->onlyFromThisDashboard !== false) {
    
        
    $buffer = 'onlyFromThisDashboard(';
        $arg0 =\var_export($input->onlyFromThisDashboard, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onlyInTimeRange !== false) {
    
        
    $buffer = 'onlyInTimeRange(';
        $arg0 =\var_export($input->onlyInTimeRange, true);
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
            if ($input->limit !== 10) {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showUser !== true) {
    
        
    $buffer = 'showUser(';
        $arg0 =\var_export($input->showUser, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showTime !== true) {
    
        
    $buffer = 'showTime(';
        $arg0 =\var_export($input->showTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showTags !== true) {
    
        
    $buffer = 'showTags(';
        $arg0 =\var_export($input->showTags, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->navigateToPanel !== true) {
    
        
    $buffer = 'navigateToPanel(';
        $arg0 =\var_export($input->navigateToPanel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->navigateBefore !== "" && $input->navigateBefore !== "10m") {
    
        
    $buffer = 'navigateBefore(';
        $arg0 =\var_export($input->navigateBefore, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->navigateAfter !== "" && $input->navigateAfter !== "10m") {
    
        
    $buffer = 'navigateAfter(';
        $arg0 =\var_export($input->navigateAfter, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

