package worker

import (
	"fmt"
	"math/rand"
	"time"
)

type logJSON struct {
	Host                    string    `json:"host"`                       // fake-node.test-node.net
	Source                  string    `json:"source"`                     // /var/log/nginx/splunk.test-node.com_access.log
	SourceType              string    `json:"source_type"`                // nginx_access
	BytesSent               uint      `json:"bytes_sent"`                 // 1261
	Charset                 string    `json:"charset"`                    // UTF-8
	Cipher                  string    `json:"cipher"`                     // TLSv1.2/ECDHE-RSA-AES256-GCM-SHA384
	Conn                    uint32    `json:"conn"`                       // 112258461
	ConnReq                 uint32    `json:"conn_req"`                   // 69
	EventType               []string  `json:"eventtype"`                  // 50
	HttpHost                string    `json:"http_host"`                  // fake-opensearch.test-node.com
	HttpReferer             string    `json:"http_referer"`               // -
	HttpUserAgent           string    `json:"http_user_agent"`            // Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36
	HttpXForwardedFor       string    `json:"http_x_forwarded_for"`       // -
	HttpXRequestID          string    `json:"http_request_id"`            // 79228903fbbb83360e51b89b380e3a45
	MaxLines                uint32    `json:"max_lines"`                  // 5
	Offset                  uint32    `json:"offset"`                     // 0
	OutputMode              string    `json:"output_mode"`                // json
	RC                      string    `json:"r_c"`                        // OK
	RemoteAddr              string    `json:"remote_addr"`                // 10.0.0.1
	RemotePort              uint32    `json:"remote_port"`                // 50933
	RemoteUser              string    `json:"remote_user"`                // -
	Request                 string    `json:"request"`                    // GET /en-GB/splunkd/__raw/servicesNS/nobody/search/search/jobs/1697441858.733597_C0CB2F3E-F05D-4110-82C5-C3ADF057936E?output_mode=json&_=1697441857574 HTTP/2.0
	RequestLength           uint32    `json:"request_length"`             // 132
	RequestTime             float32   `json:"request_time"`               // 0.0005
	RTT                     uint32    `json:"rtt"`                        // 70957
	ServerAddr              string    `json:"server_addr"`                // 10.0.0.2
	ServerName              string    `json:"server_name"`                // fake-opensearch.test-node.com
	ServerPort              uint32    `json:"server_port"`                // 443
	SSLSN                   string    `json:"ssl_s_n"`                    // fake-opensearch.test-node.com
	SSLSessionReused        string    `json:"ssl_session_reused"`         // -
	Status                  uint32    `json:"status"`                     // 200
	TimeCombined            time.Time `json:"time_combined"`              // 07:37:44.611
	UnixCategory            string    `json:"unix_category"`              // all_hosts
	UnixGroup               string    `json:"unix_group"`                 // default
	UpstreamAddr            string    `json:"upstream_addr"`              // 10.0.0.1:8000
	UpstreamCacheStatus     string    `json:"upstream_cache_status"`      // -
	UpstreamHeaderTime      float32   `json:"upstream_header_time"`       // 0.000
	UpstreamHttpContentType string    `json:"upstream_http_content_type"` // application/json; charset=UTF-8
	UpstreamName            string    `json:"upstream_name"`              // splunk
	UpstreamResponseTime    float32   `json:"upstream_response_time"`     // 0.006
	UpstreamStatus          uint32    `json:"upstream_status"`            // 200
	Time                    time.Time `json:"_time"`                      // 2023-10-16T09:37:44.000+02:00
	Index                   string    `json:"index"`                      // cds_prod_utlraman
	LineCount               uint32    `json:"linecount"`                  // 1
	Punct                   string    `json:"punct"`                      // [//:::_+]_="-"_="..."_=""_=""_="..."_=""_="-"_="-"
	SplunkServer            string    `json:"splunk_server"`              // fake1.test-node.net
}

func GenerateJSON() logJSON {
	return logJSON{
		Host:                    generateHost(),
		Source:                  generateSource(),
		SourceType:              "nginx_access",
		BytesSent:               1261,
		Charset:                 "UTF-8",
		Cipher:                  "TLSv1.2/ECDHE-RSA-AES256-GCM-SHA384",
		Conn:                    generateUint(),
		ConnReq:                 generateUint(),
		EventType:               []string{"nginx-all-logs", "nginx_ta_data", ""},
		HttpHost:                generateHost(),
		HttpReferer:             "-",
		HttpUserAgent:           generateUserAgent(),
		HttpXForwardedFor:       "-",
		HttpXRequestID:          "79228903fbbb83360e51b89b380e3a45",
		MaxLines:                5,
		Offset:                  0,
		OutputMode:              "json",
		RC:                      "OK",
		RemoteAddr:              generateIPAddress(),
		RemotePort:              generatePort(),
		RemoteUser:              "-",
		Request:                 generateRequest(),
		RequestLength:           132,
		RequestTime:             generateFloat(),
		RTT:                     generateUint(),
		ServerAddr:              generateIPAddress(),
		ServerName:              generateHost(),
		ServerPort:              443,
		SSLSN:                   generateHost(),
		SSLSessionReused:        "-",
		Status:                  200,
		TimeCombined:            generateTimeStamp(),
		UnixCategory:            "all_hosts",
		UnixGroup:               "default",
		UpstreamAddr:            generateIPAddress(),
		UpstreamCacheStatus:     "-",
		UpstreamHeaderTime:      generateFloat(),
		UpstreamHttpContentType: "application/json; charset=UTF-8",
		UpstreamName:            "splunk",
		UpstreamResponseTime:    generateFloat(),
		UpstreamStatus:          200,
		Time:                    generateTimeStamp(),
		Index:                   "cds_prod_ultraman",
		LineCount:               1,
		Punct:                   "[//:::_+]_=\"-\"_=\"...\"_=\"\"_=\"\"_=\"...\"_=\"\"_=\"-\"_=\"-\"",
		SplunkServer:            generateHost(),
	}

}

func generateHost() string {
	return "fake1.test-node.net"
}

func generateSource() string {
	return ""
}

func generateUint() uint32 {
	return rand.Uint32()
}

func generateUserAgent() string {
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
}

func generateIPAddress() string {
	return fmt.Sprintf("10.0.%d.%d", rand.Intn(254)+1, rand.Intn(254)+1)
}

func generatePort() uint32 {
	return rand.Uint32()
}

func generateRequest() string {
	return "GET /en-GB/splunkd/__raw/servicesNS/nobody/search/search/jobs/1697441858.733597_C0CB2F3E-F05D-4110-82C5-C3ADF057936E?output_mode=json&_=1697441857574 HTTP/2.0"
}

func generateFloat() float32 {
	return rand.Float32()
}

func generateTimeStamp() time.Time {
	return time.Now()
}
