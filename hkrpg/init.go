package hkrpg

import (
	"errors"
	"sync"

	"mhy-sdk/constant"
	"mhy-sdk/db"
	"mhy-sdk/logger"

	"github.com/gin-gonic/gin"
)

type HkRpg struct {
	syncRWMutex           sync.RWMutex
	ClientConfigMap       map[string]*constant.ClientConfig
	ClientRegionConfigMap map[string][]*constant.ClientRegionConfig
	RegionMap             map[string]*constant.RegionConfig
}

func (h *HkRpg) NewRouter(router *gin.Engine) {
	router.GET("/query_dispatch", h.queryDispatchHandler)
	router.GET("/query_dispatch/gucooing/sdk", h.queryDispatchHandler)
}

func NewHkRpg() *HkRpg {
	h := &HkRpg{
		syncRWMutex: sync.RWMutex{},
	}
	h.NewConf()

	return h
}

func (h *HkRpg) NewConf() {
	h.syncRWMutex.Lock()
	defer h.syncRWMutex.Unlock()
	// 拉取客户端配置
	for _, v := range db.QueryClientConfig() {
		if h.ClientConfigMap == nil {
			h.ClientConfigMap = make(map[string]*constant.ClientConfig)
		}
		h.ClientConfigMap[v.ClientVersion] = v
	}
	// 拉取区服配置
	if h.RegionMap == nil {
		h.RegionMap = make(map[string]*constant.RegionConfig)
	}
	for _, k := range db.QueryRegionConfig() {
		h.RegionMap[k.RegionName] = k
	}
	// 拉取客户端区服配置
	for _, k := range h.ClientConfigMap {
		if h.ClientRegionConfigMap == nil {
			h.ClientRegionConfigMap = make(map[string][]*constant.ClientRegionConfig)
		}
		regionList := db.QueryClientRegionConfigByName(k.ClientVersion)
		if len(regionList) == 0 {
			logger.Error("[ClientVersion:%s]没有找到此客户端配置的区服", k.ClientVersion)
			continue
		}
		for _, region := range regionList {
			if h.RegionMap[region.RegionName] == nil {
				logger.Error("[RegionName:%s]没有找到此区服", region.RegionName)
				continue
			}
		}
		h.ClientRegionConfigMap[k.ClientVersion] = regionList
	}
}

func (h *HkRpg) GetClientConfigMap(cl string) (*constant.ClientConfig, error) {
	h.syncRWMutex.RLock()
	defer h.syncRWMutex.RUnlock()
	c, ok := h.ClientConfigMap[cl]
	if !ok {
		logger.Error("[ClientVersion:%s]没有找到此客户端配置的区服", cl)
		return nil, errors.New("没有找到客户端配置")
	}
	return c, nil
}

func (h *HkRpg) GetClientRegionConfigList(cl string) ([]*constant.ClientRegionConfig, error) {
	crList, ok := h.ClientRegionConfigMap[cl]
	if !ok {
		return nil, errors.New("没有找到区服配置")
	}
	if len(crList) == 0 {
		return nil, errors.New("没有找到区服配置")
	}
	return crList, nil
}

func (h *HkRpg) GetRegionInfo(name string) *constant.RegionConfig {
	h.syncRWMutex.RLock()
	defer h.syncRWMutex.RUnlock()
	region, ok := h.RegionMap[name]
	if !ok {
		return nil
	}
	return region
}
