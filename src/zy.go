package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseZy() map[string]string {
	data := `电影天堂,http://caiji.dyttzyapi.com/api.php/provide/vod
			如意,https://cj.rycjapi.com/api.php/provide/vod
			暴风,https://bfzyapi.com/api.php/provide/vod
			天涯,https://tyyszy.com/api.php/provide/vod
			小猫咪,https://zy.xiaomaomi.cc/api.php/provide/vod
			非凡影视,http://ffzy5.tv/api.php/provide/vod
			黑木耳,https://json.heimuer.xyz/api.php/provide/vod
			360,https://360zy.com/api.php/provide/vod
			爱奇艺,https://www.iqiyizyapi.com/api.php/provide/vod
			卧龙,https://wolongzyw.com/api.php/provide/vod
			华为吧,https://cjhwba.com/api.php/provide/vod
			极速,https://jszyapi.com/api.php/provide/vod
			豆瓣,https://dbzy.tv/api.php/provide/vod
			魔爪,https://mozhuazy.com/api.php/provide/vod
			魔都,https://www.mdzyapi.com/api.php/provide/vod
			最大,https://api.zuidapi.com/api.php/provide/vod
			无尽,https://api.wujinapi.me/api.php/provide/vod
			旺旺短剧,https://wwzy.tv/api.php/provide/vod
			爱坤,https://ikunzyapi.com/api.php/provide/vod
			量子,https://cj.lziapi.com/api.php/provide/vod/
			樱花,https://m3u8.apiyhzy.com/api.php/provide/vod
	`

	data += `创艺影视,https://www.30dian.cn/api.php/provide/vod/
			天空,https://api.tiankongapi.com/api.php/provide/vod/
			看看,https://zy.hikan.xyz/api.php/provide/vod/
			量子,http://cj.lziapi.com/api.php/provide/vod/
			闪电,http://sdzyapi.com/api.php/provide/vod/
			影图,https://cj.vodimg.top/api.php/provide/vod/
			飞速,https://www.feisuzyapi.com/api.php/provide/vod/
			TVB云播,http://www.tvyb02.com/api.php/provide/vod/
			飘零影院,https://p2100.net/api.php/provide/vod/
			光速,https://api.guangsuapi.com/api.php/provide/vod/
			速影,https://xn--k5d-suyingtvcom-lc40a84t7o9i3urako0c.suyingok.com/inc/apijson.php
			百度,https://api.apibdzy.com/api.php/provide/vod/
			明帝影视,https://ys.md214.cn/api.php/provide/vod/
			789盘,https://www.rrvipw.com/api.php/provide/vod/
			卧龙,https://collect.wolongzyw.com/api.php/provide/vod/
			U酷,https://api.ukuapi.com/api.php/provide/vod/
			1080库,https://api.1080zyku.com/inc/api_mac10_all.php
			步步高,https://api.yparse.com/api/json
			飘花电影,http://www.zzrhgg.com/api.php/provide/vod/at/xml
			映迷,https://www.inmi.app/api.php/provide/vod/at/xml
			CK,http://www.feifei67.com/api.php/provide/vod/at/xml
			飞速,https://www.feisuzy.com/api.php/provide/vod/at/xml
			忆梦,http://anltv.cn/api.php/provide/vod/at/xml
			6U云,http://zy.ataoju.com/inc/api.php
			八戒,http://cj.bajiecaiji.com/inc/api.php
			韩剧,http://www.hanjuzy.com/inc/api.php
			鸡婆,https://www.jipo.tv/api.php/provide/vod
			优质库,https://hdzyk.com/inc/api.php/inc/ldg_api_all.php
			非凡,http://cj.ffzyapi.com/api.php/provide/vod/at/xml
			天堂,http://vipmv.cc/api.php/provide/vod/at/xml/
			新浪,http://api.xinlangapi.com/xinlangapi.php/provide/vod/at/xml
			明帝,http://cj.md214.cn/api.php/provide/vod/at/xml
			红牛,https://www.hongniuzy2.com/api.php/provide/vod/at/xml/
			快车,https://caiji.kczyapi.com/api.php/provide/vod/at/xml
			金鹰,https://jyzyapi.com/provide/vod/at/xml/
			77韩剧,https://www.77hanju.com/api.php/provide/vod/at/xml/
			天翼,https://www.911ysw.top/tianyi.php/provide/vod/at/xml
			FF9,https://www.ff9.top/api.php/provide/vod/at/xml/
			39影视,https://www.39kan.com/api.php/provide/vod/at/xml
			6U,http://www.6uzy.cc/inc/apijson_vod.php
			速博,https://subocaiji.com/api.php/provide/vod/
			4000,http://4000zy.com/api.php/provide/vod/
			虎牙,https://www.huyaapi.com/api.php/provide/vod/
			小猫咪,http://zy.xiaomaomi.cc/api.php/provide/vod/
			可可,https://caiji.kekezyapi.com/api.php/provide/vod/
			42,https://www.42.la/api.php/provide/vod/
			花旗,https://seacms.huaqi.live/zyapi.php
			蜜蜂影视,https://www.beeyao.com/api.php/provide/vod/
	`

	re := regexp.MustCompile(`(.*),.*(https?://.*api.php/provide/vod)`)
	matches := re.FindAllStringSubmatch(data, -1)
	ret := map[string]string{}

	for _, v := range matches {
		name := strings.TrimSpace(v[1])
		url := strings.TrimSpace(v[2])

		oldName, ok := ret[url]
		if ok {
			fmt.Println("跳过已存在资源", name, oldName)
			continue
		}

		ret[url] = name
	}

	return ret
}
