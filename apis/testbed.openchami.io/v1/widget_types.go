// Copyright © 2026 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package v1

import (
	"context"
	"github.com/openchami/fabrica/pkg/fabrica"
)

// Widget represents a widget resource
type Widget struct {
	APIVersion string           `json:"apiVersion" yaml:"apiVersion"`
	Kind       string           `json:"kind" yaml:"kind"`
	Metadata   fabrica.Metadata `json:"metadata" yaml:"metadata"`
	Spec       WidgetSpec       `json:"spec" yaml:"spec" validate:"required"`
	Status     WidgetStatus     `json:"status,omitempty" yaml:"status,omitempty"`
}

// WidgetSpec defines the desired state of Widget
type WidgetSpec struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty" validate:"max=200"`
	// Add your spec fields here
}

// WidgetStatus defines the observed state of Widget
type WidgetStatus struct {
	Phase   string `json:"phase,omitempty" yaml:"phase,omitempty"`
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	Ready   bool   `json:"ready" yaml:"ready"`
	// Add your status fields here
}

// Validate implements custom validation logic for Widget
func (r *Widget) Validate(ctx context.Context) error {
	// Add custom validation logic here
	// Example:
	// if r.Spec.Description == "forbidden" {
	//     return errors.New("description 'forbidden' is not allowed")
	// }

	return nil
}

// GetKind returns the kind of the resource
func (r *Widget) GetKind() string {
	return "Widget"
}

// GetName returns the name of the resource
func (r *Widget) GetName() string {
	return r.Metadata.Name
}

// GetUID returns the UID of the resource
func (r *Widget) GetUID() string {
	return r.Metadata.UID
}

// IsHub marks this as the hub/storage version
func (r *Widget) IsHub() {}
