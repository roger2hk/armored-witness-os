# Copyright 2022 The Armored Witness OS authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BUILD_EPOCH := $(shell date -u "+%s")
BUILD_TAGS = linkramsize,linkramstart,disable_fr_auth,linkprintk
REV = $(shell git rev-parse --short HEAD 2> /dev/null)
GIT_SEMVER_TAG ?= $(shell (git describe --tags --exact-match --match 'v*.*.*' 2>/dev/null || git describe --match 'v*.*.*' --tags 2>/dev/null || git describe --tags 2>/dev/null || echo -n v0.0.${BUILD_EPOCH}+`git rev-parse HEAD`) | tail -c +2 )
SRK_HASH ?= 

PROTOC ?= $(shell which protoc)

TAMAGO_SEMVER = $(shell [ -n "${TAMAGO}" -a -x "${TAMAGO}" ] && ${TAMAGO} version | sed 's/.*go\([0-9]\.[0-9]*\.[0-9]*\).*/\1/')
MINIMUM_TAMAGO_VERSION=1.24.1

SHELL = /usr/bin/env bash

ifeq ("${DEBUG}","1")
	BUILD_TAGS := ${BUILD_TAGS},debug
endif
ifeq ("${FAKE_STORAGE}", "1")
	BUILD_TAGS := ${BUILD_TAGS},fake_storage
endif
ifeq ("${FAKE_RPMB}", "1")
	BUILD_TAGS := ${BUILD_TAGS},fake_rpmb
endif

APP := ""
TEXT_START = 0x80010000 # ramStart (defined in mem.go under relevant tamago/soc package) + 0x10000

ifeq ("${BEE}","1")
	TEXT_START := 0x10010000
	BUILD_TAGS := ${BUILD_TAGS},bee
endif

GOENV := GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 GOOS=tamago GOARM=7 GOARCH=arm
ENTRY_POINT := _rt0_arm_tamago
QEMU ?= qemu-system-arm -machine mcimx6ul-evk -cpu cortex-a7 -m 512M \
        -nographic -monitor none -serial null -serial stdio \
        -net nic,model=imx.enet,netdev=net0 -netdev tap,id=net0,ifname=tap0,script=no,downscript=no \
        -semihosting

ARCH = "arm"

GOFLAGS = -tags ${BUILD_TAGS} -trimpath -buildvcs=false -buildmode=exe \
	-ldflags "-T ${TEXT_START} -E ${ENTRY_POINT} -R 0x1000 \
		-X 'main.Revision=${REV}' \
		-X 'main.Version=${GIT_SEMVER_TAG}' \
		-X 'main.SRKHash=${SRK_HASH}' \
		-X 'main.LogVerifier=$(shell test ${LOG_PUBLIC_KEY} && cat ${LOG_PUBLIC_KEY})' \
		-X 'main.LogOrigin=${LOG_ORIGIN}' \
		-X 'main.AppletManifestVerifier=$(shell test ${APPLET_PUBLIC_KEY} && cat ${APPLET_PUBLIC_KEY})' \
		-X 'main.OSManifestVerifier1=$(shell test ${OS_PUBLIC_KEY1} && cat ${OS_PUBLIC_KEY1})' \
		-X 'main.OSManifestVerifier2=$(shell test ${OS_PUBLIC_KEY2} && cat ${OS_PUBLIC_KEY2})'"

.PHONY: clean qemu qemu-gdb

#### primary targets ####

all: trusted_os_embedded_applet witnessctl

# This target is only used for dev builds, since the proto definitions may
# change in development and require re-compilation of protos.
trusted_os: APP=trusted_os
trusted_os: DIR=$(CURDIR)/trusted_os
trusted_os: check_embed_env create_dummy_applet proto elf manifest

trusted_os_embedded_applet: APP=trusted_os
trusted_os_embedded_applet: DIR=$(CURDIR)/trusted_os
trusted_os_embedded_applet: check_embed_env copy_applet proto elf manifest imx
trusted_os_embedded_applet:

