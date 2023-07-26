package otm

// import (
// 	"errors"
// 	"fmt"

// 	"gopkg.in/yaml.v3"
// )

// // Validates whether the model complies with the schema. If not valid, the error message provides a hint
// func (tm OpenThreatModel) Validate() (bool, error) {
// 	if outcome, err := tm.hasRequiredTopLevel(); err != nil {
// 		return outcome, err
// 	}
// 	if outcome, err := tm.areTopLevelRepresentationsWellFormed(); err != nil {
// 		return outcome, err
// 	}
// 	if outcome, err := tm.areTopLevelRepresentationsWellFormed(); err != nil {
// 		return outcome, err
// 	}
// 	if outcome, err := tm.areAssetsWellFormed(); err != nil {
// 		return outcome, err
// 	}
// 	if outcome, err := tm.areComponentsWellFormed(); err != nil {
// 		return outcome, err
// 	}

// 	if outcome, err := tm.areDataflowsValid(); err != nil {
// 		return outcome, err
// 	}

// 	if outcome, err := tm.areTrustZonesValid(); err != nil {
// 		return outcome, err
// 	}

// 	if outcome, err := tm.areThreatsValid(); err != nil {
// 		return outcome, err
// 	}

// 	if outcome, err := tm.areMitigationsValid(); err != nil {
// 		return outcome, err
// 	}

// 	return true, nil
// }

// func (tm OpenThreatModel) hasRequiredTopLevel() (bool, error) {
// 	if tm.OTMVersion == "" {
// 		return false, errors.New("model does not have required otmVersion")
// 	}

// 	if tm.Project.ID == "" || tm.Project.Name == "" {
// 		return false, errors.New("model must have project name and id specified")

// 	}
// 	return true, nil
// }

// func (tm OpenThreatModel) areTopLevelRepresentationsWellFormed() (bool, error) {

// 	for _, rep := range tm.Representations {
// 		if rep.Name == "" || rep.ID == "" || rep.Type == "" {
// 			additionalComment := getAdditionalComment(rep)
// 			return false, fmt.Errorf("representations must have name, id and type specified.%s", additionalComment)
// 		}

// 		if rep.Type == "diagram" && rep.Size == nil {
// 			additionalComment := getAdditionalComment(rep)
// 			return false, fmt.Errorf("representation of type: diagram must have a size.%s", additionalComment)
// 		}

// 		if rep.Type == "code" && rep.Repository == nil {
// 			additionalComment := getAdditionalComment(rep)
// 			return false, fmt.Errorf("representation of type: code must have a repository.%s", additionalComment)
// 		}
// 	}

// 	return true, nil
// }

// func (tm OpenThreatModel) areAssetsWellFormed() (bool, error) {

// 	for _, asset := range tm.Assets {
// 		if asset.Name == "" || asset.ID == "" {
// 			additionalComment := getAdditionalComment(asset)
// 			return false, fmt.Errorf("assets must have name, id and risk specified.%s", additionalComment)
// 		}
// 	}

// 	return true, nil

// }

// func (tm OpenThreatModel) areComponentsWellFormed() (bool, error) {

// 	for _, component := range tm.Components {
// 		if component.Name == "" || component.ID == "" || component.Type == "" { //|| component.Parent == nil {
// 			additionalComment := getAdditionalComment(component)
// 			return false, fmt.Errorf("components must have name, id, type and parent specified.%s", additionalComment)
// 		}

// 		if component.Parent.IsComponent() && component.Parent.IsTrustZone() {
// 			additionalComment := getAdditionalComment(component)
// 			return false, fmt.Errorf("components can have at most one parent.%s", additionalComment)
// 		}

// 		for _, rep := range component.Representations {
// 			if diag, valid := rep.(DiagramRepresentation); valid {
// 				if diag.Representation == "" || diag.ID == "" { //|| diag.Position == nil || diag.Size == nil {
// 					additionalComment := getAdditionalComment(diag)
// 					return false, fmt.Errorf("diagram representations must have representation, id, position and size specified.%s", additionalComment)
// 				}
// 			} else if code, valid := rep.(CodeRepresentation); valid {
// 				if code.Representation == "" || code.ID == "" {
// 					additionalComment := getAdditionalComment(code)
// 					return false, fmt.Errorf("code representations must have representation and id specified.%s", additionalComment)
// 				}
// 			}
// 		}

