package deanonymization

import (
	"github.com/s-rah/onionscan/config"
	"github.com/s-rah/onionscan/report"
	"strconv"
)

func CommonCorrelations(osreport *report.OnionScanReport, anonreport *report.AnonymityReport, osc *config.OnionScanConfig) {

	// SSH
	if osreport.SSHKey != "" {
		osc.Database.InsertRelationship(osreport.HiddenService, "ssh", osreport.SSHKey)
	}

	if osreport.SSHBanner != "" {
		osc.Database.InsertRelationship(osreport.HiddenService, "ssh", osreport.SSHBanner)
	}

	// FTP
	if osreport.FTPBanner != "" {
		osc.Database.InsertRelationship(osreport.HiddenService, "ftp", osreport.FTPBanner)
	}

	// SMTP
	if osreport.SMTPBanner != "" {
		osc.Database.InsertRelationship(osreport.HiddenService, "smtp", osreport.SMTPBanner)
	}
	
	// Adding all Crawl Ids to Common Correlations (this is a bit of a hack to make the webui nicer)
	for _, crawlId := range osreport.Crawls {
	        osc.Database.InsertRelationship(osreport.HiddenService, "crawl", strconv.Itoa(crawlId))
	}

}
