//  code generated by oapi sdk gen
package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/feishu/oapi-sdk-go/core"
	"github.com/feishu/oapi-sdk-go/httpclient"
	"github.com/feishu/oapi-sdk-go/service/approval/v4"
	"github.com/feishu/oapi-sdk-go/service/bitable/v1"
	"github.com/feishu/oapi-sdk-go/service/calendar/v4"
	"github.com/feishu/oapi-sdk-go/service/contact/v3"
	"github.com/feishu/oapi-sdk-go/service/docx/v1"
	"github.com/feishu/oapi-sdk-go/service/drive_explorer/v2"
	"github.com/feishu/oapi-sdk-go/service/im/v1"
	"github.com/feishu/oapi-sdk-go/service/sheets/v3"
)

type Client struct {
	Calendar      *calendar.CalendarService
	Approval      *approval.ApprovalService
	Contact       *contact.ContactService
	Im            *im.ImService
	Docx          *docx.DocxService
	Bitable       *bitable.BitableService
	DriveExplorer *drive_explorer.DriveExplorerService
	Sheets        *sheets.SheetsService
}

type ClientOptionFunc func(config *core.Config)

func WithAppType(appType core.AppType) ClientOptionFunc {
	return func(config *core.Config) {
		config.AppType = appType
	}
}

func WithDisableTokenCache() ClientOptionFunc {
	return func(config *core.Config) {
		config.EnableTokenCache = false
	}
}

func WithLogger(logger core.Logger) ClientOptionFunc {
	return func(config *core.Config) {
		config.Logger = logger
	}
}

func WithDomain(domain string) ClientOptionFunc {
	return func(config *core.Config) {
		config.Domain = domain
	}
}

func WithLogLevel(logLevel core.LogLevel) ClientOptionFunc {
	return func(config *core.Config) {
		config.LogLevel = logLevel
	}
}

func WithTokenCache(cache core.Cache) ClientOptionFunc {
	return func(config *core.Config) {
		config.TokenCache = cache
	}
}

func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(config *core.Config) {
		config.HelpDeskId = helpdeskID
		config.HelpDeskToken = helpdeskToken
		if helpdeskID != "" && helpdeskToken != "" {
			config.HelpdeskAuthToken = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", helpdeskID, helpdeskToken)))
		}
	}
}

func WithReqTimeout(reqTimeout time.Duration) ClientOptionFunc {
	return func(config *core.Config) {
		config.ReqTimeout = reqTimeout
	}
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) *Client {
	// 构建配置
	config := &core.Config{
		Domain:           FeishuDomain,
		AppId:            appId,
		AppSecret:        appSecret,
		EnableTokenCache: true,
		AppType:          core.AppTypeCustom,
	}
	for _, option := range options {
		option(config)
	}

	// 构建日志器
	core.NewLogger(config)

	// 构建缓存
	core.NewCache(config)

	// 创建httpclient
	httpClient := httpclient.NewHttpClient(config)

	// 创建sdk-client，并初始化服务
	client := &Client{}
	initService(client, httpClient, config)
	return client
}

func initService(client *Client, httpClient *http.Client, config *core.Config) {
	client.Calendar = calendar.NewService(httpClient, config)
	client.Approval = approval.NewService(httpClient, config)
	client.Contact = contact.NewService(httpClient, config)
	client.Im = im.NewService(httpClient, config)
	client.Docx = docx.NewService(httpClient, config)
	client.Bitable = bitable.NewService(httpClient, config)
	client.DriveExplorer = drive_explorer.NewService(httpClient, config)
	client.Sheets = sheets.NewService(httpClient, config)

}

var FeishuDomain = "https://open.feishu.cn"
