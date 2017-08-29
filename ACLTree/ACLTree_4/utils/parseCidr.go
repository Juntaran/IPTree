/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/24 15:23 
  */

package utils

import (
	"bytes"
	"errors"
)

func ParseCidr4(cidr string) (uint32, uint32, error) {
	// 192.168.1.1/16 此时的 ip 为192.168.1.1  ip-[]byte类型

	cidrByte := []byte(cidr)
	if bytes.IndexByte(cidrByte, '.') <= 0 {
		return 0, 0, errors.New("Error: Bad Cidr: " + cidr)
	}

	var mask uint32
	p := bytes.IndexByte(cidrByte, '/')
	for _, i := range cidrByte[p+1:] {
		if i < '0' || i > '9' {
			mask = 0
			break
		}
		// 192.168.1.1/16 此时的 mask 为16
		mask = mask*10 + uint32(i-'0')
	}
	// 左移后的 mask
	mask = 0xffffffff << (32 - mask)

	// 还需左移处理 ip
	ipstr := cidrByte[:p]
	var pointNum uint32 = 0
	var ip uint32 = 0
	var oct uint32 = 0
	for i := 0; i < len(ipstr); i++ {
		if ipstr[i] == '.' {
			pointNum ++
			ip = ip << 8 + oct
			oct = 0
			continue
		}
		if ipstr[i] <= '9' || ipstr[i] >= '0' {
			oct = oct * 10 + uint32(ipstr[i] - '0')
			if oct > 255 {
				return 0, 0, errors.New("Error: Bad IP: " + cidr[:p])
			}
			continue
		}
		return 0, 0, errors.New("Error: Bad IP: " + cidr[:p])
	}

	if pointNum != 3 {
		return 0, 0, errors.New("Error: Bad IP: " + cidr[:p])
	}
	ip = ip << 8 + oct
	return ip, mask, nil
}
