package policyinsights

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2RemediationId{}

// Providers2RemediationId is a struct representing the Resource ID for a Providers 2 Remediation
type Providers2RemediationId struct {
	ManagementGroupsNamespace ManagementGroupsNamespaceType
	ManagementGroupId         string
	RemediationName           string
}

// NewProviders2RemediationID returns a new Providers2RemediationId struct
func NewProviders2RemediationID(managementGroupsNamespace ManagementGroupsNamespaceType, managementGroupId string, remediationName string) Providers2RemediationId {
	return Providers2RemediationId{
		ManagementGroupsNamespace: managementGroupsNamespace,
		ManagementGroupId:         managementGroupId,
		RemediationName:           remediationName,
	}
}

// ParseProviders2RemediationID parses 'input' into a Providers2RemediationId
func ParseProviders2RemediationID(input string) (*Providers2RemediationId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2RemediationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2RemediationId{}

	if v, ok := parsed.Parsed["managementGroupsNamespace"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'managementGroupsNamespace' was not found in the resource id %q", input)
		}

		managementGroupsNamespace, err := parseManagementGroupsNamespaceType(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ManagementGroupsNamespace = *managementGroupsNamespace
	}

	if id.ManagementGroupId, ok = parsed.Parsed["managementGroupId"]; !ok {
		return nil, fmt.Errorf("the segment 'managementGroupId' was not found in the resource id %q", input)
	}

	if id.RemediationName, ok = parsed.Parsed["remediationName"]; !ok {
		return nil, fmt.Errorf("the segment 'remediationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviders2RemediationIDInsensitively parses 'input' case-insensitively into a Providers2RemediationId
// note: this method should only be used for API response data and not user input
func ParseProviders2RemediationIDInsensitively(input string) (*Providers2RemediationId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2RemediationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2RemediationId{}

	if v, ok := parsed.Parsed["managementGroupsNamespace"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'managementGroupsNamespace' was not found in the resource id %q", input)
		}

		managementGroupsNamespace, err := parseManagementGroupsNamespaceType(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ManagementGroupsNamespace = *managementGroupsNamespace
	}

	if id.ManagementGroupId, ok = parsed.Parsed["managementGroupId"]; !ok {
		return nil, fmt.Errorf("the segment 'managementGroupId' was not found in the resource id %q", input)
	}

	if id.RemediationName, ok = parsed.Parsed["remediationName"]; !ok {
		return nil, fmt.Errorf("the segment 'remediationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviders2RemediationID checks that 'input' can be parsed as a Providers 2 Remediation ID
func ValidateProviders2RemediationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2RemediationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Remediation ID
func (id Providers2RemediationId) ID() string {
	fmtString := "/providers/%s/managementGroups/%s/providers/Microsoft.PolicyInsights/remediations/%s"
	return fmt.Sprintf(fmtString, string(id.ManagementGroupsNamespace), id.ManagementGroupId, id.RemediationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Remediation ID
func (id Providers2RemediationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ConstantSegment("managementGroupsNamespace", PossibleValuesForManagementGroupsNamespaceType(), "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupIdValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticRemediations", "remediations", "remediations"),
		resourceids.UserSpecifiedSegment("remediationName", "remediationValue"),
	}
}

// String returns a human-readable description of this Providers 2 Remediation ID
func (id Providers2RemediationId) String() string {
	components := []string{
		fmt.Sprintf("Management Groups Namespace: %q", string(id.ManagementGroupsNamespace)),
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Remediation Name: %q", id.RemediationName),
	}
	return fmt.Sprintf("Providers 2 Remediation (%s)", strings.Join(components, "\n"))
}
