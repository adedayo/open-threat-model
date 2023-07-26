package otm

import (
	_ "embed"

	otm "github.com/adedayo/open-threat-model/pkg/model"
)

var (
	//go:embed api_version.txt
	version string // deployed version will be set by the api_version value from openAPI spec
)

func GetVersion() string {
	return version
}

func MakeModel() otm.OpenThreatModel {

	riskReduction := 75
	desc := "This is a test project"
	owner := "Mr. Threat Mod El"
	ownerContact := "threat@mod.el"
	assetRiskComment := "Confidentiality is critical"
	creditCardRecord := "Customer Credit Card record"
	arch_diagrem_desc := "This is an archiecture diagram"
	tz := "trust-zone-id"
	dataFlowDesc := "A description"
	bidirectional := false
	componentDesc := ""
	lineNo := 100
	tm := otm.OpenThreatModel{
		OtmVersion: "0.1.0",
		Project: otm.Project{
			Name:         "Project Name",
			Id:           "project_id",
			Description:  &desc,
			Owner:        &owner,
			OwnerContact: &ownerContact,
			Tags:         []string{"firstProject"},
			Attributes:   &map[string]string{"attr1": "attr value"},
		},
		Representations: []otm.Representation{
			{
				Name:        "Architecture Diagram",
				Id:          "architecture-diagram",
				Type:        "diagram",
				Description: &arch_diagrem_desc,
				// Size:        &otm.Size{},
				// Repository:  &otm.Repository{},
				Attributes: &map[string]string{"key": "value"},
			},
		},
		Assets: []otm.Asset{
			{
				Name:        "Credit Card",
				Id:          "cc-data",
				Description: &creditCardRecord,
				Risk: &otm.AssetRisk{
					Confidentiality: 100,
					Integrity:       50,
					Availability:    50,
					Comment:         &assetRiskComment,
				},
				Attributes: &map[string]string{"X": "Y"},
			},
		},
		Components: []otm.Component{
			{
				Name:        "Name",
				Id:          "",
				Description: &componentDesc,
				Tags:        []string{"x", "y", "z"},
				Parent: otm.Parent{
					TrustZone: otm.TrustZoneParent{
						TrustZone: tz,
					},
				},
				Representations: []otm.ComponentRepresentationsInner{
					{
						CodeRepresentation: &otm.CodeRepresentation{
							Name:       "Name",
							Id:         "ID",
							Type:       "code",
							Attributes: &map[string]string{"a": "b"},
							// File:           "",
							// Line:           0,
							// CodeSnippet:    "",
						},
					},
					{
						DiagramRepresentation: &otm.DiagramRepresentation{
							Name:       "Name",
							Id:         "ID",
							Type:       "diagram",
							Attributes: &map[string]string{"a": "b"},

							Size: otm.Size{
								Width:  10,
								Height: 10,
							},
						},
					},
					{
						ThreatModelRepresentation: &otm.ThreatModelRepresentation{
							Name: "Name",
							Type: "threat-model",
							Id:   "threat-rep-id",
						},
					},
				},
				Assets: []otm.AssetInstance{
					{
						Processed: []string{"q"},
						Stored:    []string{"r"},
					},
				},
				Threats: []otm.ThreatInstance{
					{
						Threat: "",
						State:  "",
						Mitigations: []otm.MitigationInstance{
							{
								Mitigation: "",
								State:      "",
							},
						},
					},
				},
				Attributes: &map[string]string{"a": "b"},
			},
		},
		Dataflows: []otm.Dataflow{
			{
				Name:          "Name",
				Id:            "ID",
				Description:   &dataFlowDesc,
				Tags:          []string{"tag"},
				Bidirectional: &bidirectional,
				Source:        "",
				Destination:   "",
				Assets:        []string{"q"},
				Threats: []otm.ThreatInstance{
					{
						Threat: "",
						State:  "",
						Mitigations: []otm.MitigationInstance{
							{
								Mitigation: "",
								State:      "",
							},
						},
					},
				},
				Attributes: &map[string]string{"a": "b"},
			},
		},
		TrustZones: []otm.TrustZone{
			{
				Name:        "Name",
				Id:          "ID",
				Description: &desc,
				Risk: otm.TrustZoneRisk{
					TrustRating: 20,
				},
				Parent: &otm.Parent{},
				Representations: []otm.TrustZoneRepresentationsInner{
					{

						Representation: "diagram",
						Id:             "rep-id",
						Name:           &desc,
						Position: &otm.Position{
							X: 10,
							Y: 25,
						},
						Size:       &otm.Size{},
						Attributes: &map[string]string{"a": "b"},
					},
					{
						Representation: "code",
						Id:             "rep-id",
						Name:           &desc,
						File:           &owner,
						Line:           &lineNo,
						CodeSnippet:    &assetRiskComment,
						Attributes:     &map[string]string{"a": "b"},
					},
					{

						Representation: "threat-model",
						Id:             "threat-model-id",
						Name:           &desc,
						Attributes:     &map[string]string{"a": "b"},
					},
				},
				Attributes: &map[string]string{"a": "b"},
			},
		},
		Threats: []otm.Threat{
			{
				Name:        "Threat 1",
				Id:          "threat-1",
				Description: &desc,
				Categories:  []string{"m"},
				Cwes:        []string{"x"},
				Risk: otm.ThreatRisk{
					Likelihood:        10,
					LikelihoodComment: &arch_diagrem_desc,
					Impact:            0,
					ImpactComment:     &arch_diagrem_desc,
				},
				Tags:       []string{"tag1", "tag2"},
				Attributes: &map[string]string{"a": "b"},
			},
		},
		Mitigations: []otm.Mitigation{
			{
				Name:          "Name",
				Id:            "ID",
				Description:   &desc,
				RiskReduction: riskReduction,
				Attributes:    &map[string]string{"a": "b"},
			},
		},
	}

	return tm
}

