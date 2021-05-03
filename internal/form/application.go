package form

/**
 * appBusinessDomain
 *  1:支付 2:网商 3:门户 4:运营
 * appAuthScope
 *  1:公司内部 2:团队内部 3:个人
 * */
type Application struct {
	AppOwner          string `binding:"Required"`
	ApplicationName   string `binding:"Required"`
	AppBusinessDomain string `binding:"Required"`
	AppAuthScope      string `binding:"Required"`
	AuthMembers       []int  `binding:"Required"`
	AppDescription    string `binding:"Required"`
}
