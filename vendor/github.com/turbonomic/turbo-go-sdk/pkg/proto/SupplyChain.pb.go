// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SupplyChain.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TemplateDTO_TemplateType int32

const (
	TemplateDTO_BASE      TemplateDTO_TemplateType = 0
	TemplateDTO_EXTENSION TemplateDTO_TemplateType = 1
)

var TemplateDTO_TemplateType_name = map[int32]string{
	0: "BASE",
	1: "EXTENSION",
}
var TemplateDTO_TemplateType_value = map[string]int32{
	"BASE":      0,
	"EXTENSION": 1,
}

func (x TemplateDTO_TemplateType) Enum() *TemplateDTO_TemplateType {
	p := new(TemplateDTO_TemplateType)
	*p = x
	return p
}
func (x TemplateDTO_TemplateType) String() string {
	return proto.EnumName(TemplateDTO_TemplateType_name, int32(x))
}
func (x *TemplateDTO_TemplateType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TemplateDTO_TemplateType_value, data, "TemplateDTO_TemplateType")
	if err != nil {
		return err
	}
	*x = TemplateDTO_TemplateType(value)
	return nil
}
func (TemplateDTO_TemplateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type Provider_ProviderType int32

const (
	// HOSTING is a To One relationship toward the provider, and it enforces containment.
	// This means that if the provider is removed, then every contained consumer will also be removed.
	Provider_HOSTING Provider_ProviderType = 0
	// LAYERED_OVER is a To Many relationship toward the provider, without containment.
	Provider_LAYERED_OVER Provider_ProviderType = 1
)

var Provider_ProviderType_name = map[int32]string{
	0: "HOSTING",
	1: "LAYERED_OVER",
}
var Provider_ProviderType_value = map[string]int32{
	"HOSTING":      0,
	"LAYERED_OVER": 1,
}

func (x Provider_ProviderType) Enum() *Provider_ProviderType {
	p := new(Provider_ProviderType)
	*p = x
	return p
}
func (x Provider_ProviderType) String() string {
	return proto.EnumName(Provider_ProviderType_name, int32(x))
}
func (x *Provider_ProviderType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Provider_ProviderType_value, data, "Provider_ProviderType")
	if err != nil {
		return err
	}
	*x = Provider_ProviderType(value)
	return nil
}
func (Provider_ProviderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{2, 0} }

