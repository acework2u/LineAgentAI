package router

import (
	"github.com/gin-gonic/gin"
	"linechat/handler"
)

type ReportRouter struct {
	reportHandler *handler.ReportHandler
}

func NewReportRouter(reportHandler *handler.ReportHandler) *ReportRouter {
	return &ReportRouter{reportHandler: reportHandler}
}

func (r *ReportRouter) ReportRouter(rg *gin.RouterGroup) {

	router := rg.Group("/reports")
	router.GET("", r.reportHandler.GetReports)
	router.GET("/members", r.reportHandler.GetExportMemberToExcelReport)
	router.GET("/events", r.reportHandler.GetExportEventToExcelReport)
	router.GET("/event", r.reportHandler.GetExportEventsToExcelReport)
	router.GET("/event-with-clinic", r.reportHandler.GetExportEventsByClinicToExcel)
	router.GET("/raw-event-with-clinic", r.reportHandler.GetExportEventsByClinic)

	//rg.GET("")
	//rg.GET("/reports/members/excel", r.reportHandler.GetExportMemberToExcelReport)
	//rg.GET("/reports/events/excel", r.reportHandler.GetExportEventToExcelReport)
	//rg.GET("/reports/event", r.reportHandler.GetExportEventsToExcelReport)
	//rg.GET("/reports/events-clinic", r.reportHandler.GetExportEventsByClinicToExcel)

}
