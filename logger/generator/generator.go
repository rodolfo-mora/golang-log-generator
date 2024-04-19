package generator

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// [15/Nov/2023:12:50:26 +0000] http_x_forwarded_for="-" server_addr="85.184.97.2" server_port="443" rtt="87958" remote_addr="172.104.208.16" remote_port="9664" remote_user="-" upstream_cache_status="-" server_name="test-node.net" ssl_s_n="test-node.net" ssl_session_reused="." ssl_protocol="TLSv1.3" cipher="TLSv1.3/TLS_AES_256_GCM_SHA384" ssl_curves="X25519:prime256v1:secp384r1:secp521r1" status="403" upstream_addr="-" upstream_connect_time="-" upstream_response_time="-" upstream_status="-" upstream_http_content_type="-" http_host="test-nodet.net" request="GET /docker-compose.yml HTTP/1.1" uri="/docker-compose.yml" request_time="0.000" request_length="308" r_c="OK" conn="5015904" conn_req="1" bytes_sent="1010" http_user_agent="Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.2117.157 Safari/537.36" http_referer="-" http_accept_language="en" http_accept_encoding="gzip" http_accept="*/*" md5_semi_ja3="TLSv1.3:4da515f0709b53ac7979f777b18bccf4" md5_semi_ja3_headers="TLSv1.3:9616ef01925bb7f224d04473023e7394" http_cache_control="-" upstream_http_cache_control="-" pid="4389" time_combined="12:50:26.160" tmp_ddos_rl="" is_block="" http_x_request_id="699bafd850a9610a8ff64f511e4c3477" upstream_header_time="-" internal="" upstream_name="-" limit_req_status="-" njc="-"

type logJSON struct {
	Host                    string    `json:"host"`                       // fake1.test-node.net
	Source                  string    `json:"source"`                     // /var/log/nginx/test-node.com_access.log
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

// [15/Nov/2023:19:21:49 +0000] http_x_forwarded_for="-" server_addr="85.184.97.44" server_port="443" rtt="39989" remote_addr="83.97.73.87" remote_port="47340" remote_user="-" upstream_cache_status="-" server_name="xns.unibet.dk" ssl_s_n="-" ssl_session_reused="." ssl_protocol="TLSv1.3" cipher="TLSv1.3/TLS_AES_256_GCM_SHA384" ssl_curves="X25519:prime256v1:secp384r1:secp521r1" status="444" upstream_addr="-" upstream_connect_time="-" upstream_response_time="-" upstream_status="-" upstream_http_content_type="-" http_host="85.184.97.44" request="GET /geoserver HTTP/1.1" uri="/geoserver" request_time="0.000" request_length="261" r_c="" conn="5042016" conn_req="1" bytes_sent="0" http_user_agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36" http_referer="http://85.184.97.44:80/geoserver" http_accept_language="-" http_accept_encoding="gzip" http_accept="-" md5_semi_ja3="TLSv1.3:4da515f0709b53ac7979f777b18bccf4" md5_semi_ja3_headers="TLSv1.3:972d4397bf86d9b9af9eb6782c440584" http_cache_control="-" upstream_http_cache_control="-" pid="4388" time_combined="19:21:49.994" tmp_ddos_rl="" is_block="" http_x_request_id="ef28e66609a74832b9b7267a905e4a26" upstream_header_time="-" internal="" upstream_name="-" limit_req_status="-" njc="-"

func GenerateString() string {
	return fmt.Sprintf("http_x_forwarded_for=\"-\" server_addr=\"%v\" server_port=\"443\" rtt=\"39989\" remote_addr=\"%v\" remote_port=\"47340\" remote_user=\"-\" upstream_cache_status=\"-\" upstream_cache_status=\"-\" server_name=\"%v\" ssl_s_n=\"-\" ssl_session_reused=\".\" ssl_protocol=\"TLSv1.3\" cipher=\"TLSv1.3/TLS_AES_256_GCM_SHA384\" ssl_curves=\"X25519:prime256v1:secp384r1:secp521r1\" status=\"%v\" upstream_addr=\"-\" upstream_connect_time=\"-\" upstream_response_time=\"-\" upstream_status=\"-\" upstream_http_content_type=\"-\" http_host=\"%v\" request=\"GET /geoserver HTTP/1.1\" uri=\"/geoserver\" request_time=\"0.000\" request_length=\"261\" r_c=\"\" conn=\"5042016\" conn_req=\"1\" bytes_sent=\"0\" http_user_agent=\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36\" http_referer=\"http://85.184.97.44:80/geoserver\" http_accept_language=\"-\" http_accept_encoding=\"gzip\" http_accept=\"-\" md5_semi_ja3=\"TLSv1.3:%x\" md5_semi_ja3_headers=\"TLSv1.3:%x\" http_cache_control=\"-\" upstream_http_cache_control=\"-\" pid=\"4388\" time_combined=\"19:21:49.994\" tmp_ddos_rl=\"\" is_block=\"\" http_x_request_id=\"%x\" upstream_header_time=\"-\" internal=\"\" upstream_name=\"-\" limit_req_status=\"-\" njc=\"-\"",
		generateIPAddress(),
		generateIPAddress(),
		generateHost(),
		generateStatus(),
		generateIPAddress(),
		generateMD5(),
		generateMD5(),
		generateMD5())
}

func generateStatus() string {
	statusCodes := []string{
		"200",
		"301",
		"404",
		"400",
		"500",
	}
	return statusCodes[rand.Intn(len(statusCodes))]
}

func generateMD5() []byte {
	h := md5.New()
	return h.Sum([]byte(fmt.Sprintf("%v", rand.Intn(10000)+1)))
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