//
// The TemplateDTO message represents entity types (templates) that the probe expects to
// discover in the target. For the probe to load in Operations Manager, it must discover
// entity types that are valid members of the supply chain, and these entities must have
// valid buy/sell relationships. Specifying the set of templates for a probe serves to
// validate that the specific entities the probe discovers and sends to Operations Manager do
// indeed match the entity descriptions the probe is expected to discover.
//
// Specify entity type by setting an EntityType value to the templateClass field.
//
// An entity can maintain a list of commodities that it sells.
//
// An entity can maintain a map of commodities bought (TemplateCommodity objects). Each map key is
// an instance of Provider. For each provider, the map entry is a list of the commodities bought
// from that provider.
//
// The templateType can be either {@code Base} or
// Extension (see TemplateType).
//
// A Base template indicates the initial representation
// of an entity, which means this probe performs the primary discovery of the entity and places it in the market.
// Note that there can be more than one probe that discovers the same Base entity. The template has a
// templatePriority setting that resolves such a collision. The template with the highest priority value
// wins, and discoveries made for the lower-priority template are ignored.
//
// An extension template adds data to already discovered entities. This is a way to extend the
// commodities managed by a base template.
//
type TemplateDTO struct {
	// The type of entity that the template represents. See EntityType
	// for the available types.
	TemplateClass *EntityDTO_EntityType `protobuf:"varint,1,req,name=templateClass,enum=common_dto.EntityDTO_EntityType" json:"templateClass,omitempty"`
	// The template type (Base or Extension), used during the validation process.
	TemplateType *TemplateDTO_TemplateType `protobuf:"varint,2,req,name=templateType,enum=common_dto.TemplateDTO_TemplateType" json:"templateType,omitempty"`
	// The priority of a Base template. For equivalent Base templates, Operations Manager uses the highest-priority
	// template, and discards discovered data from lower-priority Base templates.
	TemplatePriority *int32 `protobuf:"varint,3,req,name=templatePriority" json:"templatePriority,omitempty"`
	// This entity's list of {@link TemplateCommodity} items that it provides.
	CommoditySold []*TemplateCommodity `protobuf:"bytes,5,rep,name=commoditySold" json:"commoditySold,omitempty"`
	// The commodities bought from the different providers.
	// This Map contains the commodities bought where:
	CommodityBought []*TemplateDTO_CommBoughtProviderProp `protobuf:"bytes,6,rep,name=commodityBought" json:"commodityBought,omitempty"`
	// A map that defines the entity types that will be providers or consumers for this template entity.
	// The entry key is an entity type, from the EntityType enumeration. There can only be
	// one instance of each entity type in this map. The entry value is an instance of
	// ExternalEntityLink. Each entity link describes an entity type in the supply chain,
	// and the commodities it buys from or sells to the template entity.
	ExternalLink []*TemplateDTO_ExternalEntityLinkProp `protobuf:"bytes,7,rep,name=externalLink" json:"externalLink,omitempty"`
	// Each set represents a case where the entity must buy one commodity of the set ( a logical or of the set)
	// Note, the entity may buy more than one of the commodities in the set.
	CommBoughtOrSet  []*TemplateDTO_CommBoughtProviderOrSet `protobuf:"bytes,8,rep,name=commBoughtOrSet" json:"commBoughtOrSet,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *TemplateDTO) Reset()                    { *m = TemplateDTO{} }
func (m *TemplateDTO) String() string            { return proto.CompactTextString(m) }
func (*TemplateDTO) ProtoMessage()               {}
func (*TemplateDTO) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *TemplateDTO) GetTemplateClass() EntityDTO_EntityType {
	if m != nil && m.TemplateClass != nil {
		return *m.TemplateClass
	}
	return EntityDTO_SWITCH
}

func (m *TemplateDTO) GetTemplateType() TemplateDTO_TemplateType {
	if m != nil && m.TemplateType != nil {
		return *m.TemplateType
	}
	return TemplateDTO_BASE
}

func (m *TemplateDTO) GetTemplatePriority() int32 {
	if m != nil && m.TemplatePriority != nil {
		return *m.TemplatePriority
	}
	return 0
}

func (m *TemplateDTO) GetCommoditySold() []*TemplateCommodity {
	if m != nil {
		return m.CommoditySold
	}
	return nil
}

func (m *TemplateDTO) GetCommodityBought() []*TemplateDTO_CommBoughtProviderProp {
	if m != nil {
		return m.CommodityBought
	}
	return nil
}

func (m *TemplateDTO) GetExternalLink() []*TemplateDTO_ExternalEntityLinkProp {
	if m != nil {
		return m.ExternalLink
	}
	return nil
}

func (m *TemplateDTO) GetCommBoughtOrSet() []*TemplateDTO_CommBoughtProviderOrSet {
	if m != nil {
		return m.CommBoughtOrSet
	}
	return nil
}

// In some cases, an entity may buy one commodity or another, but it must buy one of the two
// This set represents the set of commodities where the entity must buy one of these.
// It could be that the set contains multiple commodities from the same provider - where only
// one of these will be bought.  Or it could be that there are multiple provider types and the
// entity must buy one.  However, for this set, the entity is only required to buy one of the
// commodities.
type TemplateDTO_CommBoughtProviderOrSet struct {
	CommBought       []*TemplateDTO_CommBoughtProviderProp `protobuf:"bytes,1,rep,name=commBought" json:"commBought,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *TemplateDTO_CommBoughtProviderOrSet) Reset()         { *m = TemplateDTO_CommBoughtProviderOrSet{} }
func (m *TemplateDTO_CommBoughtProviderOrSet) String() string { return proto.CompactTextString(m) }
func (*TemplateDTO_CommBoughtProviderOrSet) ProtoMessage()    {}
func (*TemplateDTO_CommBoughtProviderOrSet) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 0}
}

func (m *TemplateDTO_CommBoughtProviderOrSet) GetCommBought() []*TemplateDTO_CommBoughtProviderProp {
	if m != nil {
		return m.CommBought
	}
	return nil
}