// 		for _, threat := range component.Threats {
// 			if threat.State == "" || threat.Threat == "" {
// 				additionalComment := getAdditionalComment(threat)
// 				return false, fmt.Errorf("component threats must have state and threat specified.%s", additionalComment)
// 			}

// 			for _, m := range threat.Mitigations {
// 				if m.Mitigation == "" || m.State == "" {
// 					additionalComment := getAdditionalComment(m)
// 					return false, fmt.Errorf("threat mitigations must have state and mitigation specified.%s", additionalComment)
// 				}
// 			}
// 		}
// 	}

// 	return true, nil
// }

// func (tm OpenThreatModel) areDataflowsValid() (bool, error) {

// 	for _, df := range tm.DataFlows {
// 		if df.Name == "" || df.ID == "" || df.Source == "" || df.Destination == "" {
// 			additionalComment := getAdditionalComment(df)
// 			return false, fmt.Errorf("dataflows must have a name, id, source and destination specified.%s", additionalComment)
// 		}

// 		for _, threat := range df.Threats {
// 			if threat.State == "" || threat.Threat == "" {
// 				additionalComment := getAdditionalComment(threat)
// 				return false, fmt.Errorf("dataflow threats must have state and threat specified.%s", additionalComment)
// 			}

// 			for _, m := range threat.Mitigations {
// 				if m.Mitigation == "" || m.State == "" {
// 					additionalComment := getAdditionalComment(m)
// 					return false, fmt.Errorf("threat mitigations must have state and mitigation specified.%s", additionalComment)
// 				}
// 			}
// 		}
// 	}

// 	return true, nil
// }

// func (tm OpenThreatModel) areTrustZonesValid() (bool, error) {
// 	for _, tz := range tm.TrustZones {
// 		if tz.Name == "" || tz.ID == "" { //|| tz.Risk == nil {
// 			additionalComment := getAdditionalComment(tz)
// 			return false, fmt.Errorf("trustZones must have name, id and risk specified.%s", additionalComment)
// 		}

// 		if tz.Parent != nil {
// 			if tz.Parent.IsComponent() && tz.Parent.IsTrustZone() {
// 				additionalComment := getAdditionalComment(tz)
// 				return false, fmt.Errorf("trustZone can have at most one parent.%s", additionalComment)
// 			}
// 		}

// 		for _, rep := range tz.Representations {
// 			if diag, valid := rep.(DiagramRepresentation); valid {
// 				if diag.Representation == "" || diag.ID == "" { //|| diag.Position == nil || diag.Size == nil {
// 					additionalComment := getAdditionalComment(diag)
// 					return false, fmt.Errorf("diagram representations must have representation, id, position and size specified.%s", additionalComment)
// 				}
// 			} else if code, valid := rep.(CodeRepresentation); valid {
// 				if code.Representation == "" || code.ID == "" {
// 					additionalComment := getAdditionalComment(code)
// 					return false, fmt.Errorf("code representations must have representation and id specified.%s", additionalComment)
// 				}
// 			}
// 		}

// 	}
// 	return true, nil
// }

// func (tm OpenThreatModel) areThreatsValid() (bool, error) {

// 	for _, threat := range tm.Threats {
// 		if threat.Name == "" || threat.ID == "" { //|| threat.Risk == nil {
// 			additionalComment := getAdditionalComment(threat)
// 			return false, fmt.Errorf("threats must have name, id and risk specified.%s", additionalComment)
// 		}
// 	}
// 	return true, nil
// }

// func (tm OpenThreatModel) areMitigationsValid() (bool, error) {

// 	for _, m := range tm.Mitigations {
// 		if m.ID == "" || m.Name == "" { //|| m.RiskReduction == nil {
// 			additionalComment := getAdditionalComment(m)
// 			return false, fmt.Errorf("mitigations must have name, id and riskReduction specified.%s", additionalComment)
// 		}
// 	}

// 	return true, nil
// }

// func getAdditionalComment(object interface{}) (comment string) {
// 	b, err := yaml.Marshal(object)
// 	if err != nil {
// 		comment = fmt.Sprintf(" This does not: %s", string(b))
// 	}
// 	return
// }
