package hkrpg

import (
	"encoding/base64"
	"fmt"
	"time"

	"mhy-sdk/logger"
	cp "mhy-sdk/protocol/client"

	"github.com/gin-gonic/gin"
	pb "google.golang.org/protobuf/proto"
)

func (h *HkRpg) queryDispatchHandler(c *gin.Context) {
	version := c.Query("version")
	logger.Info("[ADDR:%s][Version:%s]query_dispatch", c.Request.RemoteAddr, version)
	dispatch := new(cp.Dispatch)

	defer func() {
		resp, err := pb.Marshal(dispatch)
		if err != nil {
			logger.Error("pb marshal DispatchRegionsData error: %v", err)
			return
		}
		resp64 := base64.StdEncoding.EncodeToString(resp)
		c.String(200, resp64)
	}()

	client, err := h.GetClientConfigMap(version)
	if err != nil {
		dispatch.Retcode = 4
		dispatch.Msg = err.Error()
		return
	}
	if time.Now().After(client.StopBeginTime.Time) && time.Now().Before(client.StopEndTime.Time) {
		dispatch.Retcode = 4
		dispatch.Msg = fmt.Sprintf(client.StopTips, client.StopEndTime.Time.Unix())
		dispatch.TopSeverRegionName = client.StopUrl
		return
	}

	crList, err := h.GetClientRegionConfigList(version)
	if err != nil {
		dispatch.Retcode = 4
		dispatch.Msg = err.Error()
		return
	}

	serverList := make([]*cp.RegionInfo, 0)
	for _, v := range crList {
		region := h.GetRegionInfo(v.RegionName)
		if region == nil {
			continue
		}
		server := &cp.RegionInfo{
			Name:        region.Title,
			Title:       region.RegionName,
			EnvType:     region.SdkEnv,
			DispatchUrl: region.DispatchUrl,
		}
		serverList = append(serverList, server)
	}

	dispatch.RegionList = serverList
}
