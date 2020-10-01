package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	addr string = "0.0.0.0:8080"
)

/*
Upload object
*/
type Upload struct {
	Data string
}

// Hàm xử lý tác vụ "upload" nhận dữ liệu từ máy khách (client) qua REST API,
// dữ liệu trao đổi thông qua parameter đầu vào c là 1 con trỏ có kiểu dữ liệu là *gin.Context
func upload(c *gin.Context) {

	// Khởi tạo một cấu trúc cục bộ chỉ được sử dụng trong hàm upload này
	// Giá trị của trường 'Code' phải được thống nhất giữa server-client (ví dụ : trường hợp thành công, thống nhất Code == 0 )
	type uploadResp struct {
		Code    int
		Message string
	}

	// Khởi tạo biến 'request' với kiểu dữ liệu là một cấu trúc toàn cục được định nghĩa trước là Upload và giá trị mặc định empty ( request == {"Data":""} )
	// Khởi tạo biến 'response' với kiểu dữ liệu là một cấu trúc cục bộ uploadResp và giá trị mặc định ( response == {"Code":0, "Message":""} )
	var (
		request  Upload
		response uploadResp
	)

	// Lấy dữ liệu từ body và truyền cho biến 'request'
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Cách khác cũng lấy được dữ liệu tương tự/ tuy nhiên sẽ có một chút khác biệt - đọc trong godoc
	// if err := c.BindJSON(&request); err != nil {
	// 	return
	// }

	// Kiểm tra dữ liệu có hợp lệ hay không
	if request.Data == "" { // if len(request.Data) == 0 {
		response.Code = 40
		response.Message = "Data field is invalid"
		c.JSON(http.StatusOK, response)
		return
	}

	// Các bước xử lý trung gian - không cần quan tâm bước này, chỉ quan tâm giá trị của 'err' (kiểm tra, lưu vào database ...)
	_, err := func() (int, error) {
		// Thao tác xử lý
		return 0, nil
	}()

	// Nếu quá trình xử lý không hoàn thành (lỗi) trả về lỗi
	if err != nil {
		c.JSON(http.StatusOK, uploadResp{
			Code:    53,
			Message: "Exception",
		})
		return
	}

	// Nếu không xảy ra bất cứ lỗi gì - quá trình xử lý hoàn thành tốt đẹp - trả về thành công (nếu thành công, StatusCode luôn = 200 )
	c.JSON(http.StatusOK, uploadResp{
		Code:    0,
		Message: "Successfully",
	})

}

func main() {
	// Khởi tạo cấu hình gin-gonic mặc định
	router := gin.Default()
	// Khởi tạo 1 REST API (method: POST, path: "/api/upload", action: upload function)
	router.POST("/api/upload", upload)
	// Khởi chạy máy chủ với các cấu hình đã thiết lập (phía trên) tại địa chỉ 'addr' ("0.0.0.0:8080")
	router.Run(addr)
}