type TemplateDTO_CommBoughtProviderProp struct {
	// Provider entity type created by the probe
	Key *Provider `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	// The list of commodities bought from the provider specified as key.
	Value []*TemplateCommodity `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
	// Specifies if the provider is optional or not.
	IsOptional       *bool  `protobuf:"varint,3,opt,name=isOptional,def=0" json:"isOptional,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TemplateDTO_CommBoughtProviderProp) Reset()         { *m = TemplateDTO_CommBoughtProviderProp{} }
func (m *TemplateDTO_CommBoughtProviderProp) String() string { return proto.CompactTextString(m) }
func (*TemplateDTO_CommBoughtProviderProp) ProtoMessage()    {}
func (*TemplateDTO_CommBoughtProviderProp) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 1}
}

const Default_TemplateDTO_CommBoughtProviderProp_IsOptional bool = false

func (m *TemplateDTO_CommBoughtProviderProp) GetKey() *Provider {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TemplateDTO_CommBoughtProviderProp) GetValue() []*TemplateCommodity {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TemplateDTO_CommBoughtProviderProp) GetIsOptional() bool {
	if m != nil && m.IsOptional != nil {
		return *m.IsOptional
	}
	return Default_TemplateDTO_CommBoughtProviderProp_IsOptional
}

type TemplateDTO_ExternalEntityLinkProp struct {
	Key              *EntityDTO_EntityType `protobuf:"varint,1,req,name=key,enum=common_dto.EntityDTO_EntityType" json:"key,omitempty"`
	Value            *ExternalEntityLink   `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *TemplateDTO_ExternalEntityLinkProp) Reset()         { *m = TemplateDTO_ExternalEntityLinkProp{} }
func (m *TemplateDTO_ExternalEntityLinkProp) String() string { return proto.CompactTextString(m) }
func (*TemplateDTO_ExternalEntityLinkProp) ProtoMessage()    {}
func (*TemplateDTO_ExternalEntityLinkProp) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 2}
}

func (m *TemplateDTO_ExternalEntityLinkProp) GetKey() EntityDTO_EntityType {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return EntityDTO_SWITCH
}

func (m *TemplateDTO_ExternalEntityLinkProp) GetValue() *ExternalEntityLink {
	if m != nil {
		return m.Value
	}
	return nil
}

type TemplateCommodity struct {
	CommodityType *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=commodityType,enum=common_dto.CommodityDTO_CommodityType" json:"commodityType,omitempty"`
	Key           *string                     `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	// Type of the commodity, that charges this one. This must be on of the commodities from
	// the entity (template) is expected to buy. So, this is a link between bought and sold
	// commodity of the same entity
	ChargedBy        []CommodityDTO_CommodityType `protobuf:"varint,3,rep,name=chargedBy,enum=common_dto.CommodityDTO_CommodityType" json:"chargedBy,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *TemplateCommodity) Reset()                    { *m = TemplateCommodity{} }
func (m *TemplateCommodity) String() string            { return proto.CompactTextString(m) }
func (*TemplateCommodity) ProtoMessage()               {}
func (*TemplateCommodity) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *TemplateCommodity) GetCommodityType() CommodityDTO_CommodityType {
	if m != nil && m.CommodityType != nil {
		return *m.CommodityType
	}
	return CommodityDTO_CLUSTER
}

func (m *TemplateCommodity) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *TemplateCommodity) GetChargedBy() []CommodityDTO_CommodityType {
	if m != nil {
		return m.ChargedBy
	}
	return nil
}

