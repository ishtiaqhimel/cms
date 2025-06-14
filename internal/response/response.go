package response

import (
	"net/http"

	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

type Response struct {
	Success      bool        `json:"success"`
	Message      string      `json:"message,omitempty"`
	Hash         *string     `json:"hash,omitempty"`
	Count        *int        `json:"count,omitempty"`
	PageSize     *int        `json:"page_size,omitempty"`
	PreviousPage *int        `json:"previous_page,omitempty"`
	NextPage     *int        `json:"next_page,omitempty"`
	CurrentPage  *int        `json:"current_page,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func RespondSuccess(msg string, data interface{}) (int, Response) {
	return http.StatusOK, Response{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func RespondSuccessWithHash(msg string, hash string, data interface{}) (int, Response) {
	return http.StatusOK, Response{
		Success: true,
		Message: msg,
		Hash:    &hash,
		Data:    data,
	}
}

func RespondSuccessForList(msg string, count, pageSize, curPage int, data interface{}) (int, Response) {
	var prev, next *int
	if curPage > 1 {
		prev = utils.ToP(curPage - 1)
	}
	if count > pageSize*curPage {
		next = utils.ToP(curPage + 1)
	}
	return http.StatusOK, Response{
		Success:      true,
		Message:      msg,
		Data:         data,
		Count:        &count,
		PageSize:     &pageSize,
		PreviousPage: prev,
		NextPage:     next,
		CurrentPage:  &curPage,
	}
}

func RespondCreated(msg string, data interface{}) (int, Response) {
	return http.StatusCreated, Response{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func RespondSuccessWithNoContent(msg string) (int, Response) {
	return http.StatusOK, Response{
		Success: true,
		Message: msg,
		Data:    nil,
	}
}