// type OpenThreatModel struct {
// 	OTMVersion      string           `yaml:"otmVersion" json:"otmVersion"`
// 	Project         Project          `yaml:"project" json:"project"`
// 	Representations []Representation `yaml:"representations,omitempty" json:"representations,omitempty"`
// 	Assets          []Asset          `yaml:"assets,omitempty" json:"assets,omitempty"`
// 	Components      []Component      `yaml:"components,omitempty" json:"components,omitempty"`
// 	DataFlows       []DataFlow       `yaml:"dataflows,omitempty" json:"dataflows,omitempty"`
// 	TrustZones      []TrustZone      `yaml:"trustZones,omitempty" json:"trustZones,omitempty"`
// 	Threats         []Threat         `yaml:"threats,omitempty" json:"threats,omitempty"`
// 	Mitigations     []Mitigation     `yaml:"mitigations,omitempty" json:"mitigations,omitempty"`
// }

// // Gets model entity's name by its id. If no entity has specified id, then return value of found is set to false
// func (tm OpenThreatModel) GetNameByID(id string) (name string, found bool) {

// 	if tm.Project.ID == id {
// 		return tm.Project.Name, true
// 	}

// 	for _, x := range tm.TrustZones {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	for _, x := range tm.Components {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	for _, x := range tm.Assets {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	for _, x := range tm.DataFlows {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	for _, x := range tm.Threats {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	for _, x := range tm.Mitigations {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	for _, x := range tm.Representations {
// 		if x.ID == id {
// 			return x.Name, true
// 		}
// 	}
// 	return
// }

