package domain

// AliWebDavConfig USAGE:
//    aliyundrive-webdav [FLAGS] [OPTIONS] --refresh-token <refresh-token>
// USAGE:
// aliyundrive-webdav [FLAGS] [OPTIONS] --refresh-token <refresh-token>
//
// FLAGS:
// -I, --auto-index    Automatically generate index.html
// -h, --help          Prints help information
// --no-trash      Delete file permanently instead of trashing it
// --read-only     Enable read only mode
// -V, --version       Prints version information
//
// OPTIONS:
// -W, --auth-password <auth-password>          WebDAV authentication password [env: WEBDAV_AUTH_PASSWORD=]
// -U, --auth-user <auth-user>                  WebDAV authentication username [env: WEBDAV_AUTH_USER=]
// --cache-size <cache-size>                Directory entries cache size [default: 1000]
// --cache-ttl <cache-ttl>                  Directory entries cache expiration time in seconds [default: 600]
// --domain-id <domain-id>                  Aliyun PDS domain id
// --host <host>                            Listen host [env: HOST=]  [default: 0.0.0.0]
// -p, --port <port>                            Listen port [env: PORT=]  [default: 8080]
// -S, --read-buffer-size <read-buffer-size>
// Read/download buffer size in bytes, defaults to 10MB [default: 10485760]
//
// -r, --refresh-token <refresh-token>          Aliyun drive refresh token [env: REFRESH_TOKEN=]
// --root <root>                            Root directory path [default: /]
// -w, --workdir <workdir>                      Working directory, refresh_token will be stored in there if specified
type AliWebDavConfig struct {

	// flags
	AutoIndex string `json:"auto_index"`
	NoTrash   string `json:"no_trash"`
	ReadOnly  string `json:"read_only"`
	Version   string `json:"version"`

	// option
	AuthPassword        string `json:"auth_password"`
	AuthUser            string `json:"auth_user"`
	CacheSize           string `json:"cache_size"`
	CacheTtl            string `json:"cache_ttl"`
	DomainId            string `json:"domain_id"`
	Host                string `json:"host"`
	Port                string `json:"port"`
	ReadBuffSize        string `json:"read_buff_size"`
	SyncRefreshToken    string `json:"sync_refresh_token"`
	RefreshToken        string `json:"refresh_token"`
	Root                string `json:"root"`
	WorkDir             string `json:"omitempty"`
	WebAccess           bool   `json:"web_access"`
	WebPost             int    `json:"web_post"`
	StartWithOsRetryNum string `json:"start_with_os_retry_num"`
}
