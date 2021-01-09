#!/bin/bash 

SCRIPT_DIR=$(cd $(dirname $0); pwd)

source ${SCRIPT_DIR}/../../bashenv
source ${SCRIPT_DIR}/../../bashfunc

install() {
	if [ ${OS} == 'mac' ]; then 
		command="brew install starship"
	        execute "${command}"
	elif [ ${OS} == 'linux' ]; then
		if [ ${DISTRIBUTION} == 'ubuntu' ] || [ ${DISTRIBUTION} == 'debian' ]; then
		         command="sudo apt install -y cargo libssl-dev"
	                 execute "${command}"

			 command="export PATH=${PATH}:${HOME}/.cargo/bin"
	                 execute "${command}"
			 
		         command="cargo -y install starship"
	                 execute "${command}"
		fi
	fi
}

setup() {
	local config_dir="${HOME}/.config"
	if [ ! -d "${config_dir}" ]; then
		command="mkdir -p ${config_dir}"
	        execute "${command}"
	fi

	command="cp -a ${SCRIPT_DIR}/starship.toml ${config_dir}"
	execute "${command}"
}

install
setup

