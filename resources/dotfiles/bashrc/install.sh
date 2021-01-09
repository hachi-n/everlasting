#!/bin/bash 

set -eu

SCRIPT_DIR=$(cd $(dirname $0); pwd)
source ${SCRIPT_DIR}/../../bashenv
source ${SCRIPT_DIR}/../../bashfunc

install() {
	local resource_bashrc_path=""
	if [ ${OS} == 'mac' ]; then 
		resource_bashrc_path="${SCRIPT_DIR}/${OS}_bashrc"
	elif [ ${OS} == 'linux' ]; then
		if [ ${DISTRIBUTION} == 'ubuntu' ] || [ ${DISTRIBUTION} == 'debian' ]; then
		         resource_bashrc_path="${SCRIPT_DIR}/bashrc/${DISTRIBUTION}_bashrc"
		fi
	fi
        if [ -e "${resource_bashrc_path}" ]; then  
            command="cp -a ${resource_bashrc_path} ${BASHRC_PATH}"
	    execute "${command}"

	    echo "copied!!!"
	fi
	echo "please exec \"resh\""
}

install
