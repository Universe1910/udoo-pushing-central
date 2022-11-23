import service from "@/utils/request";

// @Tags Campaign
// @Summary Create Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "Create Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaign/createCampaign [post]
export const createCampaign = (data) => {
  return service({
    url: "/campaign/createCampaign",
    method: "post",
    data,
  });
};

// @Tags Campaign
// @Summary Delete Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "Delete Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaign/deleteCampaign [delete]
export const deleteCampaign = (data) => {
  return service({
    url: "/campaign/deleteCampaign",
    method: "delete",
    data,
  });
};

// @Tags Campaign
// @Summary Debug Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "Debug Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully Debug"}"
// @Router /campaign/debugCampaign [post]
export const debugCampaign = (data) => {
  return service({
    url: "/campaign/debugCampaign",
    method: "post",
    data,
  });
};

// @Tags Campaign
// @Summary Delete Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /campaign/deleteCampaign [delete]
export const deleteCampaignByIds = (data) => {
  return service({
    url: "/campaign/deleteCampaignByIds",
    method: "delete",
    data,
  });
};

// @Tags Campaign
// @Summary 更新Campaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Campaign true "更新Campaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /campaign/updateCampaign [put]
export const updateCampaign = (data) => {
  return service({
    url: "/campaign/updateCampaign",
    method: "put",
    data,
  });
};

// @Tags Campaign
// @Summary 用idSearchCampaign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Campaign true "用idSearchCampaign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /campaign/findCampaign [get]
export const findCampaign = (params) => {
  return service({
    url: "/campaign/findCampaign",
    method: "get",
    params,
  });
};

// @Tags Campaign
// @Summary 分页获取Campaign列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Campaign列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /campaign/getCampaignList [get]
export const getCampaignList = (params) => {
  return service({
    url: "/campaign/getCampaignList",
    method: "get",
    params,
  });
};
