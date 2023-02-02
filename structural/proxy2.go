/*
https://refactoringguru.cn/design-patterns/proxy
本例演示如何使用代理模式在第三方腾讯视频 （TencentVideo， 代码示例中记为 TV） 程序库中添加缓存。
*/
package main

import "fmt"

// 抽象的服务接口
type ThirdPartyTVLib interface {
	listVideos()
	getVideoInfo(id string)
	downloadVideo(id string)
}

// 具体的接口实现类
type ThirdPartyTVClass struct{}

func (tptc *ThirdPartyTVClass) listVideos() {
	fmt.Println("列举当前所有视频")
}

func (tptc *ThirdPartyTVClass) getVideoInfo(id string) {
	fmt.Printf("获取%s视频信息\n", id)
}

func (tptc *ThirdPartyTVClass) downloadVideo(id string) {
	fmt.Printf("下载%s视频\n", id)
}

// 代理类
type Proxy struct {
	tptl  ThirdPartyTVLib
	cache bool
	num   int
}

func NewProxy(tptl ThirdPartyTVLib) ThirdPartyTVLib {
	return &Proxy{
		tptl: tptl,
	}
}

func (p *Proxy) listVideos() {
	p.tptl.listVideos()
}

func (p *Proxy) getVideoInfo(id string) {
	p.tptl.getVideoInfo(id)
}

func (p *Proxy) downloadVideo(id string) {
	p.addNum()
	if p.cache {
		fmt.Println("使用缓存的视频数据")
		return
	}
	p.tptl.downloadVideo(id)
}

func (p *Proxy) addNum() {
	p.num++
	if p.num > 1 {
		p.cache = true
	}
}

func main() {
	tptl := &ThirdPartyTVClass{}
	proxy := NewProxy(tptl)
	proxy.listVideos()
	proxy.getVideoInfo("three body")
	proxy.downloadVideo("three body")

	for i := 0; i < 3; i++ {
		proxy.downloadVideo("three body")
	}
}
