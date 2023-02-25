package types

import (
	"sort"
	"time"
)

type AlertStatus string

const (
	AlertStatusResolved AlertStatus = "resolved"
	AlertStatusFiring   AlertStatus = "firing"
)

type WebhookPayload struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   uint64            `json:"truncatedAlerts"`
	Status            AlertStatus       `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []WebhookAlert    `json:"alerts"`
}

type WebhookAlert struct {
	Status       AlertStatus       `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

func (p *WebhookPayload) LabelKeysExceptCommon() []string {
	m := map[string]struct{}{}
	for _, alert := range p.Alerts {
		for k := range alert.Labels {
			m[k] = struct{}{}
		}
	}

	var keys []string
	for k := range m {
		if _, ok := p.CommonLabels[k]; ok {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func (p *WebhookPayload) AnnotationKeysExceptCommon() []string {
	m := map[string]struct{}{}
	for _, alert := range p.Alerts {
		for k := range alert.Annotations {
			m[k] = struct{}{}
		}
	}

	var keys []string
	for k := range m {
		if _, ok := p.CommonAnnotations[k]; ok {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