// The Provider class creates a template entity that sells commodities to a
// consumer template.
//
// Each Provider instance has a templateClass to define the entity type, which is expressed
// as a member of the EntityType enumeration.
//
// A provider can have one of two types of relationship with the consumer entity -
// HOSTING or LAYERED_OVER (see ProviderType):
//
// HOSTING is a One Provider/Many Consumers relationship, where the provider contains the consumer.
// This means that if the provider is removed, then every consumer it contains will also be removed.
// For example, a PhysicalMachine contains many VirtualMachines. If you remove the PhysicalMachine
// entity, then its contained VMs will also be removed. You should move VMs off of a host before removing it.
//
// LAYERED_OVER is a Many/Many relationship, with no concept of containment. For example, many VMs
// can share more than one datastore. For LayeredOver relationships, you must specify max and min limits
// to determine how many providers can be layered over the given type of consumer. These values are set in the
// cardinalityMax and cardinalityMin members of this class.
type Provider struct {
	// The type of entity that the provider represents. See {@link Entity}
	// for the available types.
	TemplateClass *EntityDTO_EntityType `protobuf:"varint,1,req,name=templateClass,enum=common_dto.EntityDTO_EntityType" json:"templateClass,omitempty"`
	// ProviderType specifies the type of relationship between the provider and the consumer
	ProviderType *Provider_ProviderType `protobuf:"varint,2,req,name=providerType,enum=common_dto.Provider_ProviderType" json:"providerType,omitempty"`
	// For LAYERED_OVER providers, the maximum number of providers allowed for the consumer.
	CardinalityMax *int32 `protobuf:"varint,3,req,name=cardinalityMax" json:"cardinalityMax,omitempty"`
	// For LAYERED_OVER providers, the minimum number of providers allowed for the consumer.
	CardinalityMin   *int32 `protobuf:"varint,4,req,name=cardinalityMin" json:"cardinalityMin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Provider) Reset()                    { *m = Provider{} }
func (m *Provider) String() string            { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()               {}
func (*Provider) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *Provider) GetTemplateClass() EntityDTO_EntityType {
	if m != nil && m.TemplateClass != nil {
		return *m.TemplateClass
	}
	return EntityDTO_SWITCH
}

func (m *Provider) GetProviderType() Provider_ProviderType {
	if m != nil && m.ProviderType != nil {
		return *m.ProviderType
	}
	return Provider_HOSTING
}

func (m *Provider) GetCardinalityMax() int32 {
	if m != nil && m.CardinalityMax != nil {
		return *m.CardinalityMax
	}
	return 0
}

func (m *Provider) GetCardinalityMin() int32 {
	if m != nil && m.CardinalityMin != nil {
		return *m.CardinalityMin
	}
	return 0
}

// ExternalEntityLink is a subclass of {@link EntityLink} that
// describes the buy/sell relationship between an entity discovered by the probe, and
// an external entity.
//
// An external entity is one that exists in the
// Operations Manager topology, but has not been discovered by the probe.
// Operations Manager uses this link to stitch discovered entities into the
// existing topology that's managed by the Operations Manager market. This external
// entity can be a provider or a consumer. The ExternalEntityLink object
// contains a full description of the relationship between the external entity and
// the node entity.
// This description includes the entity types for the buyer and seller, the ProviderType
// (the relationship type for the provider, either HOSTING or LAYERED_OVER}),
// and the list of commodities bought from the provider.
//
// To enable stitching, the external link includes a map of {@code probeEntityDef} items
// and a list of ServerEntityPropertyDef items. These work together to identify which
// external entity to stitch together with the probe's discovered entity. The {@code probeEntityDef}
// items store data to identify the appropriate external entity. For example, a discovered application
// can store the IP address of the hosting VM.
//
// The ServerEntityPropertyDef items
// tell Operations Manager how to find identifying information in the external entities.
// For example, the discovered application stores IP address of the hosting VM. Operations Manager
// will use the ServerEntityPropertyDef to test the current VMs for a matching IP address.
type ExternalEntityLink struct {
	// Consumer entity in the link
	BuyerRef *EntityDTO_EntityType `protobuf:"varint,1,req,name=buyerRef,enum=common_dto.EntityDTO_EntityType" json:"buyerRef,omitempty"`
	// Provider entity in the link
	SellerRef *EntityDTO_EntityType `protobuf:"varint,2,req,name=sellerRef,enum=common_dto.EntityDTO_EntityType" json:"sellerRef,omitempty"`
	// Provider relationship type
	Relationship *Provider_ProviderType `protobuf:"varint,3,req,name=relationship,enum=common_dto.Provider_ProviderType" json:"relationship,omitempty"`
	// The list of commodities the consumer entity buys from the provider entity.
	CommodityDefs []*ExternalEntityLink_CommodityDef `protobuf:"bytes,4,rep,name=commodityDefs" json:"commodityDefs,omitempty"`
	// Commodity key
	Key *string `protobuf:"bytes,5,opt,name=key" json:"key,omitempty"`
	// If one of the entity is to be found outside the probe
	HasExternalEntity *bool `protobuf:"varint,6,opt,name=hasExternalEntity" json:"hasExternalEntity,omitempty"`
	// Map of the name and description of the property belonging to the entity instances
	// discovered by the probe.
	ProbeEntityPropertyDef []*ExternalEntityLink_EntityPropertyDef `protobuf:"bytes,7,rep,name=probeEntityPropertyDef" json:"probeEntityPropertyDef,omitempty"`
	// The meta data representing the property definition of the external entity.
	// The value of the property is used for matching the entity instances.
	ExternalEntityPropertyDefs []*ServerEntityPropDef `protobuf:"bytes,8,rep,name=externalEntityPropertyDefs" json:"externalEntityPropertyDefs,omitempty"`
	// if the provider can replace a placeholder entity created outside of the probe,
	// give a list of EntityTypes it can replace.  For example, a LogicalPool can replace
	// a DiskArray or LogicalPool created by another probe.  The replaced entity must be
	// marked REPLACEABLE by the probe that creates it.
	ReplacesEntity   []EntityDTO_EntityType `protobuf:"varint,9,rep,name=replacesEntity,enum=common_dto.EntityDTO_EntityType" json:"replacesEntity,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ExternalEntityLink) Reset()                    { *m = ExternalEntityLink{} }
func (m *ExternalEntityLink) String() string            { return proto.CompactTextString(m) }
func (*ExternalEntityLink) ProtoMessage()               {}
func (*ExternalEntityLink) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ExternalEntityLink) GetBuyerRef() EntityDTO_EntityType {
	if m != nil && m.BuyerRef != nil {
		return *m.BuyerRef
	}
	return EntityDTO_SWITCH
}

