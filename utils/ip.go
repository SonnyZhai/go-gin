package utils

import (
	"errors"
	"math"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

// DetectIP 检测请求的真实 IP 地址，若为本地 IP 地址，返回：local+ip 地址，若为公网 IP 地址，返回: public+ip 地址
func DetectIP(c *gin.Context) string {
	clientIP := getClientIP(c)
	//本地ip
	if clientIP == "::1" {
		return "local " + "127.0.0.1"
	}

	if isLocalIP(clientIP) {
		return "local " + clientIP
	}

	return "public " + clientIP
}

// ConvertIP 将 IP 地址格式转换，根据传入的参数转换
// input 是需要转换的 IP 地址，target 是期望转换后的格式类型
func ConvertIP(input interface{}, target string) (interface{}, error) {
	switch v := input.(type) {
	case string:
		// 如果输入是字符串，尝试转换为数值或 net.IP
		if target == "uint" {
			return ipString2Long(v)
		} else if target == "net.IP" {
			ip := net.ParseIP(v)
			if ip == nil {
				return nil, errors.New("不正确的 IP 地址字符串")
			}
			return ip, nil
		}
		return nil, errors.New("不支持的目标类型")
	case uint:
		// 如果输入是数值，尝试转换为字符串或 net.IP
		if target == "string" {
			return long2IPString(v)
		} else if target == "net.IP" {
			return long2IP(v)
		}
		return nil, errors.New("不支持的目标类型")
	case net.IP:
		// 如果输入是 net.IP，尝试转换为数值或字符串
		if target == "uint" {
			return ip2Long(v)
		} else if target == "string" {
			return v.String(), nil
		}
		return nil, errors.New("不支持的目标类型")
	default:
		return nil, errors.New("不支持的输入类型")
	}
}

// getClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func getClientIP(c *gin.Context) string {
	// 尝试从 X-Forwarded-For 头中获取 IP
	if ip := c.Request.Header.Get("X-Forwarded-For"); ip != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，取第一个非空的 IP 地址
		ips := strings.Split(ip, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if ip != "" {
				return ip
			}
		}
	}

	// 尝试从 X-Real-IP 头中获取 IP
	if ip := c.Request.Header.Get("X-Real-IP"); ip != "" {
		return strings.TrimSpace(ip)
	}

	// 最后从 RemoteAddr 中获取 IP
	ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// isLocalIP 检测 IP 地址是否是内网地址
func isLocalIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	if parsedIP.IsLoopback() {
		return true
	}

	ip4 := parsedIP.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// ipString2Long 把 IP 字符串转为数值
func ipString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("不正确的 IP 地址")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// long2IPString 把数值转为 IP 字符串
func long2IPString(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("超出 IPv4 地址范围")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

// ip2Long 把 net.IP 转为数值
func ip2Long(ip net.IP) (uint, error) {
	b := ip.To4()
	if b == nil {
		return 0, errors.New("不正确的 IP 地址")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// long2IP 把数值转为 net.IP
func long2IP(i uint) (net.IP, error) {
	if i > math.MaxUint32 {
		return nil, errors.New("超出 IPv4 地址范围")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip, nil
}
