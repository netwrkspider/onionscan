package steps

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/s-rah/onionscan/report"
)

type SimpleReportWriter struct {
        reportFile string
        asJSON  bool
        width int
}

func (srw *SimpleReportWriter) Init(outputFile string, asJSON bool, width int) {
        srw.reportFile = outputFile
        srw.asJSON = asJSON
        srw.width = width
}

func (srw *SimpleReportWriter) Do(r *report.OnionScanReport) {
	var report_str string
	var err error
	if srw.asJSON {
		//report_str, err = r.SimpleReport.Serialize()
	} else {
		//report_str, err = r.SimpleReport.Format(srw.width)
	}
	if err != nil {
		log.Printf("Could not generate report")
		return
	}
	
	reportFile := r.HiddenService + "." + srw.reportFile

	if len(reportFile) > 0 {
		f, err := os.Create(reportFile)

		for err != nil {
			log.Printf("Cannot create report file: %s...trying again in 5 seconds...", err)
			time.Sleep(time.Second * 5)
			f, err = os.Create(reportFile)
		}

		defer f.Close()

		f.WriteString(report_str)
	} else {
		fmt.Print(report_str)
	}
}
