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
func (r *ReportHandler) GetExportEventsToExcelReport(c *gin.Context) {

	file := xlsx.NewFile()

	eh, err := file.AddSheet("Events")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}
	// Event report
	//headerRow := sh.AddRow()
	//headerRow.AddCell().SetString("")
	//headerRow.AddCell().SetString("Title")
	//headerRow.AddCell().SetString("Description")
	//headerRow.AddCell().SetString("Start Date")
	//headerRow.AddCell().SetString("Start Time")
	//headerRow.AddCell().SetString("End Date")
	//headerRow.AddCell().SetString("End Time")
	eh.AddRow().AddCell().SetString("โครงการ")
	eh.AddRow().AddCell().SetString("รายละเอียด")
	eh.AddRow().AddCell().SetString("วันที่")
	eh.AddRow().AddCell().SetString("สถานที่")
	eh.AddRow().AddCell().SetString("")
	eh.AddRow().AddCell().SetString("ประธานโครงการ")
	eh.AddRow().AddCell().SetString("ประธานดำเนินการโครงการแพทย์")
	eh.AddRow().AddCell().SetString("ประธานกิจกรรมหน่วยแพทย์อาสา")
	eh.AddRow().AddCell().SetString("ฝ่ายบริการทางการแพทย์ (คลินิกบริการทางการแพทย์)")

	// Set custom detail
	theCell, _ := eh.Cell(0, 1)
	theCell.Value = "โครงการแพทย์อาสาตรวจสุขภาพพระสงฆ์ "
	theCell, _ = eh.Cell(1, 1)
	theCell.Value = "โครงการแพทย์อาสาตรวจสุขภาพพระสงฆ์ถวายเป็นพระกศล แด่สมเด็จพระสงฆราช สกลมหาสังฆปริณายก"
	theCell, _ = eh.Cell(2, 1)
	theCell.Value = "ในวันอาทิตย์ที่ 9 กุมภาพันธ์ 2568"
	theCell, _ = eh.Cell(3, 1)
	theCell.Value = "ณ วัดราชบพิธสถิตมหาสีมารามราชวรวิหาร"

	// Set with col
	newCol := xlsx.NewColForRange(1, 4)
	newCol.SetWidth(35)
	colStyle := xlsx.NewStyle()
	colStyle.Alignment.Horizontal = "left"
	colStyle.Font.Name = "TH Sarabun New"
	colStyle.Font.Size = 16
	colStyle.Font.Bold = true
	colStyle.ApplyAlignment = true
	colStyle.ApplyFont = true
	newCol.SetStyle(colStyle)
	eh.SetColParameters(newCol)

	// Add new tab sheet
	sh, err := file.AddSheet("Members")

	myStyle := xlsx.NewStyle()
	myStyle.Alignment.Horizontal = "right"
	myStyle.Fill.FgColor = "FFFFFF00"
	myStyle.Fill.PatternType = "solid"
	myStyle.Font.Name = "Georgia"
	myStyle.Font.Size = 11
	myStyle.Font.Bold = true
	myStyle.ApplyAlignment = true
	myStyle.ApplyFill = true
	myStyle.ApplyFont = true
	// creating a column that relates to worksheet columns A thru E (index 1 to 5)
	newColumn := xlsx.NewColForRange(1, 5)
	//newColumn.SetWidth(12.5)
	newColumn.SetWidth(15)
	//// we defined a style above, so let's assign this style to all cells of the column
	//newColumn.SetStyle(myStyle)
	//// now associate the sheet with this column
	sh.SetColParameters(newColumn)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
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
func (r *ReportHandler) GetExportEventsByClinicToExcel(c *gin.Context) {
	eventId := c.Query("eventId")
	if eventId == "" {
		c.JSON(400, gin.H{"message": "event id is required"})
		return
	}
	clinicReport, err := r.reportService.ExportClinicReport(eventId)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// Excel
	file := xlsx.NewFile()

	eh, err := file.AddSheet("Events")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}
	// Event report
	//headerRow := sh.AddRow()
	//headerRow.AddCell().SetString("")
	//headerRow.AddCell().SetString("Title")
	//headerRow.AddCell().SetString("Description")
	//headerRow.AddCell().SetString("Start Date")
	//headerRow.AddCell().SetString("Start Time")
	//headerRow.AddCell().SetString("End Date")
	//headerRow.AddCell().SetString("End Time")
	eh.AddRow().AddCell().SetString("โครงการ")
	eh.AddRow().AddCell().SetString("รายละเอียด")
	eh.AddRow().AddCell().SetString("วันที่")
	eh.AddRow().AddCell().SetString("สถานที่")
	eh.AddRow().AddCell().SetString("")
	eh.AddRow().AddCell().SetString("ประธานโครงการ")
	eh.AddRow().AddCell().SetString("ประธานดำเนินการโครงการแพทย์")
	eh.AddRow().AddCell().SetString("ประธานกิจกรรมหน่วยแพทย์อาสา")
	eh.AddRow().AddCell().SetString("ฝ่ายบริการทางการแพทย์ (คลินิกบริการทางการแพทย์)")

	// Set custom detail
	theCell, _ := eh.Cell(0, 1)
	theCell.Value = "โครงการแพทย์อาสาตรวจสุขภาพพระสงฆ์ "
	theCell, _ = eh.Cell(1, 1)
	theCell.Value = "โครงการแพทย์อาสาตรวจสุขภาพพระสงฆ์ถวายเป็นพระกศล แด่สมเด็จพระสงฆราช สกลมหาสังฆปริณายก"
	theCell, _ = eh.Cell(2, 1)
	theCell.Value = "ในวันอาทิตย์ที่ 9 กุมภาพันธ์ 2568"
	theCell, _ = eh.Cell(3, 1)
	theCell.Value = "ณ วัดราชบพิธสถิตมหาสีมารามราชวรวิหาร"

	// Set with col
	newCol := xlsx.NewColForRange(1, 4)
	newCol.SetWidth(35)
	colStyle := xlsx.NewStyle()
	colStyle.Alignment.Horizontal = "left"
	colStyle.Font.Name = "TH Sarabun New"
	colStyle.Font.Size = 16
	colStyle.Font.Bold = true
	colStyle.ApplyAlignment = true
	colStyle.ApplyFont = true
	newCol.SetStyle(colStyle)
	eh.SetColParameters(newCol)

	// Members
	sh, err := file.AddSheet("Members")
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	// set Header
	headerRow := sh.AddRow()
	headerRow.AddCell().SetString("clinic")
	headerRow.AddCell().SetString("Name")
	headerRow.AddCell().SetString("Last Name")
	headerRow.AddCell().SetString("Phone")
	headerRow.AddCell().SetString("Email")
	headerRow.AddCell().SetString("Position")
	headerRow.AddCell().SetString("Organization")
	headerRow.AddCell().SetString("Course")
	headerRow.AddCell().SetString("Member Type")

	c.JSON(200, gin.H{"message": clinicReport})

}
