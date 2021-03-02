package request

type DetailCampaignRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}