witnessctl: check_tamago
	@echo "building armored-witness control tool"
	@cd $(CURDIR)/cmd/witnessctl && GOPATH="${BUILD_GOPATH}" ${TAMAGO} build -v \
		-ldflags "-s -w -X 'main.Revision=${REV}'" \
		-o $(CURDIR)/bin/witnessctl

# This target builds the Trusted OS without signing it as it is intended to be
# used by the GCP build process and signed there.
trusted_os_release: APP=trusted_os
trusted_os_release: DIR=$(CURDIR)/trusted_os
trusted_os_release: check_embed_env create_dummy_applet elf

## Targets for managing a local serverless log instance for dev/testing FT related bits.

## log_initialise initialises the log stored under ${LOG_STORAGE_DIR}.
log_initialise:
	echo "(Re-)initialising log at ${LOG_STORAGE_DIR}"
	go run github.com/transparency-dev/serverless-log/cmd/integrate@a56a93b5681e5dc231882ac9de435c21cb340846 \
		--storage_dir=${LOG_STORAGE_DIR} \
		--origin=${LOG_ORIGIN} \
		--private_key=${LOG_PRIVATE_KEY} \
		--public_key=${LOG_PUBLIC_KEY} \
		--initialise

## log_os adds the trusted_os_manifest file created during the build to the dev FT log.
log_os: LOG_STORAGE_DIR=$(DEV_LOG_DIR)/log
log_os: LOG_ARTEFACT_DIR=$(DEV_LOG_DIR)/artefacts
log_os: ARTEFACT_HASH=$(shell sha256sum ${CURDIR}/bin/trusted_os.elf | cut -f1 -d" ")
log_os:
	@if [ "${LOG_PRIVATE_KEY}" == "" -o "${LOG_PUBLIC_KEY}" == "" ]; then \
		echo "You need to set LOG_PRIVATE_KEY and LOG_PUBLIC_KEY variables"; \
		exit 1; \
	fi
	@if [ "${DEV_LOG_DIR}" == "" ]; then \
		echo "You need to set the DEV_LOG_DIR variable"; \
		exit 1; \
	fi

	@if [ ! -f ${LOG_STORAGE_DIR}/checkpoint ]; then \
		make log_initialise LOG_STORAGE_DIR="${LOG_STORAGE_DIR}" ; \
	fi
	go run github.com/transparency-dev/serverless-log/cmd/sequence@a56a93b5681e5dc231882ac9de435c21cb340846 \
		--storage_dir=${LOG_STORAGE_DIR} \
		--origin=${LOG_ORIGIN} \
		--public_key=${LOG_PUBLIC_KEY} \
		--entries=${CURDIR}/bin/trusted_os_manifest
	-go run github.com/transparency-dev/serverless-log/cmd/integrate@a56a93b5681e5dc231882ac9de435c21cb340846 \
		--storage_dir=${LOG_STORAGE_DIR} \
		--origin=${LOG_ORIGIN} \
		--private_key=${LOG_PRIVATE_KEY} \
		--public_key=${LOG_PUBLIC_KEY}
	@mkdir -p ${LOG_ARTEFACT_DIR}
	cp ${CURDIR}/bin/trusted_os.elf ${LOG_ARTEFACT_DIR}/${ARTEFACT_HASH}


#### ARM targets ####

imx: $(APP).imx
elf: $(APP).elf
manifest: $(APP)_manifest

