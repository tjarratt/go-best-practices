#!/usr/bin/env bash

export GO15VENDOREXPERIMENT=1
ginkgo -r -race -randomizeSuites -randomizeAllSpecs -failOnPending./...