func (m *ExternalEntityLink) GetSellerRef() EntityDTO_EntityType {
	if m != nil && m.SellerRef != nil {
		return *m.SellerRef
	}
	return EntityDTO_SWITCH
}

func (m *ExternalEntityLink) GetRelationship() Provider_ProviderType {
	if m != nil && m.Relationship != nil {
		return *m.Relationship
	}
	return Provider_HOSTING
}

func (m *ExternalEntityLink) GetCommodityDefs() []*ExternalEntityLink_CommodityDef {
	if m != nil {
		return m.CommodityDefs
	}
	return nil
}

func (m *ExternalEntityLink) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ExternalEntityLink) GetHasExternalEntity() bool {
	if m != nil && m.HasExternalEntity != nil {
		return *m.HasExternalEntity
	}
	return false
}

func (m *ExternalEntityLink) GetProbeEntityPropertyDef() []*ExternalEntityLink_EntityPropertyDef {
	if m != nil {
		return m.ProbeEntityPropertyDef
	}
	return nil
}

func (m *ExternalEntityLink) GetExternalEntityPropertyDefs() []*ServerEntityPropDef {
	if m != nil {
		return m.ExternalEntityPropertyDefs
	}
	return nil
}

func (m *ExternalEntityLink) GetReplacesEntity() []EntityDTO_EntityType {
	if m != nil {
		return m.ReplacesEntity
	}
	return nil
}

type ExternalEntityLink_CommodityDef struct {
	Type             *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=type,enum=common_dto.CommodityDTO_CommodityType" json:"type,omitempty"`
	HasKey           *bool                       `protobuf:"varint,2,opt,name=hasKey,def=0" json:"hasKey,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ExternalEntityLink_CommodityDef) Reset()         { *m = ExternalEntityLink_CommodityDef{} }
func (m *ExternalEntityLink_CommodityDef) String() string { return proto.CompactTextString(m) }
func (*ExternalEntityLink_CommodityDef) ProtoMessage()    {}
func (*ExternalEntityLink_CommodityDef) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3, 0}
}

const Default_ExternalEntityLink_CommodityDef_HasKey bool = false

func (m *ExternalEntityLink_CommodityDef) GetType() CommodityDTO_CommodityType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CommodityDTO_CLUSTER
}

func (m *ExternalEntityLink_CommodityDef) GetHasKey() bool {
	if m != nil && m.HasKey != nil {
		return *m.HasKey
	}
	return Default_ExternalEntityLink_CommodityDef_HasKey
}

// Holds a property for the probe's discovered entity that Operations Manager can use to stitch the discovered entity
// into the Operations Manager topology. Each property contains a property name and a description.
//
// The property name specifies which property of the discovered entity you want to match. The discovered
// entity's DTO contains the list of properties and values for that entity. This link must include a property that matches a
// named property in the DTO. Note that the SDK includes builders for different types of entities.
// These builders add properties to the entity DTO, giving them names from the {@link SupplyChainConstants} enumeration.
// However, you can use arbitrary names for these properties, so long as the named property is declared in the
// entity DTO.
//
// The properties you create here match the property names in the target DTO.
// For example, the {link ApplicationBuilder} adds an IP address as a property named {@code SupplyChainConstants.IP_ADDRESS}.
// To match the application IP address in this link, add a property to the link with the same name. By doing that,
// the stitching process can access the value that is set in the discovered entity's DTO.
//
// The property description is an arbitrary string to describe the purpose of this property. This is useful
// when you print out the link via a {@code toString()} method.
type ExternalEntityLink_EntityPropertyDef struct {
	// An entity property name
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// An arbitrary description
	Description      *string `protobuf:"bytes,2,req,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalEntityLink_EntityPropertyDef) Reset()         { *m = ExternalEntityLink_EntityPropertyDef{} }
func (m *ExternalEntityLink_EntityPropertyDef) String() string { return proto.CompactTextString(m) }
func (*ExternalEntityLink_EntityPropertyDef) ProtoMessage()    {}
func (*ExternalEntityLink_EntityPropertyDef) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{3, 1}
}

