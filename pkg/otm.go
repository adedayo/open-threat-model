package otm

type OpenThreatModel struct {
	OTMVersion      string           `yaml:"otmVersion" json:"otmVersion"`
	Project         Project          `yaml:"project" json:"project"`
	Representations []Representation `yaml:"representations,omitempty" json:"representations,omitempty"`
	Assets          []Asset          `yaml:"assets,omitempty" json:"assets,omitempty"`
	Components      []Component      `yaml:"components,omitempty" json:"components,omitempty"`
	DataFlows       []DataFlow       `yaml:"dataflows,omitempty" json:"dataflows,omitempty"`
	TrustZones      []TrustZone      `yaml:"trustZones,omitempty" json:"trustZones,omitempty"`
	Threats         []Threat         `yaml:"threats,omitempty" json:"threats,omitempty"`
	Mitigations     []Mitigation     `yaml:"mitigations,omitempty" json:"mitigations,omitempty"`
}

type Project struct {
	Name         string     `yaml:"name" json:"name"`
	ID           string     `yaml:"id" json:"id"`
	Description  string     `yaml:"description,omitempty" json:"description,omitempty"`
	Owner        string     `yaml:"owner" json:"owner"`
	OwnerContact string     `yaml:"ownerContact" json:"ownerContact"`
	Attributes   Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type Mitigation struct {
	Name          string     `yaml:"name" json:"name"`
	ID            string     `yaml:"id" json:"id"`
	Description   string     `yaml:"description,omitempty" json:"description,omitempty"`
	RiskReduction int        `yaml:"riskReduction" json:"riskReduction"`
	Attributes    Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type Threat struct {
	Name        string     `yaml:"name" json:"name"`
	ID          string     `yaml:"id" json:"id"`
	Description string     `yaml:"description,omitempty" json:"description,omitempty"`
	Categories  []string   `yaml:"categories,omitempty" json:"categories,omitempty"`
	CWEs        []string   `yaml:"cwes,omitempty" json:"cwes,omitempty"`
	Risk        ThreatRisk `yaml:"risk" json:"risk"`
	Tags        []string   `yaml:"tags,omitempty" json:"tags,omitempty"`
	Attributes  Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type ThreatRisk struct {
	Likelihood        int    `yaml:"likelihood" json:"likelihood"`
	LikelihoodComment string `yaml:"likelihoodComment,omitempty" json:"likelihoodComment,omitempty"`
	Impact            int    `yaml:"impact" json:"impact"`
	ImpactComment     string `yaml:"impactComment,omitempty" json:"impactComment,omitempty"`
}

type TrustZone struct {
	Name            string                    `yaml:"name" json:"name"`
	ID              string                    `yaml:"id" json:"id"`
	Description     string                    `yaml:"description,omitempty" json:"description,omitempty"`
	Risk            TrustZoneRisk             `yaml:"risk" json:"risk"`
	Parent          Parent                    `yaml:"parent,omitempty" json:"parent,omitempty"`
	Representations []RepresentationInterface `yaml:"representations,omitempty" json:"representations,omitempty"`
	Attributes      Attributes                `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type TrustZoneRisk struct {
	TrustRating int `yaml:"trustRating" json:"trustRating"`
}

type DataFlow struct {
	Name          string           `yaml:"name" json:"name"`
	ID            string           `yaml:"id" json:"id"`
	Description   string           `yaml:"description,omitempty" json:"description,omitempty"`
	Tags          []string         `yaml:"tags,omitempty" json:"tags,omitempty"`
	Bidirectional bool             `yaml:"bidirectional" json:"bidirectional"`
	Source        string           `yaml:"source" json:"source"`
	Destination   string           `yaml:"destination" json:"destination"`
	Assets        []string         `yaml:"assets,omitempty" json:"assets,omitempty"`
	Threats       []ThreatInstance `yaml:"threats,omitempty" json:"threats,omitempty"`
	Attributes    Attributes       `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type Representation struct {
	Name        string      `yaml:"name" json:"name"`
	ID          string      `yaml:"id" json:"id"`
	Type        string      `yaml:"type" json:"type"`
	Description string      `yaml:"description,omitempty" json:"description,omitempty"`
	Size        *Dimension  `yaml:"size,omitempty" json:"size,omitempty"`
	Repository  *Repository `yaml:"repository,omitempty" json:"repository,omitempty"`
	Attributes  Attributes  `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type RepresentationInterface interface {
}

type DiagramRepresentation struct {
	RepresentationInterface `yaml:"-" json:"-"`
	Name                    string     `yaml:"name" json:"name"`
	ID                      string     `yaml:"id" json:"id"`
	Description             string     `yaml:"description,omitempty" json:"description,omitempty"`
	Representation          string     `yaml:"representation" json:"representation"`
	Type                    string     `yaml:"type" json:"type"`
	Attributes              Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
	Position                Position   `yaml:"position" json:"position"`
	Size                    Dimension  `yaml:"size" json:"size"`
}

type CodeRepresentation struct {
	RepresentationInterface `yaml:"-" json:"-"`
	Name                    string      `yaml:"name" json:"name"`
	ID                      string      `yaml:"id" json:"id"`
	Description             string      `yaml:"description,omitempty" json:"description,omitempty"`
	Representation          string      `yaml:"representation" json:"representation"`
	Type                    string      `yaml:"type" json:"type"`
	Attributes              Attributes  `yaml:"attributes,omitempty" json:"attributes,omitempty"`
	Repository              *Repository `yaml:"repository,omitempty" json:"repository,omitempty"`
	File                    string      `yaml:"file" json:"file"`
	Line                    int         `yaml:"line" json:"line"`
	CodeSnippet             string      `yaml:"codeSnippet" json:"codeSnippet"`
}

type Asset struct {
	Name        string     `yaml:"name" json:"name"`
	ID          string     `yaml:"id" json:"id"`
	Description string     `yaml:"description,omitempty" json:"description,omitempty"`
	Risk        AssetRisk  `yaml:"risk" json:"risk"`
	Attributes  Attributes `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type Component struct {
	Name            string                    `yaml:"name" json:"name"`
	ID              string                    `yaml:"id" json:"id"`
	Description     string                    `yaml:"description,omitempty" json:"description,omitempty"`
	Tags            []string                  `yaml:"tags,omitempty" json:"tags,omitempty"`
	Parent          Parent                    `yaml:"parent,omitempty" json:"parent,omitempty"`
	Representations []RepresentationInterface `yaml:"representations,omitempty" json:"representations,omitempty"`
	Assets          *AssetInstance            `yaml:"assets,omitempty" json:"assets,omitempty"`
	Threats         []ThreatInstance          `yaml:"threats,omitempty" json:"threats,omitempty"`
	Attributes      Attributes                `yaml:"attributes,omitempty" json:"attributes,omitempty"`
}

type ThreatInstance struct {
	Threat      string               `yaml:"threat" json:"threat"`
	State       string               `yaml:"state" json:"state"`
	Mitigations []MitigationInstance `yaml:"mitigations,omitempty" json:"mitigations,omitempty"`
}

type MitigationInstance struct {
	Mitigation string `yaml:"mitigation" json:"mitigation"`
	State      string `yaml:"state" json:"state"`
}

type AssetInstance struct {
	Processed []string `yaml:"processed,omitempty" json:"processed,omitempty"`
	Stored    []string `yaml:"stored,omitempty" json:"stored,omitempty"`
}
type Parent struct {
	TrustZone string `yaml:"trustZone,omitempty" json:"trustZone,omitempty"`
	Component string `yaml:"component,omitempty" json:"component,omitempty"`
}

type AssetRisk struct {
	Confidentiality int    `yaml:"confidentiality" json:"confidentiality"`
	Integrity       int    `yaml:"integrity" json:"integrity"`
	Availability    int    `yaml:"availability" json:"availability"`
	Comment         string `yaml:"comment,omitempty" json:"comment,omitempty"`
}

type Dimension struct {
	Width  int `yaml:"width" json:"width"`
	Height int `yaml:"height" json:"height"`
}

type Position struct {
	X int `yaml:"x" json:"x"`
	Y int `yaml:"y" json:"y"`
}

type Repository struct {
	Url string `yaml:"url" json:"url"`
}

type Attributes map[string]string
