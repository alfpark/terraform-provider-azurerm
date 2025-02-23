package volumegroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VolumeGroupId{}

// VolumeGroupId is a struct representing the Resource ID for a Volume Group
type VolumeGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	VolumeGroupName   string
}

// NewVolumeGroupID returns a new VolumeGroupId struct
func NewVolumeGroupID(subscriptionId string, resourceGroupName string, netAppAccountName string, volumeGroupName string) VolumeGroupId {
	return VolumeGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		VolumeGroupName:   volumeGroupName,
	}
}

// ParseVolumeGroupID parses 'input' into a VolumeGroupId
func ParseVolumeGroupID(input string) (*VolumeGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(VolumeGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VolumeGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NetAppAccountName, ok = parsed.Parsed["netAppAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'netAppAccountName' was not found in the resource id %q", input)
	}

	if id.VolumeGroupName, ok = parsed.Parsed["volumeGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'volumeGroupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVolumeGroupIDInsensitively parses 'input' case-insensitively into a VolumeGroupId
// note: this method should only be used for API response data and not user input
func ParseVolumeGroupIDInsensitively(input string) (*VolumeGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(VolumeGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VolumeGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NetAppAccountName, ok = parsed.Parsed["netAppAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'netAppAccountName' was not found in the resource id %q", input)
	}

	if id.VolumeGroupName, ok = parsed.Parsed["volumeGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'volumeGroupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVolumeGroupID checks that 'input' can be parsed as a Volume Group ID
func ValidateVolumeGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVolumeGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Volume Group ID
func (id VolumeGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/volumeGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.VolumeGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Volume Group ID
func (id VolumeGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountValue"),
		resourceids.StaticSegment("staticVolumeGroups", "volumeGroups", "volumeGroups"),
		resourceids.UserSpecifiedSegment("volumeGroupName", "volumeGroupValue"),
	}
}

// String returns a human-readable description of this Volume Group ID
func (id VolumeGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Volume Group Name: %q", id.VolumeGroupName),
	}
	return fmt.Sprintf("Volume Group (%s)", strings.Join(components, "\n"))
}