// type Project struct {
// 	Name         string      `yaml:"name" json:"name"`
// 	ID           string      `yaml:"id" json:"id"`
// 	Description  string      `yaml:"description,omitempty" json:"description,omitempty"`
// 	Owner        string      `yaml:"owner" json:"owner"`
// 	OwnerContact string      `yaml:"ownerContact" json:"ownerContact"`
// 	Tags         []string    `yaml:"tags,omitempty" json:"tags,omitempty"`
// 	Attributes   *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type Mitigation struct {
// 	Name          string      `yaml:"name" json:"name"`
// 	ID            string      `yaml:"id" json:"id"`
// 	Description   string      `yaml:"description,omitempty" json:"description,omitempty"`
// 	RiskReduction int         `yaml:"riskReduction,omitempty" json:"riskReduction,omitempty"`
// 	Attributes    *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type Threat struct {
// 	Name        string      `yaml:"name" json:"name"`
// 	ID          string      `yaml:"id" json:"id"`
// 	Description string      `yaml:"description,omitempty" json:"description,omitempty"`
// 	Categories  []string    `yaml:"categories,omitempty" json:"categories,omitempty"`
// 	CWEs        []string    `yaml:"cwes,omitempty" json:"cwes,omitempty"`
// 	Risk        ThreatRisk  `yaml:"risk,omitempty" json:"risk,omitempty"`
// 	Tags        []string    `yaml:"tags,omitempty" json:"tags,omitempty"`
// 	Attributes  *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type ThreatRisk struct {
// 	Likelihood        int    `yaml:"likelihood" json:"likelihood"`
// 	LikelihoodComment string `yaml:"likelihoodComment,omitempty" json:"likelihoodComment,omitempty"`
// 	Impact            int    `yaml:"impact" json:"impact"`
// 	ImpactComment     string `yaml:"impactComment,omitempty" json:"impactComment,omitempty"`
// }

// type TrustZone struct {
// 	Name            string                    `yaml:"name" json:"name"`
// 	ID              string                    `yaml:"id" json:"id"`
// 	Type            string                    `yaml:"type" json:"type"`
// 	Description     string                    `yaml:"description,omitempty" json:"description,omitempty"`
// 	Risk            TrustZoneRisk             `yaml:"risk" json:"risk"`
// 	Parent          *Parent                   `yaml:"parent,omitempty" json:"parent,omitempty"`
// 	Representations []RepresentationInterface `yaml:"representations,omitempty" json:"representations,omitempty"`
// 	Attributes      *Attributes               `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type TrustZoneRisk struct {
// 	TrustRating int `yaml:"trustRating" json:"trustRating"`
// }

// type DataFlow struct {
// 	Name          string           `yaml:"name" json:"name"`
// 	ID            string           `yaml:"id" json:"id"`
// 	Description   string           `yaml:"description,omitempty" json:"description,omitempty"`
// 	Tags          []string         `yaml:"tags,omitempty" json:"tags,omitempty"`
// 	Bidirectional *bool            `yaml:"bidirectional,omitempty" json:"bidirectional,omitempty"`
// 	Source        string           `yaml:"source" json:"source"`
// 	Destination   string           `yaml:"destination" json:"destination"`
// 	Assets        []string         `yaml:"assets,omitempty" json:"assets,omitempty"`
// 	Threats       []ThreatInstance `yaml:"threats,omitempty" json:"threats,omitempty"`
// 	Attributes    *Attributes      `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type Representation struct {
// 	Name        string      `yaml:"name" json:"name"`
// 	ID          string      `yaml:"id" json:"id"`
// 	Type        string      `yaml:"type" json:"type"`
// 	Description string      `yaml:"description,omitempty" json:"description,omitempty"`
// 	Size        *Dimension  `yaml:"size,omitempty" json:"size,omitempty"`
// 	Repository  *Repository `yaml:"repository,omitempty" json:"repository,omitempty"`
// 	Attributes  *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type RepresentationInterface interface {
// 	GetRepresentationType() string
// }

// type DiagramRepresentation struct {
// 	// RepresentationInterface `yaml:"-" json:"-"`
// 	Representation string      `yaml:"representation" json:"representation"`
// 	Name           string      `yaml:"name" json:"name"`
// 	ID             string      `yaml:"id" json:"id"`
// 	Position       Position    `yaml:"position,omitempty" json:"position,omitempty"`
// 	Size           Dimension   `yaml:"size,omitempty" json:"size,omitempty"`
// 	Attributes     *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// func (DiagramRepresentation) GetRepresentationType() string {
// 	return "diagram"
// }

