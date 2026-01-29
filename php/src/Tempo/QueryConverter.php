<?php

namespace Grafana\Foundation\Tempo;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Tempo\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->query !== null && $input->spec->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->spec->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->search !== null && $input->spec->search !== "") {
    
        
    $buffer = 'search(';
        $arg0 =\var_export($input->spec->search, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->serviceName !== null && $input->spec->serviceName !== "") {
    
        
    $buffer = 'serviceName(';
        $arg0 =\var_export($input->spec->serviceName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->spanName !== null && $input->spec->spanName !== "") {
    
        
    $buffer = 'spanName(';
        $arg0 =\var_export($input->spec->spanName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->minDuration !== null && $input->spec->minDuration !== "") {
    
        
    $buffer = 'minDuration(';
        $arg0 =\var_export($input->spec->minDuration, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->maxDuration !== null && $input->spec->maxDuration !== "") {
    
        
    $buffer = 'maxDuration(';
        $arg0 =\var_export($input->spec->maxDuration, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->serviceMapQuery !== null) {
    
        
    $buffer = 'serviceMapQuery(';
        switch (true) {
            case is_string($input->spec->serviceMapQuery):
                $disjunctionserviceMapQuery =\var_export($input->spec->serviceMapQuery, true);
                $arg0 = $disjunctionserviceMapQuery;
                break;
            case is_array($input->spec->serviceMapQuery):
                $tmpdisjunctionserviceMapQuery = [];
        foreach ($input->spec->serviceMapQuery as $arg1) {
        $tmpserviceMapQueryarg1 =\var_export($arg1, true);
        $tmpdisjunctionserviceMapQuery[] = $tmpserviceMapQueryarg1;
        }
        $disjunctionserviceMapQuery = "[" . implode(", \n", $tmpdisjunctionserviceMapQuery) . "]";
                $arg0 = $disjunctionserviceMapQuery;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->serviceMapIncludeNamespace !== null) {
    
        
    $buffer = 'serviceMapIncludeNamespace(';
        $arg0 =\var_export($input->spec->serviceMapIncludeNamespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->limit !== null) {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->spec->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->spss !== null) {
    
        
    $buffer = 'spss(';
        $arg0 =\var_export($input->spec->spss, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && count($input->spec->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->spec->filters as $arg1) {
        $tmpfiltersarg1 = \Grafana\Foundation\Tempo\TraceqlFilterConverter::convert($arg1);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->groupBy !== null && count($input->spec->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->spec->groupBy as $arg1) {
        $tmpgroupByarg1 = \Grafana\Foundation\Tempo\TraceqlFilterConverter::convert($arg1);
        $tmparg0[] = $tmpgroupByarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->tableType !== null) {
    
        
    $buffer = 'tableType(';
        $arg0 ='\Grafana\Foundation\Tempo\SearchTableType::fromValue("'.$input->spec->tableType.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->step !== null && $input->spec->step !== "") {
    
        
    $buffer = 'step(';
        $arg0 =\var_export($input->spec->step, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->exemplars !== null) {
    
        
    $buffer = 'exemplars(';
        $arg0 =\var_export($input->spec->exemplars, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->datasource !== null) {
    
        
    $buffer = 'oldDatasource(';
        $arg0 ='(new \Grafana\Foundation\Common\DataSourceRef('.(($input->spec->datasource->type !== null) ? 'type: '.\var_export($input->spec->datasource->type, true).', ' : '').''.(($input->spec->datasource->uid !== null) ? 'uid: '.\var_export($input->spec->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Tempo\TempoQuery && $input->spec->metricsQueryType !== null) {
    
        
    $buffer = 'metricsQueryType(';
        $arg0 ='\Grafana\Foundation\Tempo\MetricsQueryType::fromValue("'.$input->spec->metricsQueryType.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

