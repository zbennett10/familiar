package main

// TODO: Support more.. read these in from a JSON file or external API....
type CampaignSetting int

const (
	EBERRON CampaignSetting = iota
	CUSTOM  CampaignSetting = iota
)

type Edition int

const (
	THIRD          Edition = iota
	THREE_DOT_FIVE Edition = iota
	FOURTH         Edition = iota
	FIFTH          Edition = iota
)

type Campaign struct {
	setting CampaignSetting
	edition Edition
}

func NewCampaign() *Campaign {
	return &Campaign{}
}