func (m *ExternalEntityLink_EntityPropertyDef) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ExternalEntityLink_EntityPropertyDef) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*TemplateDTO)(nil), "common_dto.TemplateDTO")
	proto.RegisterType((*TemplateDTO_CommBoughtProviderOrSet)(nil), "common_dto.TemplateDTO.CommBoughtProviderOrSet")
	proto.RegisterType((*TemplateDTO_CommBoughtProviderProp)(nil), "common_dto.TemplateDTO.CommBoughtProviderProp")
	proto.RegisterType((*TemplateDTO_ExternalEntityLinkProp)(nil), "common_dto.TemplateDTO.ExternalEntityLinkProp")
	proto.RegisterType((*TemplateCommodity)(nil), "common_dto.TemplateCommodity")
	proto.RegisterType((*Provider)(nil), "common_dto.Provider")
	proto.RegisterType((*ExternalEntityLink)(nil), "common_dto.ExternalEntityLink")
	proto.RegisterType((*ExternalEntityLink_CommodityDef)(nil), "common_dto.ExternalEntityLink.CommodityDef")
	proto.RegisterType((*ExternalEntityLink_EntityPropertyDef)(nil), "common_dto.ExternalEntityLink.EntityPropertyDef")
	proto.RegisterEnum("common_dto.TemplateDTO_TemplateType", TemplateDTO_TemplateType_name, TemplateDTO_TemplateType_value)
	proto.RegisterEnum("common_dto.Provider_ProviderType", Provider_ProviderType_name, Provider_ProviderType_value)
}

