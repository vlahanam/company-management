package requests

import "github.com/vlahanam/company-management/common"

type ListUserRequest struct {
	common.Paging
	KeyWord    *string `json:"keyword,omitempty"`
	CompanyID  *int64  `json:"company_id,omitempty"`
	PositionID *int64  `json:"position_id,omitempty"`
}
