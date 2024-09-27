package main

import "fmt"

type Desc struct {
	fqName string 
	help string
variableLabels string
constLabelPairs []LabelPair
err error}

type Labels map[string]string

type LabelPair struct {
	Name *string
	Value *string
}

func Foo () {
	wrapDesc(&Desc{},"",Labels{})
}

type V2 struct {}

func (V2) NewDesc(string,string,string,Labels) *Desc {return &Desc{}}

func wrapDesc(desc *Desc, prefix string, labels Labels) *Desc {
	constLabels := Labels{}
	for _, lp := range desc.constLabelPairs {
		constLabels[*lp.Name] = *lp.Value
	}
	for ln, lv := range labels {
		if _, alreadyUsed := constLabels[ln]; alreadyUsed {
			return &Desc{
				fqName:          desc.fqName,
				help:            desc.help,
				variableLabels:  desc.variableLabels,
				constLabelPairs: desc.constLabelPairs,
				err:             fmt.Errorf("attempted wrapping with already existing label name %q", ln),
			}
		}
		constLabels[ln] = lv
	}
	newDesc := V2{}.NewDesc(prefix+desc.fqName, desc.help, desc.variableLabels, constLabels)
	if desc.err != nil {
		newDesc.err = desc.err
	}
	return newDesc
}


