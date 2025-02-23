package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPostion
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPostion{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPostion) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) AdjustAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.AddShares(qtdShares)
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPostion {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

type InvestorAssetPostion struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPostion {
	return &InvestorAssetPostion{
		AssetID: assetID,
		Shares:  shares,
	}
}

func (iap *InvestorAssetPostion) AddShares(qtd int) {
	iap.Shares += qtd
}
