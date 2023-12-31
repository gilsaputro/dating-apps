package partner

import (
	"context"
	"encoding/json"
	"fmt"
	"gilsaputro/dating-apps/internal/handler/utilhttp"
	"gilsaputro/dating-apps/internal/service/partner"
	"gilsaputro/dating-apps/internal/service/user"
	"log"
	"net/http"
	"strings"
	"time"
)

// LikedHistoryHandler is func handler for get current partner
func (h *PartnerHandler) LikedHistoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.timeoutInSec)*time.Second)
	defer cancel()

	var err error
	var response utilhttp.StandardResponse
	var code int = http.StatusOK

	defer func() {
		response.Code = code
		if err == nil {
			response.Message = "success"
		} else {
			response.Message = err.Error()
		}

		data, errMarshal := json.Marshal(response)
		if errMarshal != nil {
			log.Println("[LikedHistoryHandler]-Error Marshal Response :", err)
			code = http.StatusInternalServerError
			data = []byte(`{"code":500,"message":"Internal Server Error"}`)
		}
		utilhttp.WriteResponse(w, data, code)
	}()

	var userID int
	var ok bool
	userID, ok = r.Context().Value("id").(int)
	if !ok {
		code = http.StatusInternalServerError
		err = fmt.Errorf("Internal Server Error")
		return
	}

	var isVerified bool
	isVerified, ok = r.Context().Value("isverified").(bool)
	if !ok {
		code = http.StatusInternalServerError
		err = fmt.Errorf("Internal Server Error")
		return
	}

	errChan := make(chan error, 1)
	var partnerInfo []partner.PartnerServiceInfo
	go func(ctx context.Context) {
		partnerInfo, err = h.service.GetListLikedPartner(partner.PartnerServiceRequest{
			UserID:     userID,
			IsVerified: isVerified,
		})
		errChan <- err
	}(ctx)

	select {
	case <-ctx.Done():
		code = http.StatusGatewayTimeout
		err = fmt.Errorf("Timeout")
		return
	case err = <-errChan:
		if err != nil {
			if err == user.ErrUserNameNotExists || err == user.ErrPasswordIsIncorrect || strings.Contains(err.Error(), "not found") {
				code = http.StatusNotFound
				err = fmt.Errorf("Invalid Username or Password")
			} else {
				code = http.StatusInternalServerError
			}
			return
		}
	}

	response = mapListResponse(partnerInfo)
}

func mapListResponse(result []partner.PartnerServiceInfo) utilhttp.StandardResponse {
	var res utilhttp.StandardResponse
	var list []PartnerResponse
	for _, data := range result {
		list = append(list, PartnerResponse{
			PartnerID:   data.PartnerID,
			Fullname:    data.Fullname,
			Status:      data.Fullname,
			CreatedDate: data.CreatedDate,
		})
	}

	if len(list) > 0 {
		res.Data = list
	}

	return res
}
