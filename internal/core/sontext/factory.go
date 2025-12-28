package sontext

import (
	"reflect"

	"github.com/go-summer/internal/core/pebble"
)

type sontext struct {
	pebbles          []pebble.Pebble
	typeToPebble     map[reflect.Type][]pebble.Pebble
	buildingWrappers []*pebbleBuildingWrapper
}

var sontextVal = sontext{
	typeToPebble: make(map[reflect.Type][]pebble.Pebble),
}

type pebbleBuildingWrapper struct {
	p             pebble.Pebble
	autowireSpecs []pebble.AutowireSpec
}

func Register(p pebble.Pebble) {
	if p == nil {
		return
	}

	metadata := p.Metadata()
	interfaces := metadata.Types()

	for _, ifaceType := range interfaces {
		if ifaceType == nil {
			continue
		}
		pebbles := sontextVal.typeToPebble[ifaceType]
		if pebbles == nil {
			pebbles = make([]pebble.Pebble, 0)
		}
		pebbles = append(pebbles, p)
		sontextVal.typeToPebble[ifaceType] = pebbles
	}
}

func Build(p pebble.Pebble, autowireSpecs ...pebble.AutowireSpec) {
	bw := &pebbleBuildingWrapper{p, autowireSpecs}
	sontextVal.buildingWrappers = append(sontextVal.buildingWrappers, bw)

	buildAllPebbles()
}

func buildAllPebbles() {
	for _, bw := range sontextVal.buildingWrappers {
		if bw.p.Metadata().IsReady() {
			continue
		}

		allFieldsSet := true
		for _, as := range bw.autowireSpecs {
			autowireCandidate := findPebbleCandidate(as.Type(), as.Name())
			if autowireCandidate == nil {
				allFieldsSet = false
			} else {
				as.Inject(autowireCandidate)
			}
		}
		if allFieldsSet {
			bw.p.Metadata().Ready()
		}

	}
}

func findPebbleCandidate(ttype reflect.Type, name string) pebble.Pebble {
	allCandidates := sontextVal.typeToPebble[ttype]
	if allCandidates == nil || len(allCandidates) == 0 {
		return nil
	}

	if len(allCandidates) == 1 {
		return allCandidates[0]
	}

	if name != "" {
		for _, candidate := range allCandidates {
			if candidate.Metadata().Name() == name {
				return candidate
			}
		}
	}

	// TODO add primary boolean to pebble metadata
	return nil
}
