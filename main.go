package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

// SendingIntervalPolicy enum of all mail sending policy
type SendingIntervalPolicy int

const (
	_ SendingIntervalPolicy = iota
	// FixedTime will send email with the corresponding interval set in config file
	FixedTime
	// GradualIncrease sending interval increase everytime
	// For example we have interval set to 1h
	// Then after we send the first alert mail
	// We send the second mail 2h later after the first mail, the third mail 4h after the second mail, the fourth mail 8h after the third mail, etc
	GradualIncrease
	// DoNotBotherMeAgain do not send any more mail about the same event after the first mail
	DoNotBotherMeAgain
)

type SegmentDownConf struct {
	AlertInterval       time.Duration         `yaml:"alertInterval,omitempty"`
	AlertIntervalPolicy SendingIntervalPolicy `yaml:"alertIntervalPolicy,omitempty"`
	SendRecovery        bool                  `yaml:"recoveryAlert,omitempty"`
	RefreshCache        time.Duration         `yaml:"cacheRefreshSchedule,omitempty"`
}

type ThresholdConf struct {
	ExceedPercent  int           `yaml:"exceedPercent"`
	ExceedDuration time.Duration `yaml:"exceedDuration"`
}
type ConfigSpecs struct {
	SegmentDown            SegmentDownConf `yaml:"segmentDown,omitempty"`
	SegmentAvgCPUThreshold ThresholdConf   `yaml:"segmentAvgCPUThreshold"`
	MasterCPUThreshold     ThresholdConf   `yaml:"masterCPUThreshold"`
	SegmentAvgMemThreshold ThresholdConf   `yaml:"segmentAvgMemThreshold"`
	MasterMemThreshold     ThresholdConf   `yaml:"masterMemThreshold"`
}

type AlertConf struct {
	Version string      `yaml:"gpccVersion,omitempty"`
	Kind    string      `yaml:"kind,omitempty"`
	Spec    ConfigSpecs `yaml:"spec,omitempty"`
}

func (c *AlertConf) getConf() *AlertConf {

	yamlFile, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Printf("Get err while loading config file   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	a := AlertConf{}
	a.getConf()
	fmt.Printf("%v", a.Spec)
}