proto:
	@echo "generating protobuf classes"
	-rm -f $(CURDIR)/api/*.pb.go
	PATH=$(shell go env GOPATH | awk -F":" '{print $$1"/bin"}'):${PATH} ${PROTOC} --proto_path=$(CURDIR)/api --go_out=$(CURDIR)/api api.proto

$(APP).bin: CROSS_COMPILE=arm-none-eabi-
$(APP).bin: $(APP).elf
	$(CROSS_COMPILE)objcopy -j .text -j .rodata -j .shstrtab -j .typelink \
	    -j .itablink -j .gopclntab -j .go.buildinfo -j .noptrdata -j .data \
	    -j .bss --set-section-flags .bss=alloc,load,contents \
	    -j .noptrbss --set-section-flags .noptrbss=alloc,load,contents \
	    $(CURDIR)/bin/$(APP).elf -O binary $(CURDIR)/bin/$(APP).bin

$(APP).imx: $(APP).bin $(APP).dcd
	echo "## disabling TZASC bypass in DCD for pre-DDR initialization ##"; \
	chmod 644 $(CURDIR)/bin/$(APP).dcd; \
	echo "DATA 4 0x020e4024 0x00000001  # TZASC_BYPASS" >> $(CURDIR)/bin/$(APP).dcd; \
	mkimage -n $(CURDIR)/bin/$(APP).dcd -T imximage -e $(TEXT_START) -d $(CURDIR)/bin/$(APP).bin $(CURDIR)/bin/$(APP).imx
	# Copy entry point from ELF file
	dd if=$(CURDIR)/bin/$(APP).elf of=$(CURDIR)/bin/$(APP).imx bs=1 count=4 skip=24 seek=4 conv=notrunc

$(APP).dcd: check_tamago
$(APP).dcd: GOMODCACHE=$(shell ${TAMAGO} env GOMODCACHE)
$(APP).dcd: TAMAGO_PKG=$(shell grep "github.com/usbarmory/tamago v" go.mod | awk '{print $$1"@"$$2}')
$(APP).dcd: dcd

#### utilities ####

# Various strings need to be embedded into the binary, keys, log info, etc. check they are present.
check_embed_env:
	@if [ "${LOG_ORIGIN}" == "" ]; then \
		echo 'You need to set the LOG_ORIGIN variable'; \
		exit 1; \
	fi
	@if [ "${LOG_PUBLIC_KEY}" == "" ] || [ ! -f "${LOG_PUBLIC_KEY}" ]; then \
		echo 'You need to set the LOG_PUBLIC_KEY variable to a valid note verifier key path'; \
		exit 1; \
	fi
	@if [ "${APPLET_PUBLIC_KEY}" == "" ] || [ ! -f "${APPLET_PUBLIC_KEY}" ]; then \
		echo 'You need to set the APPLET_PUBLIC_KEY variable to a valid note verifier key path'; \
		exit 1; \
	fi
	@if [ "${OS_PUBLIC_KEY1}" == "" ] || [ ! -f "${OS_PUBLIC_KEY1}" ]; then \
		echo 'You need to set the OS_PUBLIC_KEY1 variable to a valid note verifier key path'; \
		exit 1; \
	fi
	@if [ "${OS_PUBLIC_KEY2}" == "" ] || [ ! -f "${OS_PUBLIC_KEY2}" ]; then \
		echo 'You need to set the OS_PUBLIC_KEY2 variable to a valid note verifier key path'; \
		exit 1; \
	fi

copy_applet: LOG_URL=file://$(DEV_LOG_DIR)/log/
copy_applet:
	@if [ "${APPLET_PATH}" == "" ]; then \
		echo 'You need to set the APPLET_PATH variable to a valid path for the directory holding applet elf and proof bundle files (e.g. path to armored-witness-applet/bin)'; \
		exit 1; \
	fi
	mkdir -p ${CURDIR}/trusted_os/assets
	cp ${APPLET_PATH}/trusted_applet.elf ${CURDIR}/trusted_os/assets/
	cp ${APPLET_PATH}/trusted_applet_manifest ${CURDIR}/trusted_os/assets/
	go run ./cmd/proofbundle \
		--log_origin=${LOG_ORIGIN} \
		--log_url=${LOG_URL} \
		--log_pubkey_file=${LOG_PUBLIC_KEY} \
		--manifest_pubkey_file=${APPLET_PUBLIC_KEY} \
		--manifest_file=${CURDIR}/trusted_os/assets/trusted_applet_manifest \
		--applet_file=${CURDIR}/trusted_os/assets/trusted_applet.elf \
		--output_file=${CURDIR}/trusted_os/assets/trusted_applet.proofbundle

create_dummy_applet:
	mkdir -p $(DIR)/assets
	rm -f $(DIR)/assets/trusted_applet.elf && touch $(DIR)/assets/trusted_applet.elf
	rm -f $(DIR)/assets/trusted_applet.proofbundle && touch $(DIR)/assets/trusted_applet.proofbundle

check_tamago:
	@if [ "${TAMAGO}" == "" ] || [ ! -f "${TAMAGO}" ]; then \
		echo 'You need to set the TAMAGO variable to a compiled version of https://github.com/usbarmory/tamago-go'; \
		exit 1; \
	fi
	@if [ "$(shell printf '%s\n' ${MINIMUM_TAMAGO_VERSION} ${TAMAGO_SEMVER} | sort -V | head -n1 )" != "${MINIMUM_TAMAGO_VERSION}" ]; then \
		echo "You need TamaGo >= ${MINIMUM_TAMAGO_VERSION}, found ${TAMAGO_SEMVER}" ; \
		exit 1; \
	fi

dcd:
	cp -f $(GOMODCACHE)/$(TAMAGO_PKG)/board/usbarmory/mk2/imximage.cfg $(CURDIR)/bin/$(APP).dcd

clean:
	@rm -fr $(CURDIR)/bin/* $(CURDIR)/trusted_os/assets/* $(CURDIR)/qemu.dtb

qemu: trusted_os_embedded_applet
	$(QEMU) -kernel $(CURDIR)/bin/trusted_os.elf

qemu-gdb: GOFLAGS := $(GOFLAGS:-w=)
qemu-gdb: GOFLAGS := $(GOFLAGS:-s=)
qemu-gdb: trusted_os_embedded_applet
	$(QEMU) -kernel $(CURDIR)/bin/trusted_os.elf -S -s

#### application target ####

$(APP).elf: check_tamago
	cd $(DIR) && $(GOENV) $(TAMAGO) build $(GOFLAGS) -o $(CURDIR)/bin/$(APP).elf

$(APP)_manifest:
	@if [ "${OS_PRIVATE_KEY1}" == "" ] || [ ! -f "${OS_PRIVATE_KEY1}" ]; then \
		echo 'You need to set the OS_PRIVATE_KEY1 variable to a valid note signing key path'; \
		exit 1; \
	fi
	@if [ "${OS_PRIVATE_KEY2}" == "" ] || [ ! -f "${OS_PRIVATE_KEY2}" ]; then \
		echo 'You need to set the OS_PRIVATE_KEY2 variable to a valid note signing key path'; \
		exit 1; \
	fi
	# Create manifest
	@echo ---------- Manifest --------------
	go run github.com/transparency-dev/armored-witness/cmd/manifest@561c0b09a2cc48877a8c9e59c3fbf7ffc81cdd4d \
		create \
		--git_tag=${GIT_SEMVER_TAG} \
		--git_commit_fingerprint="${REV}" \
		--firmware_file=${CURDIR}/bin/$(APP).elf \
		--firmware_type=TRUSTED_OS \
		--tamago_version=${TAMAGO_SEMVER} \
		--private_key_file=${OS_PRIVATE_KEY1} \
		--output_file=${CURDIR}/bin/${APP}_manifest
	@echo ----------------------------------
	# Now counter sign with OS_PRIVATE_KEY2
	go run github.com/transparency-dev/armored-witness/cmd/manifest@561c0b09a2cc48877a8c9e59c3fbf7ffc81cdd4d \
		create \
		--git_tag=${GIT_SEMVER_TAG} \
		--git_commit_fingerprint="${REV}" \
		--firmware_file=${CURDIR}/bin/$(APP).elf \
		--firmware_type=TRUSTED_OS \
		--tamago_version=${TAMAGO_SEMVER} \
		--private_key_file=${OS_PRIVATE_KEY2} | tail -1 >> ${CURDIR}/bin/${APP}_manifest
