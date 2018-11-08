package registration

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/pkg/supplychain"
)

var (
	transactionType proto.CommodityDTO_CommodityType = proto.CommodityDTO_TRANSACTION
	fakeKey                                          = "key-placeholder"

	transactionTemplateCommWithKey *proto.TemplateCommodity = &proto.TemplateCommodity{Key: &fakeKey, CommodityType: &transactionType}
)

type SupplyChainFactory struct{}

func (f *SupplyChainFactory) CreateSupplyChain() ([]*proto.TemplateDTO, error) {
	// Virtual application supply chain template
	vAppSupplyChainNode, err := f.buildVirtualApplicationSupplyBuilder()
	if err != nil {
		return nil, err
	}
	glog.V(4).Infof("supply chain node : %++v", vAppSupplyChainNode)

	//builder.SetPriority(-1)
	//builder.SetTemplateType(proto.TemplateDTO_BASE)
	//builder.SetTemplateType(proto.TemplateDTO_EXTENSION)

	supplyChainBuilder := supplychain.NewSupplyChainBuilder()
	supplyChainBuilder.Top(vAppSupplyChainNode)

	return supplyChainBuilder.Create()
}

func (f *SupplyChainFactory) buildVirtualApplicationSupplyBuilder() (*proto.TemplateDTO, error) {
	vAppSupplyChainNodeBuilder := supplychain.NewSupplyChainNodeBuilder(proto.EntityDTO_VIRTUAL_APPLICATION)
	vAppSupplyChainNodeBuilder = vAppSupplyChainNodeBuilder.
		Sells(transactionTemplateCommWithKey)
	return vAppSupplyChainNodeBuilder.Create()
}
