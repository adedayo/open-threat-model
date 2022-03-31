package otm_test

import (
	"testing"

	otm "github.com/adedayo/open-threat-model/pkg"
)

func TestParse(t *testing.T) {

}

func makeModel() otm.OpenThreatModel {

	tm := otm.OpenThreatModel{
		OTMVersion: "0.1.0",
		Project: otm.Project{
			Name:         "Project Name",
			ID:           "project_id",
			Description:  "This is a test project",
			Owner:        "Mr. Threat Mod El",
			OwnerContact: "threat@mod.el",
			Attributes:   otm.Attributes{"attr1": "attr value"},
		},
		Representations: []otm.Representation{
			{
				Name:        "Architecture Diagram",
				ID:          "architecture-diagram",
				Type:        "diagram",
				Description: "This is an archiecture diagram",
				Attributes: otm.Attributes{
					"key": "value",
				},
			},
		},
		Assets: []otm.Asset{
			{
				Name:        "Credit Card",
				ID:          "cc-data",
				Description: "Customer Credit Card record",
				Risk: otm.AssetRisk{
					Confidentiality: 100,
					Integrity:       50,
					Availability:    50,
					Comment:         "Confidentiality is critical",
				},
				Attributes: otm.Attributes{"X": "Y"},
			},
		},
		Components: []otm.Component{
			{
				Name:        "Name",
				ID:          "",
				Description: "",
				Tags:        []string{"x", "y", "z"},
				Parent: otm.Parent{
					TrustZone: "",
					Component: "",
				},
				Representations: []otm.RepresentationInterface{
					otm.CodeRepresentation{
						Name:           "Name",
						ID:             "ID",
						Description:    "A description",
						Representation: "",
						Type:           "",
						Attributes:     otm.Attributes{"a": "b"},
						Repository:     &otm.Repository{Url: ""},
						File:           "",
						Line:           0,
						CodeSnippet:    "",
					},
					otm.DiagramRepresentation{
						Name:           "Name",
						ID:             "ID",
						Description:    "A description",
						Representation: "",
						Type:           "",
						Attributes:     otm.Attributes{"a": "b"},
						Position:       otm.Position{X: 12, Y: 29},
						Size:           otm.Dimension{Width: 0, Height: 0},
					},
				},
				Assets: &otm.AssetInstance{
					Processed: []string{"q"},
					Stored:    []string{"r"},
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
				Attributes: otm.Attributes{"a": "b"},
			},
		},
		DataFlows: []otm.DataFlow{
			{
				Name:          "Name",
				ID:            "ID",
				Description:   "A description",
				Tags:          []string{"tag"},
				Bidirectional: false,
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
				Attributes: otm.Attributes{"a": "b"},
			},
		},
		TrustZones: []otm.TrustZone{
			{
				Name:            "Name",
				ID:              "ID",
				Description:     "A description",
				Risk:            otm.TrustZoneRisk{},
				Parent:          otm.Parent{},
				Representations: []otm.RepresentationInterface{},
				Attributes:      otm.Attributes{"a": "b"},
			},
		},
		Threats: []otm.Threat{
			{
				Name:        "Name",
				ID:          "",
				Description: "",
				Categories:  []string{"m"},
				CWEs:        []string{"x"},
				Risk: otm.ThreatRisk{
					Likelihood:        0,
					LikelihoodComment: "",
					Impact:            0,
					ImpactComment:     "",
				},
				Tags:       []string{"t"},
				Attributes: otm.Attributes{"a": "b"},
			},
		},
		Mitigations: []otm.Mitigation{
			{
				Name:          "Name",
				ID:            "ID",
				Description:   "A description",
				RiskReduction: 0,
				Attributes:    otm.Attributes{"a": "b"},
			},
		},
	}

	return tm
}