func init() { proto.RegisterFile("SupplyChain.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 823 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x8e, 0xe2, 0x36,
	0x14, 0xde, 0xf0, 0x37, 0x70, 0x60, 0x66, 0xc1, 0xaa, 0xa6, 0x11, 0xd2, 0x76, 0x29, 0x6a, 0xa7,
	0xa8, 0x3f, 0x69, 0x45, 0x7b, 0xb5, 0xaa, 0x2a, 0x2d, 0x90, 0x76, 0x46, 0x9d, 0xc2, 0xd4, 0xa0,
	0x6a, 0xf7, 0x6a, 0xe4, 0x21, 0x66, 0xb1, 0x36, 0xc4, 0x91, 0x63, 0xd0, 0xe6, 0xb6, 0x4f, 0xd1,
	0x8b, 0xbe, 0x47, 0xd5, 0xc7, 0xe8, 0x1b, 0x55, 0x71, 0x02, 0x71, 0x08, 0x33, 0xcb, 0x54, 0xbd,
	0x73, 0xec, 0xef, 0xfb, 0xce, 0xf1, 0xf1, 0x39, 0x9f, 0x02, 0xad, 0xe9, 0xda, 0xf7, 0xdd, 0x70,
	0xb8, 0x24, 0xcc, 0xb3, 0x7c, 0xc1, 0x25, 0x47, 0x30, 0xe7, 0xab, 0x15, 0xf7, 0x6e, 0x1d, 0xc9,
	0xdb, 0x4f, 0x87, 0x6a, 0x3d, 0x9a, 0x4d, 0xe2, 0xc3, 0xee, 0x3f, 0x27, 0x50, 0x9f, 0xd1, 0x95,
	0xef, 0x12, 0x49, 0x47, 0xb3, 0x09, 0xfa, 0x11, 0x4e, 0x65, 0xf2, 0x39, 0x74, 0x49, 0x10, 0x98,
	0x46, 0xa7, 0xd0, 0x3b, 0xeb, 0x77, 0xac, 0x54, 0xc4, 0xb2, 0x3d, 0xc9, 0x64, 0x18, 0x69, 0xc4,
	0xab, 0x59, 0xe8, 0x53, 0x9c, 0xa5, 0xa1, 0x4b, 0x68, 0x6c, 0x37, 0xa2, 0x63, 0xb3, 0xa0, 0x64,
	0x3e, 0xd1, 0x65, 0xb4, 0xb0, 0xbb, 0xb5, 0x92, 0xca, 0x30, 0xd1, 0xe7, 0xd0, 0xdc, 0x7e, 0xdf,
	0x08, 0xc6, 0x05, 0x93, 0xa1, 0x59, 0xec, 0x14, 0x7a, 0x65, 0x9c, 0xdb, 0x47, 0x43, 0x38, 0x55,
	0x01, 0x1c, 0x26, 0xc3, 0x29, 0x77, 0x1d, 0xb3, 0xdc, 0x29, 0xf6, 0xea, 0xfd, 0x67, 0x87, 0xc2,
	0x0e, 0xb7, 0x40, 0x9c, 0xe5, 0xa0, 0x57, 0xf0, 0x74, 0xb7, 0x31, 0xe0, 0xeb, 0x37, 0x4b, 0x69,
	0x56, 0x94, 0x8c, 0x75, 0x5f, 0xf6, 0x91, 0x54, 0x8c, 0xbc, 0x11, 0x7c, 0xc3, 0x1c, 0x2a, 0x6e,
	0x04, 0xf7, 0xf1, 0xbe, 0x0c, 0xc2, 0xd0, 0xa0, 0xef, 0x24, 0x15, 0x1e, 0x71, 0xaf, 0x99, 0xf7,
	0xd6, 0x3c, 0x79, 0x58, 0xd6, 0x4e, 0xb0, 0x71, 0x95, 0x23, 0x86, 0x92, 0xcd, 0x68, 0xa0, 0xd7,
	0x71, 0xb6, 0x71, 0x84, 0x89, 0x98, 0x52, 0x69, 0x56, 0x95, 0xec, 0xd7, 0xc7, 0x67, 0xab, 0x68,
	0x78, 0x5f, 0xa7, 0xcd, 0xe0, 0xc3, 0x7b, 0xb0, 0x68, 0x0c, 0x90, 0xa2, 0x4d, 0xe3, 0x3f, 0x95,
	0x47, 0x53, 0x68, 0xff, 0x69, 0xc0, 0xf9, 0x61, 0x18, 0xba, 0x80, 0xe2, 0x5b, 0x1a, 0xaa, 0x3e,
	0xac, 0xf7, 0x3f, 0xd0, 0x63, 0x6c, 0x61, 0x38, 0x02, 0xa0, 0x6f, 0xa1, 0xbc, 0x21, 0xee, 0x3a,
	0x6a, 0xb5, 0x23, 0xde, 0x3c, 0xc6, 0xa2, 0x4f, 0x01, 0x58, 0x30, 0xf1, 0x25, 0xe3, 0x1e, 0x71,
	0xcd, 0x62, 0xc7, 0xe8, 0x55, 0x5f, 0x94, 0x17, 0xc4, 0x0d, 0x28, 0xd6, 0x0e, 0xda, 0xbf, 0x1b,
	0x70, 0x7e, 0xf8, 0x35, 0x50, 0x3f, 0x4d, 0xef, 0x98, 0x31, 0x51, 0xa9, 0x7e, 0x97, 0xa6, 0x1a,
	0x5d, 0xea, 0xa3, 0x0c, 0x2b, 0x17, 0x26, 0xc9, 0xb5, 0xfb, 0x19, 0x34, 0xf4, 0x31, 0x41, 0x55,
	0x28, 0x0d, 0x5e, 0x4e, 0xed, 0xe6, 0x13, 0x74, 0x0a, 0x35, 0xfb, 0xd5, 0xcc, 0x1e, 0x4f, 0xaf,
	0x26, 0xe3, 0xa6, 0xd1, 0xfd, 0xdb, 0x80, 0x56, 0xee, 0xc6, 0xe8, 0x5a, 0x9b, 0x0d, 0x35, 0x92,
	0x71, 0xca, 0x17, 0x7a, 0xf0, 0x1d, 0x7a, 0xfb, 0x6c, 0x3b, 0x34, 0xce, 0x92, 0x51, 0x33, 0xbe,
	0x76, 0xa1, 0x63, 0xf4, 0x6a, 0xf1, 0xa5, 0x46, 0x50, 0x9b, 0x2f, 0x89, 0x78, 0x43, 0x9d, 0x41,
	0x34, 0xa0, 0xc5, 0x47, 0x68, 0xa7, 0xc4, 0xee, 0x1f, 0x05, 0xa8, 0x6e, 0xdf, 0xf5, 0x7f, 0x33,
	0x23, 0x1b, 0x1a, 0x7e, 0xa2, 0xa9, 0x99, 0xd1, 0xc7, 0x87, 0x7a, 0x69, 0xb7, 0x88, 0x9d, 0x48,
	0xa7, 0xa1, 0x0b, 0x38, 0x9b, 0x13, 0xe1, 0x30, 0x8f, 0xb8, 0x4c, 0x86, 0xbf, 0x90, 0x77, 0x89,
	0x0f, 0xed, 0xed, 0xee, 0xe3, 0x98, 0x67, 0x96, 0xf2, 0x38, 0xe6, 0x75, 0xbf, 0x82, 0x86, 0x1e,
	0x0d, 0xd5, 0xe1, 0xe4, 0x72, 0x32, 0x9d, 0x5d, 0x8d, 0x7f, 0x6a, 0x3e, 0x41, 0x4d, 0x68, 0x5c,
	0xbf, 0x7c, 0x6d, 0x63, 0x7b, 0x74, 0x3b, 0xf9, 0xcd, 0xc6, 0x4d, 0xa3, 0xfb, 0x57, 0x05, 0x50,
	0xbe, 0x3b, 0xd0, 0xf7, 0x50, 0xbd, 0x5b, 0x87, 0x54, 0x60, 0xba, 0x38, 0xba, 0x3e, 0x3b, 0x06,
	0xfa, 0x01, 0x6a, 0x01, 0x75, 0xdd, 0x98, 0x5e, 0x38, 0x92, 0x9e, 0x52, 0xa2, 0xd2, 0x0a, 0xea,
	0x92, 0x68, 0x4e, 0x82, 0x25, 0xf3, 0x55, 0x45, 0x8e, 0x2b, 0xad, 0x4e, 0x43, 0xbf, 0x6a, 0xcd,
	0x39, 0xa2, 0x8b, 0xc0, 0x2c, 0xa9, 0x21, 0xfe, 0xe2, 0xe1, 0xc9, 0xd0, 0x7a, 0x8a, 0x2e, 0x70,
	0x56, 0x61, 0xdb, 0xa1, 0xe5, 0xb4, 0x43, 0xbf, 0x84, 0xd6, 0x92, 0x04, 0x59, 0x19, 0xb3, 0x12,
	0xcd, 0x3c, 0xce, 0x1f, 0xa0, 0x25, 0x9c, 0xfb, 0x82, 0xdf, 0xd1, 0xf8, 0x33, 0x9a, 0x75, 0x2a,
	0x94, 0x74, 0x62, 0xdb, 0xdf, 0xbc, 0x27, 0xb7, 0x1c, 0x0f, 0xdf, 0xa3, 0x87, 0x6e, 0xa1, 0x4d,
	0x33, 0x7c, 0xed, 0x30, 0x48, 0xdc, 0xfc, 0xb9, 0x1e, 0x6d, 0x4a, 0xc5, 0x86, 0x8a, 0x14, 0x1b,
	0x89, 0x3f, 0x20, 0x81, 0x2e, 0xe1, 0x4c, 0x50, 0xdf, 0x25, 0x73, 0x1a, 0x24, 0xb7, 0xae, 0xa9,
	0xf9, 0x7c, 0xff, 0x4b, 0xef, 0xf1, 0xda, 0x0c, 0x1a, 0x7a, 0xcd, 0xd1, 0x0b, 0x28, 0xc9, 0xc7,
	0x7b, 0x89, 0xe2, 0xa0, 0x67, 0x50, 0x59, 0x92, 0xe0, 0xe7, 0xc4, 0x45, 0x76, 0xbe, 0x9b, 0x6c,
	0xb6, 0xaf, 0xa0, 0x95, 0x2f, 0x15, 0x82, 0x92, 0x47, 0x56, 0x71, 0xbc, 0x1a, 0x56, 0x6b, 0xd4,
	0x81, 0xba, 0x43, 0x83, 0xb9, 0x60, 0xca, 0xad, 0x55, 0x13, 0xd7, 0xb0, 0xbe, 0x35, 0xb0, 0xe0,
	0xf9, 0x9c, 0xaf, 0xac, 0xcd, 0x4a, 0xae, 0xc5, 0x1d, 0xb7, 0x22, 0x67, 0x58, 0x70, 0xb1, 0x4a,
	0xb2, 0xb5, 0x1c, 0xc9, 0x07, 0x75, 0xed, 0xbf, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69,
	0x47, 0x98, 0xbb, 0x45, 0x09, 0x00, 0x00,
}