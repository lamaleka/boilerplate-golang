package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/lamaleka/boilerplate-golang/common/constant"
	"github.com/lamaleka/boilerplate-golang/common/enum"
	"github.com/lamaleka/boilerplate-golang/common/rest"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeploymentMode() enum.DeploymentMode {
	if _, err := os.Open(enum.Local.File()); err == nil {
		return enum.Local
	}
	if _, err := os.Open(enum.Dev.File()); err == nil {
		return enum.Dev
	}
	if _, err := os.Open(enum.Prod.File()); err == nil {
		return enum.Prod
	}
	return enum.Local
}
func CalculatePageSize(dataLength int, pageSize int) int {
	if dataLength < pageSize {
		return dataLength
	}
	return pageSize
}

func CalculateTotalPage(totalData int64, pageSize int) int {
	totalPages := int(totalData) / pageSize
	if int(totalData)%pageSize != 0 {
		totalPages++
	}
	return totalPages
}

func BindAndValidate(c echo.Context, payload interface{}) error {
	if err := c.Bind(payload); err != nil {
		return err
	}
	err := c.Validate(payload)
	if err != nil {
		return err
	}
	return nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err != nil
}

//	func IsValidUUIDs(u ...string) (bool, string) {
//		for i := 0; i < len(u); i++ {
//			_, err := uuid.Parse(u[i])
//			if err != nil {
//				return false, u[i]
//			}
//		}
//		return true, ""
//	}
func ParseStringParamsToInt(statusStr string) []int {
	if statusStr == "" {
		return nil
	}
	parts := strings.Split(statusStr, ",")
	statusValues := make([]int, 0, len(parts))

	for _, p := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			continue
		}
		statusValues = append(statusValues, val)
	}
	if len(statusValues) == 0 {
		return nil
	}
	return statusValues
}
func ParseLimitAndOffset(pageSize string, pageNumber string) *rest.DefaultParams {
	pageSizeInt, errPageSizeInt := strconv.Atoi(pageSize)
	if errPageSizeInt != nil {
		pageSizeInt = constant.DefaultPageSize
	}
	pageNumberInt, errPageNumberInt := strconv.Atoi(pageNumber)
	if errPageNumberInt != nil {
		pageNumberInt = constant.DefaultPageNumber
	}
	if pageSizeInt <= 0 {
		pageSizeInt = constant.DefaultPageSize
	}
	if pageSizeInt > 100 {
		pageSizeInt = 100
	}
	offset := pageNumberInt*pageSizeInt - pageSizeInt
	params := &rest.DefaultParams{
		Offset: offset,
		Pagination: &rest.Pagination{
			PageSize:   pageSizeInt,
			PageNumber: pageNumberInt,
		},
	}
	return params
}

func RoundTo3(f float64) float64 {
	return math.Round(f*1000) / 1000
}

func RoundTo2(f float64) float64 {
	return float64(math.Round(float64(f)*100) / 100)
}

func MapStringSearchQuery(conditions *[]string, args *[]interface{}) func(column string, value string) {
	return func(column string, value string) {
		if value != "" {
			*conditions = append(*conditions, fmt.Sprintf("LOWER(%s) LIKE LOWER(?)", column))
			*args = append(*args, fmt.Sprintf("%%%s%%", value))
		}
	}
}
func MapNumericSearchQuery(conditions *[]string, args *[]interface{}) func(column string, value *int) {
	return func(column string, value *int) {
		if value != nil {
			*conditions = append(*conditions, fmt.Sprintf("%s = ?", column))
			*args = append(*args, value)
		}
	}
}
func MapDateSearchQuery(conditions *[]string, args *[]interface{}) func(column string, date string) {
	return func(column string, date string) {
		if date != "" {
			*conditions = append(*conditions, fmt.Sprintf("%s = ?", column))
			*args = append(*args, date)
		}
	}
}

func MapDateRangeSearchQuery(conditions *[]string, args *[]interface{}) func(column string, startDate string, endDate string) {
	return func(column string, startDate string, endDate string) {
		if startDate != "" && endDate != "" {
			*conditions = append(*conditions, fmt.Sprintf("%s BETWEEN ? AND ?", column))
			*args = append(*args, startDate, endDate)
		}
	}
}

func GetTotalData(db *gorm.DB, tableName string, totalData *int64) {
	db.Raw(`
        SELECT ISNULL(SUM(row_count), 0) AS total_rows
        FROM sys.dm_db_partition_stats
        WHERE object_id = OBJECT_ID(?)
        AND index_id < 2
    `, tableName).Scan(totalData)
}

func GetMimeType(extension string) string {
	extension = strings.ToLower(extension)
	if !strings.HasPrefix(extension, ".") {
		extension = "." + extension
	}
	mimeTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".bmp":  "image/bmp",
		".webp": "image/webp",
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}

	mimeType, exists := mimeTypes[extension]
	if exists {
		return mimeType
	}

	return "application/octet-stream" // Default MIME type for unknown extensions
}
