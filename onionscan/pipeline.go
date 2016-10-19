package onionscan

import (
        "github.com/s-rah/onionscan/report"
)

type Pipeline struct {
	Steps []PipelineStep
}

type PipelineStep interface {
	Do(*report.OnionScanReport)
}

func (p *Pipeline) AddStep(step PipelineStep) {
        p.Steps = append(p.Steps, step)
}

func (p *Pipeline) Execute(r *report.OnionScanReport) {
        for _,step := range p.Steps {
                step.Do(r)
        }
}