// type CodeRepresentation struct {
// 	// RepresentationInterface `yaml:"-" json:"-"`
// 	Representation string `yaml:"representation" json:"representation"`
// 	Name           string `yaml:"name" json:"name"`
// 	ID             string `yaml:"id" json:"id"`
// 	// Repository     *Repository `yaml:"repository,omitempty" json:"repository,omitempty"`
// 	File        string      `yaml:"file" json:"file"`
// 	Line        int         `yaml:"line" json:"line"`
// 	CodeSnippet string      `yaml:"codeSnippet" json:"codeSnippet"`
// 	Attributes  *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// func (CodeRepresentation) GetRepresentationType() string {
// 	return "code"
// }

// type ThreatModelRepresentation struct {
// 	// RepresentationInterface `yaml:"-" json:"-"`
// 	Representation string      `yaml:"representation" json:"representation"`
// 	Name           string      `yaml:"name" json:"name"`
// 	ID             string      `yaml:"id" json:"id"`
// 	Attributes     *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// func (ThreatModelRepresentation) GetRepresentationType() string {
// 	return "threat-model"
// }

// type Asset struct {
// 	ID          string      `yaml:"id" json:"id"`
// 	Name        string      `yaml:"name" json:"name"`
// 	Description string      `yaml:"description,omitempty" json:"description,omitempty"`
// 	Risk        AssetRisk   `yaml:"risk,omitempty" json:"risk,omitempty"`
// 	Attributes  *Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type Component struct {
// 	Name            string                    `yaml:"name" json:"name"`
// 	ID              string                    `yaml:"id" json:"id"`
// 	Type            string                    `yaml:"type" json:"type"`
// 	Description     string                    `yaml:"description,omitempty" json:"description,omitempty"`
// 	Tags            []string                  `yaml:"tags,omitempty" json:"tags,omitempty"`
// 	Parent          Parent                    `yaml:"parent,omitempty" json:"parent,omitempty"`
// 	Representations []RepresentationInterface `yaml:"representations,omitempty" json:"representations,omitempty"`
// 	Assets          *AssetInstance            `yaml:"assets,omitempty" json:"assets,omitempty"`
// 	Threats         []ThreatInstance          `yaml:"threats,omitempty" json:"threats,omitempty"`
// 	Attributes      *Attributes               `yaml:"attributes,omitempty" json:"attributes,omitempty"`
// }

// type ThreatInstance struct {
// 	Threat      string               `yaml:"threat" json:"threat"`
// 	State       string               `yaml:"state" json:"state"`
// 	Mitigations []MitigationInstance `yaml:"mitigations,omitempty" json:"mitigations,omitempty"`
// }

// type MitigationInstance struct {
// 	Mitigation string `yaml:"mitigation" json:"mitigation"`
// 	State      string `yaml:"state" json:"state"`
// }

// type AssetInstance struct {
// 	Processed []string `yaml:"processed,omitempty" json:"processed,omitempty"`
// 	Stored    []string `yaml:"stored,omitempty" json:"stored,omitempty"`
// }

// // Parent ID. A parent can be a TrustZone or a component
// type Parent struct {
// 	TrustZone *string `yaml:"trustZone,omitempty" json:"trustZone,omitempty"`
// 	Component *string `yaml:"component,omitempty" json:"component,omitempty"`
// }

// func (p Parent) IsTrustZone() bool {
// 	return p.TrustZone != nil
// }

// func (p Parent) IsComponent() bool {
// 	return p.Component != nil
// }

// func (p Parent) GetID() string {
// 	if p.IsTrustZone() {
// 		return *p.TrustZone
// 	}

// 	if p.IsComponent() {
// 		return *p.Component
// 	}
// 	return ""
// }

// type AssetRisk struct {
// 	Confidentiality int     `yaml:"confidentiality" json:"confidentiality"`
// 	Integrity       int     `yaml:"integrity" json:"integrity"`
// 	Availability    int     `yaml:"availability" json:"availability"`
// 	Comment         *string `yaml:"comment,omitempty" json:"comment,omitempty"`
// }

// type Dimension struct {
// 	Width  int `yaml:"width" json:"width"`
// 	Height int `yaml:"height" json:"height"`
// }

// type Position struct {
// 	X int `yaml:"x" json:"x"`
// 	Y int `yaml:"y" json:"y"`
// }

// type Repository struct {
// 	Url string `yaml:"url" json:"url"`
// }

// type Attributes map[string]string
