package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx/v3"
	"linechat/services"
	"time"
)

type ReportHandler struct {
	reportService services.ReportService
}

func NewReportHandler(reportService services.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

func (r *ReportHandler) GetExportMemberToExcelReport(c *gin.Context) {

	// Create a new Excel file
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Members")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	members, err := r.reportService.ExportMemberReport()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// Add a data rows

	// Add headers
	headerRow := sheet.AddRow()
	headerRow.AddCell().SetString("Member ID")
	headerRow.AddCell().SetString("Name")
	headerRow.AddCell().SetString("Last Name")
	headerRow.AddCell().SetString("Phone")
	headerRow.AddCell().SetString("Email")
	headerRow.AddCell().SetString("Position")
	headerRow.AddCell().SetString("Organization")
	headerRow.AddCell().SetString("Course")
	headerRow.AddCell().SetString("Member Type")
	headerRow.AddCell().SetString("Extra Info")
	headerRow.AddCell().SetString("Registered Date")
	headerRow.AddCell().SetString("Line Name")

	// Set Title Report
	rowTitle := sheet.AddRow()
	rowTitle.AddCell().SetString("Report Members")

	// Add data rows
	for _, member := range members {
		row := sheet.AddRow()
		row.AddCell().SetString(member.MemberId)
		row.AddCell().SetString(member.Name)
		row.AddCell().SetString(member.LastName)
		row.AddCell().SetString(member.Phone)
		row.AddCell().SetString(member.Email)
		row.AddCell().SetString(member.Position)
		row.AddCell().SetString(member.Organization)
		row.AddCell().SetString(member.Course)
		row.AddCell().SetString(member.MemberType)
		row.AddCell().SetString(member.ExtraInfo)
		row.AddCell().SetString(member.RegisteredDate)
		row.AddCell().SetString(member.LineName)
	}

	// Set header for file download
	fileNamesWithCurrentTime := fmt.Sprintf("members-report%s.xlsx", time.Now().Format("2006-01-02 15:04:05"))
	c.Header("Content-Disposition", "attachment; filename="+fileNamesWithCurrentTime)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Transfer-Encoding", "binary")
	// write the file to the response writer
	err = file.Write(c.Writer)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

}
func (r *ReportHandler) GetExportEventToExcelReport(c *gin.Context) {

	result, err := r.reportService.ExportEventReport()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// crate a new Excel file
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Events")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// Add  a data rows
	// add a header
	headerRow := sheet.AddRow()
	headerRow.AddCell().SetString("Event ID")
	headerRow.AddCell().SetString("Title")
	headerRow.AddCell().SetString("Description")
	headerRow.AddCell().SetString("Start Date")
	headerRow.AddCell().SetString("Start Time")
	headerRow.AddCell().SetString("End Date")
	headerRow.AddCell().SetString("End Time")
	headerRow.AddCell().SetString("Location")
	headerRow.AddCell().SetString("Member ID")
	headerRow.AddCell().SetString("Name")
	headerRow.AddCell().SetString("Last Name")
	headerRow.AddCell().SetString("Phone")
	headerRow.AddCell().SetString("Email")
	headerRow.AddCell().SetString("Position")
	headerRow.AddCell().SetString("Organization")
	headerRow.AddCell().SetString("Course")
	headerRow.AddCell().SetString("Member Type")
	headerRow.AddCell().SetString("Extra Info")
	headerRow.AddCell().SetString("Registered Date")
	headerRow.AddCell().SetString("Line Name")

	// add data rows
	for _, event := range result {
		for _, member := range event.Members {
			registerDateStr := time.Unix(member.RegisterDate, 0).Format("2006-01-02 15:04:05")
			row := sheet.AddRow()
			row.AddCell().SetString(event.EventId)
			row.AddCell().SetString(event.Title)
			row.AddCell().SetString(event.Description)
			row.AddCell().SetString(event.StartDate)
			row.AddCell().SetString(event.StartTime)
			row.AddCell().SetString(event.EndDate)
			row.AddCell().SetString(event.EndTime)
			row.AddCell().SetString(event.Location)
			row.AddCell().SetString(member.LineId)
			row.AddCell().SetString(member.Name)
			row.AddCell().SetString(member.LastName)
			row.AddCell().SetString(member.Phone)
			row.AddCell().SetString(member.Email)
			row.AddCell().SetString(member.Position)
			row.AddCell().SetString(member.Organization)
			row.AddCell().SetString(member.Course)
			row.AddCell().SetString(member.Med)
			row.AddCell().SetString(member.MedExtraInfo)
			row.AddCell().SetString(registerDateStr)
			row.AddCell().SetString(member.LineName)

		}
	}

	// Set header for file download
	fileNamesWithCurrentTime := fmt.Sprintf("events-report%s.xlsx", time.Now().Format("2006-01-02 15:04:05"))
	c.Header("Content-Disposition", "attachment; filename="+fileNamesWithCurrentTime)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Transfer-Encoding", "binary")
	// Write the file to response Writer
	err = file.Write(c.Writer)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

}
