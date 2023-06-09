package listener

import (
	"github.com/crosect/cc-go/config"
	"github.com/crosect/cc-go/pubsub"
	"github.com/crosect/cc-go/web/constant"
	"github.com/crosect/cc-go/web/event"
	"github.com/crosect/cc-go/web/log"
	"github.com/crosect/cc-go/web/properties"
	"strings"
)

type RequestCompletedLogListener struct {
	appProps         *config.AppProperties
	httpRequestProps *properties.HttpRequestLogProperties
}

func NewRequestCompletedLogListener(
	appProps *config.AppProperties,
	httpRequestProps *properties.HttpRequestLogProperties,
) pubsub.Subscriber {
	return &RequestCompletedLogListener{
		appProps:         appProps,
		httpRequestProps: httpRequestProps,
	}
}

func (r RequestCompletedLogListener) Supports(e pubsub.Event) bool {
	_, ok := e.(*event.RequestCompletedEvent)
	return ok
}

func (r RequestCompletedLogListener) Handle(e pubsub.Event) {
	if r.httpRequestProps.Disabled {
		return
	}
	ev := e.(*event.RequestCompletedEvent)
	if payload, ok := ev.Payload().(*event.RequestCompletedMessage); ok {
		// TODO Should remove context path in the highest filter
		if r.isDisabled(payload.Method, r.removeContextPath(payload.Uri, r.appProps.Path)) {
			return
		}
		log.Infow([]interface{}{constant.ContextReqMeta, r.makeHttpRequestLog(payload)}, "finish router")
	}
}

func (r RequestCompletedLogListener) isDisabled(method string, uri string) bool {
	for _, urlMatching := range r.httpRequestProps.AllDisabledUrls() {
		if urlMatching.Method != "" && urlMatching.Method != method {
			continue
		}
		if urlMatching.UrlRegexp() != nil && urlMatching.UrlRegexp().MatchString(uri) {
			return true
		}
	}
	return false
}

func (r RequestCompletedLogListener) removeContextPath(uri string, contextPath string) string {
	uri = strings.TrimPrefix(uri, contextPath)
	return "/" + strings.TrimLeft(uri, "/")
}

func (r RequestCompletedLogListener) makeHttpRequestLog(message *event.RequestCompletedMessage) *log.HttpRequestLog {
	return &log.HttpRequestLog{
		LoggingContext: log.LoggingContext{
			UserId:            message.UserId,
			DeviceId:          message.DeviceId,
			DeviceSessionId:   message.DeviceSessionId,
			CorrelationId:     message.CorrelationId,
			TechnicalUsername: message.TechnicalUsername,
		},
		Status:         message.Status,
		ExecutionTime:  message.ExecutionTime.Milliseconds(),
		RequestPattern: message.Mapping,
		RequestPath:    message.Uri,
		Method:         message.Method,
		Query:          message.Query,
		Url:            message.Url,
		RequestId:      message.CorrelationId,
		CallerId:       message.CallerId,
		ClientIp:       message.ClientIpAddress,
		UserAgent:      message.UserAgent,
	}
}
